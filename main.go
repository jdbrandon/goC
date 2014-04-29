package main
/*
extern int callc(void);
*/
import "C"
import "fmt"

func callc() int32

func main() {
   fmt.Println("Hello from go src")
   i := int(C.callc())
   fmt.Println("Back in go result was ", i)
}
