#!/bin/bash

echo "Generating protobuf files ..."
buf generate proto

pkgjson=$(cat package.json | jq '.exports = {"./package.json": {"default": "./package.json"}}')

echo -e "package proto\n" > ./proto/all.go

echo "Creating es+d.ts barrel files"
for pkg in ./gen/es/tkd/**/*; do 
    # remove old barrel files
    rm $pkg/index.js $pkg/index.d.ts || true

    entrypoint="${pkg/gen\/es\/tkd\//}"

    pkgjson=$(echo "$pkgjson" | jq ".exports += { \"${entrypoint}\": { \"types\": \"${pkg}/index.d.ts\", \"default\": \"${pkg}/index.js\", \"esm\": \"${pkg}/index.js\", \"esm2022\": \"${pkg}/index.js\" } }" )

    echo "${pkg}" 

    echo "import _ \"github.com/tierklinik-dobersberg/apis${pkg/.\/gen\/es/\/gen\/go}\"" >> ./proto/all.go

    # generate new index.js
    for file in $pkg/*.js ; do 
        echo "export * from './$(basename $file)'" >> $pkg/index.js
    done

    # generate new index.d.ts
    for file in $pkg/*.d.ts ; do 
        echo "export * from './$(basename ${file/.d.ts//})'"  >> $pkg/index.d.ts
    done
done

if [[ "$1" != "" ]]; then
    echo "Setting new version: $1"
    pkgjson=$(echo "$pkgjson" | jq ".version = \"$1\"")
fi

echo "Creating package.json exports"

echo "$pkgjson" > ./package.json

if [[ "$1" != "" && "$2" != "" ]]; then
    echo "Releasing  ...."
    git add .
    git commit -m "$2" --no-gpg
    git tag "v$1"
    git push && git push --tags
    npm publish
fi