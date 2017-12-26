// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfiguration(t *testing.T) {
	templateFile, err := ioutil.TempFile("", "configuration")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(templateFile.Name())
	data := []byte(`# configuration file

currency = "EUR"

[portfolio]
BTC = "0.013"
DOGE = "150"
ETH = "0.145"
LTC = "0.123"

`)
	err = ioutil.WriteFile(templateFile.Name(), data, 0700)
	if err != nil {
		t.Fatal(err)
	}
	configuration, err := LoadFileConfig(templateFile.Name())
	assert.NoError(t, err)

	fmt.Printf("Configuration : %#v\n", configuration)
	assert.Equal(t, "EUR", configuration.Currency)
	assert.Equal(t, "0.013", configuration.Portfolio["bitcoin"])
	assert.Equal(t, "150", configuration.Portfolio["dogecoin"])
	assert.Equal(t, "0.145", configuration.Portfolio["ethereum"])
	assert.Equal(t, "0.123", configuration.Portfolio["litecoin"])
}
