package main

import (
	"context"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/robfig/cron/v3"
	"github.com/ysmood/kit"
)

func main() {
	if !isLoggedIn() {
		login()
	}

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

	kit.Retry(context.Background(), kit.CountSleeper(10), func() (bool, error) {
		page := browser.Page("https://www.v2ex.com/mission/daily")
		defer page.Close()

		// 这里不可以太快，否则会触发 v2ex 的反爬虫机制
		kit.Sleep(5)

		if !page.Has("[value='领取 X 铜币']") {
			return true, nil
		}

		wait := page.WaitRequestIdle()
		page.Element("[value='领取 X 铜币']").Click()
		wait()

		page.Screenshot("")
		return false, nil
	})
}

func newBrowser(headless bool) *rod.Browser {
	url := launcher.New().Headless(headless).UserDataDir("tmp/user").Launch()
	return rod.New().ControlURL(url).Connect()
}
