/*
 * Copyright © 2022 Atomist, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/atomist-skills/go-skill"
)

func processCommitMessage(ctx context.Context, req skill.RequestContext, commit OnCommit) skill.Status {

	regex, _ := regexp.Compile(`(?s)^fix:.*mistybugz:(\d+).*`)
	match_result := regex.FindStringSubmatch(commit.Message)
	if len(match_result) > 0 {
		bug_id := match_result[1]

		msg := fmt.Sprintf("Found bug id %s in commit", bug_id)
		fmt.Println(msg)

		err := createBug(ctx, req, commit, bug_id, "mistybugz")

		if err != nil {
			msg := fmt.Sprintf("Failed to create bug with id %s from commit message", bug_id)
			fmt.Println(msg)
			return skill.Status{
				State:  skill.Failed,
				Reason: msg,
			}
		} else {
			return skill.Status{
				State:  skill.Completed,
				Reason: msg,
			}
		}

	} else {
		msg := "No bug id found in commit"
		fmt.Println(msg)
		return skill.Status{
			State:  skill.Completed,
			Reason: msg,
		}
	}

}

func createBug(ctx context.Context, req skill.RequestContext, commit OnCommit, bugId string, bugSystem string) error {
	transaction := skill.NewTransaction()

	transaction.AddEntities(LinkedBugEntity{
		Entity: transaction.MakeEntity("bug"),
		Commit: GitCommitEntity{
			Entity: transaction.MakeEntity("git/commit"),
			Sha:    commit.Sha,
			Repo: GitRepoEntity{
				Entity:   transaction.MakeEntity("git/repo"),
				SourceId: commit.Repo.SourceId,
				Url:      commit.Repo.Org.Url,
			},
			Url: commit.Repo.Org.Url,
		},

		ID:     bugId,
		System: bugSystem,
	})

	err := req.Transact(transaction.Entities())

	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Transacted bug with id %s for system %s\n", bugId, bugSystem)
	fmt.Println(msg)
	req.Log.Infof(msg)
	return err
}
