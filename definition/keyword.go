package definition

import "github.com/jinzhu/gorm"

type KeywordTable struct {
	gorm.Model
	Keyword     string `gorm:"size:255"` //long key
	Weight      uint
	RelateKeyID uint
}

type RelateKeyTable struct {
	gorm.Model
	RelateKey string `gorm:"size:255"` //long key
	LastKeyID uint
}

type KeywordQueTable struct {
	gorm.Model
	KeywordID uint
}

//长尾词-长尾词序号-栏目-文章标题-地址-域名-文章ID-被访问次数-是否收录
type KeyWordArticleTable struct {
	gorm.Model
	Keyword   string `gorm:"size:255"` //long key
	RelateID  uint
	Title     string `gorm:"size:255"`
	CrawAddr  string
	Host      string
	ArticleID uint
	Access    uint
	BaiduRec  bool
}

type Sites struct {
	gorm.Model
	Host string `gorm:"size:255"`
}
