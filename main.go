package main

import "time"

import "fmt"
import "github.com/go-rod/rod"


func main() {
    page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
    fmt.Println("Navigating to page")
    page.MustWaitStable()
    simulateScroll(page)
    fmt.Println(
        page.MustEval("() => document.title"),	
	)
}


func simulateScroll(page *rod.Page) {
	sleep := func() { time.Sleep(time.Millisecond * 200) }

	for i := 0; i < 3; i++ {
		page.Mouse.Scroll(0, -300, 10)
		sleep()

		page.Mouse.Scroll(0, 300, 10)
		sleep()
	}
}