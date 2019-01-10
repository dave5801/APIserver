# API Server Project for Upbound

This is an API Endpoint written in the Go programming language. The purpose of this project is a code assessment for [Upbound](https://upbound.io/). It was a great opportunity to learn the Go programming language.

The purpose of this project is take in a group of .yml files, check each individual file for App configuration data. Each .yml file must have properly formated fields in order to be considered valid. The validation process checks for empty fields, invalid email and urls.

Once validated the following endpoint is displayed.

```
[{"Title":"Valid App 1","Version":"0.0.1","Maintainers":[{"email":"firstmaintainer@hotmail.com","name":"firstmaintainer app1"},{"email":"secondmaintainer@gmail.com","name":"secondmaintainer app1"}],"Company":"Random Inc.","Website":"https://website.com","Source":"https://github.com/random/repo","License":"Apache-2.0","Description":"### Interesting Title\nSome application content, and description"},{"Title":"Valid App 2","Version":"1.0.1","Maintainers":[{"email":"apptwo@hotmail.com","name":"AppTwo Maintainer"}],"Company":"Upbound Inc.","Website":"https://upbound.io","Source":"https://github.com/upbound/repo","License":"Apache-2.0","Description":"### Why app 2 is the best\nBecause it simply is..."}]
```


### Prerequisites

The Go Programming language, 

The following go packages:

gopkg.in/yaml.v2
github.com/gorilla/mux

### Installing

Open Terminal

Create Project Directory.

```
cd <Directory-Name>
```

Install Go. See [installation instructions](https://golang.org/doc/install)

Clone my Repository.

```
git clone https://github.com/dave5801/APIserver.git
```

Install dependencies:

```
go get gopkg.in/yaml.v2
go get -u github.com/gorilla/mux
```

### Usage

Open Terminal.
Navigate to your project directory.

use command:
```
go run serverMain.go
```

to see api endpoint, use the following url:

```
http://localhost:8000/configs
```

### Version
This is a rough draft

### License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

### Author
David Franklin
