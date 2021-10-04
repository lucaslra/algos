package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, e error) {
	d, err := io.ReadAll(r.r)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for i := 0; i < len(d); i++ {
		switch {
		case d[i] >= 97 && d[i] < 110 || d[i] >= 65 && d[i] < 78:
			b[i] = d[i] + 13
		case d[i] >= 110 && d[i] <= 122 || d[i] >= 78 && d[i] < 90:
			b[i] = d[i] - 13
		default:
			b[i] = d[i]
		}
	}

	return len(d), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
