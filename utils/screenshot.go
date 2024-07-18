package utils

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"time"
)

func TakeScreenshot(domain string, outputPath string) error {
	// 创建Chromedp上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置超时上下文
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var buf []byte
	// 如果domain以http://或https://开头，则直接访问，否则加上http://
	var url string
	if "http://" != domain[:7] && "https://" != domain[:8] {
		url = "http://" + domain
	} else {
		url = domain
	}
	err := chromedp.Run(ctx, fullScreenshot(url, 90, &buf))
	if err != nil {
		return fmt.Errorf("无法访问 %s: %v", url, err)
	}

	// 保存截图
	if err := ioutil.WriteFile(outputPath, buf, 0644); err != nil {
		return fmt.Errorf("无法保存截图: %v", err)
	}

	return nil
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second),
		chromedp.FullScreenshot(res, quality),
	}
}
