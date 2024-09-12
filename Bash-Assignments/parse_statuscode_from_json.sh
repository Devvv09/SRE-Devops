#!/bin/bash

jq -r '.services[].endpoint' url.json | while read -r endpoint
do
	 curl --head www.google.com | sed -n '1p' | awk '{print $2}' >> statuslogs.txt 
done

	
