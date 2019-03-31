package database

import (
	"fmt"
	"strconv"

	"sync"

	"github.com/jinzhu/gorm"

	//_ "github.com/go-sql-driver/mysql"
	"../common"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbKey int
	//DBM 公共数据件
	DBM *gorm.DB
	//DBS 公共数据件
	DBS        []*gorm.DB
	checkCount int
	//  OffsetMap  map[string]string
	OffsetMap sync.Map
)

func init() {
	/*
		fmt.Println("init db start")
		DBM = GetDB("master")
		if DBM == nil{
			return
		}
		//OffsetMap = make(map[string]string)
		DBM.AutoMigrate(&definition.Videos{})
		DBM.AutoMigrate(&definition.Configs{})
		//GetDB("slave")
		fmt.Println("init db end")
	*/
}

// GetDB : 获取数据库操作
func GetDB(dbType string) *gorm.DB {
	var db *gorm.DB
	if dbType == "master" {
		var err error
		//fmt.Println("get db master")
		if DBM == nil {
			dsn := common.GetConfig("mysql", "masterDsn").String()
			fmt.Println(fmt.Sprintf("dsn is %s", dsn))
			db, err = DbConn(dsn)
			if err == nil {
				return db
			}
			fmt.Println("db connetct fail")
			return nil
		}
		//fmt.Println("get old db")
		return DBM
	}

	if dbType == "slave" {
		if len(DBS) < 1 {
			fmt.Println("slave new conn")
			slaveCount, _ := common.GetConfig("mysql", "slaveCount").Int()
			DBS = make([]*gorm.DB, slaveCount)
			for i := 0; i < slaveCount; i++ {
				dsn := "slaveDsn" + strconv.Itoa(i+1)
				dsn = common.GetConfig("mysql", dsn).String()
				fmt.Println("slave dns is : " + dsn)
				dbs, err := DbConn(dsn)
				if err == nil {
					DBS[i] = dbs
				} else {
					fmt.Println(err)
				}
			}
		}
		if len(DBS) > 0 {
			if dbKey > len(DBS)-1 || dbKey < 1 {
				dbKey = 0
			} else {
				dbKey++
			}
			return DBS[dbKey]
		}

	}
	return nil
}

//DbConn : 数据库连接
func DbConn(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		//fmt.Println(fmt.Sprint("database connect fail ,db:%s,err:%d",db,err))
		return db, err
	}
	fmt.Println("恭喜，数据库连接成功")
	//db.LogMode(true)
	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		//return "yf_" + defaultTableName
		return defaultTableName

	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(200)
	return db, err
}

/*
func CheckDb() {
	fmt.Println("checkdb 执行...")

	for i := 0; i < 1; i++ {
		if DBM != nil {

			// Raw SQL
			rows, dbmErr := DBM.Raw("select 1 from mysql.db limit 1").Rows()

			if dbmErr != nil {
				fmt.Println("master fail ! 报警处理=================================")
				fmt.Println(dbmErr)

				//panic("db fail")

				//尝试重连接
				//GetDB("master")

			} else {
				defer rows.Close()
				fmt.Println(strconv.Itoa(checkCount) + "--主数据库查询正常\n")
			}
		} else {
			fmt.Println(strconv.Itoa(checkCount) + "主数据库没连接")
		}

		checkCount++
	}

	for i := 0; i < len(DBS); i++ {
		if DBS[i] != nil {

			// Raw SQL
			rows, dbsErr := DBS[i].Raw("select 1 from mysql.db limit 1").Rows()
			if dbsErr != nil {
				fmt.Println("slave fail ! 报警处理")
				fmt.Println(dbsErr)

			} else {
				defer rows.Close()
				fmt.Printf("从数据库%d查询正常\n", i)
			}
		} else {
			fmt.Println(strconv.Itoa(len(DBS)) + "从数据库没连接")
		}
	}

}
*/
