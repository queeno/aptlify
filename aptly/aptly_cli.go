package aptly

import (
	"fmt"
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/snapshot"
	"github.com/queeno/aptlify/utils"
	"strings"
	"time"
)

// Check interface
var (
	_ Aptly = &AptlyCli{}
)

type AptlyCli struct{}

var aptlyCmd string = "aptly"

func cleanSlice(slice []string) []string {

	var clean_slice []string

	for _, elem := range slice {
		if elem != "" {
			clean_slice = append(clean_slice, elem)
		}
	}
	return clean_slice
}

func (a *AptlyCli) Gpg_add(gpg_key string) ([]string, error) {

	out, err := utils.Exec("gpg", "--no-default-keyring", "--keyring", "trustedkeys.gpg", "--keyserver", "keys.gnupg.net", "--recv-keys", gpg_key)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (a *AptlyCli) Mirror_list() ([]string, error) {

	mirrors, err := utils.Exec(aptlyCmd, "mirror", "list", "-raw")
	if err != nil {
		return mirrors, err
	}
	return mirrors, nil
}

func (a *AptlyCli) Mirror_update(mirrorName string) ([]string, error) {
	out, err := utils.Exec(aptlyCmd, "mirror", "update", mirrorName)
	return out, err
}

// mirror_create: IN: mirror AptlyMirrorStruct, OUT: []string, error
func (a *AptlyCli) Mirror_create(mirror mirror.AptlyMirrorStruct) ([]string, error) {

	filter_with_deps_cmd := ""
	filter_cmd := ""

	if utils.IsStringEmpty(mirror.Name) {
		return nil, fmt.Errorf("Missing name from mirror")
	}
	if utils.IsStringEmpty(mirror.Url) {
		return nil, fmt.Errorf("Missing url from mirror")
	}
	if utils.IsStringEmpty(mirror.Dist) {
		return nil, fmt.Errorf("Missing distribution from mirror")
	}

	component := mirror.Component
	if utils.IsStringEmpty(component) {
		component = ""
	}

	if mirror.Filter != nil {
		var filter_cmds []string
		for _, filter := range mirror.Filter {
			filter_cmds = append(filter_cmds, createAptlyMirrorFilterCommand(filter))
		}

		if len(filter_cmds) > 1 {
			filter_cmd = fmt.Sprintf("-filter=%s", strings.Join(filter_cmds, " | "))
		} else if len(filter_cmds) == 1 {
			filter_cmd = fmt.Sprintf("-filter=%s", filter_cmds[0])
		}
	}

	if mirror.FilterDeps {
		filter_with_deps_cmd = "-filter-with-deps"
	}

	args := []string{"mirror", "create", filter_cmd, filter_with_deps_cmd, mirror.Name, mirror.Url, mirror.Dist, component}
	args = cleanSlice(args)

	out, err := utils.Exec(aptlyCmd, args...)

	return out, err
}

func (a *AptlyCli) Repo_list() ([]string, error) {
	repos, err := utils.Exec(aptlyCmd, "repo", "list", "-raw")
	return repos, err
}

func (a *AptlyCli) Mirror_drop(mirrorName string) ([]string, error) {
	out, err := utils.Exec(aptlyCmd, "mirror", "drop", mirrorName)
	return out, err
}

func (a *AptlyCli) Repo_create(repoName string) ([]string, error) {
	out, err := utils.Exec(aptlyCmd, "repo", "create", repoName)
	return out, err
}

func SnapshotCreate(res snapshot.ResourceStruct) ([]string, error, string) {
	var snapNameArr = []string{res.Name, timestamp()}
	snapName := strings.Join(snapNameArr, "_")
	out, err := utils.Exec(aptlyCmd, "snapshot", "create", snapName, "from", res.Type, res.Name)
	return out, err, snapName
}

/* Supporting functions */

func createAptlyMirrorFilterCommand(filter mirror.AptlyFilterStruct) string {

	var f []string
	if !utils.IsStringEmpty(filter.Name) {
		f = append(f, fmt.Sprintf("Name (= %s )", filter.Name))
	}

	if !utils.IsStringEmpty(filter.Version) {
		f = append(f, fmt.Sprintf("$Version (= %s )", filter.Version))
	}

	f_str := ""

	if len(f) > 1 {
		f_str = fmt.Sprintf("( %s )", strings.Join(f, " , "))
	} else if len(f) == 1 {
		f_str = fmt.Sprintf("( %s )", f[0])
	}

	return f_str

}

func timestamp() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d_%02d:%02d:%02d", t.Year(), t.Month(),
		t.Day(), t.Hour(), t.Minute(), t.Second())
}
