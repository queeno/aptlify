package action

import (
	aptlylib "github.com/queeno/aptlify/aptly"
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/repo"
	"github.com/queeno/aptlify/config"
	colour "github.com/fatih/color"
	"fmt"
	"strings"
)

var aptly = aptlylib.AptlyCli{}

func (a ActionStruct) Apply(conf *config.ConfigStruct, new_state *config.ConfigStruct) {

	switch {

	case a.ChangeType == Mirror_update:
		findMirror := mirror.AptlyMirrorStruct{ Name: a.ResourceName }
		mir := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)
		out, err := aptly.Mirror_update(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("mirror %s update failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		new_state.AddMirror(mir)
		colour.Green(fmt.Sprintf("mirror %s update succeeded", a.ResourceName))

	case a.ChangeType == Mirror_create:
		findMirror := mirror.AptlyMirrorStruct{ Name: a.ResourceName }
		mirror := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)
		out, err := aptly.Mirror_create(mirror)
		if err != nil {
			msg := fmt.Sprintf("mirror %s create failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		new_state.AddMirror(mirror)
		colour.Green(fmt.Sprintf("mirror %s create succeeded", a.ResourceName))


	case a.ChangeType == Mirror_recreate:

		out, err := aptly.Mirror_drop(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("mirror %s drop failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		colour.Green(fmt.Sprintf("mirror %s drop succeeded", a.ResourceName))

		findMirror := mirror.AptlyMirrorStruct{ Name: a.ResourceName }
		mir := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)

		out, err = aptly.Mirror_create(mir)
		if err != nil {
			msg := fmt.Sprintf("mirror %s create failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		new_state.AddMirror(mir)
		colour.Green(fmt.Sprintf("mirror %s create succeeded", a.ResourceName))

	case a.ChangeType == Repo_create:
		findRepo := repo.AptlyRepoStruct { Name: a.ResourceName }
		repo := findRepo.SearchRepoInAptlyRepos(conf.Repos)
		out, err := aptly.Repo_create(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("repo %s creation failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		new_state.AddRepo(repo)
		colour.Green(fmt.Sprintf("repo %s creation succeeded", a.ResourceName))

	case a.ChangeType == Gpg_add:
		out, err := aptly.Gpg_add(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("gpg %s creation failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		new_state.AddGpg(a.ResourceName)
		colour.Green(fmt.Sprintf("gpg %s creation succeeded", a.ResourceName))

	case a.ChangeType == Snapshot_update:
		findSnapshot := snapshot.AptlySnapshotStruct { Name: a.ResourceName }
		snapshot := findSnapshot.SearchSnapshotInAptlySnapshots(conf.Snapshots)

		var inter_snapshot_names []string
		var temp_snapshot_name string

		for _, resource := snapshot.Resources {

			temp_snapshot_name = aptly.Snapshot_create(resource)

			if snapshot.Resources.Filter != nil {
					temp_snapshot_name = aptly.Snapshot_filter(resource, temp_snapshot_name)
			}

			inter_snapshot_names = apppend(inter_snapshot_names, )

		}




	case a.ChangeType == Noop:
		if a.ResourceType == mirrorType {
			findMirror := mirror.AptlyMirrorStruct{ Name: a.ResourceName }
			mir := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)
			new_state.AddMirror(mir)
		} else if a.ResourceType == repoType {
			findRepo := repo.AptlyRepoStruct { Name: a.ResourceName }
			repo := findRepo.SearchRepoInAptlyRepos(conf.Repos)
			new_state.AddRepo(repo)
		} else if a.ResourceType == gpgType {
			new_state.AddGpg(a.ResourceName)
		}

		colour.Green(fmt.Sprintf("resource %s unchanged", a.ResourceName))
	}

}
