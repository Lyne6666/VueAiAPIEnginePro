// cmd/vueaiapienginepro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"vueaiapienginepro/internal/vueaiapienginepro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := vueaiapienginepro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
