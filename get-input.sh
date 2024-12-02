#!/bin/sh
# Get the input for the current day for Advent of Code

# Extract the day from the current directory
day=$(basename $(pwd))
# Remove any leading zeros from the day
day=$(echo $day | sed 's/^0*//')
# check if day is a valid day
if [ $day -lt 1 ] || [ $day -gt 25 ]; then
    echo "Invalid day: $day"
    exit 1
fi

# Extract the year from the parent directory
year=$(basename $(dirname $(pwd)))
# check if year is a valid year
if [ $year -lt 2015 ] || [ $year -gt 2024 ]; then
    echo "Invalid year: $year"
    exit 1
fi

# Construct the URL to get the input
url="https://adventofcode.com/$year/day/$day/input"
echo "Getting input for day $day of year $year from $url"

# Fetch the input and save it to a file named input
curl -H "Cookie: session=$AOC_SESSION_COOKIE" -s $url > input