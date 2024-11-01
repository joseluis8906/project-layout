package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func New() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./configs")
	v.SetConfigName("mtx")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}

	configURL, ok := os.LookupEnv("CONFIG_URL")
	if !ok {
		return v
	}

	v.AddRemoteProvider("etcd3", configURL, "/configs/mtx.yml")
	if err := v.ReadRemoteConfig(); err != nil {
		log.Fatalf("cannot read remote config: %v", err)
	}

	return v
}
