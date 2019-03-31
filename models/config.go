package models

import (
	"../common"
	"../database"
	"../definition"
	"github.com/jinzhu/gorm"
)
func GetSub(md5 string) string {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return ""
	}
	/*
	var config definition.Configs
	db.Where(&definition.Configs{ID: 1}).First(&config)
	if config.Domain != "" {
		var video definition.Videos
		db.Where(&definition.Videos{Md5: md5}).First(&video)
		return config.Domain + "/"+video.DestPath
	}
	*/
	var config definition.Cfgs
	db.Where(&definition.Cfgs{ID: 1}).First(&config)
	if config.DomainNameSetting != "" {
		var video definition.Videos
		db.Where(&definition.Videos{Md5: md5}).First(&video)
		return config.DomainNameSetting + "/"+video.DestPath
	}

	return ""
}


func GetDomainForPlay(md5 string) string {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return ""
	}

	var config definition.Cfgs
	db.Where(&definition.Cfgs{ID: 1}).First(&config)
	if config.DomainNameSetting != "" {
		var video definition.Videos
		db.Where(&definition.Videos{Md5: md5}).First(&video)
		return config.DomainNameSetting + "/"+video.DestPath
	}

	return ""
}

func GetMaxJobs() string {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return ""
	}

	var config definition.TranscodingConfgs
	db.Where(&definition.TranscodingConfgs{ID: 1}).First(&config)
	return config.NumberOfTasks
}

func GetFragmentDuartion() string {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return ""
	}

	var config definition.TranscodingConfgs
	db.Where(&definition.TranscodingConfgs{ID: 1}).First(&config)
	return config.FragmentationDuration
}

//GetDomain  获取文件写入路径
func GetDomain() string {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return ""
	}
	var config definition.Configs
	db.Where(&definition.Configs{ID: 1}).First(&config)
	if config.Domain != "" {
		return config.Domain
	}
	return ""
}
