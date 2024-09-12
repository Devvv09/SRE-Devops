#!/bin/bash



jq -r '.[].url' image.json | while read -r url
do
	wget "$url"
done


