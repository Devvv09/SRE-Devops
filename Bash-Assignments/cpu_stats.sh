#!/bin/bash

set -x
file_name=cpu_stats.txt

iostat > $file_name
echo "Disk usage of $(pwd) is:" >> $file_name
du -h >> $file_name
echo "Free RAM:" >> $file_name
free --mega | sed -n '2p' | awk '{print $3}' >> $file_name

