package rajaongkir_go

type Config struct {
	BaseUrl     string
	Header      map[string]string
	AccountKey  string
	AccountType string
}

func NewConfig(accountKey string, accountType string) Config {
	return Config{
		BaseUrl:     GetApiURL(accountType),
		Header:      map[string]string{"key": accountKey},
		AccountKey:  accountKey,
		AccountType: accountType,
	}
}

func GetApiURL(accountType string) string {
	switch accountType {
	case AccountStarter:
		return StarterAPIUrl
	case AccountBasic:
		return BasicAPIUrl
	case AccountPro:
		return ProAPIUrl
	default:
		return StarterAPIUrl
	}
}
