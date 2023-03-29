# golang_api

this is small rest api with golang just for learning purpose and practice, trying to write everything without framework to get things more clear

now we have only 2 routes to get all accounts in db table & /account/{id} route will check if method is GET will return recored from db by ID and if DELETE then will call handleDeleteAccount method to delete recorde from db
```
router.HandleFunc("/all", makeHTTPHandleFunc(s.handleAccount))
router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleGetAccountByID)))
```
