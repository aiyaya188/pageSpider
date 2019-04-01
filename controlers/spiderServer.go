/**
* @file spiderServer.go
* @brief 处理爬虫服务端调度
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-25
 */

package controlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../common"
	"../definition"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

//SpiderMsg 客户端传来需要爬取的信息
type SpiderRequstMsg struct {
	Keyword     string `json:"keyword"`     //关键字
	NotifyAddr  string `json:"notifyAddr"`  //回调地址
	RequestType string `json:"requestType"` //请求类型：article，普通文章； wenda:问答类
	Sign        string `json:"sign"`        //签名，用于服务端数据校验
}

type SearchObj struct {
	SearchFunc definition.SearchFunc // 搜索函数
	SearchAddr string                //搜索地址
}

func (search *SearchObj) CrawAction(wd selenium.WebDriver, keyword string) ([]string, bool) {
	fmt.Println("addr:", search.SearchAddr)
	res, links := search.SearchFunc(wd, search.SearchAddr, keyword)
	fmt.Println("search end ")
	return links, res
}

//回复客户的数据包
type SpiderResponMsg struct {
	Keyword    string   `json:"keyword"`
	BaiduAddr  []string `json:"baiduAddr"`
	GoogleAddr []string `json:"googleAddr"`
	BiyingAddr []string `json:"biyingAddr"`
	WeiXinAddr []string `json:"weixinAddr"`
	Status     string   `json:"status"` //抓取状态
	Sign       string   `json:"sign"`   //签名，用于客户端数据校验
}

type SpiderServer struct {
	Searchers  map[string]*SearchObj //爬虫实体
	RequestMsg *SpiderRequstMsg
	ResponMsg  *SpiderResponMsg
	wd         selenium.WebDriver
}

//func NewSearchObj(jk)
//创建
func NewSpiderServer() *SpiderServer {

	spider := &SpiderServer{}
	spider.Searchers = make(map[string]*SearchObj)
	spider.RequestMsg = new(SpiderRequstMsg)
	spider.ResponMsg = new(SpiderResponMsg)
	spider.Searchers["baidu"] = new(SearchObj)
	spider.Searchers["weixin"] = new(SearchObj)
	spider.Searchers["biying"] = new(SearchObj)
	spider.Searchers["google"] = new(SearchObj)
	//spider.Searchers = new(map[string]*SearchObj)
	spider.Searchers["baidu"].SearchFunc = BaiduSearch
	spider.Searchers["baidu"].SearchAddr = common.GetConfig("system", "baiduAddr").String()

	spider.Searchers["google"].SearchFunc = GoSearch
	spider.Searchers["google"].SearchAddr = common.GetConfig("system", "googleAddr").String()

	spider.Searchers["biying"].SearchFunc = BiyingSearch
	spider.Searchers["biying"].SearchAddr = common.GetConfig("system", "biyingAddr").String()

	spider.Searchers["weixin"].SearchFunc = WeixinSearch
	spider.Searchers["weixin"].SearchAddr = common.GetConfig("system", "weixinAddr").String()

	wd, wdres := GetChrome()
	if wdres == false {
		return nil
	}
	spider.wd = wd
	return spider
}
func (spiderSer *SpiderServer) CrawPage(keyword string) bool {
	//var searchFunc definition.SearchFunc // 搜索函数
	//var searchAddr string                //搜索引擎地址
	baiduExit := common.GetConfig("system", "baidu").String()
	googleExit := common.GetConfig("system", "google").String()
	weixinExit := common.GetConfig("system", "weixin").String()
	biyingExit := common.GetConfig("system", "biying").String()
	var links []string
	var res bool
	var searchRes bool
	searchRes = false //有其中一个成功都把应答包设置为ok，否则设置为false
	if baiduExit == "1" {
		links, res = spiderSer.Searchers["baidu"].CrawAction(spiderSer.wd, keyword)
		if res {
			searchRes = true
			spiderSer.ResponMsg.BaiduAddr = links
		}
	}
	if googleExit == "1" {
		links, res = spiderSer.Searchers["google"].CrawAction(spiderSer.wd, keyword)
		if res {
			searchRes = true
			spiderSer.ResponMsg.GoogleAddr = links
		}
	}
	if weixinExit == "1" {
		links, res = spiderSer.Searchers["weixin"].CrawAction(spiderSer.wd, keyword)
		if res {

			searchRes = true
			spiderSer.ResponMsg.WeiXinAddr = links
		}
	}
	if biyingExit == "1" {
		links, res = spiderSer.Searchers["biying"].CrawAction(spiderSer.wd, keyword)
		if res {

			searchRes = true
			spiderSer.ResponMsg.BiyingAddr = links
		}
	}
	if searchRes {
		spiderSer.ResponMsg.Status = "ok"
	} else {
		spiderSer.ResponMsg.Status = "fail"
	}
	spiderSer.ResponMsg.Keyword = keyword
	return searchRes
	/*
		SearchFunc := spiderSer.Searchers["baidu"].SearchFunc // 搜索函数
		Target := spiderSer.Searchers["baidu"].SearchAddr     // 搜索函数
		res, links := SearchFunc(spiderSer.wd, Target, keyword)
		spiderSer.ResponMsg.Keyword = keyword
		if res {
			spiderSer.ResponMsg.BaiduAddr = links
			spiderSer.ResponMsg.Status = "ok"
			//fmt.Println("links:", links)
			//return true
		}
		spiderSer.ResponMsg.Status = "fail"
		return false
	*/
}

//实现蜘蛛搜索功能
/*
func (spiderSer *SpiderServer) CrawPage(keyword string) bool {
	SearchFunc := spiderSer.Searchers["baidu"].SearchFunc // 搜索函数
	Target := spiderSer.Searchers["baidu"].SearchAddr     // 搜索函数
	res, links := SearchFunc(spiderSer.wd, Target, keyword)
	spiderSer.ResponMsg.Keyword = keyword
	if res {
		spiderSer.ResponMsg.BaiduAddr = links
		spiderSer.ResponMsg.Status = "ok"
		//fmt.Println("links:", links)
		//return true
	}
	spiderSer.ResponMsg.Status = "fail"
	return false

}
*/
func GetChromeForDownLoad() (selenium.WebDriver, bool) {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//暂时不要
	//caps.SetLogLevel(log.Performance, log.All)
	var err error
	// 禁止加载图片，加快渲染速度
	//imagCaps := map[string]interface{}{
	//"profile.managed_default_content_settings.images": 2,
	//}
	var chromParam []string
	ifHeadless := common.GetConfig("system", "headless").String()
	if ifHeadless == "1" {
		chromParam = []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		}
	} else {
		chromParam = []string{
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		}
	}
	chromeCaps := chrome.Capabilities{
		//Prefs: imagCaps,
		Path: "",
		Args: chromParam,
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	_, err = selenium.NewChromeDriverService("chromedriver", 9515, opts...)
	resCheck := common.CheckError(err, "start chrome")
	if resCheck != true {
		return nil, false
	}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	resCheck = common.CheckError(err, "get webDriver")
	if resCheck != true {
		return nil, false
	}
	return webDriver, true
}

// 获取chrome
func GetChrome() (selenium.WebDriver, bool) {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//暂时不要
	//caps.SetLogLevel(log.Performance, log.All)
	var err error
	// 禁止加载图片，加快渲染速度
	//imagCaps := map[string]interface{}{
	//"profile.managed_default_content_settings.images": 2,
	//}
	var chromParam []string

	ifHeadless := common.GetConfig("system", "headless").String()
	if ifHeadless == "1" {
		chromParam = []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		}
	} else {
		chromParam = []string{
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		}
	}
	chromeCaps := chrome.Capabilities{
		//Prefs: imagCaps,
		Path: "",
		Args: chromParam,
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	_, err = selenium.NewChromeDriverService("chromedriver", 9515, opts...)
	resCheck := common.CheckError(err, "start chrome")
	if resCheck != true {
		return nil, false
	}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	resCheck = common.CheckError(err, "get webDriver")
	if resCheck != true {
		return nil, false
	}
	return webDriver, true
}

//CloseWd 关闭浏览器
func (spiderSer *SpiderServer) CloseWd() {
	spiderSer.wd.Close()
}

//CrawServer 处理客户端请求
func (spiderSer *SpiderServer) CrawServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//err = json.Unmarshal(b, &crawData)
	err = json.Unmarshal(b, spiderSer.RequestMsg)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	spiderSer.CrawPage(spiderSer.RequestMsg.Keyword)
	/*
		if res == false {
			spiderSer.ResponMsg.Status = "fail"
		}
	*/
	w.Header().Set("access-control-allow-origin", "*")  //允许访问所有域
	w.Header().Add("access-control-allow-headers", "*") //header的类型
	w.Header().Add("access-control-expose-headers", "*")
	w.Header().Set("content-type", "application/json")
	output, _ := json.Marshal(spiderSer.ResponMsg)
	w.Write(output)
	return
}
