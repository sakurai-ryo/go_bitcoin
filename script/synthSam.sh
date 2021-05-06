#!/bin/bash

set -eu

cd $(dirname $0)

cd ../lambda
GOOS=linux go build main.go
zip function.zip main

cd ../
cdk synth >template.yml
rm lambda/function.zip
rm lambda/main
