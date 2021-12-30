package tools

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	baseUrl    = "https://fundf10.eastmoney.com/F10DataApi.aspx?type=lsjz&"
	timeFormat = "2006-01-02"
	pageSize   = 20
	now        = time.Now()
)

func GetHtmls(code string, span int) []string {
	var htmls []string

	// start date
	from := now.AddDate(0, 0, -1 * (span - 1)).Format(timeFormat)
	// today
	to := now.Format(timeFormat)

	pageNo := 1
	for {
		param := fmt.Sprintf("code=%s&page=%d&per=%d&sdate=%s&edate=%s", code, pageNo, pageSize, from, to)
		resp, err := http.Get(baseUrl + param)
		if err != nil {
			logrus.Error(err)
			break
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logrus.Warnf("can not read body, err: %s", err.Error())
		}

		bodyString := string(bodyBytes)
		html := getHtmlFromRaw(bodyString)
		if isEmpty(html) {
			break
		}

		htmls = append(htmls, html)
		pageNo += 1
	}

	return htmls
}

func isEmpty(rawHtml string) bool {
	matched, _ := regexp.MatchString("暂无数据", rawHtml)

	return matched
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

func SetNowTimeForTest(t time.Time) {
	now = t
}
