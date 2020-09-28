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
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sitewhere/swctl/pkg/cli/output"
)

const outputFlag = "output"

// bindOutputFlag will add the output flag to the given command and bind the
// value to the given format pointer
func bindOutputFlag(cmd *cobra.Command, varRef *output.Format) {
	cmd.Flags().VarP(newOutputValue(output.Table, varRef), outputFlag, "o",
		fmt.Sprintf("prints the output in the specified format. Allowed values: %s", strings.Join(output.Formats(), ", ")))

	err := cmd.RegisterFlagCompletionFunc(outputFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var formatNames []string
		for _, format := range output.Formats() {
			if strings.HasPrefix(format, toComplete) {
				formatNames = append(formatNames, format)
			}
		}
		return formatNames, cobra.ShellCompDirectiveDefault
	})

	if err != nil {
		log.Fatal(err)
	}
}

type outputValue output.Format

func newOutputValue(defaultValue output.Format, p *output.Format) *outputValue {
	*p = defaultValue
	return (*outputValue)(p)
}

func (o *outputValue) String() string {
	// It is much cleaner looking (and technically less allocations) to just
	// convert to a string rather than type asserting to the underlying
	// output.Format
	return string(*o)
}

func (o *outputValue) Type() string {
	return "format"
}

func (o *outputValue) Set(s string) error {
	outfmt, err := output.ParseFormat(s)
	if err != nil {
		return err
	}
	*o = outputValue(outfmt)
	return nil
}
