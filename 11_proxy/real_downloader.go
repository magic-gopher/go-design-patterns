package proxy

import "fmt"

// 真实的文件下载器（实现Downloader接口）

type RealDownloader struct{}

func (r *RealDownloader) Download(url string) {
	fmt.Printf("从%s下载文件\n", url)
}
