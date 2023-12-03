package decorator

// 装饰者

// CoffeeDecorator 是装饰者基类
type CoffeeDecorator struct {
	Coffee Coffee
}

func (cd *CoffeeDecorator) GetDescription() string {
	return cd.Coffee.GetDescription()
}

func (cd *CoffeeDecorator) GetCost() float64 {
	return cd.Coffee.GetCost()
}

// MilkDecorator 是具体的装饰者类，添加牛奶
type MilkDecorator struct {
	*CoffeeDecorator
}

func (md *MilkDecorator) GetDescription() string {
	return md.CoffeeDecorator.GetDescription() + " + 牛奶"
}

func (md *MilkDecorator) GetCost() float64 {
	return md.CoffeeDecorator.GetCost() + 3.0
}

// SugarDecorator 是具体的装饰者类，添加糖
type SugarDecorator struct {
	*CoffeeDecorator
}

func (sd *SugarDecorator) GetDescription() string {
	return sd.CoffeeDecorator.GetDescription() + " + 糖"
}

func (sd *SugarDecorator) GetCost() float64 {
	return sd.CoffeeDecorator.GetCost() + 2.0
}
