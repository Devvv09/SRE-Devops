#!/bin/bash

cat nginx-dummy.log | awk {'print $9'}  
