package common

import (
	"github.com/spf13/viper"
	"log"
)

// ViperParameters /** Type for viper parameters.
// this type has Key for define key on environment file
// and ConfigFile for define environment file path
type ViperParameters struct {
	Key        string
	ConfigFile string
}

// KeyWithDefaultConfig /** set default config or customize configuration environment file
func (ViperParameters) KeyWithDefaultConfig(key string, configFile string) ViperParameters {
	if configFile == "" {
		configFile = ".env"
	}
	var VP ViperParameters
	VP.Key = key
	VP.ConfigFile = configFile
	return VP
}

func ViperEnvVariable(VP ViperParameters) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(VP.ConfigFile)

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(VP.Key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}