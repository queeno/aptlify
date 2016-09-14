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
