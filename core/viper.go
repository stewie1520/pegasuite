package core

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"pegasuite/core/internal"
	"pegasuite/global"

	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("Using gin mode %s to infer config file %s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("Using gin mode %s to infer config file %s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("Using gin mode %s to infer config file %s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else {
				config = configEnv
				fmt.Printf("Using ConfigEnv %s to get config file %s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("Using config file specified in cli flag: %s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("Using config file %s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
