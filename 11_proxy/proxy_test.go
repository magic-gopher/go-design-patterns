package proxy

import "testing"

func TestProxyPattern(t *testing.T) {
	downloader := &DownloaderProxy{}
	downloader.Download("http://example.com/file.txt")
}
