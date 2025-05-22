package main

import (
    "encoding/xml"
    "log"
)

func parseHost(d *xml.Decoder, xmlstart *xml.StartElement) string {
    var address string
    err := d.DecodeElement(address, xmlstart)
    if err != nil {
        log.Fatal(err)
    }
    return address
}
