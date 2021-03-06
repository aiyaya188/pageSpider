/**
* @file spiderCli.go
* @brief 爬虫客户端
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-28
 */

package controlers

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"../common"
	"../database"
)

//url2io解析正文
type Url2Io struct {
	//ContentAddr string
	netUtil *NetUtils //网络请求操作包
}

//利用url2io 接口获取网页正文
func (url2io *Url2Io) GetContent(contentAddr string) string {
	url := `http://api.url2io.com/article?token=Xfp3ttoyRmK0KyJy2m3X7Q&url=` + contentAddr
	data, res := url2io.netUtil.DoSimpleGet(url)
	if !res {
		return ""
	}
	return data
}

type BlockEx struct {
}

func (block *BlockEx) GetContent(contentAddr string) string {
	cmd := "php cxExtractor/pageExtract.php " + contentAddr
	data, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		common.CheckError(err, "GetContent")
		return ""
	}
	//fmt.Println("data:", string(data))
	return string(data)
}

//定义cxExtractor 获取正文
type CxExtractor struct {
	netUtil *NetUtils
}

func (cxExtractor *CxExtractor) GetContent(contentAddr string) string {
	return contentAddr
}

//Abstrator 定义提取正文的接口
type Abstrator interface {
	GetContent(contentAddr string) string
}

//SpiderCli 爬虫客户端
type SpiderCli struct {
	DB         *database.DbEngin //数据库引擎
	netUtil    *NetUtils         //网络请求操作包
	RequestMsg *SpiderRequstMsg  //请求格式包
	ResponMsg  *SpiderResponMsg  //响应包
	//Articles   *ArticleObj
	*ArticleObj
	*Utils
	Abstrator
}

//NewSpiderCli 创建爬虫客户端
func NewSpiderCli(db *database.DbEngin) *SpiderCli {
	spiderCli := new(SpiderCli)
	spiderCli.DB = db
	spiderCli.netUtil = NewNetUtils()
	spiderCli.Utils = new(Utils)
	spiderCli.ResponMsg = new(SpiderResponMsg)
	spiderCli.RequestMsg = new(SpiderRequstMsg)
	//SpiderCli.Articles = new(Article)
	//spiderCli.Articles = NewArticle(spiderCli.DB)
	spiderCli.ArticleObj = NewArticle(spiderCli.DB)
	//选择嵌入的提取正文接口
	//spiderCli.Abstrator = &Url2Io{spiderCli.netUtil}
	spiderCli.Abstrator = &BlockEx{}
	return spiderCli
}

func (cli *SpiderCli) ProcessResponse() {
	var articleSections []string
	var addr []string
	var section string
	fmt.Println("process response")
	if len(cli.ResponMsg.BaiduAddr) > 0 {
		addr, _ = cli.Random(cli.ResponMsg.BaiduAddr, 1)
		section = cli.GetContent(addr[0])
		articleSections = append(articleSections, section)
	}
	if len(cli.ResponMsg.GoogleAddr) > 0 {
		addr, _ = cli.Random(cli.ResponMsg.GoogleAddr, 1)
		section = cli.GetContent(addr[0])
		articleSections = append(articleSections, section)
	}
	if len(cli.ResponMsg.WeiXinAddr) > 0 {
		addr, _ = cli.Random(cli.ResponMsg.WeiXinAddr, 1)
		section = cli.GetContent(addr[0])
		articleSections = append(articleSections, section)
	}
	if len(cli.ResponMsg.BiyingAddr) > 0 {
		addr, _ = cli.Random(cli.ResponMsg.BiyingAddr, 1)
		section = cli.GetContent(addr[0])
		articleSections = append(articleSections, section)
	}

	fmt.Println("mix article")
	cli.MixArticle(cli.ResponMsg.Keyword, articleSections)
	fmt.Println("mix article end")
}
func (cli *SpiderCli) MixArticle(keyWord string, sections []string) {
	var article string
	article = ""
	for _, v := range sections {
		article = article + v
	}
	cli.CreateArticle(keyWord, article)
}

//CrawArticle 根据关键词爬虫文章链接
func (cli *SpiderCli) CrawArticle(keyword string) {
	addr := common.GetConfig("system", "spiderServer").String()
	cli.RequestMsg.Keyword = keyword
	cli.RequestMsg.RequestType = "article"
	body, _ := json.Marshal(cli.RequestMsg)
	respon, res := cli.netUtil.DoPost(addr, nil, string(body), true)
	if !res {
		fmt.Println("spider err")
		return
	}
	err := json.Unmarshal([]byte(respon), cli.ResponMsg)
	res = common.CheckError(err, "CrawArticle")
	if res {
		cli.ProcessResponse()
	}
	//baiduAddr := cli.ResponMsg.BaiduAddr
}

func (cli *SpiderCli) TestAbstrator() {
	data := cli.GetContent("https://blog.csdn.net/u014703817/article/details/51120742")
	fmt.Println("data:", data)
}
