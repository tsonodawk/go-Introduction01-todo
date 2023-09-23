package main

import (
	"fmt"
	"introduction/todoApp/app/controllers"
	"introduction/todoApp/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("test3@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)

}
