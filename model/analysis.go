package model

import (
	"github.com/TonyShanc/fund-strategy/tools"
)

// Analysis represents what we should send and receive in a strategy.
type Analysis struct {
	// Code is fund code
	Code string

	BorderAnalysis *BorderAnalysis

	HPMAnalysis *HPMAnalysis

	// Msg is what we should be noticed. If all results reach their
	// thresholds, msg will be sent.
	Msg string
}

type BorderAnalysis struct {
	Border   *Border
	Data     []float32
	IfEscape bool
}

// Horizontal Price Movement detection, 横盘分析
type HPMAnalysis struct {
	DetectHPM *DetectHPM
	Data      []float32
	IfHPM     bool
}

func (ana *Analysis) Analyze() bool {
	var (
		borderAnalyzeResult = true
		hpmAnalyzeResult    = true
	)
	ana.borderAnalyze()
	if ana.BorderAnalysis != nil {
		borderAnalyzeResult = ana.BorderAnalysis.IfEscape
	}

	ana.hpmAnalyze()
	if ana.HPMAnalysis != nil {
		hpmAnalyzeResult = ana.HPMAnalysis.IfHPM
	}

	return borderAnalyzeResult && hpmAnalyzeResult
}

func (ana *Analysis) borderAnalyze() {
	if ana.BorderAnalysis == nil {
		return
	}

	ba := ana.BorderAnalysis
	ba.Data = calIncomePercentages(ana.Code, ba.Border.Span)

	var sum float32
	for _, ip := range ba.Data {
		sum += ip
	}

	ba.IfEscape = (ba.Border.Max != 0 && sum >= ba.Border.Max) || (ba.Border.Min != 0 && sum <= ba.Border.Min)
}

// TODO(tonyshanc): Maybe partial scope's change matters. Check partial scope's growth instead whole scope.
func (ana *Analysis) hpmAnalyze() {
	if ana.HPMAnalysis == nil {
		return
	}

	ha := ana.HPMAnalysis
	ha.Data = calIncomePercentages(ana.Code, ha.DetectHPM.Span)

	var sum float32
	for _, ip := range ha.Data {
		sum += ip
	}

	ha.IfHPM = (sum >= ha.DetectHPM.Min && sum <= ha.DetectHPM.Max)
}

func calIncomePercentages(code string, span int) (incomePercentages []float32) {
	htmls := tools.GetHtmls(code, span)

	for _, html := range htmls {
		incomePercentages = append(incomePercentages, tools.ParseHtmlText(html)...)
	}
	return
}
