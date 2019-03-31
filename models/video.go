//Package models implement model for video
/*
* @file video.go
* @brief: 视频控制器
* @author frankie@gmail.com
* @version v1.0
* @date 2018-10-21
 */
package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"../cache"
	"../common"
	"../database"
	"../definition"
	"github.com/jinzhu/gorm"
)

/*
type Videos struct {
	gorm.Model
	//ID    		 int `gorm:"AUTO_INCREMENT"`
	Md5           string `gorm:"size:255"`
	FileName      string `gorm:"size:255"`
	FileType      string `gorm:"size:255"`
	FileSize      uint64 `gorm:"size:16"`
	Minutes       int     `gorm:"size:16"`
	Offset        uint64 `gorm:"size:16"`
	TranslatCount int   `gorm:"size:16"`
	UploadCount   int   `gorm:"size:16"`
	Status        int   `gorm:"size:16"`

}
//0:传输中，1传输完成，等待转码
const (
	VideoUploading = 0
	VideoWaitTran = 1
)
*/

/*CheckVideoType implement check video
/**
* @brief CheckVideoType 检查类型是否支持
*
* @param string 类型描述
*
* @returns  支持返回false,不支持返回true：
*/
/* ----------------------------------------------------------------------------*/
func CheckVideoType(videoType string) bool {
	if videoType == "" {
		return false
	}
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return false
	}
	var configs definition.TranscodingConfgs
	var find int
	find = 0
	db.Where(&definition.TranscodingConfgs{ID: 1}).First(&configs)
	format := configs.TranscodingFormat
	formats := strings.Split(format, "|")
	common.Log1("formats:", formats)
	for _, k := range formats {
		if k == videoType {
			find = 1
			break
		}
	}
	if find == 1 {
		common.Log1("video type check ok")
		return true
	} else {
		common.Log1("video type check fail")
		return false
	}
}

/*NewVideo --------------------------------------------------------------------------*/
/**
* @brief: NewVideo 新建视频任务,传输类型为cmd时候调用
*    <1>     如果视频不存在，插入视频数据,插入任务id，返回期望offset为0，状态码为0
     <2>     如果视频存在，并状态为传输中，返回期望offset,状态码为0
	 <3>	 如果视频存在，状态为其他，返回期望offser为filesize,状态码返回1
* @returns:
*/
/* ----------------------------------------------------------------------------*/
func NewVideo(videoSample definition.Videos) (int64, int) {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		//common.Log1("get db nill")
		common.Log("get db nill")
		return 0, 0
	}
	//var video definition.Videos
	common.Log1("new video")
	if videoSample != (definition.Videos{}) {

		cacheUpload := cache.GetCache(definition.KeyUpload + videoSample.Md5)
		common.Log1("cacheUpload:", cacheUpload)
		if cacheUpload != "" {
			temp := strings.Split(cacheUpload, "|")
			offset := temp[0]
			offsetStr, _ := strconv.ParseInt(offset, 10, 64)
			return offsetStr, 0
		} else {
			var video definition.Videos
			db.Where(&definition.Videos{Md5: videoSample.Md5}).First(&video)
			common.Log1("new video db:", video)
			if video.Md5 != "" {
				if video.Status > definition.VideoUploading {
					return video.FileSize, 1
				} else {
					PutSampleCache(video)
					return video.Offset, 0
				}

			}
			//创建任务id
			timeStr := strconv.FormatInt(time.Now().Unix(), 10)
			taskID := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(1000))
			videoSample.TaskId = taskID
			videoSample.Path = time.Now().Format("2006-01-02") + "/" + videoSample.Md5 + "/" + videoSample.Md5 + "." + videoSample.FileType
			videoSample.DestPath = time.Now().Format("2006-01-02") + "/" + videoSample.Md5 + timeStr + "/" + "index.m3u8"
			db.Create(&videoSample)
			//写入缓存
			PutSampleCache(videoSample)
			return 0, 0
		}
	}
	//common.Log1("videosamle is nill")
	common.Log("video sample is nill")
	return 0, 2
}

//PutSampleCache : 把分片信息写入cache,并且追加一个监听该key的进程，如果key失效，则把文件长度作为offset写入mysql
func PutSampleCache(videoSample definition.Videos) {
	cacheValue := strconv.FormatInt(videoSample.Offset, 10) + "|" + videoSample.Path
	cacheKey := definition.KeyUpload + videoSample.Md5
	cache.PutCache(cacheKey, cacheValue, 172800) //有效期2天
	cacheUpload := cache.GetCache(cacheKey)
	common.Log1("putsamplecache,cacheUpload,cachekey:", cacheUpload, cacheKey)
	return
}

/*GetEoffset --------------------------------------------------------------------------*/
/**
* @brief GetEoffset 查询taskid， 得到期望的偏移
*
* @param
* @param int
*
* @returns
 */
/* ----------------------------------------------------------------------------*/
func GetEoffset(videoSample definition.Videos) (int64, bool) {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return 0, false
	}
	var video definition.Videos
	db.Where(&definition.Videos{TaskId: videoSample.TaskId}).First(&video)
	if video.FileType == "" {
		return 0, false
	}
	return video.Offset, true

}

//GetTaskID
func GetTaskID(videoSample definition.Videos) (string, bool) {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return "", false
	}
	var video definition.Videos
	db.Where(&definition.Videos{Md5: videoSample.Md5}).First(&video)
	if video.TaskId == "" {
		return "", false
	}
	return video.TaskId, true

}

/*UpdateOffset --------------------------------------------------------------------------*/
/**
* @brief UpdateOffset 更新上传任务的offset
*
* @param string
* @param int64
* * @returns
 */
/* ----------------------------------------------------------------------------*/
func UpdateOffset(taskID string, newOffset int64) bool {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		return false
	}
	var video definition.Videos

	common.Log(fmt.Sprintf("update taskid: %s,new offset :%d", taskID, newOffset))
	db.Model(&video).Where("task_id=?", taskID).Update("Offset", newOffset)
	common.Log("update ok")
	return true
}

//UpdateOffsetCache  更新缓存中的偏移值,返回key
func UpadateOffsetCache(video definition.Videos, newOffset int64) string {
	if video == (definition.Videos{}) {
		return ""
	}
	key := definition.KeyUpload + video.Md5
	value := cache.GetCache(key)
	if value == "" {
		common.Log("update offset:cache key nil")
		return ""
	}
	temp := strings.Split(value, "|")
	newOffsetStr := strconv.FormatInt(newOffset, 10)
	/*newValue := temp[0] + "|" + newOffsetStr*/
	newValue := newOffsetStr + "|" + temp[1]
	cache.PutCache(key, newValue, 172800) //有效期2天
	/* offsetStr := temp[0]*/
	//offset, _ := strconv.ParseInt(offsetStr, 10, 64)
	//if offset != video.Offset {
	//common.Log(" offset is not valid")
	//return false
	/*}*/

	//path := temp[1]
	return key
}
func UpdateStatusByMd5(md5 string, status int) bool {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return false
	}
	var video definition.Videos

	common.Log(fmt.Sprintf("update md5: %v,status:%v", md5, status))
	db.Model(&video).Where("md5=?", md5).Update("status", status)
	common.Log("update status ok")
	return true

}

//Update status
func UpdateStatus(taskID string, status int) bool {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return false
	}
	var video definition.Videos

	common.Log(fmt.Sprintf("update taskid: %v,status:%v", taskID, status))
	db.Model(&video).Where("task_id=?", taskID).Update("status", status)
	common.Log("update status ok")
	return true

}

/*CheckOffset --------------------------------------------------------------------------*/
/**
* @brief CheckOffset   检查当前要写入的偏移是否与期望值一样
*
* @param definition.Videos
*
* @returns
 */
/* ----------------------------------------------------------------------------*/
func CheckOffset(video definition.Videos) bool {
	if video == (definition.Videos{}) {
		return false
	}
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return false
	}
	var videoSample definition.Videos
	db.Where(&definition.Videos{Md5: video.Md5}).First(&videoSample)
	/*db.Where(&definition.Videos{TaskId: video.TaskId}).First(&videoSample)*/
	if video.Offset == videoSample.Offset {
		return true
	}
	return false
}

//PreWrite 预写入,返回是否可写，可写返回写入路径
func PreWrite(video definition.Videos) (bool, string) {
	//taskRes := CheckTaskID(video)
	//if taskRes == false {
	//common.Log("task id is nill")
	//return false, ""
	//}
	if video == (definition.Videos{}) {
		return false, ""
	}
	key := definition.KeyUpload + video.Md5

	value := cache.GetCache(key)
	common.Log1("prewrite ,key,value", key, value)
	if value == "" {
		common.Log(video.Md5 + "please new video first	")
		return false, ""
	}
	temp := strings.Split(value, "|")
	offsetStr := temp[0]
	offset, _ := strconv.ParseInt(offsetStr, 10, 64)
	if offset != video.Offset {
		common.Log(" offset is not valid")
		return false, ""
	}
	path := temp[1]
	/*
		offsetRes := CheckOffset(video)
		if offsetRes == false {
			common.Log("write offset is not valid")
			return false, ""
		}
		path := GetVidePath(video)
		if path == "" {
			return false, ""
		}
	*/
	return true, path

}

//GetVideoPath  获取文件写入路径
func GetVidePath(video definition.Videos) string {
	if video == (definition.Videos{}) {
		return ""
	}
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return ""
	}
	var videoSample definition.Videos
	db.Where(&definition.Videos{Md5: video.Md5}).First(&videoSample)
	if videoSample.Path != "" {
		return videoSample.Path
	}
	return ""

}

//CheckTaskID 检查是否已经建立了任务ID
func CheckTaskID(video definition.Videos) bool {
	if video == (definition.Videos{}) {
		common.Log("video nil")
		return false
	}
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return false
	}
	var videoSample definition.Videos
	db.Where(&definition.Videos{Md5: video.Md5}).First(&videoSample)
	if videoSample.TaskId != "" {
		return true
	}
	return false
}

func GetWatermarkPlace() int {
	var db *gorm.DB
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return 8
	}
	var configs definition.TranscodingConfgs
	var place string
	//var err error
	db.Where(&definition.TranscodingConfgs{ID: 1}).First(&configs)
	place = configs.LeftUpperWatermark
	if place == `10:10|20:20` {
		return 0
	}
	place = configs.RightUpperWatermark
	if place == `10:10|20:20` {
		return 1
	}
	place = configs.LeftLowerWatermark
	if place == `10:10|20:20` {
		return 2
	}

	place = configs.RightLowerWatermark
	if place == `10:10|20:20` {
		return 3
	}
	return 4
}
func GetSubtitle() (string, string, string, string) {
	var db *gorm.DB
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return "", "", "", ""
	}
	var configs definition.Cfgs
	db.Where(&definition.Cfgs{ID: 1}).First(&configs)
	subTitle := configs.SubtitleSettings
	color := configs.Colour
	fontSize := configs.FontSize
	start := configs.StartingPosition
	return subTitle, color, fontSize, start
}

func GetWaitingVideos() (int, []definition.Videos) {
	var db *gorm.DB
	var videoSet []definition.Videos
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return 0, videoSet
	}
	db.Where(&definition.Videos{Status: definition.VideoWaitTran}).First(&videoSet)
	videoCount := len(videoSet)
	return videoCount, videoSet

}

func UpdatePlayTime(taskID string, newTime string) bool {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		return false
	}
	var video definition.Videos
	db.Model(&video).Where("task_id=?", taskID).Update("play_time", newTime)
	common.Log("update ok")
	return true
}
