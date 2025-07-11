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
func NewJsonlPort(np Port) JsonlPort {
    return JsonlPort {
        np.Protocol,
        np.PortId,
        np.State.State,
        np.State.Reason,
        np.Service.Name,
        np.Service.Product,
        np.Service.Version,
        np.Service.Extrainfo,
    }
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

func NewJsonl(hosts []Host) [][]byte {
    var (
	    addr string
        ports []JsonlPort
	    res [][]byte
    )

    //scorriamo gli host in []hosts
    // fare il parsing degli indirizzi
    // fare il parsing delle porte
    for _, h := range hosts {
        for _, a := range h.Addrs {
            // we only care about the ipv4 address.
            if a.AddrType == "ipv4" {
                addr = a.Addr
            }
        }
        for i := 0; i < len(h.Ports); i++ {
            ports = append(ports, NewJsonlPort(h.Ports[i]))
        }

        jsonl := &Jsonl {
            Addr: addr,
            Ports: ports,
        }
        // qui faccio il marshal dello struct jsonl in json.
        // magari ritorno solo lo struct?
        // ritornare solo struct dovrebbe essere piu' testabile
        //res = append(res, j.Marshal())
        res = append(res, jsonl.Marshal())
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
