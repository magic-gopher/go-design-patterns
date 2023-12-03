package composite

import "testing"

// Test && Benchmark

func TestOrganization(t *testing.T) {
	// 创建公司
	company := NewCompany("ABC Company")

	// 创建部门
	department1 := NewDepartment("Sales Department")
	department2 := NewDepartment("Engineering Department")

	// 创建员工
	employee1 := NewEmployee("John")
	employee2 := NewEmployee("Alice")
	employee3 := NewEmployee("Bob")

	// 添加员工到部门
	department1.AddEmployee(employee1)
	department1.AddEmployee(employee2)
	department2.AddEmployee(employee3)

	// 添加部门到公司
	company.AddDepartment(department1)
	company.AddDepartment(department2)

	// 打印公司信息
	company.Print()
}
