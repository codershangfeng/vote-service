# vote-service

A simple demo service for illustrating how service mesh (Istio) works in k8s.

A RESTful API is represented by [Swagger 2.0 (aka OpenAPI 2.0)](https://swagger.io/) ([go-swagger](https://github.com/go-swagger/go-swagger))

## How to Build Locally

### Clone the repo
```zsh
$ gh repo clone codershangfeng/vote-service
```

### Install Go
Make sure you are using go1.15. To check, run:
```zsh
$ go version
```
Or, you could use `gvm` to install and alter between versions:
```zsh
$ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
$ gvm listall
$ gvm install go1.15
```

To check your installed Go version, run:
```zsh
$ gvm list
```

To switch to a specific version, run:
```
$ gvm use go1.15 --default
```

### Build app
```zsh
$ cd app
$ make build
```
The binary will be generated into `app/bin` folder.


## How to Deploy