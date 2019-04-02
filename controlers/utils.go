/**
* @file utils.go
* @brief 工具类函数
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-31
 */

package controlers

import (
	"errors"
	"math/rand"
	"strings"
)

type Utils struct {
}

//对slice进行洗牌，随机获取

func (utils *Utils) Random(strings []string, length int) ([]string, error) {
	var res []string //需要返回对结果
	if len(strings) <= 0 {
		return nil, errors.New("the length of the parameter strings should not be less than 0")
	}

	if length <= 0 || len(strings) <= length {
		return nil, errors.New("the size of the parameter length illegal")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	for i := 0; i < length; i++ {
		res = append(res, strings[i])
	}
	return res, nil
}

//GetDbByHost 根据域名获取数据库名
func (utils *Utils) GetDbByHost(hostName string) string {
	dbname := strings.Replace(hostName, `.`, `_`, -1)
	return dbname
}
