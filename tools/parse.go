package tools

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

func ParseHtmlText(htmlText string) []float32 {
	var rst []float32

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlText))
	if err != nil {
		logrus.Warnf("can not parse html: %s, err: %s", htmlText, err.Error())
	}

	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(j int, s *goquery.Selection) {
			switch j {
			case 3:
				growth := s.Text()
				rst = append(rst, transform(growth))
			}
		})
	})

	return rst
}

func getHtmlFromRaw(raw string) string {
	r := regexp.MustCompile(`<.*>`)
	return strings.Join(r.FindAllString(raw, -1), " ")
}

// transform transforms percentage to point. e.g: 1.5% to 1.5
func transform(percentage string) float32 {
	if len(percentage) == 0 {
		return float32(0)
	}
	pointString := percentage[:len(percentage)-1]
	point, err := strconv.ParseFloat(pointString, 32)
	if err != nil {
		logrus.Warnf("transform percentage: %s err: %s", percentage, err.Error())
	}
	return float32(point)
}
