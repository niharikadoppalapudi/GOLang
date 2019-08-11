package main

import "fmt"
import "time"

var(x int
	 y int)

func main(){
	var a int=3
	fmt.Println(a)
	var m,n,o= 32.5,35,42
	fmt.Println(m,n,o)
	city,state,zipcode:="plano","TX",75093
	fmt.Println(city,state,zipcode)
	var cn,ln,id="niha","dopp",836
	fmt.Println(cn,ln,id)
x,y:=10,20
fmt.Println(x)
fmt.Println(y)
 currentTime :=time.Now()
fmt.Println(currentTime.Format(time.UnixDate))
fmt.Println(currentTime.Day())

}