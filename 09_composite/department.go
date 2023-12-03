package composite

import "fmt"

// Department 部门
type Department struct {
	name      string
	employees []Organization
}

// NewDepartment 创建新的部门
func NewDepartment(name string) *Department {
	return &Department{
		name:      name,
		employees: make([]Organization, 0),
	}
}

// GetName 获取部门名称
func (d *Department) GetName() string {
	return d.name
}

// AddEmployee 添加员工
func (d *Department) AddEmployee(employee Organization) {
	d.employees = append(d.employees, employee)
}

// Print 打印部门信息
func (d *Department) Print() {
	fmt.Printf("Department: %s\n", d.name)
	for _, employee := range d.employees {
		employee.Print()
	}
}
