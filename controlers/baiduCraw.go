package controlers

import (
	"fmt"
	"strings"
	"time"

	"../common"
	"github.com/tebeka/selenium"
)

//获取搜索结果的链接
func getResultLink(data string) string {
	preFind := `href="`
	fixFind := `" target="_blank"`
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

func BaiduSearch(webDriver selenium.WebDriver, target string, keyword string) (bool, []string) {
	webDriver.MaximizeWindow("")
	err1 := webDriver.Get(target)
	webDriver.SetImplicitWaitTimeout(30 * time.Second)
	title, _ := webDriver.Title()
	fmt.Println("title:", title)
	var resCheck bool
	resCheck = common.CheckError(err1, "open target ")
	if resCheck != true {
		return false, nil
	}
	time.Sleep(2 * time.Second)
	//robotgo.MoveMouseSmooth(323, 188)
	input, _ := webDriver.FindElement(selenium.ByXPATH, `//*[@id="kw"]`)
	//input.SendKeys("泰国哪里有好吃对东西")
	input.SendKeys(keyword)
	time.Sleep(1 * time.Second)
	souButton, _ := webDriver.FindElement(selenium.ByXPATH, `//*[@id="su"]`)
	souButton.Click()
	time.Sleep(5 * time.Second)
	var searchLink []string
	wd, _ := webDriver.FindElements(selenium.ByClassName, `c-container`)
	if len(wd) > 0 {
		for _, searchRes := range wd {
			element1, _ := searchRes.FindElement(selenium.ByClassName, `t`)
			element2, _ := element1.FindElement(selenium.ByTagName, `a`)
			eleText, _ := element2.GetAttribute("outerHTML")
			link := getResultLink(eleText)
			searchLink = append(searchLink, link)
		}
	}
	//resultLink := getResultLink(eleText)
	return true, searchLink
}
