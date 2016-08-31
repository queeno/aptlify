package aptly

import (
	"github.com/queeno/aptlify/utils"
)

type AptlyCli struct{}

var string aptlyCmd
var aptlyCmd = "/usr/local/bin/aptly"

func (a AptlyCli) Mirror_list() ([]string, error) {

	cmd := fmt.Sprintf("%s mirror list", aptlyCmd)

	out, err := utils.Exec(cmd)
	if err != nil {
		return nil, err
	}

	mirrors := utils.SplitStringToSlice(out)

	return mirrors
}

func (a AptlyCli) Mirror_update(string mirrorName) error {
	cmd := fmt.Sprintf("%s mirror update %s", aptlyCmd, mirrorName)
	out, err := utils.Exec(cmd)
	if err != nil {
		return err
	}
	return nil
}

// mirror_create: IN: name, url, dist, components, filters, filter-with-deps, OUT: error
func (a AptlyCli) Mirror_create(AptlyMirrorStruct mirror) error {

	filter_with_deps_cmd := ""
	filter_cmd := ""

	if isEmpty(mirror.Name) {
		return nil, fmt.Errorf("Missing name from mirror")
	}
	if isEmpty(mirror.Url) {
		return nil, fmt.Errorf("Missing url from mirror")
	}
	if isEmpty(mirror.Dist) {
		return nil, fmt.Errorf("Missing distribution from mirror")
	}

	component := mirror.Component
	if isEmpty(component) {
		component = ""
	}

	if mirror.Filter != nil {
		var filter_cmds []string
		for _, filter := range mirror.Filter {
			filter_cmds = append(filter_cmds, filter.createAptlyMirrorFilter())
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

	args := createStringArray("mirror", "create", filter_cmd, filter_with_deps_cmd, mirror.Name, mirror.Url, mirror.Dist, component)

	cmd := fmt.Sprintf("%s mirror create %s %s %s %s %s %s",
		aptlyCmd, filter_cmd, filter_with_deps_cmd,
		mirror.Name, mirror.Url, mirror.Dist, component)
	out, err := utils.Exec(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (a AptlyCli) Repo_list() ([]string, error) {
	cmd := fmt.Sprintf("%s repo list", aptlyCmd)

	out, err := utils.Exec(cmd)
	if err != nil {
		return nil, err
	}
	repos := utils.SplitStringToSlice(out)
	return repos
}
