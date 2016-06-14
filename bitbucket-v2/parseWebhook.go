package bitbucket_v2

import (
	"encoding/json"
	"time"
)

type PushHook struct {
	Push       PushType
	Actor      User
	Repository Repo
}

type PushType struct {
	Changes []*ChangesType
}

type ChangesType struct {
	Force     bool
	Old       TargetValue
	New       TargetValue
	Created   bool //to indicate whether the reference is new.
	Commits   []*Refs
	Truncated bool
	Closed    bool
}

type PullRequestHook struct {
	Actor       User
	PullRequest PR
	Repository  Repo
}

type PR struct {
	Id          int64
	Title       string
	Description string
	State       PullRequestState
	Author      User
	Source      struct {
		Branch     struct{ Name string }
		Commit     struct{ Hash string }
		Repository Repo
	}
	Destination struct {
		Branch     struct{ Name string }
		Commit     struct{ Hash string }
		Repository Repo
	}
	MergeCommit       struct{ Hash string } `json:"merge_commit"`
	Participants      []*User
	Reviewers         []*User
	CloseSourceBranch bool `json:"close_source_branch"`
	CloseBy           User `json:close_by`
	Reason            string
	Created           time.Time `json:"created_on"`
	Updated           time.Time `json:"updated_on"`
}

type PullRequestState string

const (
	PR_STATE_OPEN     PullRequestState = "OPEN"
	PR_STATE_MERGED   PullRequestState = "MERGED"
	PR_STATE_DECLIEND PullRequestState = "DECLIEND"
)

func ParseRepoHook(payload []byte) (hook *PushHook, err error) {
	err = json.Unmarshal(payload, hook)
	return
}

func ParsePRHook(payload []byte) (hook *PullRequestHook, err error) {
	err = json.Unmarshal(payload, hook)
	return
}
