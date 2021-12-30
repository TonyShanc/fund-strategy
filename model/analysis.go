package model

import (
	"github.com/TonyShanc/fund-strategy/tools"
)

// Analysis represents what we should send and receive in a strategy.
type Analysis struct {
	// Code is fund code
	Code string

	BorderAnalysis []BorderAnalysis

	HPMAnalysis []HPMAnalysis

	// Msg is what we should be noticed. If all results reach their
	// thresholds, msg will be sent.
	Msg string
}

type BorderAnalysis struct {
	Span     int
	Border   *Border
	IfEscape bool
}

type HPMAnalysis struct {
	Span int
	*DetectHPM
	IfHPM bool
}

func (ana *Analysis) Analyze() (ifNotice bool) {
	ana.borderAnalyze()
	for _, ba := range ana.BorderAnalysis {
		ifNotice = ifNotice && ba.IfEscape
	}

	ana.hpmAnalyze()
	for _, ha := range ana.HPMAnalysis {
		ifNotice = ifNotice && ha.IfHPM
	}

	return
}

func (ana *Analysis) borderAnalyze() {
	for _, ba := range ana.BorderAnalysis {
		incomePercentages := calIncomePercentages(ana.Code, ba.Span)

		var sum float32
		for _, ip := range incomePercentages {
			sum += ip
		}

		ba.IfEscape = (ba.Border.Max != 0 && sum >= ba.Border.Max) || (ba.Border.Min != 0 && sum <= ba.Border.Min)
	}
}

// TODO(tonyshanc): Maybe partial scope's change matters. Calculate partial scope's growth instead whole scope.
func (ana *Analysis) hpmAnalyze() {

	for _, ha := range ana.HPMAnalysis {
		incomePercentages := calIncomePercentages(ana.Code, ha.Span)

		var sum float32
		for _, ip := range incomePercentages {
			sum += ip
		}

		ha.IfHPM = (sum >= ha.Min && sum <= ha.Max)
	}
}

func calIncomePercentages(code string, span int) (incomePercentages []float32) {
	htmls := tools.GetHtmls(code, span)

	for _, html := range htmls {
		incomePercentages = append(incomePercentages, tools.ParseHtmlText(html)...)
	}
	return
}
