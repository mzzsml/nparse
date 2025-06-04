package main

import (
    //"encoding/xml"
    "encoding/json"
    "fmt"
    "log"
    "os"
)

//func unmarshalXml (f string) {
//    var n Nmaprun
//    filecontent, err := os.Open(f)
//    if err != nil {
//       log.Fatal(err)
//    }
//    err = xml.Unmarshal([]byte(filecontent), &n)
//    if err != nil {
//        log.Fatal(err)
//    }
//}

// xmlDecode decodes the Nmap XML into a Nmaprun struct.
// It takes a filename and returns a Nmaprun struct.
func xmlDecode(f string) Nmaprun {
    var n Nmaprun
    filecontent, err := os.Open(f)
    if err != nil {
        log.Fatal(err)
    }
    n = parse(filecontent)
    return n
}

// jsonEncode encodes an Nmaprun variable to JSON.
func jsonEncode(n Nmaprun) []byte {
    b, err := json.Marshal(&n)
    if err != nil {
        log.Fatal(err)
    }
    return b
}

// output prints Nparse output.
// It prints either to stdout or to a file, depending if the relative `-o` flag is set.
func output(b []byte, t string) {
    if t == "-" {
        fmt.Fprintf(os.Stdout, "%s\n", b)
    } else {
        file, err := os.Create(t)
        if err != nil {
            log.Fatal(err)
        }
        // Quando la funzione output() finisce, chiudiamo il file.
        defer file.Close()
        _, err = file.Write(b)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func main() {
    parseFlags()
}
