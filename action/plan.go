package action

import (
	"fmt"
	"errors"
	"strings"
)


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
  default:
		message = fmt.Sprintf("no case matched")
	}

	fmt.Println(message)
	return nil

}
