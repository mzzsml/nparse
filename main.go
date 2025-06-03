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

// xmlDecode decodes the Nmap XML into a Nmaprun struct.
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
        // FIX: il metodo Write vuole []byte, con jsonl noi abbiamo [][]byte
        _, err = file.Write(b)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func outputJsonl(b [][]byte, t string) {
    if t == "-" {
        for i := 0; i < len(b); i++ {
            fmt.Fprintf(os.Stdout, "%s\n", b[i])
        }
    } else {
        file, err := os.Create(t)
        if err != nil {
            log.Fatal(err)
        }
        // Quando la funzione output() finisce, chiudiamo il file.
        defer file.Close()
        // FIX: il metodo Write vuole []byte, con jsonl noi abbiamo [][]byte

        for i := 0; i < len(b); i++ {
            //_, err = file.Write(b[i])
            fmt.Fprintf(file, "%s\n", b[i])
            if err != nil {
                log.Fatal(err)
            }
        }
    }
}

func main() {
    // Define the `-jsonl` flag, and set it to false (disabled) by default.
    jsonFlag := flag.Bool("jsonl", false, "output to jsonl")
    // Define the output flag, outputs to os.Stdout by default.
    outputFlag := flag.String("o", "-", "file to output. default = Stdout")
    // In orded to be able to access the file passed as an argument,
    // we need to parse all the arguments.
    flag.Parse()
    // We can access the filename using Arg(0), since it always points to the first argument
    // that IS NOT a flag.
    filename := flag.Arg(0)

    // Decode the XML into th structs defined in types.go.
    //asd := unmarshalXml(filename)
    n := xmlDecode(filename)

    // If the flag is `-jsonl`, then use the jsonl output.
    // Else, use the regular JSON output.
    if *jsonFlag {
        // Because jsonl returns a [][]byte type, and output() wants a []byte type, i had to make a separate function.
        var j Jsonl
        //outputJsonl(jsonl(n), *outputFlag) //old: uses old jsonl function
        outputJsonl(j.encode(n), *outputFlag)
    } else {
        output(jsonEncode(n), *outputFlag)
    }
}
