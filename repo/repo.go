package repo

type AptlyRepoStruct struct {
	Name 	string	`json:"name"`
}

func (thisRepo AptlyRepoStruct) SearchRepoInAptlyRepos(repos []AptlyRepoStruct) AptlyRepoStruct {

	for _, repo := range repos {
			if repo.Name == thisRepo.Name {
				return repo
			}
	}

	return AptlyRepoStruct{}
}
