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
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/nlamirault/picsou/providers/coinmarketcap"
)

var (
	name  string
	names []string
)

type cryptoCmd struct {
	out io.Writer
}

func newCryptoCmd(out io.Writer) *cobra.Command {
	cryptoCmd := &cryptoCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "crypto",
		Short: "Manage crypto currencies.",
		Long:  "Manage crypto currencies. See List, Get, ... subcommands.",
		RunE:  nil,
	}

	listCryptoCmd := &cobra.Command{
		Use:   "list",
		Short: "Display crypto currencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := coinmarketcap.NewClient()
			if err != nil {
				return err
			}
			return cryptoCmd.listCryptoCurrencies(client)
		},
	}

	getCryptoCmd := &cobra.Command{
		Use:   "get",
		Short: "Display specific crypto currency",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(name) == 0 {
				return fmt.Errorf("missing crypto currency name")
			}
			client, err := coinmarketcap.NewClient()
			if err != nil {
				return err
			}
			return cryptoCmd.getCryptoCurrency(client, name)
		},
	}

	walletCryptoCmd := &cobra.Command{
		Use:   "wallet",
		Short: "Display specific crypto currencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	getCryptoCmd.PersistentFlags().StringVar(&name, "name", "", "Crypto currency name")
	walletCryptoCmd.PersistentFlags().StringSlice("cryptos", nil, "Cryptos' names")
	cmd.AddCommand(walletCryptoCmd)
	cmd.AddCommand(getCryptoCmd)
	cmd.AddCommand(listCryptoCmd)
	return cmd
}

func (cmd cryptoCmd) listCryptoCurrencies(client *coinmarketcap.Client) error {
	glog.V(1).Info("List crypto currencies")
	client.GetCoins("EUR", 100)
	return nil
}

func (cmd cryptoCmd) getCryptoCurrency(client *coinmarketcap.Client, name string) error {
	glog.V(1).Infof("Get crypto currency: %s", name)
	client.GetCoin(name, "EUR", 100)
	return nil
}
