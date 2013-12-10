package hq9plus

import (
  "io"
  "fmt"
)

func Hello(w io.Writer) {
  fmt.Fprint(w, "Hello, world!\n")
}
