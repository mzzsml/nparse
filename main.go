package main

import (
	//"encoding/xml"
	//"encoding/json"
	//"fmt"
	//"log"
	//"os"

	"github.com/mzzsml/nparse/lib"
)

//func unmarshalXml (f string) {
//    var n lib.Nmaprun
//    filecontent, err := os.Open(f)
//    if err != nil {
//       log.Fatal(err)
//    }
//    err = xml.Unmarshal([]byte(filecontent), &n)
//    if err != nil {
//        log.Fatal(err)
//    }
//}

func main() {
	lib.ParseFlags()
}
