package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("test")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	os.Exit(0)
}
