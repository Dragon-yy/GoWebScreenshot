# GoWebScreenshot

GoWebScreenshot 是一个用 Golang 开发的命令行工具，利用 chromedp 库从 txt 文件中读取网站域名，批量进行网站截图。该工具会处理无法正常访问或获取截图的网站，并生成相应的日志文件。

## 安装

1. 确保已安装 Go 并设置了 GOPATH。
2. 克隆本仓库到本地。
3. 运行 `go mod tidy` 安装依赖。

## 使用

1. 在 `domains.txt` 文件中添加你要截图的网站域名，每行一个。
2. 运行以下命令开始批量截图：
    ```sh
    go run main.go screenshot --file domains.txt --output output
    ```
   
3. 运行以下命令单个跑截图
   ```sh
    go run main.go screenshot -u https://www.baidu.com
   ```

截图结果将保存在指定的输出目录中，无法访问的网站将记录在日志文件中。

## 依赖

- [chromedp](https://github.com/chromedp/chromedp)
- [cobra](https://github.com/spf13/cobra)
- [excelize](https://github.com/xuri/excelize/v2)