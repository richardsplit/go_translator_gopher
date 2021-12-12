# English <-> Gopher - Go Service translator :page_with_curl:

[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/richardsplit/translator_go/blob/main/LICENSE)

## Overview

This is an HTTP  English <-> Gopher - Go Service translator.

##Docker 
-
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
