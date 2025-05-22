package types

import (
    "encoding/xml"
)

type Nmaprun struct {
    XMLName  xml.Name `xml:"nmaprun"`
    Args     string `xml:"args,attr"`
    Hosts    []Host `xml:"host"`
}

type Host struct {
    Addrs []Address`xml:"address"`
    Hostnames []string `xml:"host>hostnames"`
    Ports []Port `xml:"ports>port"`
}

type Address struct {
    Addr     string `xml:"addr,attr"`
    AddrType string `xml:"addrtype,attr"`
    Vendor   string `xml:"vendor,attr"`
}

type Port struct {
    Protocol string `xml:"protocol,attr"`
    PortId   int `xml:"portid,attr"`
    State    State `xml:"state"`
    Service  Service `xml:"service"`
}

type State struct {
    State string `xml:"state,attr"`
}

type Service struct {
    Name string `xml:"name,attr"`
    Product string `xml:"product,attr"`
    Extrainfo string `xml:"extrainfo,attr"`
}
