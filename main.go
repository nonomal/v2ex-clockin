package main

import (
	"github.com/robfig/cron/v3"
	"github.com/ysmood/kit"
	"github.com/ysmood/rod"
	"github.com/ysmood/rod/lib/launcher"
)

func main() {
	scheduler := cron.New()
	kit.E(scheduler.AddFunc("0 12 * * *", func() {
		if !isLoggedIn() {
			login()
		}
		clockIn()

		kit.Log("clocked in")
	}))
	scheduler.Run()
}

func isLoggedIn() bool {
	browser := newBrowser(true)
	defer browser.Close()

	return browser.Page("https://www.v2ex.com/signin").WaitLoad().HasMatches("a", "登出")
}

func login() {
	browser := newBrowser(false)
	defer browser.Close()

	browser.Page("https://www.v2ex.com/signin").ElementMatches("a", "登出")
}

func clockIn() {
	browser := newBrowser(true)
	defer browser.Close()

	page := browser.Page("https://www.v2ex.com/mission/daily").WaitLoad()

	if !page.Has("[value='领取 X 铜币']") {
		return
	}

	wait := page.WaitRequestIdle()
	page.Element("[value='领取 X 铜币']").Click()
	wait()
}

func newBrowser(headless bool) *rod.Browser {
	url := launcher.New().Headless(headless).UserDataDir("tmp").Launch()
	return rod.New().ControlURL(url).Connect()
}
