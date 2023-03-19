package region

type Region string

const (
	Unknown    Region = ""
	CN1               = "cn_gf01"
	CNChannel1        = "cn_qd01"
	USA               = "os_usa"
	EURO              = "os_euro"
	ASIA              = "os_asia"
	TW                = "os_cht"
)

func GetById(id string) Region {
	first := id[:1]
	switch first {
	case "1", "2", "3", "4":
		return CN1
	case "5":
		return CNChannel1
	case "6":
		return USA
	case "7":
		return EURO
	case "8":
		return ASIA
	case "9":
		return TW
	default:
		return Unknown
	}
}
