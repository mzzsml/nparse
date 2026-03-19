package nparse

import (
    "testing"
    "os"
)

func TestAddrInfo(t *testing.T) {
    addripv4 := Address{"192.168.10.10", "ipv4", ""}
    addrmac := Address{"AB:CD:EF:00:00:00", "mac", "COMPANY SPA"}
    host := Host{Addrs: []Address{addripv4, addrmac}}
    if a := host.AddrInfo("ipv4"); a != (Address{}) {
        if a != addripv4 {
            t.Errorf("AddrInfo(\"ipv4\"): got %v, want %v\n", a, addripv4)
        }
    }
    if a := host.AddrInfo("mac"); a != (Address{}) {
        if a != addrmac {
            t.Errorf("AddrInfo(\"mac\"): got %v, want %v\n", a, addrmac)
        }
    }
    if a := host.AddrInfo("Asadasdadasd"); a != (Address{}) {
        t.Errorf("AddrInfo(\"Asasdasdadsd\"): illegal address type: the Address struct should be empty but got: %v\n", a)
    }
    if a := host.AddrInfo(""); a != (Address{}) {
        t.Errorf("AddrInfo(\"\"): empty address type: the Address struct should be empty but got: %v\n", a)
    }
}

func testIPv4(t *testing.T, nmapscan *NmapScan) {
    host := nmapscan.Host("192.168.1.1")
    addr := host.AddrInfo("ipv4")
    if addr != (Address{}) && addr.Addr != "192.168.1.1" {
        t.Errorf("NewNmapScan(): ipv4 addr is %s, want: 192.168.1.1", addr.Addr)
    }
}

func TestNewNmapScan(t *testing.T) {
    file, _ := os.ReadFile("test/scan.xml")
    nmapscan, _ := NewNmapScan(file)
    if nmapscan != nil && len(nmapscan.Hosts) > 0 {
        testIPv4(t, nmapscan)
    } else {
        t.Error("NewNmapScan(): nmapscan.Hosts is length 0. Should be 1 host.")
    }
}
