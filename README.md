# English <-> Gopher - Go Service translator :page_with_curl:

## Overview

This is an HTTP  English <-> Gopher - Go Service translator.

##Docker 

# Via github workflow:
[![Docker Image CI](https://github.com/richardsplit/go_translator_gopher/actions/workflows/docker_app_docker-image.yml/badge.svg?branch=docker_app)](https://github.com/richardsplit/go_translator_gopher/actions/workflows/docker_app_docker-image.yml)

```
# Build docker image from docker_app_docker-image.yml

- docker build . --tag ghrc.io/go_translation_gopher/docker_app/gopher:latest   

# Create container locally 
- docker run -it -p 8081:8081 ghrc.io/go_translation_gopher/docker_app/gopher:late   

```



# Locally 
-Docker should be installed 

Docker file is in go_translation_gopher main dir
run commands:
```
docker build -t <nameofimage> -f Dockerfile .    
#container 
 docker run -it -p localport:containerport  <nameofcreatedimage> 
   example:
 docker run -it -p 8081:8081  <nameofcreatedimage>   
```
