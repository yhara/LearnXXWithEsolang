package hq9plus

import (
  "io"
)

func Run(src string, w io.Writer) {
  for _, c := range src {
    switch c {
    case 'H':
      Hello(w)
    case 'Q':
      Quine(src, w)
    case '9':
      NinetyNineBottlesOfBeer(w)
    case '+':
      Increment()
    }
  }
}
