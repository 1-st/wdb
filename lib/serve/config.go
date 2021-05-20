package serve

import (
	"encoding/xml"
	"log"
	"os"
	"wdb/lib/embed"
)

func init() {
	if os.Getenv("WDB_MODEL") != "" {
		conf := os.Getenv("WDB_CONFIG")
		f, err := os.Open(conf)
		if err != nil {
			log.Fatal("找不到config.xml文件")
		}
		encoder := xml.NewDecoder(f)
		err = encoder.Decode(ConfigBody)
		if err != nil {
			log.Fatal("config.xml有错误")
		}
	} else {

		f,err:= embed.Assets().Open("config.xml")
		if err != nil {
			log.Fatal("assets找不到config.xml文件"+err.Error())
		}
		encoder := xml.NewDecoder(f)
		err = encoder.Decode(ConfigBody)
		if err != nil {
			log.Fatal("config.xml有错误")
		}
	}
}
