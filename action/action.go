package action

/* This should represent an atomic change to aptly

- State of change
- Resource to be changed
- Response
*/

import (
	"errors"
	"fmt"
	"strings"
	"github.com/queeno/aptlify/config"
)

type ActionStruct struct {
	ResourceName	string
	ChangeType   int
	changeReason []string
}

func (a ActionStruct) isEmpty() bool {
	if a.ChangeType == 0 && a.changeReason == nil {
		return true
	}
	return false
}

func (a *ActionStruct) AddReasonToAction(reason string) {
	if reason == "" {
		a.changeReason = nil
	}
	a.changeReason = append(a.changeReason, reason)
}

func (c ActionStruct) Apply() {


}

func (a ActionStruct) Plan() error {

	var message string

	if a.isEmpty() {
		return errors.New("Uninitialised action")
	}

	switch {
	case a.ChangeType == Mirror_update:
		message = fmt.Sprintf("+mirror %s will be updated. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == Mirror_create:
		message = fmt.Sprintf("+mirror %s will be created. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == Mirror_recreate:
		message = fmt.Sprintf("+/-mirror %s will be recreated. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == Repo_create:
		message = fmt.Sprintf("+repo %s will be created. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == Gpg_add:
		message = fmt.Sprintf("+gpg key %s will be added. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == Noop:
		message = fmt.Sprintf("resource unchanged: %s", a.ResourceName)
	}

	fmt.Println(message)
	return nil

}

func CreateActions(config *config.ConfigStruct, state *config.ConfigStruct) []ActionStruct {

	mirrorActions := createMirrorActions(config.Mirrors, state.Mirrors)
	repoActions := createRepoActions(config.Repos, state.Repos)
	gpgActions := createGpgActions(config.Gpg_keys, state.Gpg_keys)

	return append(append(mirrorActions, repoActions...), gpgActions...)
}
