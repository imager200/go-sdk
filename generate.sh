#!/bin/bash



curl -o imager200.json https://www.imager200.io/imager200.json 

oapi-codegen --config=oapi-codegen.yaml imager200.json 

rm imager200.json

