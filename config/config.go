package config

type Config struct {
	Github
}

type Github struct {
	Username string `env:"USERNAME,required"`
	Password string `env:"PASSWORD,required"`
}
