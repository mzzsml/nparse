package main

import (
    //"encoding/xml"
    "encoding/json"
    "flag"
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

// decodeXml decodes the Nmap XML into a Nmaprun struct.
func decodeXml(f string) Nmaprun {
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

func main() {
    // In orded to be able to access the file passed as an argument,
    // we need to parse all the arguments.
    flag.Parse()
    // We can access the filename with Arg(0), since it always points to the first argument
    // that IS NOT a flag.
    filename := flag.Arg(0)
    //asd := unmarshalXml(filename)
    n := decodeXml(filename)

    b := jsonEncode(n)
    fmt.Fprintf(os.Stdout, "%s\n", b)

    //jsonl
    fmt.Fprintf(os.Stdout, "JSONL: %s\n", jsonl(n))
}
