package nparse

import (
    "encoding/xml"
    "encoding/json"
    "log"
    "os"
    "fmt"
)

type NmapScan struct {
    Args  string `xml:"args,attr" json:"args"`
    Hosts []Host `xml:"host" json:"hosts"`
}

func NewNmapScan(data []byte) (n *NmapScan, err error) {
    err = xml.Unmarshal(data, &n)
    if err != nil {
        log.Fatalf("error unmarshaling xml: %s", err)
        return n, err
    }
    return n, nil
}

// jsonEncode encodes an NmapScan variable to JSON.
func (n *NmapScan) Json() []byte {
    b, err := json.Marshal(&n)
    if err != nil {
        log.Fatal(err)
    }
    return b
}

// jsonl parses an NmapScan object and turns it into a Jsonl object.
func (n *NmapScan) GetHosts() []Host {
    var hosts []Host
    // scorriamo tutti gli host in NmapScan e li mettiamo in Hosts
    for _, h := range n.Hosts {
        hosts = append(hosts, h)
    }
    return hosts
}

type Host struct {
    Addrs     []Address `xml:"address" json:"addrs"`
    Hostnames []string  `xml:"host>hostnames" json:"hostnames"`
    Ports     []Port    `xml:"ports>port" json:"ports"`
}

type Address struct {
    Addr     string `xml:"addr,attr" json:"addr"`
    AddrType string `xml:"addrtype,attr" json:"addrType"`
    Vendor   string `xml:"vendor,attr" json:"vendor"`
}

type Port struct {
    Protocol string  `xml:"protocol,attr" json:"protocol"`
    PortId   int     `xml:"portid,attr" json:"portId"`
    State    State   `xml:"state" json:"state"`
    Service  Service `xml:"service" json:"service"`
}

type State struct {
    State  string `xml:"state,attr" json:"state"`
    Reason string `xml:"reason,attr" json:"reason"`
}

type Service struct {
    Name      string `xml:"name,attr" json:"name"`
    Product   string `xml:"product,attr" json:"product"`
    Version   string `xml:"version,attr" json:"version"`
    Extrainfo string `xml:"extrainfo,attr" json:"extrainfo"`
}

func ParseFile(p string) (n *NmapScan, err error) {
    fs, err := os.Stat(p)
    if fs == nil &&  err != nil {
        return
    }
    file, err := os.Open(p)
    if err != nil {
        return
    }
    b := make([]byte, fs.Size())
    _, err = file.Read(b)
    if err != nil {
        return
    }
    n, err = NewNmapScan(b)
    if err != nil {
        return
    }
    return
}

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
