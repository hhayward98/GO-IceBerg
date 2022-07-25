package main
  
import (
    "fmt"
    "strings"
)
  
// Main function
func main() {
  
    // Creating and initializing strings
    str1 := "1-1=0:errr;"
    str2 := "WHERE 1=1; 'INSERT INTO users ..... "
                        // DELETE* FROM users
  
    fmt.Println("Original strings")
    fmt.Println("String 1: ", str1)
    fmt.Println("String 2: ", str2)
  
    // Checking the string present or not
    //  Using Contains() function
    res1 := strings.Contains(str1, "=")
    res2 := strings.Contains(str1, "-")
    res3 := strings.Contains(str1, ";")
    res4 := strings.Contains(str1, ":")
    res5 := strings.Contains(str1, "'")
    // res6 := strings.Contains(str1, '"')

    res21 := strings.Contains(str2, "=")
    res22 := strings.Contains(str2, "-")
    res23 := strings.Contains(str2, ";")
    res24 := strings.Contains(str2, ":")
    res25 := strings.Contains(str2, "'")
  
    // Displaying the result
    fmt.Println("\nResult 1: ", res1)
    fmt.Println("Result 2: ", res2)
    fmt.Println("Result 2: ", res3)
    fmt.Println("Result 2: ", res4)
    fmt.Println("Result 2: ", res5)
    // fmt.Println("Result 2: ", res6)
    fmt.Println("\nResult 1: ", res21)
    fmt.Println("Result 2: ", res22)
    fmt.Println("Result 2: ", res23)
    fmt.Println("Result 2: ", res24)
    fmt.Println("Result 2: ", res25)
}