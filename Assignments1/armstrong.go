package main
// fmt package provides the function to print anything
import "fmt"

// CREATE A FUNCTION TO CHECK FOR ARMSTRONG NUMBER
func CHECKARMSTRONG(number int) bool {

   // declare
 the variables
   fmt.Printf("\nEntered Number =%d\n", number)
   temp := 0
   remainder := 0
   var result int = 0
   
   // initialize the variables
   temp = number
   
   // Use of For Loop
   // here we calculate the sum of the cube of each digit of the
   // given number to check the given number is Armstrong or not
   for {
      remainder = temp % 10
      result += remainder * remainder * remainder
      temp /= 10
      if temp == 0 {
         break // Break Statement used to stop the loop
      }
   }
   // If satisfies Armstrong condition
   if result == number {
      fmt.Printf("%d is an Armstrong number.", number)
   } else {
      fmt.Printf("%d is not an Armstrong number.", number)
   }
   return true
   // print the result
}
func main() {
   fmt.Println("GOLANG PROGRAM TO CHECK ARMSTRONG NUMBER")
   // calling the function CHECKARMSTRONG()
   CHECKARMSTRONG(153)gp
   CHECKARMSTRONG(351)
}