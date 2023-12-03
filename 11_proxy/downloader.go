package proxy

// 下载接口（抽象）

// Downloader 下载接口
type Downloader interface {
	// Download URL下载
	Download(url string)
}
