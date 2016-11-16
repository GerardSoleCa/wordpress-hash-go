# Wordpress Hasher for Golang

[![Build Status](https://travis-ci.org/GerardSoleCa/wordpress-hash-go.svg?branch=master)](https://travis-ci.org/GerardSoleCa/wordpress-hash-go)

Implementation of Wordpress hashing system for Golang

## Installation

	go get github.com/GerardSoleCa/wordpress-hash-go
	
## Usage
```go
    import wphash
    
    func main(){
        password := "thisisapassword"
        hash := wphash.HashPassword(password)
        checked := wphash.CheckPassword(password, hash)
    }
```
## Tests
	
	go test

## License

[MIT](https://github.com/GerardSoleCa/wordpress-hash-go/blob/master/LICENSE)
