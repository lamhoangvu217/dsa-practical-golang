package common

import (
	"bufio"
	"dsa-practicing/LinkedList"
	"dsa-practicing/Pointers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FeatureMenu struct {
	Features []Feature
	Prompt   string
}

type Feature struct {
	Key  string
	Name string
}

func (fm *FeatureMenu) String() string {
	var builder strings.Builder
	builder.WriteString("=== Feature List ===\n")
	for _, feature := range fm.Features {
		fmt.Fprintf(&builder, "%s: %s\n", feature.Key, feature.Name)
	}
	return builder.String()
}

func (fm *FeatureMenu) SelectFeature() (uint64, error) {
	fmt.Print(fm.Prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input: %w", err)
	}

	input = strings.TrimSpace(input)
	selection, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input: please enter a number")
	}

	if selection < 1 || selection > uint64(len(fm.Features)) {
		return 0, fmt.Errorf("invalid selection: please choose between 1 and %d", len(fm.Features))
	}

	return selection, nil
}

func ExecuteLinkedListFeature(selection uint64) error {
	switch selection {
	case 1:
		LinkedList.MusicPlayer()
	case 2:
		LinkedList.BrowserHistory()
	case 3:
		LinkedList.MergeTwoSortedLists()
	case 4:
		LinkedList.MultiFeatures()
	default:
		return fmt.Errorf("invalid LinkedList feature selection")
	}
	return nil
}

func ExecutePointerFeature(selection uint64) error {
	switch selection {
	case 1:
		Pointers.BasicDeclaration()
	case 2:
		Pointers.PointerToStruct()
	case 3:
		Pointers.PointerArrays()
	default:
		return fmt.Errorf("invalid Pointer feature selection")
	}
	return nil
}
