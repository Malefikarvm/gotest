package cntrllrs

import(
  "fmt"
)

type goclass struct{
  a int
  b string
}

func NewGoclass() *goclass{
  _ = fmt.Println
  return &goclass{}
}
