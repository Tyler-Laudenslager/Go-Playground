package main

import 
(
   "fmt"
   "os"
   "strconv"
)

func main() {
   var s, sep string
   for index, value := range os.Args[1:] {
       //chAr to Integer => Itoa
       s += sep + strconv.Itoa(index) + " " + value
       sep = "\n"
   }
   fmt.Println(s)
}
