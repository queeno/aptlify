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

func (a ActionStruct) Apply(conf *config.ConfigStruct, newState *config.ConfigStruct) {

	switch {

	case a.ChangeType == MirrorUpdate:
		findMirror := mirror.AptlyMirrorStruct{Name: a.ResourceName}
		mir := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)
		out, err := aptly.MirrorUpdate(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("mirror %s update failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		newState.AddMirror(mir)
		colour.Green(fmt.Sprintf("mirror %s update succeeded", a.ResourceName))

	case a.ChangeType == MirrorCreate:
		findMirror := mirror.AptlyMirrorStruct{Name: a.ResourceName}
		mirror := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)
		out, err := aptly.MirrorCreate(mirror)
		if err != nil {
			msg := fmt.Sprintf("mirror %s create failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		newState.AddMirror(mirror)
		colour.Green(fmt.Sprintf("mirror %s create succeeded", a.ResourceName))

	case a.ChangeType == MirrorRecreate:

		out, err := aptly.MirrorDrop(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("mirror %s drop failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		colour.Green(fmt.Sprintf("mirror %s drop succeeded", a.ResourceName))

		findMirror := mirror.AptlyMirrorStruct{Name: a.ResourceName}
		mir := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)

		out, err = aptly.MirrorCreate(mir)
		if err != nil {
			msg := fmt.Sprintf("mirror %s create failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		newState.AddMirror(mir)
		colour.Green(fmt.Sprintf("mirror %s create succeeded", a.ResourceName))

	case a.ChangeType == RepoCreate:
		findRepo := repo.AptlyRepoStruct{Name: a.ResourceName}
		repo := findRepo.SearchRepoInAptlyRepos(conf.Repos)
		out, err := aptly.RepoCreate(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("repo %s creation failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		newState.AddRepo(repo)
		colour.Green(fmt.Sprintf("repo %s creation succeeded", a.ResourceName))

	case a.ChangeType == GpgAdd:
		out, err := aptly.GpgAdd(a.ResourceName)
		if err != nil {
			msg := fmt.Sprintf("gpg %s creation failed", a.ResourceName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}
		newState.AddGpg(a.ResourceName)
		colour.Green(fmt.Sprintf("gpg %s creation succeeded", a.ResourceName))

	case a.ChangeType == SnapshotUpdate:
		findSnapshot := snap.AptlySnapshotStruct{Name: a.ResourceName}
		snapshot := findSnapshot.SearchSnapshotInAptlySnapshots(conf.Snapshots)

		fmt.Println(snapshot.Name)
		fmt.Println(fmt.Sprintf("Snapshot update has been passed revision: %05d", a.SnapshotRevision))

		var interSnapshotNames []string
		var tempSnapshotName string
		var delSnapshotName string
		var combinedSnapshotName string
		var out []string
		var err error

		for _, resource := range snapshot.Resources {

			out, err, tempSnapshotName = aptly.SnapshotCreate(resource)
			if err != nil {
				msg := fmt.Sprintf("snapshot %s creation failed", tempSnapshotName)
				colour.Red(msg)
				fmt.Println(strings.Join(out, " "))
				return
			}

			if resource.Filter != nil {
				delSnapshotName = tempSnapshotName
				out, err, tempSnapshotName = aptly.SnapshotFilter(resource, tempSnapshotName)
				if err != nil {
					msg := fmt.Sprintf("snapshot %s filter failed", tempSnapshotName)
					colour.Red(msg)
					fmt.Println(strings.Join(out, " "))
					return
				}
				out, err = aptly.SnapshotDrop(delSnapshotName, true)
				if err != nil {
					msg := fmt.Sprintf("snapshot %s drop failed", delSnapshotName)
					colour.Red(msg)
					fmt.Println(strings.Join(out, " "))
				}
			}

			interSnapshotNames = append(interSnapshotNames, tempSnapshotName)
		}

		textRevision := fmt.Sprintf("%05d", a.SnapshotRevision)
		combinedSnapNameArr := []string{snapshot.Name, textRevision}
		combinedSnapshotName = strings.Join(combinedSnapNameArr, "_")
		out, err = aptly.SnapshotMerge(combinedSnapshotName, interSnapshotNames)
		if err != nil {
			msg := fmt.Sprintf("snapshot %s merge failed", combinedSnapshotName)
			colour.Red(msg)
			fmt.Println(strings.Join(out, " "))
			return
		}

		for _, s := range interSnapshotNames {
			out, err = aptly.SnapshotDrop(s, true)
			if err != nil {
				msg := fmt.Sprintf("snapshot %s drop failed", s)
				colour.Red(msg)
				fmt.Println(strings.Join(out, " "))
			}
		}

		s := snap.AptlySnapshotStruct{Name: snapshot.Name, Revision: a.SnapshotRevision}
		newState.AddSnapshot(s)

		colour.Green(fmt.Sprintf("combined snapshot created %s at revision %d", snapshot.Name, a.SnapshotRevision))

	case a.ChangeType == Noop:
		if a.ResourceType == mirrorType {
			findMirror := mirror.AptlyMirrorStruct{Name: a.ResourceName}
			mir := findMirror.SearchMirrorInAptlyMirrors(conf.Mirrors)
			newState.AddMirror(mir)
		} else if a.ResourceType == repoType {
			findRepo := repo.AptlyRepoStruct{Name: a.ResourceName}
			repo := findRepo.SearchRepoInAptlyRepos(conf.Repos)
			newState.AddRepo(repo)
		} else if a.ResourceType == gpgType {
			newState.AddGpg(a.ResourceName)
		}

		colour.Green(fmt.Sprintf("resource %s unchanged", a.ResourceName))
	}

}
