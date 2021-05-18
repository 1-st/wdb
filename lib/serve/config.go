package serve

import (
	"encoding/xml"
	"log"
	"os"
)

func init() {
	f, err := os.Open("./data/config.xml")
	if err != nil {
		log.Fatal("找不到config.xml文件")
	}
	encoder := xml.NewDecoder(f)
	err = encoder.Decode(ConfigBody)
	if err!=nil{
		log.Fatal("config.xml有错误")
	}
}
