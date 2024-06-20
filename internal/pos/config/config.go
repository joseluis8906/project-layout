package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func New() *viper.Viper {
	v := viper.New()
	v.AddRemoteProvider("etcd3", os.Getenv("CONFIG_URL"), "/configs/pos.yml")
	v.SetConfigType("yml")
	if err := v.ReadRemoteConfig(); err != nil {
		log.Fatalf("cannot read remote config: %v", err)
	}

	return v
}
