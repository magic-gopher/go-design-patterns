package proxy

import "fmt"

// 代理实例

type DownloaderProxy struct {
	realDownloader *RealDownloader
}

func (p *DownloaderProxy) Download(url string) {
	// 在下载文件之前执行一些额外的操作，比如权限验证和缓存检查
	fmt.Println("进行权限验证...")
	fmt.Println("检查缓存...")

	// 使用真正的文件下载器下载文件
	if p.realDownloader == nil {
		p.realDownloader = &RealDownloader{}
	}
	p.realDownloader.Download(url)
}
