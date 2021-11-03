This project have been created for learning purposes.

It will contain all hometask works, and also some functionality related to testing theoretical materials.

Project contain hasher package. It contain two functions - **HashPassword** and **CheckPasswordHash**.

This package contain functionality to work with passwords.

Currently, we have two functions to use **HashPassword** and **CheckPasswordHash**.

HashPassword provide functionality to create hash from password string and return password hash or error. 

Example: `hash := HashPassword(password1)`

CheckPasswordHash provide functionality to check password hash and return bool true if password hash is equal to hash. 

Example: `isHashValid := CheckPasswordHash(password1, hash1)`

The entry point placed into **main.go** file.

So now you can run this project with next command - go run main.go
