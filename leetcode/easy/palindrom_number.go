package main
import  "fmt"
import "strconv"



func main(){
	fmt.Println(isPalindrome(1001))
}

func isPalindrome(x int) bool {
    palindrom := strconv.Itoa(x)   
        for i := range palindrom {
            if(palindrom[i] != palindrom[(len(palindrom)-1) - i]){
                return false
		} 
	}
	return true
}
