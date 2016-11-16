# Wordpress Hasher for Golang

[![Godoc reference](https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667)](https://godoc.org/github.com/GerardSoleCa/wordpress-hash-go)  [![Build Status](https://travis-ci.org/GerardSoleCa/wordpress-hash-go.svg?branch=master)](https://travis-ci.org/GerardSoleCa/wordpress-hash-go) [![Code Climate](https://codeclimate.com/github/GerardSoleCa/wordpress-hash-go/badges/gpa.svg)](https://codeclimate.com/github/GerardSoleCa/wordpress-hash-go)

Implementation of Wordpress hashing system for Golang

## Installation

	go get github.com/GerardSoleCa/wordpress-hash-go
	
## Usage
```go
    import "github.com/GerardSoleCa/wordpress-hash-go"
    
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
