package lib

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type JsonlPort struct {
    Protocol  string `json:"protocol"`
    PortId    int    `json:"port"`
    State     string `json:"state"`
    Reason    string `json:"reason"`
    Service   string `json:"service"`
    Product   string `json:"product"`
    Version string `json:"version"`
    Extrainfo string `json:"extrainfo"`
}

// Prende come parametro una porta di Nmaprun, e la converte in JsonlPort
func (p JsonlPort) ParsePort(np Port) JsonlPort {
    p.Protocol = np.Protocol
    p.PortId = np.PortId
    p.State = np.State.State
    p.Reason = np.State.Reason
    p.Service = np.Service.Name
    p.Product = np.Service.Product
    p.Version = np.Service.Version
    p.Extrainfo = np.Service.Extrainfo

    return p
}

type Jsonl struct {
    Addr  string      `json:"addr"`
    Ports []JsonlPort `json:"ports"`
}

func (j *Jsonl) Marshal() []byte {
    marshaled, err := json.Marshal(j)
    if err != nil {
        log.Fatal(err)
    }
    return marshaled
}

func (j *Jsonl) ParseHosts(hosts []Host) [][]byte {
    var res [][]byte
    //scorriamo gli host in []hosts
    // fare il parsing degli indirizzi
    // fare il parsing delle porte
    for _, h := range hosts {
        for _, a := range h.Addrs {
            if a.AddrType == "ipv4" {
                j.Addr = a.Addr
            }
        }
        var ports []JsonlPort
        var port JsonlPort
        for _, p := range h.Ports {
            ports = append(ports, port.ParsePort(p))
        }
        j.Ports = ports
        res = append(res, j.Marshal())
    }
    return res
}

// ouputJson takes a slice of marshaled json objects ([]byte) and writes each of them
// in the specified file. If no file is specified it prints the json objects to stdout.
func PrintHosts(b [][]byte, t string) {
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
