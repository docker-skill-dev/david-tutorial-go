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

	"github.com/atomist-skills/go-skill"
	"github.com/atomist-skills/go-skill/util"
)

func handleGitPush(ctx context.Context, req skill.RequestContext) skill.Status {
	result := req.Event.Context.Subscription.Result[0]
	commit := util.Decode[OnCommit](result[0])

	fmt.Printf("\nNew commit to repo %s\n", commit.Repo.Name)
	fmt.Printf("revision: %s\n", commit.Sha)
	fmt.Printf("message:  %s\n", commit.Message)

	return skill.Status{
		State:  skill.Completed,
		Reason: "Handled Git push",
	}
}
