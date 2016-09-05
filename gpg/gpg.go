package gpg

import (
	"github.com/queeno/aptlify/action"
)

type AptlyGpgStruct struct {
	Fingerprint 	[]string     `json:"fingerprint"`
}

func CreateGpgActions(configGpgs AptlyGpgStruct, stateGpgs AptlyGpgStruct) []action.ActionStruct {

	var actions = []action.ActionStruct{}

	for _, fprnt := range configGpgs.Fingerprint {
		actions = append(actions, stateGpgs.Compare(fprnt))
	}

	return actions

}

func (a AptlyGpgStruct) Compare (b string) action.ActionStruct {

	var ac = action.ActionStruct{ResourceName: b, ChangeType: action.Gpg_add }
	ac.AddReasonToAction("GPG key not found")


	for _, gpg := range a.Fingerprint {
		if gpg == b {
			ac.ResourceName = b
			ac.ChangeType = action.Noop
			ac.AddReasonToAction("")
		}
	}

	return ac

}
