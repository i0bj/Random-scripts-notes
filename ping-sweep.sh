#!/bin/bash
#Enter the first 3 octets of the range you want to scan

echo "ping sweep of "$1.0 "has started"
for ip in {1..255}
  do
    ping -c1 -w1 > /dev/null 2> /dev/null
    if [ $? -ne 0 ]
    then
      echo "${ip} is up!"
    else
      echo "${ip} is down"
    fi
done
