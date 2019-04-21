#!/bin/zsh

rm -rf opt dist/*
yarn build
mkdir -p opt
cp -R dist opt/dist
cp .env.production opt/.env
# requires https://github.com/FiloSottile/homebrew-musl-cross
CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" -o ./opt/boilerplate main.go
#GOOS=linux GOARCH=amd64 go build -o ./opt/boilerplate main.go
tree opt
ssh vultr 'rm -rf ~/boilerplate/*'
tar -czvf boilerplate.tar.gz opt/*
scp ./boilerplate.tar.gz boilerplate@vultr:/boilerplate/boilerplate/boilerplate.tar.gz
ssh vultr 'cd boilerplate && tar -xzvf boilerplate.tar.gz && rm boilerplate.tar.gz'
rm boilerplate.tar.gz
rm -rf opt
ssh vultr 'mkdir ~/boilerplate/opt/database'
scp ./database/db.db boilerplate@vultr:~/boilerplate/opt/database/db.db
ssh vultr 'chmod 666 ~/boilerplate/opt/database/db.db && docker restart boilerplate && tree ~/boilerplate'