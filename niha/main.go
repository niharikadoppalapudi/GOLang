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

func RegisterUser() {
	strArray := []string{"user1", "user2"}
	for k := 0; k <= len(strArray); k++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter FirstName: ")
		FName, _ := reader.ReadString('\n')
		fmt.Scan(&FName)
		fmt.Println("Enter LastName: ")
		LName, _ := reader.ReadString('\n')
		fmt.Scan(&LName)
		//user = UserInfo{FName: FName, LName: LName}
		Mk := make(map[UserInfo]string)
		strArray[k] = &UserInfo{FName, LName}
		Mk[*strArray[k]] = "Balance Amount:"
		fmt.Println(Mk)

	}
}
func main() {
	//var user UserInfo
L:
	for {
		var i int
		fmt.Println("Choose any option")
		fmt.Scan(&i)
		switch i {
		case 1:

			fmt.Println("Register by entering the below details:")
			RegisterUser()

		case 2:
			fmt.Println("View the user profile:")
			//fmt.Println(Mk)

		case 3:
			fmt.Println("Modify/edit the profile: ")

		case 4:
			fmt.Println("Transfer the amount: ")

		case 5:
			break L
		}

	}

}
