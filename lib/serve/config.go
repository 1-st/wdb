package serve

import (
	"encoding/xml"
	"log"
	"os"
)

func init() {
	conf := ""
	if os.Getenv("WDB_MODEL") == "" {
		conf = "./data/config.xml"
	} else {
		conf = os.Getenv("WDB_CONFIG")
	}
	f, err := os.Open(conf)
	if err != nil {
		log.Fatal("找不到config.xml文件")
	}
	encoder := xml.NewDecoder(f)
	err = encoder.Decode(ConfigBody)
	if err != nil {
		log.Fatal("config.xml有错误")
	}
}
