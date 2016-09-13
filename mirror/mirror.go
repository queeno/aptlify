package mirror

type AptlyFilterStruct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (f AptlyFilterStruct) IsEmpty() bool {
	if f.Name == "" {
		return true
	}
	return false
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

func (thisMirror AptlyMirrorStruct) SearchMirrorInAptlyMirrors(Mirrors []AptlyMirrorStruct) AptlyMirrorStruct {

	for _, mirror := range Mirrors {
		if mirror.Name == thisMirror.Name {
			return mirror
		}
	}

	return AptlyMirrorStruct{}
}

func (a AptlyMirrorStruct) IsEmpty() bool {
	if a.Name == "" {
		return true
	}
	return false
}

/*func mirrorExists(mirrorName string) bool {
	mirror, _ := context.CollectionFactory().RemoteRepoCollection().ByName(mirrorName)

	if mirror == nil {
		return false
	}
	return true
}

func repoExists(repoName string) bool {
	repo, _ := context.CollectionFactory().LocalRepoCollection().ByName(repoName)

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

	fStr := ""

	if len(f) > 1 {
		fStr = fmt.Sprintf("( %s )", strings.Join(f, " , "))
	} else if len(f) == 1 {
		fStr = fmt.Sprintf("( %s )", f[0])
	}

	return fStr

}

func genCreateAptlyRepoCmd(repoName string) ([]string, error) {
	var cmd []string
	if repoName == "" {
		return cmd, fmt.Errorf("Missing mirror name")
	}
	cmd = createStringArray("repo", "create", repoName)
	fmt.Println(cmd)
	return cmd, nil
}

func (mirror *AptlyMirrorStruct) genCreateAptlyMirrorCmd() ([]string, error) {

	filterWithDepsCmd := ""
	filterCmd := ""

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

		var filterCmds []string
		for _, filter := range mirror.Filter {
			filterCmds = append(filterCmds, filter.createAptlyMirrorFilter())
		}

		if len(filterCmds) > 1 {
			filterCmd = fmt.Sprintf("-filter=%s", strings.Join(filterCmds, " | "))
		} else if len(filter_cmds) == 1 {
			filterCmd = fmt.Sprintf("-filter=%s", filterCmds[0])
		}
	}

	if mirror.FilterDeps {
		filterWithDepsCmd = "-filter-with-deps"
	}

	args := createStringArray("mirror", "create", filterCmd, filterWithDepsCmd, mirror.Name, mirror.Url, mirror.Dist, component)

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
	var cmdCreate []string
	var e error

	for _, repoName := range c.Repos {
		if !repoExists(repoName) {
			cmdCreate, e = genCreateAptlyRepoCmd(repoName)
			if e != nil {
				return nil, e
			}
			commands = append(commands, cmdCreate)
		}
	}
	return commands, nil
}

func (c *aptlySetupConfigStruct) genCreateAndUpdateMirrorCmds() ([][]string, error) {

	var commands [][]string
	var cmdCreate []string
	var cmdUpdate []string
	var e error

	for _, mirror := range c.Mirrors {

		if !mirrorExists(mirror.Name) {
			cmdCreate, e = mirror.genCreateAptlyMirrorCmd()
			if e != nil {
				return nil, e
			}
			commands = append(commands, cmdCreate)
		}

		cmdUpdate, e = mirror.genUpdateAptlyMirrorCmd()
		if e != nil {
			return nil, e
		}
		commands = append(commands, cmdUpdate)
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

func RetrieveState(confMirrors []AptlyMirrorStruct) error {

ctx.Logging.Info.Println("retrieving mirror information from aptly")

a := &aptly.AptlyCli{}

if aptlyMirrors, err := a.MirrorList(); err != nil {
	return err
}

ctx.Logging.Info.Println("retrieving aptlify mirror configuration")
confMirrors := context.Config().Mirrors

for _, mirror := range confMirrors {
		if isMirrorInAptly(mirror, aptlyMirrors) {

		}
}

}*/
