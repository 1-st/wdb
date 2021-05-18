package similar

import (
	"fmt"
	"github.com/sajari/word2vec"
	"log"
	"os"
	"wdb/lib/serve"
)

func init() {
	fmt.Println("加载Word2vec模型中......")
	r, _ := os.Open("./data/GoogleNews-vectors-negative300.bin")
	var err error
	// Load the model from an io.Reader (i.e. a file).
	serve.Model, err = word2vec.FromReader(r)
	if err != nil {
		log.Fatalf("模型出错: %v", err)
	}
	log.Println("加载Word2vec模型成功!")
}

func GetNetworkSimilarity(a, b string) (float32, error) {
	ea := word2vec.Expr{}
	ea.Add(1, a)
	eb := word2vec.Expr{}
	eb.Add(1, b)
	sim, err := serve.Model.Cos(ea, eb)
	if err != nil {
		return -1, err
	}
	return sim, nil
}
