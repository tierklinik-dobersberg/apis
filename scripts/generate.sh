#!/bin/bash

echo "Generating protobuf files ..."
buf generate proto

pkgjson=$(cat package.json | jq '.exports = {"./package.json": {"default": "./package.json"}}')

echo "Creating es+d.ts barrel files"
for pkg in ./gen/es/tkd/**/*; do 
    # remove old barrel files
    rm $pkg/index.js $pkg/index.d.ts || true

    entrypoint="${pkg/gen\/es\/tkd\//}"

    pkgjson=$(echo "$pkgjson" | jq ".exports += { \"${entrypoint}\": { \"types\": \"${pkg}/index.d.ts\", \"default\": \"${pkg}/index.js\" } }" )

    echo "${pkg}" 
    # generate new index.js
    for file in $pkg/*.js ; do 
        echo "export * from './$(basename $file)'" >> $pkg/index.js
    done

    # generate new index.d.ts
    for file in $pkg/*.d.ts ; do 
        echo "export * from './$(basename ${file/.d.ts//})'"  >> $pkg/index.d.ts
    done
done

echo "Creating package.json exports"

echo "$pkgjson" > ./package.json
