package action

import (
  "github.com/queeno/aptlify/mirror"
)

func UpdateMirrors(mirrors []mirror.AptlyMirrorStruct) []ActionStruct {

  var actions = []ActionStruct{}
  var action = ActionStruct{}

  for _, mirror := range mirrors {
    action = ActionStruct{ResourceName: mirror.Name, ChangeType: Mirror_update, ResourceType: mirrorType}
    action.AddReasonToAction("update_mirror_requested")
    actions = append(actions, action)
  }

  return actions

}


func createMirrorActions(configMirrors []mirror.AptlyMirrorStruct, stateMirrors []mirror.AptlyMirrorStruct) []ActionStruct {

	var actions = []ActionStruct{}

	for _, configMirror := range configMirrors {
		actions = append(actions, compareMirrors(configMirror, configMirror.SearchMirrorInAptlyMirrors(stateMirrors)))
	}

	return actions

}

func compareMirrors (a mirror.AptlyMirrorStruct, b mirror.AptlyMirrorStruct) ActionStruct {

	var ac = ActionStruct{ResourceName: a.Name, ChangeType: Noop, ResourceType: mirrorType }

	if b.IsEmpty() {
		ac.ChangeType = Mirror_create
		ac.AddReasonToAction("new_mirror")
		ac.ResourceType = mirrorType
		return ac
	}

	if a.Url != b.Url {
		ac.AddReasonToAction("url")
		ac.ChangeType = Mirror_recreate
	}

	if a.Dist != b.Dist {
		ac.AddReasonToAction("distribution")
		ac.ChangeType = Mirror_recreate
	}

	if a.Component != b.Component {
		ac.AddReasonToAction("component")
		ac.ChangeType = Mirror_recreate
	}

	if a.FilterDeps != b.FilterDeps {
		ac.AddReasonToAction("filter-deps")
		ac.ChangeType = Mirror_recreate
	}

	if diff, _, _ := mirror.DiffFilterSlices(a.Filter, b.Filter); diff != nil {
		ac.ChangeType = Mirror_recreate
		ac.AddReasonToAction("filter")
	}

	ac.ResourceType = mirrorType
	return ac

}
