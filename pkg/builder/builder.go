package builder

type IBuilder interface{
	SetWindowsType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

func GetBuilider(builderType string) IBuilder {

	if builderType == "normal"{
		return NewNormalBuilder()
	}

	if builderType == "igloo" {
		return NewIglooBuilder()
	}

	return nil
}