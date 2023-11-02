package rajaongkir_go

type Config struct {
	Uri         string
	AccountKey  string
	AccountType string
}

func NewConfig(accountKey string, accountType string) *Config {
	return &Config{
		Uri:         GetApiURL(accountType),
		AccountKey:  accountKey,
		AccountType: accountType,
	}
}

func GetApiURL(accountType string) string {
	switch accountType {
	case AccountBasic:
		return StarterAPIUrl
	case AccountPro:
		return ProAPIUrl
	default:
		return BasicAPIUrl
	}
}
