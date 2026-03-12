package nparse

import (
	"encoding/json"
	"encoding/xml"
)

// NmapScan is the unmarshaled version of the Nmap's XML output.
type NmapScan struct {
	Args  string `xml:"args,attr" json:"args"`
	Hosts []Host `xml:"host" json:"hosts"`
}

// NewNmapScan reads the content of the Nmap's XML file, in bytes, and returns a
// NmapScan struct, or an error.
func NewNmapScan(data []byte) (*NmapScan, error) {
	var n *NmapScan
	if err := xml.Unmarshal(data, &n); err != nil {
		return n, err
	}
	return n, nil
}

// Json simply marshals an NmapScan to JSON.
func (n *NmapScan) Json() []byte {
	b, _ := json.Marshal(&n)
	return b
}

// GetHosts simply returns a list of all the Hosts contained in a NmapScan.
func (n *NmapScan) GetHosts() []Host {
	var hosts []Host
	for _, h := range n.Hosts {
		hosts = append(hosts, h)
	}
	return hosts
}

// Host is a marshaled version of the hosts nodes in the XML.
type Host struct {
	Addrs     []Address `xml:"address" json:"addrs"`
	Hostnames []string  `xml:"host>hostnames" json:"hostnames"`
	Ports     []Port    `xml:"ports>port" json:"ports"`
}

// Address is a marshaled version of the hosts nodes in the XML.
type Address struct {
	Addr     string `xml:"addr,attr" json:"addr"`
	AddrType string `xml:"addrtype,attr" json:"addrType"`
	Vendor   string `xml:"vendor,attr" json:"vendor"`
}

// Port is a marshaled version of the hosts nodes in the XML.
type Port struct {
	Protocol string `xml:"protocol,attr" json:"protocol"`
	PortId   int    `xml:"portid,attr" json:"portId"`
	State    `xml:"state" json:"state"`
	Service  `xml:"service" json:"service"`
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
