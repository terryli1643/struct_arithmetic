package leetcode

import (
	"math"
	"testing"
	"time"
)

func Test(t *testing.T) {
	now := time.Now()
	begin := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -29)
	data := []string{
		"2021-03-05",
		"2021-03-25",
		"2021-03-26",
		"2021-03-27",
		"2021-04-03",
	}
	cell := make([]string, 30)
	for _, d := range data {
		date, _ := time.Parse("2006-01-02", d)
		i := (int(math.Ceil(date.Sub(begin).Hours()/24)) - 1) % 30
		cell[i] = d
	}

	for i, v := range cell {
		t.Logf("%s : %s\n", begin.AddDate(0, 0, i).Format("2006-01-02"), v)
	}
}
