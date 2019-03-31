package main

import (
	"net/http"
	"os"

	"./controlers"
	"./database"
	//"./models"
	//	"fmt"
	//"os/exec"
	// "bytes"
	//"io/ioutil"
	// "net/url"
)

func main() {
	/*
		keyword := "哪里算命比较准"
		keyword1 := "算命风水好"
		keyword2 := "风水怎么看"
	*/
	dbs := database.NewDbEngin()
	dbs.MysqlDb.CreateMysqlDb("www_abc_com")
	spiderCli := controlers.NewSpiderCli(dbs)
	//spiderCli.Articles.CreateArticle()
	//controlers.RequestUrl2Io("地球村")
	//return
	if len(os.Args) > 1 {
		spiderCli.CrawArticle(os.Args[1])
		return
	}
	spiderSer := controlers.NewSpiderServer()
	http.HandleFunc("/craw", spiderSer.CrawServer)
	http.ListenAndServe(":8181", nil)
	defer spiderSer.CloseWd()

	//return
	//controlers.StartCrawMulit(os.Args[1])
	//controlers.StartCraw()
	//res := controlers.RequestUrl2Io("http%3A%2F%2Fmp.weixin.qq.com%2Fs%3Fsrc%3D11%26timestamp%3D1553272950%26ver%3D1500%26signature%3DjY*0ZfEtSmUdoqi*MzNSwZE6FMniAbtA6q79pbZ0W3*EWWxZVJ6ii**jNyOitH0mAzlNrMi-v2TiaSbQKVQJkWKiCPi47VMa*rGqMh6dyQ8M3elpEvCmtyTCx5fwa4lP%26new%3D1")
	//fmt.Println("res:", res)

	//res := controlers.RequestUrl2Io(os.Args[1])
	//fmt.Println("res:", res)
	//controlers.StartCraw("https://www.baidu.com")
	//controlers.ParseArticle()

	//controlers.StartCraw("https://pigav.com/%E6%AF%8F%E6%97%A5av%E7%B7%9A%E4%B8%8A%E7%9C%8B")
	//addr := "https://pigav.com/227095/%E9%93%83%E6%9D%91%E7%88%B1%E7%90%86-%E9%93%83%E6%9D%91%E7%88%B1%E7%90%86%E7%9A%84%E5%A5%B3%E4%BA%BA%E6%BF%80%E6%83%85%E6%80%A7%E4%BA%A44%E6%80%A7%E4%BA%A4.html"
	//controlers.StartCraw(addr)
	//controlers.StartCrawAll(addr)
	//addrXvideo := "https://www.xvideos.com/video32656615/7hang.com_-_-_"
	//addrXvideo := "https://www.xvideos.com/video37242175/leggy_lesbians_athina_and_ella_hughes_ride_gigantic_double_dong_on_xxxmas"
	//controlers.StartCraw(addrXvideo)
	//http.HandleFunc("/craw", controlers.CrawVideo)
	//http.ListenAndServe(":8181", nil)
	//controlers.TestIcbc()
	//controlers.TestIe1()
	//controlers.TestCmb()
}
