#!/bin/bash

git diff HEAD^ --name-status | grep "^D" -v | sed 's/^.\t//g' | grep "\.go$" > changed_files

while read FILE; do
    go fmt $FILE
done < changed_files

rm changed_files
