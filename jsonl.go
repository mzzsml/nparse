package main

import (
    "encoding/json"
    "log"
    "fmt"
    "os"
)

// The Jsonl struct encodes the Nmap XML output into sinle-lines, self-closing JSON objects.
type Jsonl struct {
    Addr      string `json:"addr"`
    Port      int    `json:"port"`
    Protocol  string `json:"protocol"`
    State     string `json:"state"`
    Reason    string `json:"reason"`
    Service   string `json:"service"`
    Product   string `json:"product"`
    Extrainfo string `json:"extrainfo"`
}

// jsonl parses an Nmaprun object and turns it into a Jsonl object.
func (j *Jsonl) encode(n Nmaprun) (res [][]byte) {
    // vogliamo soltanto l'ipv4
    // saltiamo hostname -> gia faccio una riga diversa per ogni porta, non posso farlo anche per ogni hostname (ammesso che la macchina abbia piu' hostname
    for _, host := range n.Hosts {
        for _, address := range host.Addrs {
            if address.AddrType == "ipv4" {
                j.Addr = address.Addr
            }
        }
        for _, port := range host.Ports {
            j.Port = port.PortId
            j.Protocol = port.Protocol

            j.State = port.State.State
            j.Reason = port.State.Reason

            j.Service = port.Service.Name
            j.Product = port.Service.Product
            j.Extrainfo = port.Service.Extrainfo

            // nel punto piu' "basso" del nostro struct, dobbiamo appendere il risultato ad una slice
            // in questo modo abbiamo ogni porta per riga.
            jsonlMarshaled, err := json.Marshal(j)
            if err != nil {
                log.Fatal(err)
            }
            res = append(res, jsonlMarshaled)
        }
    }
    return res
}

// ouputJson takes a slice of marshaled json objects ([]byte) and writes each of them
// in the specified file. If no file is specified it prints the json objects to stdout.
func outputJsonl(b [][]byte, t string) {
    if t == "-" {
        for i := 0; i < len(b); i++ {
            fmt.Fprintf(os.Stdout, "%s\n", b[i])
        }
    } else {
        file, err := os.Create(t)
        if err != nil {
            log.Fatal(err)
        }
        // Close the file when `outputJsonl` returns.
        defer file.Close()
        for i := 0; i < len(b); i++ {
            //_, err = file.Write(b[i])
            fmt.Fprintf(file, "%s\n", b[i])
            if err != nil {
                log.Fatal(err)
            }
        }
    }
}
