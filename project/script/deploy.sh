#!/bin/zsh

rm -rf opt dist/*
yarn build
mkdir -p opt
cp -R dist opt/dist
cp .env.production opt/.env
# requires https://github.com/FiloSottile/homebrew-musl-cross
CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" -o ./opt/boilerplateprj main.go
#GOOS=linux GOARCH=amd64 go build -o ./opt/boilerplateprj main.go
tree opt
ssh vultr 'rm -rf ~/boilerplateprj/*'
tar -czvf boilerplateprj.tar.gz opt/*
scp ./boilerplateprj.tar.gz boilerplateprj@vultr:/boilerplateprj/boilerplateprj/boilerplateprj.tar.gz
ssh vultr 'cd boilerplateprj && tar -xzvf boilerplateprj.tar.gz && rm boilerplateprj.tar.gz'
rm boilerplateprj.tar.gz
rm -rf opt
ssh vultr 'mkdir ~/boilerplateprj/opt/database'
scp ./database/db.db boilerplateprj@vultr:~/boilerplateprj/opt/database/db.db
ssh vultr 'chmod 666 ~/boilerplateprj/opt/database/db.db && docker restart boilerplateprj && tree ~/boilerplateprj'