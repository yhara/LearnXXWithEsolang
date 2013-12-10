package hq9plus

import (
  "io"
  "fmt"
)

func NinetyNineBottlesOfBeer(w io.Writer) {
  for k:=99 ; k>=0 ; k-- {
    var before1, after, action string
    var before2 string

    switch k {
    case 0:
      before1 = "No more bottles"
      before2 = "no more bottles"
      after = "99 bottles"
    case 1:
      before1 = "1 bottle"
      after = "no more bottles"
    case 2:
      before1 = "2 bottles"
      after = "1 bottle"
    default:
      before1 = fmt.Sprintf("%d bottles", k)
      after = fmt.Sprintf("%d bottles", k-1)
    }
    if before2 == "" {
      before2 = before1
    }

    if k == 0 {
      action = "Go to the store and buy some more"
    } else {
      action = "Take one down and pass it around"
    }

    fmt.Fprintf(w, "%s of beer on the wall, %s of beer.\n",
                before1, before2)

    fmt.Fprintf(w, "%s, %s of beer on the wall.\n",
      action, after)
    if (k != 0) {
      fmt.Fprint(w, "\n")
    }
  }
}

