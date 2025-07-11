package lib

import (
	"flag"
)

const (
	jsonlFlagHelp  = "Get the ouput in JSON Lines format."
	outputFlagHelp = "Write the output to the specified file."
)

func ParseFlags() {
	var (
		jsonlFlag  bool
		outputFlag string
	)

	// Define both long-form flags and short-form flags.
	// JSON Lines output is disable by default.
	flag.BoolVar(&jsonlFlag, "jsonl", false, jsonlFlagHelp)
	flag.BoolVar(&jsonlFlag, "l", false, jsonlFlagHelp)
	// Output is printed on Stdout by default.
	// You're still free by specifying `-o -`.
	flag.StringVar(&outputFlag, "output", "-", outputFlagHelp)
	flag.StringVar(&outputFlag, "o", "-", outputFlagHelp)
	flag.Parse()

	// We can access the filename using Arg(0), since it always points to the first argument
	// that IS NOT a flag.
	filename := flag.Arg(0)

	// If `-jsonl` is specified, call the `outputJsonl` function.
	// The outputJsonl function then outputs the result to the value specified with the `-output` flag.
	// By default is Stdout
	if jsonlFlag {
		n := ParseFromFile(filename)
		PrintHosts(NewJsonl(n.Hosts), outputFlag)
	} else {
		// Regular JSON output.
		n := ParseFromFile(filename)
		Print(n.Encode(), outputFlag)
	}
}
