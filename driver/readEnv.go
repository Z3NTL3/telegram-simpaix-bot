/*
Work for: Simpaix.net
Author: Z3NTL3
License: GNU
*/

package driver

import (
	"log"

	"github.com/spf13/viper"
)

func ReadEnv_Variables() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
