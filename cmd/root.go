// Copyright (C) 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	goflag "flag"
	"fmt"
	"io"
	"os"

	_ "github.com/golang/glog" // init glog to get its flags
	"github.com/spf13/cobra"

	pkgcmd "github.com/nlamirault/picsou/pkg/cmd"
)

var (
	cliName     = "picsou"
	helpMessage = "Picsou - CLI to monitor cryptocurrencies"
)

func newPoseidonCommand(out io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   cliName,
		Short: "command-line tool to monitor cryptocurrencies",
		Long:  `The command-line tool to manage cryptocurrencies.`,
	}
	rootCmd.AddCommand(
		newVersionCmd(out, helpMessage),
		// newCompletionCommand(out, completionExample),
	)
	cobra.EnablePrefixMatching = true

	// add glog flags
	rootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	// https://github.com/kubernetes/dns/pull/27/files
	goflag.CommandLine.Parse([]string{})

	return rootCmd
}

func Execute() {
	cmd := newPoseidonCommand(os.Stdout)
	if err := cmd.Execute(); err != nil {
		fmt.Println(pkgcmd.RedOut(err))
		os.Exit(1)
	}
}
