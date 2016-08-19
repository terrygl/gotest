package main
import "fmt"
import "math"
func main(){
// var (v1 int
	// v2 string
 // 	v3 [10]int
	//  v4 []int //数组切片
	//  v5 struct{
	// 	f int
	// }
 // 	v6 * int
 // 	v7 map[string] int
 // 	v8 func(a int ) int 
// )
// v1 =10
//initialize
// var v1 int =10
// var v2 = 10
// v3 :=10
// swap variables
// v9 :=20
// fmt.Printf("%d,%d\n",v1,v9)
// v1,v9 = v9,v1
// fmt.Printf("%d,%d\n",v1,v9)
//多个返回值
// _, _, d :=GetName()
// fmt.Printf("%s\n",d);
// 预定义常量
// const (
// 	Sunday = iota //初始值为0，自己增加一
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// 	numberOfDays
// )	
// fmt.Printf("%d\n",Monday)
//浮点数
// var fvalue1 float32
// fvalue1 =12.0
// fmt.Printf("%f\n",fvalue1)
//浮点数比较
// var  fvalue1 float64 = 23.0123412341234
// var fvalue2 float64 =44.03242141
// var p float64 = 0.000000001
// fmt.Printf("%f\n",IsEqual(fvalue1,fvalue2,p))
// 复数
// var value1 complex64 
// value1 = 3.2 +12i
// value2 := 3.2 +12i
// value3 := complex(3.2 ,12)
// 字符串
// var str string
// str = "hello world"
// ch := str[0]
// fmt.Printf("The length of \"%s\" is %d \n",str,len(str))
// fmt.Printf("The first character is %c",ch)
// 遍历字符串
str := "hello world"
n := len(str)
for i:=0;i<n;i++{
	ch := str[i]
	fmt.Println(i,ch)
	fmt.Printf("%c\n",ch)
}


}
func IsEqual(f1,f2,p float64) bool{
	return math.Abs(f1-f2)<p
}
func GetName() (a,b,c string){
	return "May" , "Chan" ,"Chibi Maruko"
}
