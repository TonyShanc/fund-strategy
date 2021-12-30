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
	Border            *Border
	IncomePercentages []float32
	IfEscape          bool
}

// Horizontal Price Movement detection, 横盘检测
type HPMAnalysis struct {
	DetectHPM         *DetectHPM
	IncomePercentages []float32
	IfHPM             bool
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
	ba.IncomePercentages = calIncomePercentages(ana.Code, ba.Border.Span)

	var sum float32
	for _, ip := range ba.IncomePercentages {
		sum += ip
	}

	ba.IfEscape = (ba.Border.Max != 0 && sum >= ba.Border.Max) || (ba.Border.Min != 0 && sum <= ba.Border.Min)
}

// TODO(tonyshanc): Maybe partial scope's change matters. Check partial scope's growth instead of whole scope.
func (ana *Analysis) hpmAnalyze() {
	if ana.HPMAnalysis == nil {
		return
	}

	ha := ana.HPMAnalysis
	ha.IncomePercentages = calIncomePercentages(ana.Code, ha.DetectHPM.Span)

	var sum float32
	for _, ip := range ha.IncomePercentages {
		sum += ip
		if sum <= ha.DetectHPM.Min && sum >= ha.DetectHPM.Max {
			return
		}
	}

	ha.IfHPM = true
}

func calIncomePercentages(code string, span int) (incomePercentages []float32) {
	htmls := tools.GetHtmls(code, span)

	for _, html := range htmls {
		incomePercentages = append(incomePercentages, tools.ParseHtmlText(html)...)
	}
	return
}
