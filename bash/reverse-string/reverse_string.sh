#!/usr/bin/env bash

for ((i=${#1}; i>=0; i--)); do
    printf "${1:$i:1}";
done
