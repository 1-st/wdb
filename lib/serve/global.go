package serve

import (
	"github.com/dyatlov/gostardict/stardict"
	"github.com/sajari/word2vec"
)

var(
	Dict *stardict.Dictionary
	Model *word2vec.Model
	ConfigBody =new(Cbody)
	DB =new(Cdb)
)
