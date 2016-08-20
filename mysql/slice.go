package main
import "fmt"

func main(){
	// var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	// // 基于数组创建一个切片
	// var mySlice []int = myArray[:5]

	// fmt.Println("Elements of myArray")

	// for _,v :=range myArray {
	// 	fmt.Print(v," ")
	// }


	// fmt.Println("\nElements of mySlice")

	// for _,v:=range mySlice{
	// 	fmt.Print(v," ")
	// }
	// fmt.Println()
	// 直接创建
	// 创建长度为5的数组切片，初始值为0
// mySlice1 := make([]int ,5)
// 创建长度为5的数组切片，初始值为0，预留长度为10
mySlice2 := make([]int ,5,10)
// 直接创建并初始化包含5个元素的数组切片
// mySlice3 := []int{1,2,3,4,5}
for i,v := range mySlice2{
	fmt.Println("mySlice[",i,"] = ",v)
}


}