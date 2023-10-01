package abstractFactory

type FactoryInterface interface {
	ChooseFruit(name string) ProductInterface
}

type ProductInterface interface {
	Fruit() string
}

type HainanApple struct {}
type WuhanApple struct {}
type HainanBanana struct {}
type WuhanBanana struct {}

type WuhanFactory struct {}
type HainanFactory struct {}

func (*HainanApple) Fruit() string {
	return "HainanApple"
}

func (*WuhanApple) Fruit() string {
	return "WuhanApple"
}

func (*HainanBanana) Fruit() string {
	return "HainanBanana"
}

func (*WuhanBanana) Fruit() string {
	return "WuhanBanana"
}

func (*WuhanFactory) ChooseFruit(name string) ProductInterface {
	switch name {
	case "apple":
		return &WuhanApple{}
	case "banana":
		return &WuhanBanana{}
	default:
		return nil
	}
}

func (*HainanFactory) ChooseFruit(name string) ProductInterface {
	switch name {
	case "apple":
		return &HainanApple{}
	case "banana":
		return &HainanBanana{}
	default:
		return nil
	}
}

func GetFactory(name string) FactoryInterface {
	switch name {
	case "hainan":
		return &HainanFactory{}
	case "wuhan":
		return &WuhanFactory{}
	default:
		return nil
	}
}
