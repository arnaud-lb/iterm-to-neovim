package main

import (
	"fmt"
	"log"
	"os"

	plist "github.com/DHowett/go-plist"
)

type ColorSchemeColor struct {
	R float64 `plist:"Red Component"`
	G float64 `plist:"Green Component"`
	B float64 `plist:"Blue Component"`
}

type ColorScheme map[string]ColorSchemeColor

func main() {

	result := make(ColorScheme)

	decoder := plist.NewDecoder(os.Stdin)
	err := decoder.Decode(&result)
	if err != nil {
		log.Fatalf("Failed parsing the color scheme: %v", err)
	}

	for i := 0; i <= 15; i++ {
		k := fmt.Sprintf("Ansi %d Color", i)

		def, ok := result[k]
		if !ok {
			log.Printf("%s is not defined", k)
			continue
		}

		r := int(def.R * 255)
		g := int(def.G * 255)
		b := int(def.B * 255)

		hex := fmt.Sprintf("#%02x%02x%02x", r, g, b)

		fmt.Fprintf(os.Stdout, "let g:terminal_color_%d = '%s'\n", i, hex)
	}
}
