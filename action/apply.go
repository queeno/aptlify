package action

import (
	"fmt"
	colour "github.com/fatih/color"
	aptlylib "github.com/queeno/aptlify/aptly"
	"github.com/queeno/aptlify/config"
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/repo"
	snap "github.com/queeno/aptlify/snapshot"
	"strings"
)

var aptly = aptlylib.AptlyCli{}

func (a ActionStruct) Apply(conf *config.ConfigStruct, new_state *config.ConfigStruct) {

	switch {

	case a.ChangeType == Mirror_update:
		findMirror := mirror.AptlyMirrorStruct{Name: a.ResourceName}
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
		findMirror := mirror.AptlyMirrorStruct{Name: a.ResourceName}
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

		findMirror := mirror.AptlyMirrorStruct{Name: a.ResourceName}
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
		findRepo := repo.AptlyRepoStruct{Name: a.ResourceName}
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
		findSnapshot := snap.AptlySnapshotStruct{Name: a.ResourceName}
		snapshot := findSnapshot.SearchSnapshotInAptlySnapshots(conf.Snapshots)

		fmt.Println(fmt.Sprintf("Snapshot update has been passed revision: %05d", a.SnapshotRevision))

		var inter_snapshot_names []string
		var temp_snapshot_name string
		var del_snapshot_name string
		var combinedSnapshotName string
		var out []string
		var err error

		for _, resource := range snapshot.Resources {

			out, err, temp_snapshot_name = aptly.SnapshotCreate(resource)
			if err != nil {
				msg := fmt.Sprintf("snapshot %s creation failed", temp_snapshot_name)
				colour.Red(msg)
				fmt.Println(strings.Join(out, " "))
				return
			}

			if resource.Filter != nil {
				del_snapshot_name = temp_snapshot_name
				out, err, temp_snapshot_name = aptly.SnapshotFilter(resource, temp_snapshot_name)
				if err != nil {
					msg := fmt.Sprintf("snapshot %s filter failed", temp_snapshot_name)
					colour.Red(msg)
					fmt.Println(strings.Join(out, " "))
					return
				}
				out, err = aptly.SnapshotDrop(del_snapshot_name, true)
				if err != nil {
					msg := fmt.Sprintf("snapshot %s drop failed", del_snapshot_name)
					colour.Red(msg)
					fmt.Println(strings.Join(out, " "))
				}
			}

			inter_snapshot_names = append(inter_snapshot_names, temp_snapshot_name)
		}

		textRevision := fmt.Sprintf("%05d", a.SnapshotRevision)
		fmt.Println(textRevision)
		combinedSnapNameArr := []string{snapshot.Name, textRevision}
		combinedSnapshotName = strings.Join(combinedSnapNameArr, "_")
		out, err = aptly.SnapshotMerge(combinedSnapshotName, inter_snapshot_names)
		if err != nil {
			msg := fmt.Sprintf("snapshot %s merge failed", combinedSnapshotName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}

		for _, s := range inter_snapshot_names {
			out, err = aptly.SnapshotDrop(s, true)
			if err != nil {
				msg := fmt.Sprintf("snapshot %s drop failed", s)
				colour.Red(msg)
				fmt.Println(strings.Join(out, " "))
			}
		}

		s := snap.AptlySnapshotStruct{Name: combinedSnapshotName, Revision: a.SnapshotRevision}
		new_state.AddSnapshot(s)

		colour.Green(fmt.Sprintf("combined snapshot created %s at revision %d", combinedSnapshotName, a.SnapshotRevision))

	case a.ChangeType == Noop:
		if a.ResourceType == mirrorType {
			findMirror := mirror.AptlyMirrorStruct{Name: a.ResourceName}
			mir := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)
			new_state.AddMirror(mir)
		} else if a.ResourceType == repoType {
			findRepo := repo.AptlyRepoStruct{Name: a.ResourceName}
			repo := findRepo.SearchRepoInAptlyRepos(conf.Repos)
			new_state.AddRepo(repo)
		} else if a.ResourceType == gpgType {
			new_state.AddGpg(a.ResourceName)
		}

		colour.Green(fmt.Sprintf("resource %s unchanged", a.ResourceName))
	}

}
