# English <-> Gopher - Go Service translator :page_with_curl:

~

[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/richardsplit/translator_go/blob/main/LICENSE)

## Overview

This is an HTTP  English <-> Gopher - Go Service translator.

Gophers are friendly creatures but it’s not that easy to communicate with them. They have their own language and they don’t understand English. You can use this service to translate to their language.

## Setup and run locally

```sh
go mod vendor
go run main.go -port=<input_port>
# go run main.go
```
type the port which the application should use


## Docker

```sh
Go to branch "docker_app" 
https://github.com/richardsplit/go_translator_gopher/tree/docker_app
```

## Run tests

### Prerequisites

* [ginkgo](http://onsi.github.io/ginkgo/)

```sh
go generate ./...
ginkgo -v
```

## Sample calls locally

```sh
curl -X POST http://127.0.0.1:8081/word -d '{"english-word": "subliminal"}'
```
Example output:
{"gopher-word":"ubliminalsogo"}


```sh
curl -X POST http://127.0.0.1:8081/sentence -d '{"english-sentence": "Ever have that feeling where you are not sure if you are awake or dreaming?"}'
```

Example output:
{"gopher-sentence":"gEver avehogo atthogo eelingfogo herewogo ouyogo gare otnogo uresogo gif ouyogo gare gawake gor reamingdogo?"}


```sh
curl http://127.0.0.1:your_port/history
```

## Sample call through POSTMAN
GET/history
```sh
http://127.0.0.1:your_port/history
```
Example Output:
{
    "history": [
        {
            "gopher": "ophergogo"
        },
        {
            "neo": "eonogo"
        }
    ]
}


POST/word
```sh
http://127.0.0.1:your_port/word

Body raw example:
{"english-word":"neo"}
```
Example Output:
{
    "gopher-word": "eonogo"
}


POST/sentence
```sh
http://127.0.0.1:your_port/sentence


Body raw example:
{"english-sentence": "english-sentence":"We're not here because we're free. We're here because we're not free."}
```

Example Output:
{
    "gopher-sentence": "ereWogo otnogo erehogo ecausebogo erewogo reefogo. ereWogo erehogo ecausebogo erewogo otnogo reefogo."
}


