/**
* @file dbEngin.go
* @brief 数据库引擎
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-28
 */
package database

type DbEngin struct {
	MysqlDb *MysqlDb
}

//NewDbEngin 创建数据库引擎
//需要注意并发访问
func NewDbEngin() *DbEngin {
	db := new(DbEngin)
	db.MysqlDb = NewMysqlDb()
	return db
}
