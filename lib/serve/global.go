package serve

import (
	"github.com/1-st/gostardict/stardict"
	"github.com/sajari/word2vec"
	"wdb/lib/constant"
)

var (
	Dict       *stardict.Dictionary
	Model      *word2vec.Model = nil
	ConfigBody                 = new(constant.Cbody)
	DB                         = new(constant.Cdb)

)
