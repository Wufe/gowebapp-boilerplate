#!/bin/zsh

files=(
	'package.json'
	'main.go'
	'glide.yaml'
	'go.mod'
	'infrastructure/database.go'
	'presentation/home-handler.go'
	'services/home-service.go'
	'project/script/deploy.sh'
)

for file in $files; do
	sed -i -e s/$1/$2/g $file
done