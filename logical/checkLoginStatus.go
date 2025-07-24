package logical

import "fmt"

func CheckLoginStatus() {

	username := "admin"
	password := "1234"

	isValid := username == "admin" && password == "1234"
	fmt.Println("Login success:", isValid)

}
