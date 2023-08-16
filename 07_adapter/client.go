package adapter

import "fmt"

// Client 客户端
type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(computer Computer) {
	fmt.Println("客户端将Lightning连接器插入计算机。")
	computer.InsertIntoLightningPort()
}
