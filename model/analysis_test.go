package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/TonyShanc/fund-strategy/tools"
	"github.com/stretchr/testify/assert"
)

func TestAnalyze(t *testing.T) {
	var (
		tests = []struct {
			Ana    Analysis
			Now    time.Time
			Expect bool
		}{
			{
				// 2021-12-29
				Now: time.Date(2021, time.December, 29, 0, 0, 0, 0, time.Local),
				Ana: Analysis{
					Code: "400015",
					BorderAnalysis: &BorderAnalysis{
						Border: &Border{
							Span: 3,
							Min:  2,
						},
					},
					HPMAnalysis: &HPMAnalysis{
						DetectHPM: &DetectHPM{
							Span: 30,
							Min:  -20,
							Max:  -10,
						},
					},
					Msg: "东方新能源汽车主题混合上个月在负20个点和负10个点之间振荡，最近三天跌幅超三个点，请关注",
				},

				Expect: false,
			},
			{
				// 2021-12-28
				Now: time.Date(2021, time.December, 28, 0, 0, 0, 0, time.Local),
				Ana: Analysis{
					Code: "400015",
					BorderAnalysis: &BorderAnalysis{
						Border: &Border{
							Span: 365,
							Max:  50,
						},
					},
					Msg: "东方新能源混合去年涨幅超过50个点，请关注",
				},
				Expect: true,
			},
		}
	)

	for _, tt := range tests {
		tools.SetNowTimeForTest(tt.Now)
		fmt.Println(tt.Now)
		assert.Equal(t, tt.Expect, tt.Ana.Analyze())
	}
}
