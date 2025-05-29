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
    // Define the `-jsonl` flag, and set it to false (disabled) by default.
    jsonFlag := flag.Bool("jsonl", false, "output to jsonl")
    // In orded to be able to access the file passed as an argument,
    // we need to parse all the arguments.
    flag.Parse()
    // We can access the filename using Arg(0), since it always points to the first argument
    // that IS NOT a flag.
    filename := flag.Arg(0)

    // Decode the XML into th structs defined in types.go.
    //asd := unmarshalXml(filename)
    n := decodeXml(filename)

    // If the flag is `-jsonl`, then use the jsonl output.
    // Else, use the regular JSON output.
    if *jsonFlag {
        fmt.Fprintf(os.Stdout, "%s\n", jsonl(n))
    } else {
        fmt.Fprintf(os.Stdout, "%s\n", jsonEncode(n))
    }
}
