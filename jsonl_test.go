package main

import "testing"

func TestEncode(t *testing.T) {
    var test1 Nmaprun
    var host1 Host
    var address1 Address
    var port1 Port

    port1 = Port{
        Protocol: "tcp",
        PortId: 22,
        State: State{
            State: "up",
            Reason: "reason",
        },
        Service: Service{
            Name: "ssh",
            Product: "openssh",
            Extrainfo: "version 2",
        },
    }

    address1 = Address{
        Addr: "192.168.1.1",
        AddrType: "ipv4",
        Vendor: "cisco",
    }

    host1 = Host{
        Addrs: []Address{address1},
        Hostnames: []string{"test1"},
        Ports:[]Port{port1},
    }

    test1 = Nmaprun{
        Args: "nmap -O",
        Hosts: []Host{host1},
    }

    want := `{"addr":"192.168.1.1","port":22,"protocol":"tcp","state":"up","reason":"reason","service":"ssh","product":"openssh","extrainfo":"version 2"}`
    var jsonl1 Jsonl
    jsonl := jsonl1.encode(test1)
    for i := 0; i < len(jsonl); i++ {
        if string(jsonl[i]) != want {
            t.Errorf("want: %s\ngot: %s\n", want, jsonl[i])
        }
    }
}
