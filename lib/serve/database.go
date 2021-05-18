package serve

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

func init(){
		f, err := os.Open("./database.xml")
		if err != nil {
			log.Fatalln("当前目录找不到数据库!")
		}
		encoder := xml.NewDecoder(f)
		err = encoder.Decode(DB)
		if err!=nil{
			log.Fatalln("database.xml有错误")
		}
}

func Save(){
	origin,err:= ioutil.ReadFile("./database.xml")
	if err!=nil{
		log.Fatalln("原始数据库不能读取!")
	}
	if err := os.Remove("./database.xml");err!=nil{
		log.Fatalln("原始数据库不能删除!")
	}
	f,err:= os.Create("./database.xml")
	if err!=nil{
		ioutil.WriteFile("./database.xml",origin,0777)
		log.Fatalln("不能创建新数据库文件!")
	}
	e:=xml.NewEncoder(f)
	if err:=e.Encode(DB);err!=nil{
		log.Fatalln("反序列化数据库失败!")
	}
}
