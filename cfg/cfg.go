package cfg

type JWT struct {
	SecretKey string `env:"SECRET"`
}
