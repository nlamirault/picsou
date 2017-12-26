// Copyright (C) 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

package coinmarketcap

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"

	"github.com/nlamirault/picsou/version"
)

var (
	userAgent = fmt.Sprintf("picsou/%s", version.Version)
)

func fetchCoin(uri string) ([]byte, error) {
	glog.V(2).Infof("URI: %s", uri)

	client := &http.Client{}
	r, _ := http.NewRequest("GET", uri, nil)
	r.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(r)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("Http request to %s failed: %s", r.URL, err.Error())
	}
	glog.V(2).Infof("HTTP Status: %s", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, fmt.Errorf("Reading the body: %s", err.Error())
	}
	return body, nil
}
