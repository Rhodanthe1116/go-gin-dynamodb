# Go Gin Boilerplate

## Deploy

```bash
export USERNAME=
docker build -t $USERNAME/covid-tracker-user .
docker push $USERNAME/covid-tracker-user
docker build -t $USERNAME/covid-tracker-store .
docker push $USERNAME/covid-tracker-store
docker build -t $USERNAME/covid-tracker-record .
docker push $USERNAME/covid-tracker-record
```

> A starter project with Golang, Gin and DynamoDB

[![Build Status][travis-image]][travis-url]
[![codebeat badge](https://codebeat.co/badges/ed248580-942c-4ffc-919f-d3681d28a799)](https://codebeat.co/projects/github-com-vsouza-go-gin-boilerplate)
[![Go Version][go-image]][go-url]
[![License][license-image]][license-url]
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate?ref=badge_shield)


Golang Gin boilerplate with DynamoDB resource. Supports multiple configuration environments.

![](header.jpg)

This project use a [DynamoDB base docker image](https://github.com/vsouza/docker-dynamoDB-local).

Setup DynamoDB dependency:  `docker run -d -p 8080:8080 vsouza/dynamo-local --port 8080`

### Boilerplate structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Installation

```sh
make deps
```

## Usage example

`curl -s ${TARGET_URL}/health`

## Development setup

Running DynamoDB on Docker Image:

check this project: [vsouza/docker-dynamoDB-local](https://github.com/vsouza/docker-dynamoDB-local)

## Release History

* 0.0.1
    * Configuration by environment, Auth and Log middlewares, User entity.

## Meta

Vinicius Souza – [@iamvsouza](https://twitter.com/iamvsouza) – hi@vsouza.com

Distributed under the MIT license. See [License](https://vsouza.mit-license.org) for more information.

[https://github.com/vsouza](https://github.com/vsouza)

[go-image]: https://img.shields.io/badge/Go--version-1.9-blue.svg
[go-url]: https://golang.org/doc/go1.9
[travis-image]: https://travis-ci.org/vsouza/go-gin-boilerplate.svg?branch=master
[travis-url]: https://travis-ci.org/vsouza/go-gin-boilerplate
[license-image]: https://img.shields.io/badge/License-MIT-blue.svg
[license-url]: https://vsouza.mit-license.org

## APIs

```
TARGET_URL='http://3.224.156.234.sslip.io'
USER_PHONE='0912345678'
USER_PASS='878787'
STORE_PHONE='0987654321'
STORE_PASS='87878787'
STORE_NAME='giver'
STORE_ADDR='taipei'
curl -s -X POST ${TARGET_URL}/auth/user/signup -d "{\"phone\": \"${USER_PHONE}\", \"password\": \"${USER_PASS}\"}"
user_token=$(curl -s -X POST ${TARGET_URL}/auth/user/login -d "{\"phone\": \"${USER_PHONE}\", \"password\": \"${USER_PASS}\"}" | tee /dev/tty | sed 's/{"token":"\([^"]*\)"}/\1/')
curl -s -X POST ${TARGET_URL}/auth/store/signup -d "{\"phone\": \"${STORE_PHONE}\", \"password\": \"${STORE_PASS}\", \"name\": \"${STORE_NAME}\", \"address\": \"${STORE_ADDR}\"}"
store_token=$(curl -s -X POST ${TARGET_URL}/auth/store/login -d "{\"phone\": \"${STORE_PHONE}\", \"password\": \"${STORE_PASS}\"}" | tee /dev/tty | sed 's/{"token":"\([^"]*\)"}/\1/')
qrcode=$(curl -s -X GET ${TARGET_URL}/auth/store/profile -H "Authorization: Bearer ${store_token}" | tee /dev/tty | grep -o '"qrcode":"[^"]*"' | sed 's/"qrcode":"\([^"]*\)"/\1/')
curl -s -X POST ${TARGET_URL}/records -H "Authorization: Bearer ${user_token}" -d "{\"store_id\": \"${qrcode}\"}"
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate?ref=badge_large)

