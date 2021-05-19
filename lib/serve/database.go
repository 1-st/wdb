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

func SaveDB()bool{
	origin,err:= ioutil.ReadFile("./database.xml")
	if err!=nil{
		log.Println("原始数据库不能读取!")
		return false
	}
	if err := os.Remove("./database.xml");err!=nil{
		log.Fatalln("原始数据库不能删除!")
		return false
	}
	f,err:= os.Create("./database.xml")
	if err!=nil{
		ioutil.WriteFile("./database.xml",origin,0777)
		log.Println("不能创建新数据库文件!")
		return false
	}
	e:=xml.NewEncoder(f)
	if err:=e.Encode(DB);err!=nil{
		log.Println("反序列化数据库失败!")
		return false
	}
	return true
}
