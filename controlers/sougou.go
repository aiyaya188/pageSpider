package controlers

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func SogouSearch(webDriver selenium.WebDriver, target string) bool {
	webDriver.MaximizeWindow("")
	err1 := webDriver.Get(target)
	webDriver.SetImplicitWaitTimeout(30 * time.Second)
	if err1 != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err1))
		return false
	}
	time.Sleep(2 * time.Second)
	input, _ := webDriver.FindElement(selenium.ByXPATH, `//*[@id="query"]`)
	input.SendKeys(os.Args[2])
	time.Sleep(1 * time.Second)
	souButton, _ := webDriver.FindElement(selenium.ByXPATH, `//*[@id="stb"]`)
	souButton.Click()
	webDriver.SetImplicitWaitTimeout(30 * time.Second)
	//time.Sleep(5 * time.Second)
	wd, _ := webDriver.FindElements(selenium.ByClassName, `vrTitle`)
	fmt.Println("wd len:", len(wd))
	//element1, _ := wd[0].FindElement(selenium.ByClassName, `t`)
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
}
