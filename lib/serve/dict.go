package serve

import (
	"github.com/dyatlov/gostardict/stardict"
)

func main() {
	var err error
	// init dictionary with path to dictionary files and name of dictionary
	Dict, err = stardict.NewDictionary("./stardict-oxford-gb-2.4.2.tar.bz2", "oxford")
	if err != nil {
		panic(err)
	}
}
