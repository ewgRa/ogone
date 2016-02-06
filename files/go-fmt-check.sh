#!/bin/bash

git diff HEAD^ --name-status | grep "^D" -v | sed 's/^.\t//g' | grep "\.go$" > changed_files

while read FILE; do
    echo $FILE
done < changed_files

rm changed_files

git diff --exit-code --quiet

exit $?