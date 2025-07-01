package lib

import (
    "encoding/json"
    "encoding/xml"
    "io"
    "log"
    "os"
)

type Nmaprun struct {
    Args  string `xml:"args,attr" json:"args"`
    Hosts []Host `xml:"host" json:"hosts"`
}

// jsonEncode encodes an Nmaprun variable to JSON.
func (n *Nmaprun) Encode() []byte {
    b, err := json.Marshal(&n)
    if err != nil {
        log.Fatal(err)
    }
    return b
}

// jsonl parses an Nmaprun object and turns it into a Jsonl object.
func (n *Nmaprun) GetHosts() []Host {
    var hosts []Host
    // scorriamo tutti gli host in Nmaprun e li mettiamo in Hosts
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

// xmlDecode decodes the Nmap XML into a Nmaprun struct.
// It takes a filename and returns a Nmaprun struct.
func XmlDecode(f string) Nmaprun {
    var n Nmaprun
    filecontent, err := os.Open(f)
    if err != nil {
        log.Fatal(err)
    }
    n = Parse(filecontent)
    return n
}

// parse does the parsing of the Nmap XML, using Nmaprun struct.
func Parse(file io.Reader) (n Nmaprun) {
    d := xml.NewDecoder(file)

    // Iterate through the entire stream.
    for {
        // Get the current token.
        t, err := d.Token()
        // The Token() method returns io.EOF after it returns the last token.
        // At that point, we can exit the loop.
        if err == io.EOF {
            break
        }
        // If the token is a StartElement type, and if it is equal to "nmaprun"
        // decode that element into the Nmapstruct (n).
        switch i := t.(type) {
        case xml.StartElement:
            if i.Name.Local == "nmaprun" {
                d.DecodeElement(&n, &i)
            }
        }
    }
    return n
}
