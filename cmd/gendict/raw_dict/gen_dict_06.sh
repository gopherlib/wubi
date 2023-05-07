#!/usr/bin/env bash

CURRENT_DIR=$(cd "$(dirname "$0")";pwd)

go run ${CURRENT_DIR}/../main.go -dict ${CURRENT_DIR}/dict_06.txt -pkg dict -var Dict06 \
    -pattern '^\d+\t(?<code>[a-zA-Z]+)\t(?<char>[一-龥〇])\t\d+' > "${CURRENT_DIR}/../../../dict/dict_06.go"
