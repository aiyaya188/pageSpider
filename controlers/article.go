/**
* @file article.go
* @brief 文章操作
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-28
 */
package controlers

import (
	"../database"
	"github.com/jinzhu/gorm"
)

//ArticleTable 文章数据表
type Article struct {
	gorm.Model
	Link           string `gorm:"size:255"`        //标题
	Title          string `gorm:"size:255"`        //标题
	ArticleContent string `gorm:"type:mediumtext"` //正文
	Summary        string `gorm:"size:600"`        //摘要
	Status         int    //文章状态
	ReadTimes      int    //阅读次数
	ForumId        uint   //版块ID
}

type ArticleObj struct {
	db           *database.DbEngin //创建函数需要注入
	articleTable *Article
}

//NewAdsrticle 创建文章
func NewArticle(db *database.DbEngin) *ArticleObj {

	articleObj := new(ArticleObj)
	articleObj.db = db
	articleObj.articleTable = new(Article)
	return articleObj
}

//CreateArticle  把文章写入数据库
func (article *ArticleObj) CreateArticle(title string, content string) {
	article.articleTable.Title = title
	article.articleTable.ArticleContent = content
	//article.articleTable.Summary = string(content[100:201])
	article.articleTable.Summary = title
	article.articleTable.Status = 1
	article.articleTable.ReadTimes = 123
	article.articleTable.ForumId = 1
	hostName := "www_abc_com"
	db := article.db.MysqlDb.GetDb(hostName)
	if db.NewRecord(&article.articleTable) {
		db.Create(&article.articleTable)
	}
}
