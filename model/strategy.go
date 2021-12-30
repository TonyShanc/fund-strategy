package model

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type SubStrategy struct {
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
	Span int
	// Notice: default zero if unset.
	Min float32 `yaml:"min"`
	// Notice: default zero if unset.
	Max float32 `yaml:"max"`
}

// Border represent min and max income percentage we set.
type Border struct {
	Span int
	// Notice: Zero is invalid value.
	Min float32
	// Notice: Zero is invalid value.
	Max float32
}

var (
	myStrategies *MyStrategies
)

func LoadStrategy() *MyStrategies {
	if myStrategies != nil {
		return myStrategies
	}

	yamlFile, err := ioutil.ReadFile("example/strategy.yaml")
	if err != nil {
		log.Fatalf("read yaml file err: %v", err)
	}
	if err := yaml.Unmarshal(yamlFile, &myStrategies); err != nil {
		log.Fatalf("yaml unmarshal err: %v", err)
	}
	return myStrategies
}

func (s *Strategy) GenAnalysis() []Analysis {
	anaSlice := []Analysis{}

	for _, subStra := range s.SubStrategies {
		subAna := Analysis{
			Code: s.Code,
			Msg:  subStra.Msg,
		}

		if subStra.Border != nil {
			subAna.BorderAnalysis = &BorderAnalysis{
				Border: subStra.Border,
			}
		}

		if subStra.DetectHPM != nil {
			subAna.HPMAnalysis = &HPMAnalysis{
				DetectHPM: subStra.DetectHPM,
			}
		}

		// reduce meaningless request
		if subStra.Msg == "" {
			continue
		}

		anaSlice = append(anaSlice, subAna)
	}

	return anaSlice
}
