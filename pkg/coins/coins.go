// Copyright (C) 2017-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

package coins

// Coin represents a crypto currency
type Coin struct {
	Name    string
	Percent float64
	Money   float64
}

// Wallet represents a wallet of crypto currencies
type Wallet struct {
	Coins map[string]Coin
	Total float64
}