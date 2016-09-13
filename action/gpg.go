package action

import (
	"github.com/queeno/aptlify/gpg"
)

func createGpgActions(configGpgs gpg.AptlyGpgStruct, stateGpgs gpg.AptlyGpgStruct) []ActionStruct {

	var actions = []ActionStruct{}

	for _, fprnt := range configGpgs.Fingerprint {
		actions = append(actions, compareGpg(stateGpgs, fprnt))
	}

	return actions

}

func compareGpg(a gpg.AptlyGpgStruct, b string) ActionStruct {

	var ac = ActionStruct{ResourceName: b, ChangeType: GpgAdd, ResourceType: gpgType}
	ac.AddReasonToAction("GPG key not found")

	for _, gpg := range a.Fingerprint {
		if gpg == b {
			ac.ResourceName = b
			ac.ResourceType = gpgType
			ac.ChangeType = Noop
			ac.AddReasonToAction("")
		}
	}

	return ac

}
