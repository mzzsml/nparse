package lib

import (
    "encoding/xml"
    "log"
    "os"
)

func ParseFromFile(file string) Nmaprun {
	filecontent, err := os.Open(file)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	// In order for us to read the exact number of bytes,
	// we can use Stat() and Size() to get the exact file size.
	fs, err := filecontent.Stat()
	if err != nil {
		log.Fatalf("error getting file stats: %s", err)
	}
	b := make([]byte, fs.Size())
	_, err = filecontent.Read(b)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	n, err := NewNmaprun(b)
	if err != nil {
		log.Fatalf("error unmarshaling xml: %s", err)
	}
	return n
}

func NewNmaprun(data []byte) (Nmaprun, error) {
	var n Nmaprun
	err := xml.Unmarshal(data, &n)
	if err != nil {
		log.Fatalf("error unmarshaling xml: %s", err)
		// cannot return nil as a Nmaprun value,
		// so we return n, whatever it is.
		return n, err
	}
	return n, nil
}
