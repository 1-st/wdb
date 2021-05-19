package similar

import (
	"fmt"
	"testing"
	"time"
)

func TestDiff(t *testing.T){
	for{
		time.Sleep(time.Second)
		fmt.Println(GetDiffSimilarity("super","superb"))
	}
}
