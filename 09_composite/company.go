package composite

import "fmt"

// Company 公司
type Company struct {
	name        string
	departments []Organization
}

// NewCompany 创建新的公司
func NewCompany(name string) *Company {
	return &Company{
		name:        name,
		departments: make([]Organization, 0),
	}
}

// GetName 获取公司名称
func (c *Company) GetName() string {
	return c.name
}

// AddDepartment 添加部门
func (c *Company) AddDepartment(department Organization) {
	c.departments = append(c.departments, department)
}

// Print 打印公司信息
func (c *Company) Print() {
	fmt.Printf("Company: %s\n", c.name)
	for _, department := range c.departments {
		department.Print()
	}
}
