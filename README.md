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
go run main.go -port=<input_port>
# go run main.go
```
type the port which the application should use

##Docker 
-
-Docker should be installed 

Docker file is in go_translation_gopher main dir
run commands:
```
docker build -t <nameofimage> -f Dockerfile .    
#container 
 docker run -it -p 8081:8081  <nameofcreatedimage>   
```
