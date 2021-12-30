package main

import (
	"fmt"

	"github.com/TonyShanc/fund-strategy/model"
)

func main() {
	stras := model.LoadStrategy()
	for _, substra := range stras.Strategies {
		for _, ana := range substra.GenAnalysis() {
			if ifNotice := ana.Analyze(); ifNotice {
				fmt.Println(ana.Msg)
			}
		}
	}
}
