package util

import (
	"math"
	"time"
)

func GetScore(times []time.Time) float64 {
	if len(times) == 0 {
		return 0
	}
	var percents = make([]float64, len(times))
	var now = time.Now()
	for k, v := range times {
		m := float64(now.Unix()-v.Unix()) / 60
		percents[k] = 1 - 0.294807*math.Pow(m, 0.116552501)
	}
	score := percents[len(percents)-1]
	for i := len(percents) - 2; i >= 0; i-- {
		score += math.Pow(1-score, 1.5) * percents[i]
	}
	if score > 0.60{
		score = math.Sqrt(1-math.Pow(score-1,2))
	}
	return score
}
