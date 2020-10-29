#!/usr/bin/env bash

git config  user.name "liuxuewen"

git config  user.email "liiuxuewen@gmail.com"

git config --list

git status

git add .


if [[ ! -z $1 ]]; then
  git commit -m "$1"
  echo $1

else
  git commit -m "fix"
  echo "fix"
fi


git push
