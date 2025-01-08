package main

import (
	"dsa-practicing/common"
	"fmt"
)

//	func FeatureSelectString(featuresInput []common.Feature) string {
//		var featureSelects strings.Builder
//		featureSelects.WriteString("=== Feature List ===\n")
//		for _, feature := range featuresInput {
//			featureSelects.WriteString(fmt.Sprintf("%s: %s\n", feature.Key, feature.Name))
//		}
//		result := featureSelects.String()
//		return result
//	}
func main() {
	mainMenu := &common.FeatureMenu{
		Features: features,
		Prompt:   "Select a feature you want to use: ",
	}

	linkedListMenu := &common.FeatureMenu{
		Features: linkedListSubFeatures,
		Prompt:   "Select a LinkedList sub-feature you want to use: ",
	}

	pointerMenu := &common.FeatureMenu{
		Features: pointerSubFeatures,
		Prompt:   "Select a Pointer sub-feature you want to use: ",
	}

	fmt.Print(mainMenu.String())
	mainSelection, err := mainMenu.SelectFeature()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var subMenu *common.FeatureMenu
	var executeFeature func(uint64) error

	switch mainSelection {
	case 1:
		subMenu = pointerMenu
		executeFeature = common.ExecutePointerFeature
	case 2:
		subMenu = linkedListMenu
		executeFeature = common.ExecuteLinkedListFeature
	default:
		fmt.Println("Invalid feature selection")
		return
	}

	fmt.Print(subMenu.String())
	subSelection, err := subMenu.SelectFeature()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if err := executeFeature(subSelection); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
