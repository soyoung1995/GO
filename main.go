package main
import "fmt"
 
func main() {
  var a int = 20
  var ip *int
 
  ip = &a
 
  fmt.Print("1, Address of a variable: %x\n", &a)
  fmt.Printf(" 2, Address stored in ip variable: %x\n", *ip )
  fmt.Printf("3, Value of *ip variable: %d\n", *ip )

                        
