package main

import (
  "os"
  "io/ioutil"
  "./hq9plus"
)

func main() {
  bytes, _ := ioutil.ReadAll(os.Stdin)
  hq9plus.Run(string(bytes), os.Stdout)
}
