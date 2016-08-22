package main
import "fmt"
func main(){
	/* the below is wrong
	a:=9
	if a<5 {
		return 0
	} else {
		return 1
	}
*/
	
i:=8
	switch i {
	case 0:
		fmt.Printf("0")
	case 1: 	
	fmt.Printf("1")
	case 2: 	
	fmt.Printf("2")
	case 3: 	
	fmt.Printf("3")
	default:
		fmt.Printf("default")
	}
	Num :=5
	//比较有意思的地方
	switch {
	case 0 <=Num && Num <=3:
		fmt.Printf("0-3")
		case 4 <=Num && Num <=6:
		fmt.Printf("4-6")
		case 7<=Num && Num <=9:
		fmt.Printf("7-9") 
	}
}

// func example(x int ) int{

// 	}