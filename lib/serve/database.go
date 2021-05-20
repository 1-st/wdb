package serve

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	path2 "path"
)

var FileName = "database.xml"
var Path = ""
func init() {
	if os.Getenv("WDB_DB") == "" {
		Path = "./database.xml"
	} else {
		Path = os.Getenv("WDB_DB")
		FileName = path2.Base(Path)
	}
	f, err := os.Open(Path)
	if err != nil {
		log.Fatalln("找不到数据库!")
	}
	encoder := xml.NewDecoder(f)
	err = encoder.Decode(DB)
	if err != nil {
		log.Fatalln(FileName+"有错误")
	}
}

func SaveDB() bool {
	origin, err := ioutil.ReadFile(Path)
	if err != nil {
		log.Println("原始数据库不能读取!")
		return false
	}
	if err := os.Remove(Path); err != nil {
		log.Fatalln("原始数据库不能删除!")
		return false
	}
	f, err := os.Create(Path)
	if err != nil {
		ioutil.WriteFile(Path, origin, 0777)
		log.Println("不能创建新数据库文件!")
		return false
	}
	e := xml.NewEncoder(f)
	if err := e.Encode(DB); err != nil {
		log.Println("反序列化数据库失败!")
		return false
	}
	return true
}
