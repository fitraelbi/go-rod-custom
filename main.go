package main

// import "time"

import "fmt"
import "github.com/go-rod/rod"
import "github.com/go-rod/rod/lib/launcher"

// func main() {
// 	// fmt.Println("Set proxy")
// 	// l := launcher.New().Set("proxy-server", "gw1.dataimpulse.com:10000")
// 	// controlURL, _ := l.Launch()
//     // browser := rod.New().ControlURL(controlURL).MustConnect()
// 	// fmt.Println("Set login")
// 	// go browser.MustHandleAuth("1b4723160125bb95b27f", "aca43dd13647c321")()

// 	// browser.MustIgnoreCertErrors(true)
// 	fmt.Println("Set ENV")
// 	l := launcher.MustNewManaged("").Set("proxy-server", "gw1.dataimpulse.com:10000")

// 	l.Set("disable-gpu").Delete("disable-gpu")

// 	l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
	
// 	browser := rod.New().Client(l.MustClient()).MustConnect()

// 	go browser.MustHandleAuth("1b4723160125bb95b27f", "aca43dd13647c321")()

// 	// browser.MustIgnoreCertErrors(true)

// 	fmt.Println("Set Page")
// 	// page := browser.MustPage("https://www.wikipedia.org/")
//     // fmt.Println("Navigating to page")
//     // page.MustWaitStable()
//     // simulateScroll(page)
//     // fmt.Println(
//     //     page.MustEval("() => document.title"),	
// 	// )
// 	page := browser.MustPage("http://api.ipify.org")

// 	// IP address should be the same, since we are using local
// 	// proxy, however the response signals that the proxy works
// 	println(page.MustElement("html").MustText())
// }


func main(){
	l := launcher.MustNewManaged("")
	l = l.Set("disable-gpu").Delete("disable-gpu")
	l = l.Set("proxy-server", "gw1.dataimpulse.com:10000")

	l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
	controlURL, _ := l.Launch()

	browser := rod.New().ControlURL(controlURL).MustConnect()

	fmt.Println("3. browser connected to:", controlURL)

	fmt.Println("Auth proxy")

	go browser.MustHandleAuth("1b4723160125bb95b27f", "aca43dd13647c321")()

	fmt.Println("Goto Page")
	
	page := browser.MustPage("http://api.ipify.org")

	println(page.MustElement("html").MustText())

}

// func UpdateProxy(b *rod.Browser) *rod.Browser {
// 	controlURL := GetControlURL()
// 	b.ControlURL(controlURL)
// 	return b
// }

// func InitBrowser() {
// 	controlURL := GetControlURL()
// 	browser = rod.New().ControlURL(controlURL).MustConnect()
// 	log.Println("3. browser connected to:", controlURL)
// }

// func GetControlURL() string {
// 	l := launcher.MustNewManaged("")
// 	l = l.Set("disable-gpu").Delete("disable-gpu")
// 	l = l.Set("proxy-server", "gw1.dataimpulse.com:10000")

// 	l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
// 	controlURL, _ := l.Launch()
// 	return controlURL
// }

// func simulateScroll(page *rod.Page) {
// 	sleep := func() { time.Sleep(time.Millisecond * 200) }

// 	for i := 0; i < 3; i++ {
// 		page.Mouse.Scroll(0, -300, 10)
// 		sleep()

// 		page.Mouse.Scroll(0, 300, 10)
// 		sleep()
// 	}
// }