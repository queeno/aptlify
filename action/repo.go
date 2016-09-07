package action

import (
  "github.com/queeno/aptlify/repo"
)


func createRepoActions(configRepos []repo.AptlyRepoStruct, stateRepos []repo.AptlyRepoStruct) []ActionStruct {

	var actions = []ActionStruct{}

	for _, configRepo := range configRepos {
		actions = append(actions, compareRepos(configRepo, configRepo.SearchRepoInAptlyRepos(stateRepos)))
	}

	return actions

}

func compareRepos (a repo.AptlyRepoStruct, b repo.AptlyRepoStruct) ActionStruct {

	var ac = ActionStruct{ResourceName: a.Name, ChangeType: Noop, ResourceType: repoType }

	if a.Name != b.Name {
		ac.ChangeType = Repo_create
		ac.ResourceType = repoType
		ac.AddReasonToAction("new repo")
	}

	return ac

}
