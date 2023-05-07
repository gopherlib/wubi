#!/usr/bin/env bash

CURRENT_DIR=$(cd "$(dirname "$0")";pwd)

go run ${CURRENT_DIR}/../main.go -dict ${CURRENT_DIR}/dict_86.txt -pkg dict -var Dict86 \
    -pattern '^(?<code>[a-zA-Z]+)|\s+~*(?<char>[一-龥〇])(?![一-龥〇])' > "${CURRENT_DIR}/../../../dict/dict_86.go"
