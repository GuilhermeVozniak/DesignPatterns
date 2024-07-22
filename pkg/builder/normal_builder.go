package builder

type NormalBuilder struct {
	WindowType string
	DoorType   string
	NumFloors  int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) SetWindowsType() {
	nb.WindowType = "Wooden Window"
}

func (nb *NormalBuilder) SetDoorType() {
	nb.DoorType = "Wooden Door"
}

func (nb *NormalBuilder) SetNumFloor() {
	nb.NumFloors = 2
}

func (nb *NormalBuilder) GetHouse() House {
	return House{
		DoorType:   nb.DoorType,
		WindowType: nb.WindowType,
		Floor:      nb.NumFloors,
	}
}
