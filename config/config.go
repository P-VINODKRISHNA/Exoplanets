package config

import (
	"Exoplanet/constants"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Init() (err error) {
	possiblePaths := []string{
		"./" + constants.MODULE_NAME + ".toml",
		"./conf/" + constants.MODULE_NAME + ".toml",
	}
	var configPath string
	// Try to find the config file in one of the possible paths
	for _, path := range possiblePaths {
		viper.SetConfigFile(path)
		err = viper.ReadInConfig()
		if err == nil {
			configPath = path
			break
		}
	}
	if err != nil {
		log.Println("Error loading configuration file:", err)
		return err
	}
	fmt.Println("Config file loaded from path:", configPath)
	config()
	return
}

// func Init() (err error) {

// 	HOMEPATH := "./"
// 	fmt.Println("HOMEPATH", HOMEPATH)
// 	path := HOMEPATH + constants.MODULE_NAME + ".toml"
// 	viper.SetConfigFile(path)
// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		log.Println("Error loading configuration```````````````` file: ", err)
// 		return
// 	}

// 	config()

// 	return
// }

func config() (err error) {

	// if err = logs.LogInit(); err != nil {
	// 	fmt.Println("Logger inittialisation failed:", err)
	// 	return
	// }

	// logs.LogInfo("Going to connect Redis")
	// initRedis()
	// logs.LogInfo("Redis connected sccessfully")
	// initMapCache()
	// logs.LogInfo("mapcache connected sccessfully")
	log.Println("Going to connect database")
	DBinit()
	log.Println(" connected database sccessfully")
	return
}
