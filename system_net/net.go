// Sample program to show how to write a simple version of curl using
// the io.Reader and io.Writer interface support.
package system_net

import (
	"net/http"
	"time"
	"github.com/fogetu/go_system/system_config"
)

// main is the entry point for the application.
func Get(url string) (resp *http.Response, err error) {
	system_config.Configer().String("appname")
	client := http.Client{
		Timeout: time.Duration(2000 * time.Millisecond),
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
