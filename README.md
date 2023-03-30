# golang_api

this is small rest api with golang just for who want to learn how things works in golang as web development , trying to write everything without framework to get things more clear

Routes endpoints 

```
/login
/account [Post] to create account
/account [Get] to get all accounts
/account/{id} [Get] to get account by id
/transfer [Get] to transfer Balance ``` not completed yet xD ```
```

main enirty point to start the app
```
package main

import (
	"log"
)

func main() {

	dbConn, err := PostgresConn()

	if err != nil {
		log.Fatal(err)
	}

	if err := dbConn.Init(); err != nil {
		log.Fatal(err)
	}

	server := newAPIServer(":1992", dbConn)
	server.Run()
}
)
```

Check other files to understand how code flow works, it's very simple 
this is a new line xD
