// Sample program to show how to write a simple version of curl using
// the io.Reader and io.Writer interface support.
package system_net

import (
	"crypto/tls"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"go_system/system_config"
	"io/ioutil"
	"time"
)

type NetOptions struct {
	Headers map[string]string
	Time    struct {
		connectMsTimeout   time.Duration
		readWriteMsTimeout time.Duration
	}
	Cookies            map[string]string
	InsecureSkipVerify bool
}

type EumMethod uint8

const (
	MethodPost EumMethod = iota
	MethodGet
)

func doRequest(url string, method EumMethod, params map[string]string, netOptions NetOptions) (resp string, err error) {
	var req *httplib.BeegoHTTPRequest
	if method == MethodGet {
		req = httplib.Get(url)
	} else if method == MethodPost {
		req = httplib.Post(url)
	} else {
		panic("unknow method:" + string(method))
	}
	if method == MethodPost && len(params) > 0 {
		for k, v := range params {
			req.Param(k, v)
		}
	}
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: netOptions.InsecureSkipVerify})
	timeOption := netOptions.Time
	configer := system_config.Configer()
	connectMsTimeout, _ := configer.Int64("connectMsTimeout")
	readWriteMsTimeout, _ := configer.Int64("readWriteMsTimeout")
	connectTimeout := time.Duration(connectMsTimeout)
	readWriteTimeout := time.Duration(readWriteMsTimeout)
	if timeOption.connectMsTimeout > 0 {
		connectTimeout = timeOption.connectMsTimeout
	}
	if timeOption.readWriteMsTimeout > 0 {
		readWriteTimeout = timeOption.readWriteMsTimeout
	}
	req.SetTimeout(connectTimeout*time.Millisecond, readWriteTimeout*time.Millisecond)
	for k, v := range netOptions.Headers {
		req.Header(k, v)
	}
	if len(netOptions.Cookies) > 0 {
		var cookieString = ""
		for k, v := range netOptions.Cookies {
			req.Header(k, v)
			cookieString += k + "=" + v + ";"
		}
		req.Header("Cookie", cookieString)
	}
	originResp, err := req.Response()
	if err != nil {
		logs.Error(err)
		return "", err
	}
	defer originResp.Body.Close()
	body, err := ioutil.ReadAll(originResp.Body)
	str := string(body)
	return str, err
}

func Get(url string, netOptions NetOptions) (resp string, err error) {
	return doRequest(url, MethodGet, nil, netOptions)
}

func Post(url string, params map[string]string, netOptions NetOptions) (resp string, err error) {
	return doRequest(url, MethodPost, params, netOptions)
}
