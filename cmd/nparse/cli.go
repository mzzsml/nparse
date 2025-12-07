package main

import (
    "flag"

    "github.com/mzzsml/nparse"
)

const (
    outputFlagHelp = "Write the output to the specified file."
)

func ParseFlags() {
    var (
        outputFlag string
    )

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
    n, _ := nparse.ParseFile(filename)
    nparse.Print(n.Json(), outputFlag)
}
