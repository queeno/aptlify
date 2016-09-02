package aptly

import (
	"fmt"
	"github.com/queeno/aptlify/config"
	"github.com/queeno/aptlify/utils"
	"strings"
)

// Check interface
var (
	_ Aptly = &AptlyCli{}
)

type AptlyCli struct{}

var aptlyCmd string = "aptly"

func (a *AptlyCli) Mirror_list() ([]string, error) {

	cmd := fmt.Sprintf("%s mirror list -raw", aptlyCmd)

	mirrors, err := utils.Exec(cmd)
	if err != nil {
		return nil, err
	}
	return mirrors, nil
}

func (a *AptlyCli) Mirror_update(mirrorName string) ([]string, error) {
	cmd := fmt.Sprintf("%s mirror update %s", aptlyCmd, mirrorName)
	out, err := utils.Exec(cmd)
	return out, err
}

// mirror_create: IN: mirror AptlyMirrorStruct, OUT: []string, error
func (a *AptlyCli) Mirror_create(mirror config.AptlyMirrorStruct) ([]string, error) {

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

	cmd := fmt.Sprintf("%s mirror create %s %s %s %s %s %s",
		aptlyCmd, filter_cmd, filter_with_deps_cmd,
		mirror.Name, mirror.Url, mirror.Dist, component)

	out, err := utils.Exec(cmd)
	return out, err
}

func (a *AptlyCli) Repo_list() ([]string, error) {
	cmd := fmt.Sprintf("%s repo list -raw", aptlyCmd)

	repos, err := utils.Exec(cmd)
	return repos, err
}

func (a *AptlyCli) Repo_add(repoName string) ([]string, error) {
	cmd := fmt.Sprintf("%s repo add %s", aptlyCmd, repoName)

	repos, err := utils.Exec(cmd)
	return repos, err
}

/* Supporting functions */

func createAptlyMirrorFilterCommand(filter config.AptlyFilterStruct) string {

	var f []string
	if !utils.IsStringEmpty(filter.Name) {
		f = append(f, fmt.Sprintf("Name (= %s)", filter.Name))
	}

	if !utils.IsStringEmpty(filter.Version) {
		f = append(f, fmt.Sprintf("$Version (= %s)", filter.Version))
	}

	f_str := ""

	if len(f) > 1 {
		f_str = fmt.Sprintf("( %s )", strings.Join(f, " , "))
	} else if len(f) == 1 {
		f_str = fmt.Sprintf("( %s )", f[0])
	}

	return f_str

}
