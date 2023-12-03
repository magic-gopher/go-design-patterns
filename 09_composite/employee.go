package composite

import "fmt"

// Employee 员工
type Employee struct {
	name string
}

// NewEmployee 创建新的员工
func NewEmployee(name string) *Employee {
	return &Employee{
		name: name,
	}
}

// GetName 获取员工姓名
func (e *Employee) GetName() string {
	return e.name
}

// Print 打印员工信息
func (e *Employee) Print() {
	fmt.Printf("Employee: %s\n", e.name)
}
