package main

import (
	"app/internal/user"
	"fmt"
)

func main() {

	fmt.Println("Send/receive signals ...")

	fmt.Println("STATUS: ", user.INACTIVE)

	u := user.UserSvc{}
	fmt.Println(u.ID)
	//userwf.UserWorkflow(con)

}
