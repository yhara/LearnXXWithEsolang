package hq9plus

import (
  "io"
  "fmt"
)

func Quine(src string, w io.Writer) {
  fmt.Fprint(w, src)
}
