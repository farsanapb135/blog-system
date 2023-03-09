package config

import (
	"log"

	"github.com/spf13/viper"
)

// Data will have the configuration data from config file
var Data DbConnData

type DbConnData struct {
	DbUser        string `json:"dbuser"`
	DbPass        string `json:"dbpass"`
	DbName        string `json:"dbname"`
	DbHost        string `json:"dbhost"`
	DbPort        string `json:"dbport"`
	Port          string `json:"port"`
	HTTPSCertFile string `json:"https_cert_file"`
	HTTPSKeyFile  string `json:"https_key_file"`
	ListenAddrs   string `json:"address"`
}

// SetConfiguration will extract the config data from file
func SetConfiguration() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("No configuration file loaded - using defaults", err)
	}

	viper.SetDefault("database.dbuser", "user1")
	viper.SetDefault("database.dbpass", "password1")
	viper.SetDefault("database.dbname", "blog_system")
	viper.SetDefault("database.dbhost", "localhost")
	viper.SetDefault("database.dbport", "5432")
	viper.SetDefault("listen.address", "localhost")
	viper.SetDefault("listen.port", "8080")

	Data.Port = viper.GetString("server.port")
	Data.HTTPSCertFile = viper.GetString("server.https_cert_file")
	Data.HTTPSKeyFile = viper.GetString("server.https_key_file")

	Data.DbUser = viper.GetString("database.dbuser")
	Data.DbPass = viper.GetString("database.dbpass")
	Data.DbName = viper.GetString("database.dbname")
	Data.DbHost = viper.GetString("database.dbhost")
	Data.DbPort = viper.GetString("database.dbport")

	Data.Port = viper.GetString("listen.port")
	Data.ListenAddrs = viper.GetString("listen.address")

}
