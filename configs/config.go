package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string    `mapstructure:"DB_DRIVER"`
	DBHost        string    `mapstructure:"DB_HOST"`
	DBPort        string    `mapstructure:"DB_PORT"`
	DBName        string    `mapstructure:"DB_NAME"`
	DBUser        string    `mapstructure:"DB_USER"`
	DBPassword    string    `mapstructure:"DB_PASS"`
	WebServerPort string    `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string    `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  time.Time `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func NewConfig() *conf {

	appDir, _ := os.Getwd()

	conf, _ := LoadConfig(appDir + "/cmd/server")
	return conf
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	fmt.Println(path)
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
