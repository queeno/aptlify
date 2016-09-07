package action

import (
	aptlylib "github.com/queeno/aptlify/aptly"
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/config"
	colour "github.com/fatih/color"
	"fmt"
	"errors"
	"strings"
)

var aptly = aptlylib.AptlyCli{}

func (a ActionStruct) Apply(config *config.ConfigStruct) error {

	if a.isEmpty() {
		return errors.New("Uninitialised action")
	}

	switch {

	case a.ChangeType == Mirror_update:
		out, err := aptly.Mirror_update(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("mirror %s update failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return errors.New(msg)
		}
		colour.Green(fmt.Sprintf("mirror %s update succeeded", a.ResourceName))

	case a.ChangeType == Mirror_create:
		findMirror := mirror.AptlyMirrorStruct{ Name: a.ResourceName }
		mirror := findMirror.SearchMirrorInAptlyMirrors(config.Mirrors)
		out, err := aptly.Mirror_create(mirror)
		if err != nil {
			msg := fmt.Sprintf("mirror %s create failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return errors.New(msg)
		}
		colour.Green(fmt.Sprintf("mirror %s create succeeded", a.ResourceName))


	case a.ChangeType == Mirror_recreate:

		out, err := aptly.Mirror_drop(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("mirror %s drop failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return errors.New(msg)
		}
		colour.Green(fmt.Sprintf("mirror %s drop succeeded", a.ResourceName))

		findMirror := mirror.AptlyMirrorStruct{ Name: a.ResourceName }
		mir := findMirror.SearchMirrorInAptlyMirrors(config.Mirrors)

		out, err = aptly.Mirror_create(mir)
		if err != nil {
			msg := fmt.Sprintf("mirror %s create failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return errors.New(msg)
		}
		colour.Green(fmt.Sprintf("mirror %s create succeeded", a.ResourceName))

	case a.ChangeType == Repo_create:
		out, err := aptly.Repo_create(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("repo %s creation failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return errors.New(msg)
		}
		colour.Green(fmt.Sprintf("repo %s creation succeeded", a.ResourceName))

	case a.ChangeType == Gpg_add:

	case a.ChangeType == Noop:
		colour.Green(fmt.Sprintf("resource %s unchanged"))
	}

	return nil

}
