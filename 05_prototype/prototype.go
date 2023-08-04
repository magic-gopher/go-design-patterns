package prototype

// Sheep 原型接口
type Sheep interface {
	// Eat 吃
	Eat(food string)
	// Clone 克隆
	Clone() Sheep
}

// Goat 山羊
type Goat struct {
	food string
}

// NewGoat 获取NewGoat结构体指针
func NewGoat() *Goat {
	return &Goat{}
}

func (g *Goat) Eat(food string) {
	g.food = food
}

func (g *Goat) Clone() Sheep {
	clone := *g
	return &clone
}
