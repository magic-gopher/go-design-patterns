package composite

// Organization 组织结构接口
type Organization interface {
	// GetName 获取名称
	GetName() string
	// Print 打印信息
	Print()
}
