package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 替换以下ed2k链接为您想要下载的链接
	ed2kURL := "ed2k://|file|%E6%97%A0%E8%80%BB%E5%AE%B6%E5%BA%AD.Shameless.US.S02E03.Chi_Eng.HR-HDTV.AC3.1024X576.x264-YYeTs%E4%BA%BA%E4%BA%BA%E5%BD%B1%E8%A7%86.mkv|523125707|a9ce9c25c9f3e67aaeb9c764bed7edf3|h=hmql6c5av5lowxqz7qr3ifjeav6gzi27|/"

	fileName := "example_file.txt"

	err := downloadED2K(ed2kURL, fileName)
	if err != nil {
		fmt.Println("下载失败:", err)
	} else {
		fmt.Println("下载完成!")
	}
}

func downloadED2K(url, fileName string) error {
	// 发起HTTP请求获取ed2k链接对应文件的下载链接
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建要下载的文件
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将文件内容写入到本地文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
