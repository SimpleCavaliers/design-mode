package simpleFactory

type FruitFactory interface {
	Fruit() string
}

type apple struct {}
type banana struct {}

func (*apple) Fruit() string {
	return "Apple"
}

func (*banana) Fruit() string {
	return "Banana"
}

func GetFruit(name string) FruitFactory{
	switch name {
	case "apple":
		return &apple{}
	case "banana":
		return &banana{}
	default:
		return nil
	}
}