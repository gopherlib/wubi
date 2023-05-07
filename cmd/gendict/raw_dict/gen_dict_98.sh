#!/usr/bin/env bash

CURRENT_DIR=$(cd "$(dirname "$0")";pwd)

go run ${CURRENT_DIR}/../main.go -dict ${CURRENT_DIR}/dict_98.txt -pkg dict -var Dict98 \
    -pattern '^(?<char>[一-龥〇])\s+(?<code>[a-zA-Z]+)' > "${CURRENT_DIR}/../../../dict/dict_98.go"
