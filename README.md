# golang_api

this is small rest api with golang just for learning purpose and practice, trying to write everything without framework to get things more clear

now we have only 2 routes to get all accounts in db table & get account by id
```
router.HandleFunc("/all", makeHTTPHandleFunc(s.handleAccount))
router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID))
```