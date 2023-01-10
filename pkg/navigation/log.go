package navigation

import (
	"log"
	"strconv"
	"strings"

	"github.com/DiegoAraujoJS/webdev-git-server/pkg/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const prefix = "refs/heads/"

func GetRemoteBranches() []*plumbing.Reference {
	repo := utils.GetRepository()

	remote, err := repo.Remote("origin")
	if err != nil {
		log.Fatal(err.Error())
	}

	ref_list, err := remote.List(&git.ListOptions{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var remote_branches []*plumbing.Reference
	for _, ref := range ref_list {
		if !strings.HasPrefix(ref.Name().String(), prefix) {
			continue
		}
		remote_branches = append(remote_branches, ref)
	}
	return remote_branches
}

type BranchResponse struct {
    Commit *object.Commit `json:"commit"`
    NewReference string `json:"new_reference"`
}

func GetReleaseBranchesWithTheirVersioning() []*BranchResponse {
    repo := utils.GetRepository()

    var result []*BranchResponse

    branches, err := repo.Branches()
    if err != nil {
        log.Fatal(err.Error())
    }
    for {
        branch, err := branches.Next()
        if err != nil {
            break
        }
        if strings.Contains(branch.Name().String(), "RELEASE") {
            commits_from_master := utils.GetCommitsFromBranchToMaster(branch)
            for k, c := range commits_from_master {
                version_number_string := strings.Split(branch.Name().String(), "_")[1]
                version := version_number_string + "." + strconv.Itoa(len(commits_from_master) - k)

                result = append(result, &BranchResponse{
                    Commit: c,
                    NewReference: version,
                })
            }
        }
    }
    return result
}
