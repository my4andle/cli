// Copyright © 2018 Zechen Jiang <zechen@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type resultObjectCmd struct {
	client  command.Interface
	teamID  string
	cloudID string
	out     io.Writer
	level   string
}

func newResultObjectCmd(client command.Interface, out io.Writer) *cobra.Command {
	resultObject := &resultObjectCmd{
		client: client,
		out:    out,
	}
	cmd := &cobra.Command{
		Use:     content.CmdResultObjectUse,
		Short:   content.CmdResultObjectShort,
		Long:    content.CmdResultObjectLong,
		Example: content.CmdResultObjectExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if resultObject.client == nil {
				resultObject.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			resultObject.teamID = teamID
			return resultObject.run()
		},
	}
	f := cmd.Flags()
	f.StringVar(&resultObject.cloudID, content.CmdFlagCloudIDLong, content.None, content.CmdFlagCloudIDDescription)
	f.StringVar(&resultObject.level, content.CmdFlagLevelLong, content.None, content.CmdFlagLevelDescription)
	return cmd
}

func (t *resultObjectCmd) run() error {
	res, err := t.client.ShowResultObject(t.teamID, t.cloudID, t.level)
	if err != nil {
		return err
	}
	fmt.Fprintln(t.out, util.PrettyJSON(*res))
	return nil
}
