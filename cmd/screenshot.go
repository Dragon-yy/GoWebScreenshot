package cmd

import (
	"bufio"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"log"
	"os"
	"path/filepath"
	"strings"

	"GoWebScreenshot/utils"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var (
	domainFile string
	singleURL  string
	outputFile string
)

var screenshotCmd = &cobra.Command{
	Use:   "screenshot",
	Short: "Take screenshots of websites from a domain list",
	Run: func(cmd *cobra.Command, args []string) {
		if singleURL != "" {
			domains := []string{singleURL}
			runScreenshot(domains, outputFile)
		} else if domainFile != "" {
			domains, err := readDomainsFromFile(domainFile)
			if err != nil {
				log.Fatalf("无法读取域名文件: %v", err)
			}
			runScreenshot(domains, outputFile)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	screenshotCmd.Flags().StringVarP(&domainFile, "file", "f", "", "File containing the list of domains")
	screenshotCmd.Flags().StringVarP(&singleURL, "url", "u", "", "Single URL to take a screenshot of")
	screenshotCmd.Flags().StringVarP(&outputFile, "output", "o", "output.xlsx", "XLSX file to save the screenshots")
}

func readDomainsFromFile(domainFile string) ([]string, error) {
	file, err := os.Open(domainFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain != "" {
			domains = append(domains, domain)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}

func runScreenshot(domains []string, outputFile string) {
	// 创建输出文件夹和日志文件夹
	outputDir := filepath.Dir(outputFile)
	os.MkdirAll(filepath.Join(outputDir, "logs"), os.ModePerm)

	// 创建日志文件
	logFile, err := os.OpenFile(filepath.Join(outputDir, "logs", "error.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("无法创建日志文件: %v", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)

	// 创建 Excel 文件
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "Domain")
	f.SetCellValue("Sheet1", "B1", "Screenshot")

	// 读取域名并截取截图
	for i, domain := range domains {
		//替换https://和http://，为了保存图片
		text := strings.ReplaceAll(domain, "https://", "")
		text = strings.ReplaceAll(text, "http://", "")
		//fmt.Println(text)
		screenshotFile := filepath.Join(outputDir, "imgs", text+".png")
		//fmt.Println(screenshotFile)
		err := utils.TakeScreenshot(domain, screenshotFile)
		if err != nil {
			fmt.Sprintf("无法截取 %s 的截图: %v", domain, err)
			logger.Printf("无法截取 %s 的截图: %v", domain, err)
			continue
		}

		//println(i)
		rowIndex := i + 2
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowIndex), domain)
		err = f.AddPicture("Sheet1", fmt.Sprintf("B%d", rowIndex), screenshotFile, &excelize.GraphicOptions{
			ScaleX:          0.05,
			ScaleY:          0.05,
			LockAspectRatio: true,
		})
		//err = f.AddPicture("Sheet1", fmt.Sprintf("B%d", rowIndex), screenshotFile, &excelize.GraphicOptions{ScaleX: 0.5, ScaleY: 0.5})
		if err != nil {
			fmt.Sprintf("无法插入截图 %s 到 Excel: %v", screenshotFile, err)
			logger.Printf("无法插入截图 %s 到 Excel: %v", screenshotFile, err)
		}

		fmt.Printf("%s 的截图已保存至 %s 并插入到 Excel\n", domain, screenshotFile)
	}

	// 保存 Excel 文件
	if err := f.SaveAs(outputFile); err != nil {
		log.Fatalf("无法保存 Excel 文件: %v", err)
	}
}
