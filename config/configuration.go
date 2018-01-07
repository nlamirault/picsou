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

package config

import (
	"github.com/BurntSushi/toml"
)

// Configuration holds configuration for application
type Configuration struct {
	Currency  string
	Portfolio map[string]string
}

// LoadFileConfig returns a Configuration from reading the specified file (a toml file).
func LoadFileConfig(file string) (*Configuration, error) {
	var configuration Configuration
	if _, err := toml.DecodeFile(file, &configuration); err != nil {
		return nil, err
	}
	return &configuration, nil
}
