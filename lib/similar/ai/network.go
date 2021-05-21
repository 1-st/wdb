package ai

import (
	"fmt"
	"github.com/sajari/word2vec"
	"log"
	"os"
	"wdb/lib/serve"
)

var FastBoot = false

func Load() {
	if FastBoot {
		fmt.Println("快速启动")
		return
	}
	fmt.Println("加载Word2vec模型中...")
	mod := ""
	if os.Getenv("WDB_MODEL") == "" {
		mod = "./model/GoogleNews-vectors-negative300.bin"
	} else {
		mod = os.Getenv("WDB_MODEL")
	}
	r, err := os.Open(mod)
	if err != nil {
		log.Printf("未找到模型: %v\n", err)
		return
	}
	// Load the model from an io.Reader (i.e. a file).
	serve.Model, err = word2vec.FromReader(r)
	if err != nil {
		log.Printf("未能加载模型: %v\n", err)
	}
	fmt.Println("加载Word2vec模型成功!")
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

func GetSimilarity(str string, n int) []word2vec.Match {
	e := word2vec.Expr{}
	e.Add(1, str)
	m, err := serve.Model.CosN(e, n)
	if err != nil {
		fmt.Println("Model.CosN error")
		return nil
	} else {
		return m
	}
}
