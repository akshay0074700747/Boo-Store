package configurations

import (
	"github.com/spf13/viper"
)

const (
	envPath = ".env"
)

type Configurations struct {
	UserDataPath       string
	UserBooksDataPath  string
	AdminBooksDataPath string
	Port               string
	Secret             string
}

func LoadConfigurationss() (Configurations, error) {

	var config Configurations

	//setting the path of the env
	viper.SetConfigFile(envPath)

	// reading the env file
	err := viper.ReadInConfig()
	if err != nil {
		return Configurations{}, err
	}

	// getting the values from env file
	config.UserDataPath = viper.GetString("user_data_path")
	config.UserBooksDataPath = viper.GetString("user_books_data_path")
	config.AdminBooksDataPath = viper.GetString("admin_books_data_path")
	config.Secret = viper.GetString("jwt_secret")
	config.Port = viper.GetString("port")

	return config, nil
}
