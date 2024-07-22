package builder

type IglooBuilder struct {
	WindowType string
	DoorType   string
	NumFloors  int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (ig *IglooBuilder) SetWindowsType() {
	ig.WindowType = "Snow Window"
}

func (ig *IglooBuilder) SetDoorType() {
	ig.DoorType = "Snow Door"
}

func (ig *IglooBuilder) SetNumFloor() {
	ig.NumFloors = 1
}

func (ig *IglooBuilder) GetHouse() House {
	return House{
		DoorType:   ig.DoorType,
		WindowType: ig.WindowType,
		Floor:      ig.NumFloors,
	}
}
