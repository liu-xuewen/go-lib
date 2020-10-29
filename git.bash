#!/usr/bin/env bash


git config  user.name "liuxuewen"

git config  user.email "liiuxuewen@gmail.com"

git config --list

git status

git add .

if [ $1 ]
then
  git commit -m "$1"
else
  git commit -m "fix"
fi


git push
