#!/bin/sh

# if argument --delete is passed, then remove the files
if [ "$1" = "--delete" ]; then
    find . -name input -exec git rm {} \;
# else echo a message
else
    # Remove all files named input from all subdirectories
    # do a dry run first
    find . -name input -exec echo {} \;
    echo "To delete the files, pass --delete as an argument"
fi