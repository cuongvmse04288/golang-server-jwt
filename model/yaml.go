package model

type Config struct {
	SQL SQL
	JWT JWT
}

type SQL struct {
	Username string
	Password string
	DBName   string `yaml:"db-name"`
	DBPort   string `yaml:"db-port"`
}

type JWT struct {
	UsingRSa       bool   `yaml:"using-rsa"`
	PrivateKeyPath string `yaml:"private-key-path"`
	PublicKeyPath  string `yaml:"public-key-path"`
	SigningKey     string `yaml:"signing-key"`
}
