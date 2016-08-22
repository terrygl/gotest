package main
import "fmt"

type PersonInfo struct{
	ID string
	Name string
	Address string
}

func main(){
	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

	// 往map里插入数据
	personDB["1234"] = PersonInfo{"1235","terry","changan street,..."}
	personDB["1"] = PersonInfo{"1","sofr","xixiaokou street,..."}

	// 从这个 map  查找键为“1234”的信息
	person , ok := personDB["1234"]
	if ok{
		fmt.Println("Found person" , person.Name,"with ID")
	} else{
		fmt.Println("Did not find persion with ID 1234")
	}
	
	for k := range personDB{
		fmt.Println("the key is :",k,"the value is :",personDB[k])
	}
}