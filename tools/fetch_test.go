package tools

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsWeekend(t *testing.T) {

	// Sunday
	assert.True(t, isWeekend(time.Date(2021, time.December, 26, 0, 0, 0, 0, time.UTC).Add(time.Hour*8)))

	// Saturday
	assert.True(t, isWeekend(time.Date(2021, time.December, 25, 0, 0, 0, 0, time.UTC).Add(time.Hour*8)))

	// Friday
	assert.False(t, isWeekend(time.Date(2021, time.December, 24, 0, 0, 0, 0, time.UTC).Add(time.Hour*8)))
}

func TestGetHtmls(t *testing.T) {

	var (
		tests = []struct {
			Now    time.Time
			Code   string
			Span   int
			Expect []string
		}{
			{
				Now:    time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local),
				Code:   "400015",
				Span:   1,
				Expect: nil,
			},
			{
				Now:    time.Date(2021, time.January, 2, 0, 0, 0, 0, time.Local),
				Code:   "400015",
				Span:   1,
				Expect: nil,
			},
			{
				Now:    time.Date(2021, time.December, 28, 0, 0, 0, 0, time.Local),
				Code:   "400015",
				Span:   1,
				Expect: []string{"<table class='w782 comm lsjz'><thead><tr><th class='first'>净值日期</th><th>单位净值</th><th>累计净值</th><th>日增长率</th><th>申购状态</th><th>赎回状态</th><th class='tor last'>分红送配</th></tr></thead><tbody><tr><td>2021-12-28</td><td class='tor bold'>4.2485</td><td class='tor bold'>4.7085</td><td class='tor bold red'>3.10%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr></tbody></table>"},
			},
			{
				Now:    time.Date(2021, time.December, 28, 0, 0, 0, 0, time.Local),
				Code:   "400015",
				Span:   30,
				Expect: []string{
							"<table class='w782 comm lsjz'><thead><tr><th class='first'>净值日期</th><th>单位净值</th><th>累计净值</th><th>日增长率</th><th>申购状态</th><th>赎回状态</th><th class='tor last'>分红送配</th></tr></thead><tbody><tr><td>2021-12-28</td><td class='tor bold'>4.2485</td><td class='tor bold'>4.7085</td><td class='tor bold red'>3.10%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-27</td><td class='tor bold'>4.1208</td><td class='tor bold'>4.5808</td><td class='tor bold red'>0.42%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-24</td><td class='tor bold'>4.1035</td><td class='tor bold'>4.5635</td><td class='tor bold grn'>-4.38%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-23</td><td class='tor bold'>4.2914</td><td class='tor bold'>4.7514</td><td class='tor bold grn'>-0.95%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-22</td><td class='tor bold'>4.3325</td><td class='tor bold'>4.7925</td><td class='tor bold red'>0.77%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-21</td><td class='tor bold'>4.2993</td><td class='tor bold'>4.7593</td><td class='tor bold red'>0.26%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-20</td><td class='tor bold'>4.2882</td><td class='tor bold'>4.7482</td><td class='tor bold grn'>-3.06%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-17</td><td class='tor bold'>4.4237</td><td class='tor bold'>4.8837</td><td class='tor bold grn'>-2.21%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-16</td><td class='tor bold'>4.5239</td><td class='tor bold'>4.9839</td><td class='tor bold grn'>-0.18%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-15</td><td class='tor bold'>4.5322</td><td class='tor bold'>4.9922</td><td class='tor bold grn'>-1.31%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-14</td><td class='tor bold'>4.5922</td><td class='tor bold'>5.0522</td><td class='tor bold grn'>-1.28%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-13</td><td class='tor bold'>4.6518</td><td class='tor bold'>5.1118</td><td class='tor bold grn'>-0.71%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-10</td><td class='tor bold'>4.6849</td><td class='tor bold'>5.1449</td><td class='tor bold red'>1.54%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-09</td><td class='tor bold'>4.6140</td><td class='tor bold'>5.0740</td><td class='tor bold grn'>-0.40%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-08</td><td class='tor bold'>4.6323</td><td class='tor bold'>5.0923</td><td class='tor bold red'>2.45%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-07</td><td class='tor bold'>4.5214</td><td class='tor bold'>4.9814</td><td class='tor bold grn'>-3.26%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-06</td><td class='tor bold'>4.6738</td><td class='tor bold'>5.1338</td><td class='tor bold grn'>-3.04%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-03</td><td class='tor bold'>4.8205</td><td class='tor bold'>5.2805</td><td class='tor bold red'>0.04%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-02</td><td class='tor bold'>4.8186</td><td class='tor bold'>5.2786</td><td class='tor bold red'>0.38%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-12-01</td><td class='tor bold'>4.8004</td><td class='tor bold'>5.2604</td><td class='tor bold grn'>-0.45%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr></tbody></table>",
							"<table class='w782 comm lsjz'><thead><tr><th class='first'>净值日期</th><th>单位净值</th><th>累计净值</th><th>日增长率</th><th>申购状态</th><th>赎回状态</th><th class='tor last'>分红送配</th></tr></thead><tbody><tr><td>2021-11-30</td><td class='tor bold'>4.8220</td><td class='tor bold'>5.2820</td><td class='tor bold grn'>-1.24%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr><tr><td>2021-11-29</td><td class='tor bold'>4.8825</td><td class='tor bold'>5.3425</td><td class='tor bold red'>2.67%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr></tbody></table>",
						},
			},
		}
	)

	for _, tt := range tests {
		SetNowTimeForTest(tt.Now)
		assert.Equal(t, tt.Expect, GetHtmls(tt.Code, tt.Span))
	}
}
