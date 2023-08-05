package builder

// builder

// carBuilder 车建造者接口
type carBuilder interface {
	// SetSeats 设置座位数
	SetSeats(num int)
	// SetWheel 设置车轮数
	SetWheel(num int)
	// SetColor 设置颜色
	SetColor(color string)
	// SetEngine 设置引擎
	SetEngine(engine string)
	// GetResult 返回组装好的产品结构体
	GetResult() ICar
}

// sportsCarBuilder 跑车具体构建者
type sportsCarBuilder struct {
	sc SportsCar
}

// NewSportsCarBuilder 获取sportsCarBuilder结构体指针
func NewSportsCarBuilder() *sportsCarBuilder {
	return &sportsCarBuilder{}
}

func (scb *sportsCarBuilder) SetSeats(num int) {
	scb.sc.base.seats = num
}

func (scb *sportsCarBuilder) SetWheel(num int) {
	scb.sc.base.wheel = num
}

func (scb *sportsCarBuilder) SetColor(color string) {
	scb.sc.base.color = color
}

func (scb *sportsCarBuilder) SetEngine(engine string) {
	scb.sc.base.engine = engine
}

func (scb *sportsCarBuilder) GetResult() ICar {
	return &scb.sc
}

// CarBuildEngineer 汽车组装工程师
type carBuildEngineer struct {
	cb carBuilder
}

// NewCarBuildEngineer 获取carBuildEngineer结构体指针
func NewCarBuildEngineer(builder carBuilder) *carBuildEngineer {
	cbe := &carBuildEngineer{}
	cbe.cb = builder
	return cbe
}

// BuildCar 建造汽车
func (cbe *carBuildEngineer) BuildCar(seats int, wheel int, color, engine string) ICar {
	cbe.cb.SetSeats(seats)
	cbe.cb.SetWheel(wheel)
	cbe.cb.SetColor(color)
	cbe.cb.SetEngine(engine)
	return cbe.cb.GetResult()
}
