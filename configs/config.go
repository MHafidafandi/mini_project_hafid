package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	AppPort              string `mapstructure:"APP_PORT"`
	DBUsername           string `mapstructure:"DBUSERNAME"`
	DBPassword           string `mapstructure:"DBPASSWORD"`
	DBAddress            string `mapstructure:"DBADDRESS"`
	DBName               string `mapstructure:"DBNAME"`
	JWTSecret            string `mapstructure:"JWT_SECRET"`
	GeminiAPIKey         string `mapstructure:"GEMINI_API_KEY"`
	MidtransServerKeyDev string `mapstructure:"MIDTRANS_SERVER_KEY_DEV"`
	AuthString           string `mapstructure:"AUTH_STRING"`
}

var Cfg *config

func InitConfig() {
	cfg := &config{}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading env : %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Error while unmarshal env : %v", err)
	}

	Cfg = cfg

}
