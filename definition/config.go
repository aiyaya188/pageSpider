package definition

import "github.com/tebeka/selenium"

/*
import (
	"github.com/jinzhu/gorm"
)
*/

const SignKey = "qjnnj52xt1cgr8pigd45to3p07bx25up" //用于签名
type ProcessBasename func(string) string

type SearchFuncName func(selenium.WebDriver, string) bool
type SearchFunc func(selenium.WebDriver, string, string) (bool, []string)
type Configs struct {
	ID     int    `gorm:"AUTO_INCREMENT"`
	Domain string `gorm:"size:255"`
}

type Cfgs struct {
	//gorm.Model
	ID                int    `gorm:"AUTO_INCREMENT"`
	DomainNameSetting string `gorm:"size:255"`
	SubtitleSettings  string `gorm:"size:255"`
	Colour            string `gorm:"size:255"`
	SubtitleStatus    string `gorm:"size:255"`
	FontSize          string `gorm:"size:255"`
	StartingPosition  string `gorm:"size:255"`
}

type TranscodingConfgs struct {
	//gorm.Model
	ID                    int    `gorm:"AUTO_INCREMENT"`
	NumberOfTasks         string `gorm:"size:255"`
	TranscodingFormat     string `gorm:"size:255"`
	LeftUpperWatermark    string `gorm:"size:255"`
	RightUpperWatermark   string `gorm:"size:255"`
	LeftLowerWatermark    string `gorm:"size:255"`
	RightLowerWatermark   string `gorm:"size:255"`
	WatermarkOrNot        string `gorm:"size:255"`
	FragmentationDuration string `gorm:"size:255"`
}
