# English <-> Gopher - Go Service translator :page_with_curl:

~

[![Build Status](https://travis-ci.com/adria-stef/gopher-translator-service.svg?branch=main)](https://travis-ci.com/adria-stef/gopher-translator-service)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/richardsplit/translator_go/blob/main/LICENSE)

## Overview

This is an HTTP  English <-> Gopher - Go Service translator.

Gophers are friendly creatures but it’s not that easy to communicate with them. They have their own language and they don’t understand English. You can use this service to translate to their language.

## Setup and run locally

```sh
go mod vendor
go run main.go
```
type the port which the application should use

## Run tests

### Prerequisites

* [ginkgo](http://onsi.github.io/ginkgo/)

```sh
go generate ./...
ginkgo -v
```

## Sample calls locally

```sh
curl -X POST http://127.0.0.1:your_port/word -d '{"english-word": "penguin"}'
```

```sh
curl -X POST http://127.0.0.1:your_port/sentence -d '{"english-sentence": "You either die a hero, or you live long enough to see yourself become the villain."}'
```

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
{"english-sentence": "english-sentence":"We're not here because we're free. We're here because we're not free. There's no escaping reason. No denying purpose. Because as we both know without purpose, we would not exist."}
```

Example Output:
{
    "gopher-sentence": "ereWogo otnogo erehogo ecausebogo erewogo reefogo. ereWogo erehogo ecausebogo erewogo otnogo reefogo. eresThogo onogo gescaping easonrogo. oNogo enyingdogo urposepogo. ecauseBogo gas ewogo othbogo nowkogo ithoutwogo urposepogo, ewogo ouldwogo otnogo gexist."
}


