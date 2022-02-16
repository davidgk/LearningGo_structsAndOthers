package autitos

type MotorWay interface {
	description() string
}

type Car struct {
	Info *Information
}

type Boat struct {
	Info *Information
}

func (m Boat) description() string {
	info := m.Info
	return ShowInfo(info)
}

func (m Car) description() string {
	info := m.Info
	return ShowInfo(info)
}

func ShowInfo(info Info) string {
	return info.GetShape() + " - " + info.GetFuel()
}

func (car *Car) changeType(s string) {
	(*car).Info.changeType(s)
}

type Shape interface {
	GetShape() string
}

type Power interface {
	GetFuel() string
}

type Info interface {
	Shape
	Power
	changeType(s string)
}

/**
Notar como Info interface.. al tener mas de una interface obliga a Information a implementar ambas interfaces.
Similar a Java.
*/
type Information struct {
	Type string
	Fuel string
}

func (i *Information) changeType(s string) {
	(*i).Type = s
}

func (i Information) GetShape() string {
	return i.Type
}

func (i Information) GetFuel() string {
	return i.Fuel
}
