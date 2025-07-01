package lib

import (
	"fmt"
	"log"
	"os"
)

// output prints Nparse output.
// It prints either to stdout or to a file, depending if the relative `-o` flag is set.
func Print(b []byte, t string) {
	if t == "-" {
		fmt.Fprintf(os.Stdout, "%s\n", b)
	} else {
		file, err := os.Create(t)
		if err != nil {
			log.Fatal(err)
		}
		// Quando la funzione output() finisce, chiudiamo il file.
		defer file.Close()
		_, err = file.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	}
}
