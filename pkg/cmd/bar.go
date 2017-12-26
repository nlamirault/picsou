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
	"strings"

	"github.com/leekchan/accounting"
)

const (
	barChar = "∎"
)

func DisplayWalletBars(out io.Writer, name string, money float64, walletTotal float64, ac accounting.Accounting) {
	percent := fmt.Sprintf("%.0f", (money*100)/walletTotal)
	fmt.Fprintf(out, "%s: %s, [%s%%]\t", GreenOut(name), GetMoney(ac, fmt.Sprintf("%f", money)), percent)

	var barLen int
	barLen, err := strconv.Atoi(percent)
	if err != nil {
		return
	}

	var bars string
	switch {
	case 0 < barLen && barLen < 30:
		bars = BlueOut(barChar)
	case 30 < barLen && barLen < 50:
		bars = GreenOut(barChar)
	case 50 < barLen && barLen < 70:
		bars = YellowOut(barChar)
	case 70 < barLen:
		bars = RedOut(barChar)
	}

	fmt.Fprintf(out, "%v\n", strings.Repeat(bars, barLen))
}
