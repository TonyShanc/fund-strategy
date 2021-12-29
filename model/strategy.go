package model

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type SubStrategy struct {
	Span      int        `yaml:"span"`
	Border    *Border    `yaml:"border"`
	DetectHPM *DetectHPM `yaml:"detectHPM"`
	Msg       string     `yaml:"msg"`
}

type Strategy struct {
	Code          string        `yaml:"code"`
	SubStrategies []SubStrategy `yaml:"subStrategies"`
	Msg           string        `yaml:"msg"`
}

type MyStrategies struct {
	Strategies []Strategy `yaml:"myStrategies"`
}

type DetectHPM struct {
	Min float32 `yaml:"min"`
	Max float32 `yaml:"max"`
}

// Border represent min and max income percentage we set.
type Border struct {
	// Notice: Zero is invalid value.
	Min float32
	// Notice: Zero is invalid value.
	Max float32
}

var (
	myStrategies *MyStrategies
)

func GetStrategy() *MyStrategies {
	if myStrategies != nil {
		return myStrategies
	}

	yamlFile, err := ioutil.ReadFile("../example/strategy.yaml")
	fmt.Println(string(yamlFile))
	if err != nil {
		log.Fatalf("read yaml file err: %v", err)
	}
	if err := yaml.Unmarshal(yamlFile, &myStrategies); err != nil {
		log.Fatalf("yaml unmarshal err: %v", err)
	}
	return myStrategies
}

func (s *Strategy) genAnalysis() []Analysis {
	anaSlice := []Analysis{}

	rootAnalysis := Analysis{
		Code: s.Code,
		Msg:  s.Msg,
	}

	for _, stra := range s.SubStrategies {
		subAna := Analysis{
			Code: s.Code,
			Msg:  stra.Msg,
		}

		if stra.Border != nil {
			ba := BorderAnalysis{
				Span:   stra.Span,
				Border: stra.Border,
			}
			rootAnalysis.BorderAnalysis = append(rootAnalysis.BorderAnalysis, ba)
			subAna.BorderAnalysis = append(subAna.BorderAnalysis, ba)
		}

		if stra.DetectHPM != nil {
			ha := HPMAnalysis{
				Span:      stra.Span,
				DetectHPM: stra.DetectHPM,
			}
			rootAnalysis.HPMAnalysis = append(rootAnalysis.HPMAnalysis, ha)
			subAna.HPMAnalysis = append(subAna.HPMAnalysis, ha)
		}

		// reduce meaningless request
		if stra.Msg == "" {
			continue
		}

		anaSlice = append(anaSlice, subAna)
	}

	// reduce meaningless request
	if rootAnalysis.Msg != "" {
		anaSlice = append(anaSlice, rootAnalysis)
	}

	return anaSlice
}
