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
	"strconv"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/nlamirault/picsou/config"
	pkgcmd "github.com/nlamirault/picsou/pkg/cmd"
	"github.com/nlamirault/picsou/providers/coinmarketcap"
)

var (
	configFilename string
)

type portfolioCmd struct {
	out io.Writer
}

func newPortfolioCmd(out io.Writer) *cobra.Command {
	portfolioCmd := &portfolioCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "portfolio",
		Short: "Manage portfolio.",
		Long:  "Manage portfolio.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(configFilename) == 0 {
				return fmt.Errorf("missing configuration filename")
			}

			conf, err := config.LoadFileConfig(configFilename)
			if err != nil {
				return err
			}
			client, err := coinmarketcap.NewClient()
			if err != nil {
				return err
			}
			return portfolioCmd.getPortfolio(client, conf)
		},
	}
	cmd.PersistentFlags().StringVar(&configFilename, "config", "picsou.toml", "Configuration filename")
	return cmd
}

func (cmd portfolioCmd) getPortfolio(client *coinmarketcap.Client, conf *config.Configuration) error {
	glog.V(1).Infof("Get crypto currencies portfolio: %s", conf)
	coins := []coinmarketcap.Coin{}
	wallet := map[string]float64{}

	ac := pkgcmd.GetAccounting(conf.Currency)

	for name, owned := range conf.Portfolio {
		coin, err := client.GetCoin(name, conf.Currency, 1)
		if err != nil {
			return err
		}
		coins = append(coins, coin...)
		nb, err := strconv.ParseFloat(owned, 64)
		if err != nil {
			return err
		}
		price, err := strconv.ParseFloat(coinmarketcap.GetPrice(coin[0], conf.Currency), 64)
		if err != nil {
			return err
		}
		wallet[name] = nb * price
	}
	glog.V(2).Infof("Coins: %s", coins)
	glog.V(2).Infof("Wallet: %s", wallet)
	for name, money := range wallet {
		fmt.Fprintf(cmd.out, "%s: %s\n", pkgcmd.GreenOut(name), pkgcmd.GetMoney(ac, fmt.Sprintf("%f", money)))
	}
	if err := pkgcmd.DisplayCoins(cmd.out, coins, conf.Currency); err != nil {
		return err
	}

	return nil
}
