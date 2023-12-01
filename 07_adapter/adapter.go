package adapter

import "fmt"

// 适配器
// 这个模式可以理解为转换器，让接口兼容的两个对象能够相互合作

// WindowsAdapter Windows适配器
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("适配器将Lightning接口转换为USB接口")
	w.windowMachine.insertIntoUSBPort()
}
