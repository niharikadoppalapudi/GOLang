package main

import (
	"bufio"
	"fmt"
	"os"
)

type UserInfo struct {
	FName string
	LName string
}

func main() {
	var user UserInfo
L:
	for {
		var i int
		fmt.Println("Choose any option")
		fmt.Scan(&i)
		switch i {
		case 1:
			fmt.Println("Register by entering the below details:")
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter FirstName: ")
			FName, _ := reader.ReadString('\n')
			fmt.Scanln(&FName)
			fmt.Println("Enter LastName: ")
			LName, _ := reader.ReadString('\n')
			fmt.Scanln(&LName)
			user = UserInfo{FName: FName, LName: LName}

		case 2:
			fmt.Println("View the user profile:")
			fmt.Println(user)

		case 3:
			fmt.Println("Modify/edit the profile: ")

		case 4:
			fmt.Println("Transfer the amount: ")

		case 5:
			break L
		}

	}

}
