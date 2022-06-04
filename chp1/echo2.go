package main

import (
   "os"
   "time"
   "strings"
   "fmt"
)

func main() {
  start := time.Now()
  fmt.Println(strings.Join(os.Args[:], " "))
  fmt.Printf("%.8fs elasped\n", time.Since(start).Seconds())
}
