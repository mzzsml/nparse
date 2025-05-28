package main

import (
    "encoding/xml"
    "io"
)

// parse does the parsing of the Nmap XML, using Nmaprun struct.
func parse(file io.Reader) (n Nmaprun) {
    d := xml.NewDecoder(file)

    // Iterate through all the stream.
    for {
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
