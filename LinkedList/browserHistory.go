package LinkedList

import (
	"fmt"
	"strings"
)

type WebPage struct {
	url      string
	next     *WebPage // Pointer to next url
	previous *WebPage // Pointer to previous url
}

type UrlHistory struct {
	currentUrl *WebPage
}

// NewUrlHistory creates a new url history
func NewUrlHistory() *UrlHistory {
	return &UrlHistory{
		currentUrl: nil,
	}
}
func (u *UrlHistory) AddUrl(url string) {
	newUrl := &WebPage{
		url: url,
	}
	// If url history is empty, make this the first url
	if u.currentUrl == nil {
		u.currentUrl = newUrl
		return
	}

	// Clear forward history when visiting a new page
	if u.currentUrl != nil {
		newUrl.previous = u.currentUrl
		u.currentUrl.next = newUrl
	}

	// Link the new url
	u.currentUrl = newUrl
}

func (u *UrlHistory) Back() string {
	if u.currentUrl == nil || u.currentUrl.previous == nil {
		return "You are in first url"
	}
	u.currentUrl = u.currentUrl.previous
	return fmt.Sprintf("Now viewing: %s", u.currentUrl.url)
}

func (u *UrlHistory) Forward() string {
	if u.currentUrl == nil || u.currentUrl.next == nil {
		return "You are in last url"
	}
	u.currentUrl = u.currentUrl.next
	return fmt.Sprintf("Now viewing: %s", u.currentUrl.url)
}

func (u *UrlHistory) ViewBrowserHistory() string {
	// Handle empty playlist case
	if u.currentUrl == nil {
		return "Url history is empty"
	}

	// Create a string builder for efficient string concatenation
	var result strings.Builder
	result.WriteString("=== Your Urls ===\n")

	// First, find the first url in the url list
	firstUrl := u.currentUrl
	for firstUrl.previous != nil {
		firstUrl = firstUrl.previous
	}

	// Traverse through all urls starting from the first one
	current := firstUrl
	position := 1
	for current != nil {
		// Add currently viewing indicator if this is the current url
		nowViewing := " "
		if current == u.currentUrl {
			nowViewing = "▶ " // Unicode play symbol
		}

		// Format and write this song's information
		urlInfo := fmt.Sprintf("%s%d. %s\n",
			nowViewing,
			position,
			current.url)
		result.WriteString(urlInfo)

		current = current.next
		position++
	}

	result.WriteString("==================")
	return result.String()
}
func BrowserHistory() {
	urlHistory := NewUrlHistory()
	urlHistory.AddUrl("https://chatgpt.com/")
	urlHistory.AddUrl("https://ant.design/")
	urlHistory.AddUrl("https://leetcode.com/")
	//fmt.Println(urlHistory.Forward())
	fmt.Println(urlHistory.ViewBrowserHistory())
}
