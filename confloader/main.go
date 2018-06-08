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

	tmp := viper.GetStringMap("http")
	fmt.Printf("%T %v\n", tmp["port"], tmp["port"])
	fmt.Printf("%T %v\n", tmp["user"], tmp["user"])

	os.Exit(0)
}
