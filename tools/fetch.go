package tools

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	baseUrl    = "https://fundf10.eastmoney.com/F10DataApi.aspx?type=lsjz&"
	timeFormat = "2006-01-02"
	pageSize = 20
)

func Get(span int) {

	baseUrl := "https://fundf10.eastmoney.com/F10DataApi.aspx?type=lsjz&code=400015&page=1&per=20&sdate=2021-09-10&edate=2021-12-25"

	// start date
	from := time.Now().Add(time.Hour *24 * time.Duration(span) * -1 ).Format(timeFormat)
	// today
	to := time.Now().Format(timeFormat)

	pageNo := 1
	for {
		param := fmt.Sprintf("page=%d&per=%d&sdate=%s&edate=%s", pageNo, pageSize, from, to)
		resp, err := http.Get(baseUrl + param)
		if err != nil {
			logrus.Error(err)
			break
		}
		parse(resp)

	}


}

func isWeekend(t time.Time) bool {
	t = t.UTC().Add(time.Hour * 8)
	switch t.Weekday() {
	case time.Saturday:
		return true
	case time.Sunday:
		return true
	}
	return false
}
