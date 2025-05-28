package main

import (
    "encoding/json"
    "log"
)

type Jsonl struct {
    Addr      string `json:"addr"`
    Port      int    `json:"port"`
    Protocol  string `json:"protocol"`
    State     string `json:"state"`
    Reason    string `json:"reason"`
    Service   string `json:"service"`
    Product   string `json:"product"`
    Extrainfo string `json:"extrainfo"`
}

func jsonl(n Nmaprun) (res [][]byte) {
    var jsonl Jsonl
    // otteniamo l'ip.
    // vogliamo soltanto l'ipv4
    // saltiamo hostname -> gia faccio una riga diversa per ogni porta, non posso farlo anche per ogni hostname (ammesso che la macchina abbia piu' hostname
    for _, host := range n.Hosts {
        for _, address := range host.Addrs {
            if address.AddrType == "ipv4" {
                jsonl.Addr = address.Addr
            }
        }
        for _, port := range host.Ports {
            jsonl.Port = port.PortId
            jsonl.Protocol = port.Protocol

            state := port.State
            jsonl.State = state.State
            jsonl.Reason = state.Reason

            service := port.Service
            jsonl.Service = service.Name
            jsonl.Product = service.Product
            jsonl.Extrainfo = service.Extrainfo

            // nel punto piu' "basso" del nostro struct, dobbiamo appendere il risultato ad una slice
            // in questo modo abbiamo ogni porta per riga.
            jsonlMarshaled, err := json.Marshal(jsonl)
            if err != nil {
                log.Fatal(err)
            }
            res = append(res, jsonlMarshaled)
        }
    }
    return res
}
