package serve

import (
	"fmt"
	"os"
	"wdb/lib/stardict"
)

func init() {
	var err error
	// init dictionary with path to dictionary files and name of dictionary
	dict := ""
	if os.Getenv("WDB_DICT") == "" {
		dict = "stardict-langdao-ec-gb-2.4.2"
		Dict, err = stardict.NewDictionaryFromAssets(dict, "langdao-ec-gb")
		if err != nil {
			fmt.Println("不能读取词典Assets!" + dict)
			panic(err)
		}
	} else {
		dict = os.Getenv("WDB_DICT")
		Dict, err = stardict.NewDictionary(dict, "langdao-ec-gb")
		if err != nil {
			fmt.Println("不能读取词典!" + dict)
			panic(err)
		}
	}

}
