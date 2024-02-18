#!/bin/bash -e

if [ $# != 1 ]; then
	echo "project_name is required

$ create-project.sh [project_name]
"
	exit 1
fi

project=$1
cd "$(dirname $0)/../"
template_dir="$(pwd)"

cd "${template_dir}/../"

mkdir -p "${project}"

cd "${project}"
mkdir -p "api/${project}-api/oapi"
mkdir -p "cmd/${project}-api/"
mkdir -p "internal/${project}-api"
mkdir -p scripts

cp "$template_dir/api/appname/oapi/.gitignore" "api/${project}-api/oapi/.gitignore"
test -e "api/${project}-api/spec.yaml" || cp "$template_dir/api/appname/spec.yaml" "api/${project}-api/"

if ! test -e "api/${project}-api/Makefile"; then
	cp "$template_dir/Makefile" .
	sed -i "s/@appname/${project}-api/g" Makefile
fi

test -e go.mod || go mod init \
	"$(git remote -v | grep push | awk '{print $2}' | sed -e 's/git@//g' | sed -e 's/:/\//g' | sed -e 's/.git//g')/project_examples/${project}"

go mod tidy
