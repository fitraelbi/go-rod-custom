package main

import "time"

import "fmt"
import "github.com/go-rod/rod"
import "github.com/go-rod/rod/lib/launcher"

func main() {
	// fmt.Println("Set proxy")
	// l := launcher.New().Set("proxy-server", "gw1.dataimpulse.com:10000")
	// controlURL, _ := l.Launch()
    // browser := rod.New().ControlURL(controlURL).MustConnect()
	// fmt.Println("Set login")
	// go browser.MustHandleAuth("1b4723160125bb95b27f", "aca43dd13647c321")()

	// browser.MustIgnoreCertErrors(true)
	fmt.Println("Set ENV")
	l := launcher.MustNewManaged("")

	launcher.Open(browser.ServeMonitor(""))
	l.Set("disable-gpu").Delete("disable-gpu")

	l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
	
	browser := rod.New().Client(l.MustClient()).MustConnect()

	fmt.Println("Set Page")
	page := browser.MustPage("https://www.wikipedia.org/")
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