package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func GetCookie(domain string) (result map[string]string, err error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 配置选项，启动Headless浏览器
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("disable-cache", true),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	var cookies []*network.Cookie

	tasks := chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("https://%s", domain)),
		chromedp.WaitReady(`body`),
		chromedp.WaitVisible(`body`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			cookies, err = network.GetCookies().Do(ctx)
			if err != nil {
				return err
			}

			return nil
		}),
	}

	// 执行任务
	if err = chromedp.Run(ctx, tasks); err != nil {
		return
	}

	result = make(map[string]string, len(cookies))
	for _, cookie := range cookies {
		result[cookie.Name] = cookie.Value
	}

	return
}
