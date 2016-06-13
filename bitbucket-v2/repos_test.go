package bitbucket_v2

import "testing"

var client = New("DmuycxS3kZ7xCsg5Uw", "VzSv8YSYUsEfZEkPu3GQ9y7jfbMYScxm",
	"wbBURukaFAktHqmvfL", "tkBGe33jJjEx9zhr3bQeMQy7tvgrxgQE")

func Test_ListRepos(t *testing.T) {

	// LIST of repositories
	repos, err := client.ListRepos("xuzhenglun", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if len(repos.Values) != 4 {
		t.Errorf("List of /user repositories returned empty set")
	}
}

func Test_RepoInfo(t *testing.T) {
	// FIND the named repo
	repo, err := client.RepoInfo("xuzhenglun", "test2")
	if err != nil {
		t.Error(err)
		return
	}

	if repo.Scm != "git" {
		t.Errorf("repo slug [%v]; want [%v]", repo.Scm, "git")
	}
}

func Test_Tags(t *testing.T) {
	tags, err := client.Tags("xuzhenglun", "test2", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if len(tags.Values) != 2 {
		t.Errorf("repo tags num [%v]; want [%v]", len(tags.Values), "2")
	}
}

func Test_Branch(t *testing.T) {
	branches, err := client.Branches("xuzhenglun", "test2", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if len(branches.Values) != 2 {
		t.Errorf("repo tags num [%v]; want [%v]", len(branches.Values), "2")
	}
}

func Test_Forks(t *testing.T) {
	forks, err := client.Forks("xuzhenglun", "test2", 1)

	if err != nil {
		t.Error(err)
		return
	}

	if len(forks.Values) != 1 {
		t.Errorf("repo tags num [%v]; want [%v]", len(forks.Values), "1")
	}
}
