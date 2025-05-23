package main

import (
    "encoding/xml"
    "os"
    "log"
    "fmt"
    //"bytes"
    "flag"

    "github.com/mzzsml/nmap2jsonl/types"
)

func output(s types.Nmaprun) {
    fmt.Fprintf(os.Stdout, "args: %v\n", s.Args)
    for _, h := range s.Hosts {
        fmt.Fprintf(os.Stdout, "host addrs: %v\n", h.Addrs)
        for _, port := range h.Ports {

            fmt.Fprintf(os.Stdout, "port proto: %v\n", port.Protocol)
            fmt.Fprintf(os.Stdout, "port id: %v\n", port.PortId)
            fmt.Fprintf(os.Stdout, "port status: %v\n", port.State)
            service := port.Service
            fmt.Fprintf(os.Stdout, "service name: %v\n", service.Name)
            fmt.Fprintf(os.Stdout, "service product: %v\n", service.Product)
            fmt.Fprintf(os.Stdout, "service extrainfo: %v\n", service.Extrainfo)
        }
    }
}

func main() {
    flag.Parse()

    // passiamo il nome del file, che viene passato come flag in go
    filename := flag.Arg(0)

    xmlfile, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    //os.Stdout.Write(xmlfile)

    // proviamo a leggere lo struct Nmaprun.
    var n types.Nmaprun
    xml.Unmarshal(xmlfile, &n)
    output(n)
}
