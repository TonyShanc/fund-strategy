package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStrategy(t *testing.T) {
	GetStrategy()
}

func TestGenAnalysis(t *testing.T) {
	var (
		tests = []struct {
			Stra   Strategy
			Expect []Analysis
		}{
			{
				Stra: Strategy{
					Code: "420102",
					SubStrategies: []SubStrategy{
						{
							Border: &Border{
								Span: 7,
								Min:  -1,
							},
							Msg: "天弘永利债券基金一周内跌幅超过1个点，记得加仓哟",
						},
					},
				},

				Expect: []Analysis{
					{
						Code: "420102",
						BorderAnalysis: &BorderAnalysis{

							Border: &Border{
								Span: 7,
								Min:  -1,
							},
						},
						Msg: "天弘永利债券基金一周内跌幅超过1个点，记得加仓哟",
					},
				},
			},
			{
				Stra: Strategy{
					Code: "400015",
					SubStrategies: []SubStrategy{
						{
							Border: &Border{
								Span: 365,
								Max:  50,
							},
							DetectHPM: &DetectHPM{
								Span: 182,
								Min:  -20,
								Max:  15,
							},
							Msg: "东方新能源汽车主题混合去年涨幅超过50%, 过去半年在负20个点和15个点之间振荡，值得关注",
						},
					},
				},
				Expect: []Analysis{
					{
						Code: "400015",
						BorderAnalysis: &BorderAnalysis{
							Border: &Border{
								Span: 365,
								Max:  50,
							},
						},
						HPMAnalysis: &HPMAnalysis{
							DetectHPM: &DetectHPM{
								Span: 182,
								Min:  -20,
								Max:  15,
							},
						},
						Msg: "东方新能源汽车主题混合去年涨幅超过50%, 过去半年在负20个点和15个点之间振荡，值得关注",
					},
				},
			},
		}
	)

	for _, tt := range tests {
		assert.Equal(t, tt.Expect, tt.Stra.genAnalysis())
	}
}
