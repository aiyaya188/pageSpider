/**
* @file netUtils.go
* @brief 网络操作工具
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-28
 */
package controlers

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"../common"
)

// 网络访问
type NetUtils struct {
	Addr      string       //需要访问的地址
	httpClent *http.Client //http访问
}

//SetAddr 设置地址
func (netUtil *NetUtils) SetAddr(addr string) {
	netUtil.Addr = addr
}

func NewNetUtils() *NetUtils {
	netUtil := new(NetUtils)
	netUtil.httpClent = &http.Client{}
	return netUtil
}

//DoPost 发送post请求 addr:目标地址;Header:指定头信息，如果没有赋予nil;body:发送内容;jsonFormat:是否为json格式
func (netUtil *NetUtils) DoPost(addr string, Header map[string]string, body string, jsonFormat bool) (string, bool) {
	netUtil.Addr = addr
	//var err error
	req, err := http.NewRequest("post", addr, strings.NewReader(body))
	checkResult := common.CheckError(err, "netUtil doPost")
	if !checkResult {
		return "", false
	}
	if Header != nil {
		for k, v := range Header {
			req.Header.Set(k, v)
		}
	}
	if jsonFormat {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	resp, err := netUtil.httpClent.Do(req)
	checkResult = common.CheckError(err, "netUtil doPost1")
	if !checkResult {
		return "", false
	}
	respon, err := ioutil.ReadAll(resp.Body)
	if !checkResult {
		return "", false
	}

	//fmt.Printf("response:", string(respon))
	return string(respon), true
}

//DoSimpleGet 简单的get 访问
func (netUtil *NetUtils) DoSimpleGet(addr string) (string, bool) {
	if addr == "" {
		return "", false
	}
	resp, err := http.Get(addr)
	res := common.CheckError(err, "netUtil get err")
	if !res {
		return "", false
	}

	//Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		//fmt.Println("url status:", resp.StatusCode)
		common.Log1("Do simpleget :", resp.StatusCode)
		return "", false
	}
	responData, _ := ioutil.ReadAll(resp.Body)
	return string(responData), true

}
