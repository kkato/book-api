package config

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
}

var Config ConfigList

func init() {
	Config = ConfigList{
		Port:      "8080",
		SQLDriver: "sqlite3",
		DbName:    "bookapi.db",
	}
}
