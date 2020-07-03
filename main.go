package main

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/robfig/cron/v3"
	"github.com/ysmood/kit"
)

func main() {
	clockIn()

	scheduler := cron.New()
	kit.E(scheduler.AddFunc("0 12 * * *", func() {
		clockIn()
	}))
	scheduler.Run()
}

func clockIn() {
	if !isLoggedIn() {
		login()
	}

	browser := newBrowser(true)
	defer browser.Close()

	page := browser.Page("https://www.v2ex.com/mission/daily")

	err := kit.Try(func() {
		wait := page.WaitRequestIdle()
		page.Timeout(10 * time.Second).Element("[value='领取 X 铜币']").Click()
		wait()
	})
	if err != nil {
		kit.Log("已经签过到了")
	} else {
		kit.Log("签到成功")
	}
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

func newBrowser(headless bool) *rod.Browser {
	url := launcher.New().Headless(headless).UserDataDir("tmp/user").Launch()
	return rod.New().ControlURL(url).Connect()
}
