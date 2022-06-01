package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	BindAddress       string  `mapstructure:"bind_address"`
	Port              string  `mapstructure:"listen_port"`
	ProxyProtocolPort string  `mapstructure:"proxyprotocol_port"`
	ServerLat         float64 `mapstructure:"server_lat"`
	ServerLng         float64 `mapstructure:"server_lng"`
	IPInfoAPIKey      string  `mapstructure:"ipinfo_api_key"`

	AssetsPath string `mapstructure:"assets_path"`
}

var (
	configFile   string  = ""
	loadedConfig *Config = nil
)

func init() {
	viper.SetDefault("listen_port", "8989")
	viper.SetDefault("proxyprotocol_port", "0")
	viper.SetDefault("download_chunks", 4)
	viper.SetDefault("distance_unit", "K")
	viper.SetDefault("enable_cors", false)
	viper.SetDefault("assets_path", "./assets")

	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
}

func Load(configPath string) Config {
	var conf Config

	configFile = configPath
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warnf("No config file found in search paths, using default values")
		} else {
			log.Fatalf("Error reading config: %s", err)
		}
	}

	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Error parsing config: %s", err)
	}

	loadedConfig = &conf

	return conf
}

func LoadedConfig() *Config {
	if loadedConfig == nil {
		Load(configFile)
	}
	return loadedConfig
}
