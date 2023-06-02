#!/bin/sh

source $( dirname "${BASH_SOURCE[0]}" )/../.env
export $(cut -d= -f1 $( dirname "${BASH_SOURCE[0]}" )/../.env)
go run $( dirname "${BASH_SOURCE[0]}" )/../application.go