package main

import "fmt"

func main() {
  var b int = 15
  var a int 
  /* array */
  numbers := [6] int{1,2,3,4,5}
  /* for loop execution Type 1 */
  fmt.Printf("For Loop start \n ")
  for a:= 0; a < 10 ; a++ {
      fmt.Printf(" value of a: %d \n", a)
  }
  
   fmt.Printf(" While type Loop start \n ")
   
  /* for loop execution Type 2(similar to while loop) */
  for a < b {
      a++
      fmt.Printf("value of a : %d \n", a)
  }
     fmt.Printf("Froeach type Loop start \n ")
    /* for loop execution Type 2(similar to foreach loop) */
  for i,x:= range numbers {
      fmt.Printf(" value of x= %d at index %d \n ",x,i)
  }
}
