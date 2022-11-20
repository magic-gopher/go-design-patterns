package goods

// Coffee struct
type Coffee interface {
	Info(name string) string
}

// LatteCoffee 拿铁咖啡
type LatteCoffee struct {
}

// AmericanCoffee 美式咖啡
type AmericanCoffee struct {
}

// Info 拿铁咖啡详情
func (l *LatteCoffee) Info(name string) string {
	return name
}

func (a *AmericanCoffee) Info(name string) string {
	return name
}
