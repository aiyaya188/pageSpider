package controlers

import (
	"fmt"
	"time"

	"../common"
	"github.com/tebeka/selenium"
)

func BiyingSearch(webDriver selenium.WebDriver, target string, keyword string) (bool, []string) {
	webDriver.MaximizeWindow("")
	err1 := webDriver.Get(target)
	webDriver.SetImplicitWaitTimeout(30 * time.Second)
	var resCheck bool
	resCheck = common.CheckError(err1, "open target ")
	if resCheck != true {
		return false, nil
	}
	time.Sleep(2 * time.Second)
	input, _ := webDriver.FindElement(selenium.ByXPATH, `//*[@id="sb_form_q"]`)
	input.SendKeys(keyword)
	time.Sleep(1 * time.Second)
	souButton, _ := webDriver.FindElement(selenium.ByXPATH, `//*[@id="sb_form_go"]`)
	souButton.Click()
	webDriver.SetImplicitWaitTimeout(30 * time.Second)
	//time.Sleep(5 * time.Second)
	wd, _ := webDriver.FindElements(selenium.ByClassName, `b_algo`)
	fmt.Println("wd len:", len(wd))
	var searchLink []string
	if len(wd) > 0 {
		for _, searchRes := range wd {
			element2, _ := searchRes.FindElement(selenium.ByTagName, `a`)
			resultLink, _ := element2.GetAttribute("href")
			searchLink = append(searchLink, resultLink)
		}
	}
	//resultLink := getResultLink(eleText)
	return true, searchLink

	/*
		element2, _ := wd[0].FindElement(selenium.ByTagName, `a`)
		resultLink, _ := element2.GetAttribute("href")
		fmt.Println("result link:", resultLink)
		contentJSON := RequestUrl2Io(resultLink)
		if contentJSON == "" {
			return false
		}
		title, content := ParseUrl2Io(contentJSON)
		fmt.Println("title,content:", title, content)
		time.Sleep(5 * time.Second)
		time.Sleep(10 * time.Second)
		return true
	*/
}
