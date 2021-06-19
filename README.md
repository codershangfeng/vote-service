# vote-service

A simple demo service for illustrating how service mesh (Istio) works in k8s.

A RESTful API is represented by [Swagger 2.0 (aka OpenAPI 2.0)](https://swagger.io/) ([go-swagger](https://github.com/go-swagger/go-swagger))

## How to Build Locally

### Clone the repo
```zsh
$ gh repo clone codershangfeng/vote-service
```

### Install Go
Make sure you are using go1.16. To check, run:
```zsh
$ go version
```
Or, you could use `gvm` to install and alter between versions:
```zsh
$ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
$ gvm listall
$ gvm install go1.16
```

To check your installed Go version, run:
```zsh
$ gvm list
```

To switch to a specific version, run:
```
$ gvm use go1.16 --default
```

### Install Go-Swagger

```zsh
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
```
Install swagger tool, [more details](https://goswagger.io/install.html)

### Validate Swagger Yaml

To validate the swagger file (`./api/swagger.yaml`), run:
```bash
make validate
```

### Generate Swagger

To generate server with its models and operations, run:
```bash
make gen
```

### Build app
```zsh
$ make build
```
The binary will be generated into `./bin` folder.

## Hot to Test

### Run All Test
```zsh
$ make test
```

## How to Deploy to Local Demo Kubernetes Cluster

- `docker-compose build`: build an image `vote-service` with **latest** version.
- Refer to demo k8s resources repo, and apply the resouces files in the folder `vote-service`.
