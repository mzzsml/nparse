package main

type Nmaprun struct {
	Args  string `xml:"args,attr" json:"args"`
	Hosts []Host `xml:"host" json:"hosts"`
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
	Extrainfo string `xml:"extrainfo,attr" json:"extrainfo"`
}
