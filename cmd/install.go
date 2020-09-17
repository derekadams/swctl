/**
 * Copyright © 2014-2020 The SiteWhere Authors
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

package cmd

import (
	"io"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"

	"github.com/sitewhere/swctl/cmd/require"
	"github.com/sitewhere/swctl/pkg/action"
	"github.com/sitewhere/swctl/pkg/cli/output"
	"github.com/sitewhere/swctl/pkg/install"
)

var installHelp = `
Use this command to install SiteWhere 3.0 on a Kubernetes Cluster.
This command will install:
 - SiteWhere System Namespace: sitewhere-system (default)
 - SiteWhere Custom Resources Definition.
 - SiteWhere Templates.
 - SiteWhere Operator.
 - SiteWhere Infrastructure.
`

func newInstallCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	client := action.NewInstall(cfg)
	var outfmt output.Format

	cmd := &cobra.Command{
		Use:               "install",
		Short:             "Install SiteWhere CRD and Operators",
		Long:              installHelp,
		Args:              require.NoArgs,
		ValidArgsFunction: noCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			results, err := client.Run()
			if err != nil {
				return err
			}
			return outfmt.Write(out, newInstallWriter(results))
		},
	}

	f := cmd.Flags()

	f.BoolVarP(&client.Minimal, "minimal", "m", false, "Minimal installation.")
	f.BoolVarP(&client.WaitReady, "wait", "w", false, "Wait for components to be ready before return control.")
	bindOutputFlag(cmd, &outfmt)

	return cmd
}

type installWriter struct {
}

func newInstallWriter(install *install.SiteWhereInstall) *installWriter {
	return &installWriter{}
}

func (i *installWriter) WriteTable(out io.Writer) error {
	table := uitable.New()
	table.AddRow("NAME", "NAMESPACE", "REVISION", "UPDATED", "STATUS", "CHART", "APP VERSION")
	// for _, r := range r.releases {
	// 	table.AddRow(r.Name, r.Namespace, r.Revision, r.Updated, r.Status, r.Chart, r.AppVersion)
	// }
	return output.EncodeTable(out, table)
}

func (i *installWriter) WriteJSON(out io.Writer) error {
	return output.EncodeJSON(out, i)
}

func (i *installWriter) WriteYAML(out io.Writer) error {
	return output.EncodeYAML(out, i)
}
