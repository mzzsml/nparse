package nparse

import (
    "encoding/json"
    "encoding/xml"
)

// NewNmapScan reads the content of the Nmap's XML file, in bytes, and returns a
// NmapScan struct, or an error.
func NewNmapScan(data []byte) (*NmapScan, error) {
    var n *NmapScan
    if err := xml.Unmarshal(data, &n); err != nil {
        return n, err
    }
    return n, nil
}

// NmapScan is the unmarshaled version of the Nmap's XML output.
type NmapScan struct {
    Args  string `xml:"args,attr" json:"args"`
    Hosts []Host `xml:"host" json:"hosts"`
}

// Json simply marshals an NmapScan to JSON.
func (n *NmapScan) Json() []byte {
    b, _ := json.Marshal(&n)
    return b
}

// Hosts simply returns a list of all the Hosts contained in a NmapScan struct.
func (n *NmapScan) GetHosts() []Host {
    return n.Hosts
}

// GetHost simply returns the information about the provided host.
// GetHost only accepts the host's IPv4 as parameter.
func (n *NmapScan) GetHost(addr string) Host {
    for _, host := range n.Hosts {
        for _, a := range host.Addrs {
            if a.AddrType == "ipv4" && a.Addr == addr {
                return host
            }
        }
    }
    return Host{}
}

// Host holds the host node from the XML.
type Host struct {
    Status     State      `xml:"status" json:"status"`
    Addrs      []Address  `xml:"address" json:"addrs"`
    Hostnames  []Hostname `xml:"hostnames>hostname" json:"hostnames"`
    Ports      []Port     `xml:"ports>port" json:"ports"`
    ExtraPorts ExtraPort  `xml:"ports>extraports" json:"extraports"`
    ScannedAt  string     `xml:"endtime,attr" json:"scannedAt"`
}

// AddrData returns all the information held by Address struct associated with the
// given address type ("ipv4" or "mac").
func (h Host) AddrData(addrType string) Address {
    switch addrType {
    case "ipv4":
        for _, addr := range h.Addrs {
            if addr.AddrType == "ipv4" {
                return addr
            }
        }
    case "mac":
        for _, addr := range h.Addrs {
            if addr.AddrType == "mac" {
                return addr
            }
        }
    }
    return Address{}
}

type Status struct {
    State  string `xml:"state,attr" json:"state"`
    Reason string `xml:"reason,attr" json:"reason"`
}

type Hostname struct {
    Hostaname string `xml:"name,attr" json:"hostname"`
    Type      string `xml:"type,attr" json:"type"`
}

// Address holds the address subnodes of the host nodes in the XML.
type Address struct {
    Addr     string `xml:"addr,attr" json:"addr"`
    AddrType string `xml:"addrtype,attr" json:"addrType"`
    Vendor   string `xml:"vendor,attr" json:"vendor"`
}

// Port is a marshaled version of the hosts nodes in the XML.
type Port struct {
    Protocol string  `xml:"protocol,attr" json:"protocol"`
    PortId   int     `xml:"portid,attr" json:"portId"`
    State    State   `xml:"state" json:"state"`
    Service  Service `xml:"service" json:"service"`
}

// State is a marshaled version of the hosts nodes in the XML.
type State struct {
    State  string `xml:"state,attr" json:"state"`
    Reason string `xml:"reason,attr" json:"reason"`
}

// Service is a marshaled version of the hosts nodes in the XML.
type Service struct {
    Name      string `xml:"name,attr" json:"name"`
    Product   string `xml:"product,attr" json:"product"`
    Version   string `xml:"version,attr" json:"version"`
    Extrainfo string `xml:"extrainfo,attr" json:"extrainfo"`
}

type ExtraPort struct {
    State string `xml:"state,attr" json:"state"`
    Count string `xml:"count,attr" json:"count"`
}
