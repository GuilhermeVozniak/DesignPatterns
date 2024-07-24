package factory

type GunType string

const (
	GunType_AK47   GunType = "Ak47"
	GunType_Musket GunType = "Musket"
)

func GetGun(gunType GunType) (gun IGun) {
	if gunType == GunType_AK47 {
		gun = newAK47()
		return
	}

	if gunType == GunType_Musket {
		gun = newMusket()
		return
	}

	return
}
