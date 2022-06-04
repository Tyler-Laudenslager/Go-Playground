package main

import (
   "fmt"
   "os"
   "time"
)

func main() {
   var s, sep string
   start := time.Now()
   for _, value := range os.Args[:] {
      s += sep + value 
      sep = " "
   }
   fmt.Println(s)
   fmt.Printf("%.8fs elasped\n", time.Since(start).Seconds())
}
