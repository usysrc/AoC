#!/bin/sh
# Remove all files named input from all subdirectories
# Usage: ./remove-inputs.sh --delete
# The --delete argument is optional

# if argument --delete is passed, then remove the files
if [ "$1" = "--delete" ]; then
    find . -name input -exec git rm {} \;
# else echo a message
else
    # do a dry run
    find . -name input -exec echo {} \;
    echo "To delete the files, pass --delete as an argument"
fi