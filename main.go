package main

import (
	"fmt"

	"github.com/DenisNosik/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	cfg.SetUser("MyUser")

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}
