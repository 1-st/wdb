package serve

import (
	"fmt"
	"github.com/1-st/gostardict/stardict"
	"os"
)

func init() {
	var err error
	// init dictionary with path to dictionary files and name of dictionary
	dict:= ""
	if os.Getenv("WDB_DICT")==""{
		dict = "./data/stardict-langdao-ec-gb-2.4.2"
	}else{
		dict = os.Getenv("WDB_DICT")
	}
	Dict, err = stardict.NewDictionary(dict, "langdao-ec-gb")
	if err != nil {
		fmt.Println("不能读取词典!"+dict)
		panic(err)
	}
}
