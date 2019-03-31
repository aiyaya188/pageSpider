/**
* @file helper.go
* @brief: 辅助函数
* @author frankie@gmail.com
* @version v1.0
* @date 2018-10-22
 */

package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-ini/ini"
	//"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/sessions"
	//"github.com/utrack/gin-csrf"
)

var configPath string

/* --------------------------------------------------------------------------*/
/**
* @brief: init 初始化
* @returns:
 */
/* ----------------------------------------------------------------------------*/
func init() {
	//配置文件设置
	//SetConfigPath(os.Getenv("USERVER_INI") + "/conf/config.ini")
	/*SetConfigPath("../conf/config.ini")*/
	//SetConfigPath("conf/config.ini")
	SetConfigPath("conf/config.ini")
}

/*SetConfigPath  --------------------------------------------------------------------------*/
/**
* @brief: SetConfigPath 设置配置文件路径

*
* @param: string
*
* @returns:
 */
/* ----------------------------------------------------------------------------*/
func SetConfigPath(path string) {
	configPath = path
}

/*GetConfig --------------------------------------------------------------------------*/
/**
* @brief: GetConfig 获取配置配的配置
*
* @param: string
* @param: string
*
* @returns:
 */
/* ----------------------------------------------------------------------------*/
func GetConfig(section string, key string) *ini.Key {
	cfg, _ := ini.InsensitiveLoad(configPath)
	v, _ := cfg.Section(section).GetKey(key)
	return v
}

/*Log --------------------------------------------------------------------------*/
/**
* @brief: Log 日志记录
*
* @param:
*
* @returns:
 */
/* ----------------------------------------------------------------------------*/
func Log(data string) bool {
	timeString := time.Now().Format("2006-01-02") //strings.Join(time.Now().Year,"-") //+ "ss"
	logRoot := GetConfig("system", "log").String()
	fileName := logRoot + timeString + ".log"
	//fileName := fmt.Sprintf(os.Getenv("USERVER_INI")+"/Log/%s.log", timeString) //fmt.Sprintf("./Log/%s.log",time.Now().Year+time.Now().Month+time.Now().Day())

	//common.Log1(fileName)
	logfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(data)
	return true
}
func Log1(a ...interface{}) bool {
	timeString := time.Now().Format("2006-01-02") //strings.Join(time.Now().Year,"-") //+ "ss"
	logRoot := GetConfig("system", "log").String()
	fileName := logRoot + "craw1" + timeString + ".log"
	logfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(a)
	return true
}
func Log2(a ...interface{}) bool {
	timeString := time.Now().Format("2006-01-02") //strings.Join(time.Now().Year,"-") //+ "ss"
	logRoot := GetConfig("system", "log").String()
	fileName := logRoot + "craw2" + timeString + ".log"
	logfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(a)
	return true
}
func Log3(a ...interface{}) bool {
	timeString := time.Now().Format("2006-01-02") //strings.Join(time.Now().Year,"-") //+ "ss"
	//logRoot := GetConfig("system", "log").String()
	logRoot := "Log/" //GetConfig("system", "log").String()
	fileName := logRoot + "bankspider" + timeString + ".log"
	logfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(a)
	return true
}
func LogSec(a ...interface{}) bool {
	timeString := time.Now().Format("2006-01-02") //strings.Join(time.Now().Year,"-") //+ "ss"
	//logRoot := GetConfig("system", "log").String()
	logRoot := "Log/" //GetConfig("system", "log").String()
	fileName := logRoot + "bankspider2" + timeString + ".log"
	logfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(a)
	return true
}

func LogFirst(a ...interface{}) bool {
	timeString := time.Now().Format("2006-01-02") //strings.Join(time.Now().Year,"-") //+ "ss"
	//logRoot := GetConfig("system", "log").String()
	logRoot := "Log/" //GetConfig("system", "log").String()
	fileName := logRoot + "bankspider1" + timeString + ".log"
	logfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(a)
	return true
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateMd5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func MkdirForDownload(path string) {
	exit, _ := PathExists(path)
	if !exit {
		// 创建文件夹
		err := os.MkdirAll(path, 0777)
		if err != nil {
			Log("mkdir failed!")
		}
	}
}

func MkdirForSource(pathSource string) {
	place := strings.LastIndex(pathSource, `/`)
	path := pathSource[:place]
	place1 := strings.LastIndex(path, `/`)
	path1 := path[:place1]
	// 创建文件夹
	err := os.MkdirAll(path, 0777)
	if err != nil {
		Log("mkdir failed!")
	} else {
		Log(path + "mkdir success!")
		os.Chmod(path1, 0777)
		os.Chmod(path, 0777)

	}
}

/*
func MkdirForSource(pathSource string) {
	timeString := `/` + time.Now().Format("2006-01-02") //strings.Join(time.Now().Year,"-") //+ "ss"
	path := pathSource + timeString
	exist, err := PathExists(path)
	if err != nil {
		Log("get dir error")
		return
	}
	if exist {
		return
	} else {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			Log("mkdir failed!")
		} else {
			Log(path + "mkdir success!")
		}
	}

}
*/

func ReplaceFileType(pathName string, typename string) string {
	place := strings.LastIndex(pathName, `.`)
	temp := pathName[:place] + `.` + typename
	return temp
}

func ExecuteCmd(command string) bool {
	start := time.Now()
	_, err := exec.Command("bash", "-c", command).CombinedOutput()

	if err != nil {
		fmt.Println("execut cmd:", command)
		fmt.Println("execte cmd err:", err)
		return false
	}

	fmt.Println(" exec time  ", time.Now().Sub(start).Seconds())
	return true
}
func CheckErr(err error) {
	if err != nil {
		Log1("err:", err)
		fmt.Println("err:", err)
		panic(err.Error())
	}
}
func CheckError(err error, msg string) bool {
	if err != nil {
		Log1("check msg,err :", msg, err)
		return false
	}
	return true
}

/*
func SetCookie(c *gin.Context,name string, value string, maxAge int)  {
	domain := GetConfig("cookie","domain").String()
	c.SetCookie(name,value,maxAge,"/",domain,false,true)
}

func SetTemplate(engine *gin.Engine) {
}

func SetSession(engine *gin.Engine) {

	address:=GetConfig("redis","address").String()
	sessionsecret:=GetConfig("session","sessionsecret").String()
	sessionname:=GetConfig("session","sessionname").String()

	store, _ := sessions.NewRedisStore(10, "tcp", address, "", []byte(sessionsecret))
	engine.Use(sessions.Sessions(sessionname, store))

	//csrf支持 form表单：_csrf，url参数：_csrf，Heder参数：X-CSRF-TOKEN 或 X-XSRF-TOKEN
	//忽略的请求："GET", "HEAD", "OPTIONS"
	engine.Use(csrf.Middleware(csrf.Options{
		Secret: GetConfig("session","csrfscret").String(),
		ErrorFunc: func(c *gin.Context){
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))
}*/
