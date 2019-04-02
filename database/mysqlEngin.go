/**
* @file mysqlEngin.go
* @brief mysql 引擎
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-28
 */
package database

import (
	"fmt"
	"os/exec"

	"../common"
	"github.com/jinzhu/gorm"
)

//MysqlDb mysql 引擎包
type MysqlDb struct {
	dbs map[string]*gorm.DB //gorm 引擎数组
	//connectStr string              // 数据库链接语句
}

func NewMysqlDb() *MysqlDb {
	mysqlDb := new(MysqlDb)
	mysqlDb.dbs = make(map[string]*gorm.DB)
	return mysqlDb
}

func (mysqlDb *MysqlDb) InitDbs() {

	//manageName := common.GetConfig("mysql", "managerName").String()

}

//CreateMysqlDb 根据数据库名字创建数据库链接
func (mysqlDb *MysqlDb) CreateMysqlDb(dbName string, newDb bool, migration []interface{}) bool {
	userName := common.GetConfig("mysql", "user").String()
	//fmt.Println("userName:", userName)
	userPasswd := common.GetConfig("mysql", "passwd").String()

	if newDb {
		//创建数据库
		dbCtreateStr := "mysql -u" + userName + " -p" + userPasswd + " -e " + `"create database IF NOT EXISTS ` + dbName + ` DEFAULT CHARSET utf8 COLLATE utf8_general_ci"`
		//fmt.Println("createDb:", dbCtreateStr)
		exec.Command("bash", "-c", dbCtreateStr).CombinedOutput()
	}
	//fmt.Println("userPasswd:", userPasswd)
	dbStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", userName, userPasswd, dbName)
	db, err := gorm.Open("mysql", dbStr)
	res := common.CheckError(err, "mysqlEngin:CreateMysqlDb")
	if !res {
		return false
	}
	common.Log1(dbName + ":恭喜，数据库连接成功")
	logMode := common.GetConfig("mysql", "dbLog").String()
	if logMode == "1" {
		db.LogMode(true)
	}
	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	/*
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			//return "yf_" + defaultTableName
			return defaultTableName
		}
	*/
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(200)
	//db.AutoMigrate()
	mysqlDb.dbs[dbName] = db
	return true
}

//GetDb 根据数据库名字获取数据库组件
func (mysqlDb *MysqlDb) GetDb(DbName string) *gorm.DB {
	return mysqlDb.dbs[DbName]
}
