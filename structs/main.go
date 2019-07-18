package main
 
import "fmt"

type contactInfo struct {
	email string
	zipcode int
}
type Person struct {
	fn string
	ln string
	age int
	contact contactInfo

}
func main() {
	//var Alex Person
	//fmt.Println(Alex)
	//fmt.Printf("%+v",Alex)
	//Alex = Person {fn:"Alex",ln:"Anderson",age:22}
	//fmt.Println(Alex)
	//Alex.fn="Alex"
	//Alex.ln="Anderson"
	//Alex.age=23
	//fmt.Println(Alex)
	jim := Person{
		 fn : "Jim",
		 ln : "Party",
		 age : 25,
		 contact : contactInfo {
			 email : "jim@att.com",
			 zipcode: 93000,
		 },
	}
	fmt.Println(jim)
//Person.UpdateName() 
jimp := &jim 
jimp.UpdateName("Jimmy")
fmt.Println(jim)
}

func (ppointer *Person) UpdateName(newFn string) {
(*ppointer).fn = newFn
}