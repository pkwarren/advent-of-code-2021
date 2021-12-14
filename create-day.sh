#!/bin/bash

set -x

if [ "${#}" -ne 1 ]; then
    echo "Usage: ${0} <day>" >&2
    exit 1
fi

curdir=`dirname "${0}"`
curdir=`cd "${curdir}" && pwd`
day=`printf "%02d" ${1}`
mkdir -p day${day}/testdata
touch day${day}/testdata/input{01,02}
touch day${day}/question.txt
touch day${day}/day${day}.go
touch day${day}/day${day}_test.go
