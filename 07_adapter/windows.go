package adapter

import "fmt"

type Windows struct {
}

func (w Windows) insertIntoUSBPort() {
	fmt.Println("USB连接器已插入windows机器。")
}
