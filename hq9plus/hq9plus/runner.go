package hq9plus

import (
)

func Run(src string) {
  for _, c := range src {
    switch c {
    case 'H':
      Hello()
    case 'Q':
      Quine(src)
    case '9':
      NinetyNineBottlesOfBeer()
    }
  }
}
