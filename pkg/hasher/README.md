This package contain functionality to work with passwords.

Currently, we have two functions to use HashPassword and CheckPasswordHash.

**HashPassword** provide functionality to create hash from password string and return password hash or error.
Example: `hash := HashPassword(password1)`

**CheckPasswordHash** provide functionality to check password hash and return bool true if password hash is equal to hash.
Example: `isHashValid := CheckPasswordHash(password1, hash1)`
