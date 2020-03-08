#!/usr/bin/env bash

if [[ "$#" -ne 1 || !("$1" =~ ^[0-9]+$) ]]; then
  echo "Usage: leap.sh <year>"
  exit 1
fi

year=$1
(( !(year % 4) && ( (year % 100) || !(year % 400) ) )) && echo "true" || echo "false"
