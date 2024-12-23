package main

import (
	"dsa-practicing/LinkedList"
	"fmt"
	"strings"
)

func FeatureSelectString() string {
	var featureSelects strings.Builder
	featureSelects.WriteString("=== Feature List ===\n")
	for _, feature := range features {
		featureSelects.WriteString(fmt.Sprintf("%s: %s\n", feature.key, feature.name))
	}
	result := featureSelects.String()
	return result
}
func main() {
	var featureInput uint
	fmt.Println(FeatureSelectString())
	fmt.Println("Select a feature you want to use:")
	fmt.Scanln(&featureInput)
	switch featureInput {
	case 1:
		LinkedList.MusicPlayer()
	case 2:
		LinkedList.BrowserHistory()
	case 3:
		LinkedList.MergeTwoSortedLists()
	default:
		fmt.Println("Feature is not exist")
	}
}
