package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/mzzsml/nparse"
)

func parseFile(p string) (n *nparse.NmapScan, err error) {
    fs, err := os.Stat(p)
    if fs == nil && err != nil {
        return
    }
    file, err := os.Open(p)
    if err != nil {
        return
    }
    b := make([]byte, fs.Size())
    if _, err = file.Read(b); err != nil {
        return
    }
    n, err = nparse.NewNmapScan(b)
    if err != nil {
        return
    }
    return
}

func main() {
    flag.Parse()
    n, _ := parseFile(flag.Arg(0))
    fmt.Printf("%s\n", n.Json())
}
