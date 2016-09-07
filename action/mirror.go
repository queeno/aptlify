package action

import (
  "github.com/queeno/aptlify/mirror"
)

func createMirrorActions(configMirrors []mirror.AptlyMirrorStruct, stateMirrors []mirror.AptlyMirrorStruct) []ActionStruct {

	var actions = []ActionStruct{}

	for _, configMirror := range configMirrors {
		actions = append(actions, compareMirrors(configMirror, configMirror.SearchMirrorInAptlyMirrors(stateMirrors)))
	}

	return actions

}

func compareMirrors (a mirror.AptlyMirrorStruct, b mirror.AptlyMirrorStruct) ActionStruct {

	var ac = ActionStruct{ResourceName: a.Name, ChangeType: Noop }

	if b.IsEmpty() {
		ac.ChangeType = Mirror_create
		ac.AddReasonToAction("new_mirror")
		return ac
	}

	if a.Url != b.Url {
		ac.ChangeType = Mirror_recreate
		ac.AddReasonToAction("url")
	}

	if a.Dist != b.Dist {
		ac.ChangeType = Mirror_recreate
		ac.AddReasonToAction("distribution")
	}

	if a.Component != b.Component {
		ac.ChangeType = Mirror_recreate
		ac.AddReasonToAction("component")
	}

	if a.FilterDeps != b.FilterDeps {
		ac.ChangeType = Mirror_recreate
		ac.AddReasonToAction("filter-deps")
	}

	if diff, _, _ := mirror.DiffFilterSlices(a.Filter, b.Filter); diff != nil {
		ac.ChangeType = Mirror_recreate
		ac.AddReasonToAction("filter")
	}

	return ac

}
