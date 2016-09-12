package aptly

import (
	"fmt"
	"github.com/queeno/aptlify/exec"
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

var execExec = exec.Exec

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

	out, err := execExec("gpg", "--no-default-keyring", "--keyring", "trustedkeys.gpg", "--keyserver", "keys.gnupg.net", "--recv-keys", gpg_key)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (a *AptlyCli) Mirror_list() ([]string, error) {

	mirrors, err := execExec(aptlyCmd, "mirror", "list", "-raw")
	if err != nil {
		return mirrors, err
	}
	return mirrors, nil
}

func (a *AptlyCli) Mirror_update(mirrorName string) ([]string, error) {
	out, err := execExec(aptlyCmd, "mirror", "update", mirrorName)
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

	out, err := execExec(aptlyCmd, args...)

	return out, err
}

func (a *AptlyCli) Repo_list() ([]string, error) {
	repos, err := execExec(aptlyCmd, "repo", "list", "-raw")
	return repos, err
}

func (a *AptlyCli) Mirror_drop(mirrorName string) ([]string, error) {
	out, err := execExec(aptlyCmd, "mirror", "drop", mirrorName)
	return out, err
}

func (a *AptlyCli) Repo_create(repoName string) ([]string, error) {
	out, err := execExec(aptlyCmd, "repo", "create", repoName)
	return out, err
}

func (a *AptlyCli) SnapshotCreate(res snapshot.ResourceStruct) ([]string, error, string) {
	var snapNameArr = []string{res.Name, timestamp()}
	var out []string
	var err error
	snapName := strings.Join(snapNameArr, "_")
	if res.Type == "mirror" {
		if out, err = execExec(aptlyCmd, "mirror", "update", res.Name); err != nil {
			return out, err, snapName
		}
	}
	out, err = execExec(aptlyCmd, "snapshot", "create", snapName, "from", res.Type, res.Name)
	return out, err, snapName
}

func (a *AptlyCli) SnapshotFilter(res snapshot.ResourceStruct, baseSnapName string) ([]string, error, string) {
	var snapNameArr = []string{baseSnapName, "filtered"}
	snapName := strings.Join(snapNameArr, "_")
	filterCmd := ""
	var filterCmds []string
	if res.Filter != nil {
		for _, filter := range res.Filter {
			filterCmds = append(filterCmds, createAptlyMirrorFilterCommand(filter))
		}

		if len(filterCmds) > 1 {
			filterCmd = strings.Join(filterCmds, " | ")
		} else if len(filterCmds) == 1 {
			filterCmd = filterCmds[0]
		}
	}
	out, err := execExec(aptlyCmd, "snapshot", "filter", baseSnapName, snapName, filterCmd)
	return out, err, snapName
}

func (a *AptlyCli) SnapshotDrop(snapshotName string, force bool) ([]string, error) {
	forceParam := "-force=false"
	if force {
		forceParam = "-force=true"
	}
	out, err := execExec(aptlyCmd, "snapshot", "drop", forceParam, snapshotName)
	return out, err
}

func (a *AptlyCli) SnapshotMerge(combinedName string, inputSnapshotNames []string) ([]string, error) {
	args := []string{"snapshot", "merge", combinedName}
	args = append(args, inputSnapshotNames...)
	return execExec(aptlyCmd, args...)
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
