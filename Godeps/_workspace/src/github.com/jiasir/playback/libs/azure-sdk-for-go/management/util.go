package management

import (
	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/jiasir/playback/libs/azure-sdk-for-go/core/http"
	"io/ioutil"
)

func getResponseBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
