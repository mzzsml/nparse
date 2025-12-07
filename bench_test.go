package nparse

import (
    "bytes"
    "os"
    "testing"
)

func BenchmarkDecode(b *testing.B) {
    filename := "test/nmaptest.xml"
    file, _ := os.ReadFile(filename)
    stream := bytes.NewReader(file)

    for i := 0; i < b.N; i++ {
        Parse(stream)
    }
}
