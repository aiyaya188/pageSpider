package controlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"../common"
	"github.com/tebeka/selenium"
)

//获取搜索结果的链接
func getWxResultLink(data string) string {
	preFind := `href="`
	fixFind := `" id=`
	pos1 := strings.Index(data, preFind)
	var resLink string
	resLink = ""
	if pos1 != -1 {
		pos2 := strings.Index(data, fixFind)
		if pos2 != -1 {
			resLink = data[pos1+len(preFind) : pos2]
		}
	}
	return resLink
}
func PutAritcleToLine(content string) string {
	client := &http.Client{}
	//body := `keyword=` + keyword
	body := content
	bodyLen := strconv.Itoa(len(body))
	requesUrl := "http://sites.twstgj.com/inputArticl.php"
	req, err := http.NewRequest("POST", requesUrl, strings.NewReader(body))
	if err != nil {
		fmt.Printf("dia errr")
		common.CheckErr(err)
		return ""
	}
	//req.Header.Set("Host", "perbank.abchina.com")
	req.Header.Set("Content-Length", bodyLen)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	fmt.Println("x1")
	resp, err := client.Do(req)
	if err != nil {
		common.CheckErr(err)
		//CookieNotify(definition.FAIL)
		return ""
	}

	fmt.Println("x2")
	respon, err := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(respon))
	if err != nil {
		common.CheckErr(err)
		return ""
	}

	fmt.Println("x3")
	responseData := string(respon)
	return responseData

}

func WeixinSearch(webDriver selenium.WebDriver, target string, keyword string) (bool, []string) {
	webDriver.MaximizeWindow("")
	err1 := webDriver.Get(target)
	webDriver.SetImplicitWaitTimeout(30 * time.Second)
	var resCheck bool
	resCheck = common.CheckError(err1, "open target ")
	if resCheck != true {
		return false, nil
	}
	time.Sleep(2 * time.Second)
	input, _ := webDriver.FindElement(selenium.ByXPATH, `//*[@id="query"]`)
	input.SendKeys(keyword)
	time.Sleep(1 * time.Second)
	souButton, _ := webDriver.FindElement(selenium.ByClassName, `swz`)
	souButton.Click()
	time.Sleep(10 * time.Second)
	wd, _ := webDriver.FindElements(selenium.ByClassName, `txt-box`)
	fmt.Println("wd len:", len(wd))
	var searchLink []string
	if len(wd) > 0 {
		for _, searchRes := range wd {
			element1, _ := searchRes.FindElement(selenium.ByTagName, `h3`)
			element2, _ := element1.FindElement(selenium.ByTagName, `a`)
			eleText, _ := element2.GetAttribute("outerHTML")
			resultLink := getWxResultLink(eleText)
			//fmt.Println("resultLink:", resultLink)
			res := strings.Replace(resultLink, "amp;", "", -1)

			searchLink = append(searchLink, res)
		}
	}
	//resultLink := getResultLink(eleText)
	return true, searchLink

	/*
		element1, _ := wd[0].FindElement(selenium.ByTagName, `h3`)
		fmt.Println("t1a")
		element2, _ := element1.FindElement(selenium.ByTagName, `a`)
		fmt.Println("t2")
		eleText, _ := element2.GetAttribute("outerHTML")
		fmt.Println("eleTxt:", eleText)
		resultLink := getWxResultLink(eleText)
		fmt.Println("resultLink:", resultLink)
		res := strings.Replace(resultLink, "amp;", "", -1)
		fmt.Println("res:", url.QueryEscape(res))
		//htmlData, _ := webDriver.PageSource()
		//result := PutAritcleToLine(htmlData)

		contentJSON := RequestUrl2Io(url.QueryEscape(res))
		if contentJSON == "" {
			return false
		}

		title, content := ParseUrl2Io(contentJSON)
		fmt.Println("title,content:", title, content)
		time.Sleep(10 * time.Second)
		return true
	*/
}
