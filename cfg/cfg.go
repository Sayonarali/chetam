package cfg

type JWT struct {
	SecretKey string `env:"SECRET"`
	Sms       string `env:"SMS"`
	Phone     string `env:"PHONE"`
}
