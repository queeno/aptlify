package aptly

import (
	"fmt"
	"strings"
	"time"

	"github.com/queeno/aptlify/exec"
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/snapshot"
	"github.com/queeno/aptlify/utils"
)

// Check interface
var (
	_ Aptly = &AptlyCli{}
)

type AptlyCli struct{}

var aptlyCmd string = "aptly"

var execExec = exec.Exec

var timestamp = realTimestamp

func cleanSlice(slice []string) []string {

	var cleanSlice []string

	for _, elem := range slice {
		if elem != "" {
			cleanSlice = append(cleanSlice, elem)
		}
	}
	return cleanSlice
}

func (a *AptlyCli) GpgAdd(gpg_key string) ([]string, error) {

	out, err := execExec("gpg", "--no-default-keyring", "--keyring", "trustedkeys.gpg", "--keyserver", "keys.gnupg.net", "--recv-keys", gpg_key)
	return out, err
}

func (a *AptlyCli) MirrorList() ([]string, error) {

	mirrors, err := execExec(aptlyCmd, "mirror", "list", "-raw")
	return mirrors, err
}

func (a *AptlyCli) MirrorUpdate(mirrorName string) ([]string, error) {
	out, err := execExec(aptlyCmd, "mirror", "update", mirrorName)
	return out, err
}

// mirrorCreate: IN: mirror AptlyMirrorStruct, OUT: []string, error
func (a *AptlyCli) MirrorCreate(mirror mirror.AptlyMirrorStruct) ([]string, error) {

	filterWithDepsCmd := ""
	filterCmd := ""

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
		var filterCmds []string
		for _, filter := range mirror.Filter {
			filterCmds = append(filterCmds, createAptlyMirrorFilterCommand(filter))
		}

		if len(filterCmds) > 1 {
			filterCmd = fmt.Sprintf("-filter=%s", strings.Join(filterCmds, " | "))
		} else if len(filterCmds) == 1 {
			filterCmd = fmt.Sprintf("-filter=%s", filterCmds[0])
		}
	}

	if mirror.FilterDeps {
		filterWithDepsCmd = "-filter-with-deps"
	}

	args := []string{"mirror", "create", filterCmd, filterWithDepsCmd, mirror.Name, mirror.Url, mirror.Dist, component}
	args = cleanSlice(args)

	out, err := execExec(aptlyCmd, args...)

	return out, err
}

func (a *AptlyCli) RepoList() ([]string, error) {
	repos, err := execExec(aptlyCmd, "repo", "list", "-raw")
	return repos, err
}

func (a *AptlyCli) MirrorDrop(mirrorName string) ([]string, error) {
	out, err := execExec(aptlyCmd, "mirror", "drop", mirrorName)
	return out, err
}

func (a *AptlyCli) RepoCreate(repoName string) ([]string, error) {
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
	args := []string{"snapshot", "merge", "-no-remove", combinedName}
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

	fStr := ""

	if len(f) > 1 {
		fStr = fmt.Sprintf("( %s )", strings.Join(f, " , "))
	} else if len(f) == 1 {
		fStr = fmt.Sprintf("( %s )", f[0])
	}

	return fStr

}

func realTimestamp() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d_%02d:%02d:%02d", t.Year(), t.Month(),
		t.Day(), t.Hour(), t.Minute(), t.Second())
}
