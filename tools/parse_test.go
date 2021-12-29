package tools

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseHtmlText(t *testing.T) {
	var (
		tests = []struct {
			Text string
			expect []float32
		}{
			{
				Text: `<table class='w782 comm lsjz'>
				<thead>
					<tr>
						<th class='first'>净值日期</th>
						<th>单位净值</th>
						<th>累计净值</th>
						<th>日增长率</th>
						<th>申购状态</th>
						<th>赎回状态</th>
						<th class='tor last'>分红送配</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>2020-09-02</td>
						<td class='tor bold'>2.0318</td>
						<td class='tor bold'>2.0318</td>
						<td class='tor bold grn'>-0.25%</td>
						<td>暂停申购</td>
						<td>开放赎回</td>
						<td class='red unbold'></td>
					</tr>
					<tr>
						<td>2020-09-01</td>
						<td class='tor bold'>2.0369</td>
						<td class='tor bold'>2.0369</td>
						<td class='tor bold red'>1.22%</td>
						<td>暂停申购</td>
						<td>开放赎回</td>
						<td class='red unbold'></td>
					</tr>
				</tbody>
			</table>`,

			expect: []float32{-0.25, 1.22},
			},
		}
	)

	for _, tt := range tests {
		assert.Equal(t, tt.expect, ParseHtmlText(tt.Text))
	}
}

func TestGetHtmlFromRaw(t *testing.T) {
	var(
		tests = []struct {
			Text string 
			Expect string
		}{
			{
				Text: `var apidata={ content:"<table class='w782 comm lsjz'>
				<thead>
					<tr>
						<th class='first'>净值日期</th>
						<th>单位净值</th>
						<th>累计净值</th>
						<th>日增长率</th>
						<th>申购状态</th>
						<th>赎回状态</th>
						<th class='tor last'>分红送配</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>2020-09-02</td>
						<td class='tor bold'>2.0318</td>
						<td class='tor bold'>2.0318</td>
						<td class='tor bold grn'>-0.25%</td>
						<td>暂停申购</td>
						<td>开放赎回</td>
						<td class='red unbold'></td>
					</tr>
					<tr>
						<td>2020-09-01</td>
						<td class='tor bold'>2.0369</td>
						<td class='tor bold'>2.0369</td>
						<td class='tor bold red'>1.22%</td>
						<td>暂停申购</td>
						<td>开放赎回</td>
						<td class='red unbold'></td>
					</tr>
				</tbody>
			</table>",records:2,pages:1,curpage:1};`,
				
				Expect: `<table class='w782 comm lsjz'> <thead> <tr> <th class='first'>净值日期</th> <th>单位净值</th> <th>累计净值</th> <th>日增长率</th> <th>申购状态</th> <th>赎回状态</th> <th class='tor last'>分红送配</th> </tr> </thead> <tbody> <tr> <td>2020-09-02</td> <td class='tor bold'>2.0318</td> <td class='tor bold'>2.0318</td> <td class='tor bold grn'>-0.25%</td> <td>暂停申购</td> <td>开放赎回</td> <td class='red unbold'></td> </tr> <tr> <td>2020-09-01</td> <td class='tor bold'>2.0369</td> <td class='tor bold'>2.0369</td> <td class='tor bold red'>1.22%</td> <td>暂停申购</td> <td>开放赎回</td> <td class='red unbold'></td> </tr> </tbody> </table>`,
			},
		}
	)

	for _, tt := range tests {
		assert.Equal(t, tt.Expect, getHtmlFromRaw(tt.Text))
	}
}