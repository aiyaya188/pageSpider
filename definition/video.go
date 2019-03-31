package definition

import (
	"github.com/jinzhu/gorm"
)

type CloudUpOrder struct {
	Md5  string `json:"md5"`
	Path string `json:"path"`
}

type CloudNotify struct {
	Md5    string `json:"md5"`
	Status string `json:"status"`
}

type DelVideos struct {
	Path string `json:"path"`
}

type DelVideoStatus struct {
	Status string `json:"status"`
}

type CrawVideoData struct {
	FilePath string `json:"filePath"`
	CrawAddr string `json:"crawAddr"`
}

type CrawStatus struct {
	Status string `json:"status"`
	Title  string `json:"title"`
}

type Videos struct {
	gorm.Model
	//ID    		 int `gorm:"AUTO_INCREMENT"`
	Md5      string `gorm:"size:255"`
	FileName string `gorm:"size:255"`
	// DeFileName      string `gorm:"size:255"`
	FileType      string `gorm:"size:255"`
	TaskId        string `gorm:"size:255"`
	FileSize      int64  `gorm:"size:16"`
	Minutes       int    `gorm:"size:16"`
	Offset        int64  `gorm:"size:16"`
	TranslatCount int    `gorm:"size:16"`
	UploadCount   int    `gorm:"size:16"`
	Path          string `gorm:"size:255"`
	M3u8Path      string `gorm:"size:255"`
	DestPath      string `gorm:"size:255"`
	PlayTime      string `gorm:"size:255"`
	Status        int    `gorm:"size:16"`
}
type UploadBlock struct {
	TranType string `json:"tranType"`
	Offset   string `json:"offset"`
	FileSize string `json:"FileSize"`

	FileMd5  string `json:"FileMd5"`
	FileType string `json:"FileType"`
	FileName string `json:"FileName"`
}
type ResponBlock struct {
	PlayUrl   string `json:"playUrl"`
	Offset    int64  `json:"offset"`
	BlockSize int64  `json:"blockSize"`
	Retcode   int64  `json:"retcode"`
	Errmsg    string `json:"errmsg"`
}

// retcode: 0：继续上传，1：上传成功，2：上传失败
const (
	RetContinue    = 0
	RetSucceed     = 1
	RetFailed      = 2
	RetFatal       = 3
	RetTypeInvalid = 4
)

//0:传输中，1传输完成，等待转码
const (
	VideoUploading = 0
	VideoWaitTran  = 1
	VideTranCoding = 2
	VideoToCloud   = 3 //等待上传s3
	VideSuccess    = 4
	VideoTranErr   = 5
	VideoS3Err     = 6
)
const (
	SUCCESS = "success"
	FAIL    = "fail"
)
