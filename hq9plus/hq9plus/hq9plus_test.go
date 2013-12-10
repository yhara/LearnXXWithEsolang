package hq9plus

import (
  "testing"
  "bytes"
  "io/ioutil"
)

func TestHello(t *testing.T) {
  w := bytes.NewBuffer(nil)
  Run("HH", w)
  if w.String() != "Hello, world!\nHello, world!\n" {
    t.Errorf("got %#v", w.String())
  }
}

func TestQuine(t *testing.T) {
  w := bytes.NewBuffer(nil)
  Run("QQ", w)
  if w.String() != "QQQQ" {
    t.Errorf("got %#v", w.String())
  }
}

func TestNinetyNineBottlesOfBeer(t *testing.T) {
  expected, err := ioutil.ReadFile("99bottles.txt")
  if err != nil { t.Error(err); return }

  w := bytes.NewBuffer(nil)
  Run("9", w)
  if w.String() != string(expected) {
    t.Errorf("got %#v", w.String())
  }
}

func TestIncrement(t *testing.T) {
  w := bytes.NewBuffer(nil)
  Run("++", w)
  if Count != 2 {
    t.Errorf("got %#v", Count)
  }
}

func TestMix(t *testing.T) {
  w := bytes.NewBuffer(nil)
  Run("HQ+", w)
  if w.String() != "Hello, world!\nHQ+" {
    t.Errorf("got %#v", w.String())
  }
}
