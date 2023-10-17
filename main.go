package main

import "fmt"
import "github.com/go-rod/rod"

func main() {
    page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
    page.MustWaitStable()
    page.Mouse.Scroll(0, -300, 10)
    fmt.Println(
        page.MustEval("() => document.title"),	
	)
}