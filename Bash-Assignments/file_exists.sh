#!/bin/bash

echo "What is the file name?"
read file_name

if [ -r ./$file_name ]
then
	echo "$file_name exists"
else
	echo "$file_name doesn't exists"
fi

