package config

type Configuration struct {
	PGDatabase string
	PGHost     string
	PGPort     int
	PGUser     string
	PGPassword string
	PGSchema   string
	PGSSLMode  string
}

var Config = Configuration{
	PGDatabase: "mission_data",
	PGHost:     "localhost",
	PGPort:     5432,
	PGUser:     "mission_data",
	PGPassword: "secret",
	PGSchema:   "journal",
	PGSSLMode:  "disable",
}
