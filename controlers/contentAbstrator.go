package controlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"../common"
	simplejson "github.com/bitly/go-simplejson"
)

func RequestUrl2Io(query string) string {
	url := `http://api.url2io.com/article?token=Xfp3ttoyRmK0KyJy2m3X7Q&url=` + query
	resp, err := http.Get(url)
	common.CheckErr(err)

	//Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		fmt.Println("url status:", resp.StatusCode)
		return ""
	}
	responData, _ := ioutil.ReadAll(resp.Body)
	return string(responData)
}

func ParseUrl2Io(data string) (string, string) {
	fmt.Println("url2io data", data)
	js, err := simplejson.NewJson([]byte(data))
	common.CheckErr(err)
	var title, content string
	title = js.Get("title").MustString()
	content = js.Get("content").MustString()
	return title, content

}
