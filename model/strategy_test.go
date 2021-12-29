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
							Span: 7,
							Border: &Border{
								Min: -1,
							},
						},
					},
					Msg: "天弘永利债券基金一周内跌幅超过1个点，记得加仓哟",
				},

				Expect: []Analysis{
					{
						Code: "420102",
						BorderAnalysis: []BorderAnalysis{
							{
								Span: 7,
								Border: &Border{
									Min: -1,
								},
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
							Span: 365,
							Border: &Border{
								Max: 50,
							},
							Msg: "东方新能源汽车主题混合去年涨幅超过50%，买入需谨慎",
						},
						{
							Span: 182,
							DetectHPM: &DetectHPM{
								Min: -20,
								Max: 15,
							},
						},
					},
					Msg: "东方新能源汽车主题混合去年涨幅超过50%，半年内在负20个点和15个点之间振荡，先别买入哟",
				},
				Expect: []Analysis{
					{
						Code: "400015",
						BorderAnalysis: []BorderAnalysis{
							{
								Span: 365,
								Border: &Border{
									Max: 50,
								},
							},
						},
						Msg: "东方新能源汽车主题混合去年涨幅超过50%，买入需谨慎",
					},
					{
						Code: "400015",
						BorderAnalysis: []BorderAnalysis{
							{
								Span: 365,
								Border: &Border{
									Max: 50,
								},
							},
						},
						HPMAnalysis: []HPMAnalysis{
							{
								Span: 182,
								DetectHPM: &DetectHPM{
									Min: -20,
									Max: 15,
								},
							},
						},
						Msg: "东方新能源汽车主题混合去年涨幅超过50%，半年内在负20个点和15个点之间振荡，先别买入哟",
					},
				},
			},
		}
	)

	for _, tt := range tests {
		assert.Equal(t, tt.Expect, tt.Stra.genAnalysis())
	}
}
