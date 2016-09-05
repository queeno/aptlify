package mirror

import (
	"github.com/queeno/aptlify/action"
)

type AptlyFilterStruct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (a AptlyFilterStruct) Equals(b AptlyFilterStruct) bool {
	return (a.Name == b.Name) && (a.Version == b.Version)
}


type AptlyMirrorStruct struct {
	Name       string              `json:"name"`
	Url        string              `json:"url"`
	Dist       string              `json:"dist"`
	Component  string              `json:"component"`
	Filter     []AptlyFilterStruct `json:"filter"`
	FilterDeps bool                `json:"filter-with-deps"`
}


func SearchMirrorInAptlyMirrors(thisMirror AptlyMirrorStruct, Mirrors []AptlyMirrorStruct) AptlyMirrorStruct {

	for _, mirror := range Mirrors {
			if mirror.Name == thisMirror.Name {
				return mirror
			}
	}

	return AptlyMirrorStruct{}
}

func CreateMirrorActions(configMirrors []AptlyMirrorStruct, stateMirrors []AptlyMirrorStruct) []action.ActionStruct {

	var actions = []action.ActionStruct{}

	for _, configMirror := range configMirrors {
		actions = append(actions, configMirror.Compare(SearchMirrorInAptlyMirrors(configMirror, stateMirrors)))
	}

	return actions

}

func (a AptlyMirrorStruct) Compare (b AptlyMirrorStruct) action.ActionStruct {

	var ac = action.ActionStruct{ResourceName: a.Name, ChangeType: action.Noop }

	if a.Url != b.Url {
		ac.ChangeType = action.Mirror_recreate
		ac.AddReasonToAction("url")
	}

	if a.Dist != b.Dist {
		ac.ChangeType = action.Mirror_recreate
		ac.AddReasonToAction("distribution")
	}

	if a.Component != b.Component {
		ac.ChangeType = action.Mirror_recreate
		ac.AddReasonToAction("component")
	}

	if a.FilterDeps != b.FilterDeps {
		ac.ChangeType = action.Mirror_recreate
		ac.AddReasonToAction("filter-deps")
	}

	if diff, _, _ := diffFilterSlices(a.Filter, b.Filter); diff != nil {
		ac.ChangeType = action.Mirror_recreate
		ac.AddReasonToAction("filter")
	}

	return ac

}


/*func mirrorExists(mirror_name string) bool {
	mirror, _ := context.CollectionFactory().RemoteRepoCollection().ByName(mirror_name)

	if mirror == nil {
		return false
	}
	return true
}

func repoExists(repo_name string) bool {
	repo, _ := context.CollectionFactory().LocalRepoCollection().ByName(repo_name)

	if repo == nil {
		return false
	}
	return true
}

func (filter *AptlyFilterStruct) createAptlyMirrorFilter() string {

	var f []string
	if !isEmpty(filter.Name) {
		f = append(f, fmt.Sprintf("Name (= %s)", filter.Name))
	}

	if !isEmpty(filter.Version) {
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

func genCreateAptlyRepoCmd(repo_name string) ([]string, error) {
	var cmd []string
	if repo_name == "" {
		return cmd, fmt.Errorf("Missing mirror name")
	}
	cmd = createStringArray("repo", "create", repo_name)
	fmt.Println(cmd)
	return cmd, nil
}

func (mirror *AptlyMirrorStruct) genCreateAptlyMirrorCmd() ([]string, error) {

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

	return args, nil
}

func (mirror *AptlyMirrorStruct) genUpdateAptlyMirrorCmd() ([]string, error) {

	if isEmpty(mirror.Name) {
		return nil, fmt.Errorf("Missing name from mirror")
	}

	args := createStringArray("mirror", "update", mirror.Name)
	return args, nil

}

func (c *aptlySetupConfigStruct) genCreateReposCmds() ([][]string, error) {

	var commands [][]string
	var cmd_create []string
	var e error

	for _, repo_name := range c.Repos {
		if !repoExists(repo_name) {
			cmd_create, e = genCreateAptlyRepoCmd(repo_name)
			if e != nil {
				return nil, e
			}
			commands = append(commands, cmd_create)
		}
	}
	return commands, nil
}

func (c *aptlySetupConfigStruct) genCreateAndUpdateMirrorCmds() ([][]string, error) {

	var commands [][]string
	var cmd_create []string
	var cmd_update []string
	var e error

	for _, mirror := range c.Mirrors {

		if !mirrorExists(mirror.Name) {
			cmd_create, e = mirror.genCreateAptlyMirrorCmd()
			if e != nil {
				return nil, e
			}
			commands = append(commands, cmd_create)
		}

		cmd_update, e = mirror.genUpdateAptlyMirrorCmd()
		if e != nil {
			return nil, e
		}
		commands = append(commands, cmd_update)
	}

	return commands, nil

}

func aptlyRunSetup(cmd *commander.Command, args []string) error {

	// Get setup configuration
	filename := context.Config().SetupFile

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Unable to read file: %s", err)
	}

	var c aptlySetupConfigStruct

	json.Unmarshal(f, &c)

	mirrorCommands, err := c.genCreateAndUpdateMirrorCmds()
	if err != nil {
		return err
	}

	repoCommands, err := c.genCreateReposCmds()
	if err != nil {
		return err
	}

	var commands [][]string

	if repoCommands != nil {
		for _, cmd := range repoCommands {
			commands = append(commands, cmd)
		}
	}

	if mirrorCommands != nil {
		for _, cmd := range mirrorCommands {
			commands = append(commands, cmd)
		}
	}

	if commands != nil {
		err = aptlyTaskRunCommands(commands)
	}

	return err
}

func RetrieveState(conf_mirrors []AptlyMirrorStruct) error {

ctx.Logging.Info.Println("retrieving mirror information from aptly")

a := &aptly.AptlyCli{}

if aptly_mirrors, err := a.Mirror_list(); err != nil {
	return err
}

ctx.Logging.Info.Println("retrieving aptlify mirror configuration")
conf_mirrors := context.Config().Mirrors

for _, mirror := range conf_mirrors {
		if isMirrorInAptly(mirror, aptly_mirrors) {

		}
}

}*/
