package repo

import (
	"github.com/queeno/aptlify/action"
)

type AptlyRepoStruct struct {
	Name 	string	`json:"name"`
}

func SearchRepoInAptlyRepos(thisRepo AptlyRepoStruct, repos []AptlyRepoStruct) AptlyRepoStruct {

	for _, repo := range repos {
			if repo.Name == thisRepo.Name {
				return repo
			}
	}

	return AptlyRepoStruct{}
}

func CreateRepoActions(configRepos []AptlyRepoStruct, stateRepos []AptlyRepoStruct) []action.ActionStruct {

	var actions = []action.ActionStruct{}

	for _, configRepo := range configRepos {
		actions = append(actions, configRepo.Compare(SearchRepoInAptlyRepos(configRepo, stateRepos)))
	}

	return actions

}

func (a AptlyRepoStruct) Compare (b AptlyRepoStruct) action.ActionStruct {

	var ac = action.ActionStruct{ResourceName: a.Name, ChangeType: action.Noop }

	if a.Name != b.Name {
		ac.ChangeType = action.Repo_create
		ac.AddReasonToAction("new repo")
	}

	return ac

}
