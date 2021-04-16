#!/usr/bin/python

import socket
import sys

if len(sys.argv) != 2
    print("Usage: vrfy.py <Enter txt file containing usernames>")
    sys.exit(0)
    
 with open(sys.argv[1]) as f:
        hosts = f.read().splitlines()
        
#New socekt for connection to SMTP server
for i in hosts:
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
#Connect to SMTP server
    connect = s.connect(("SMTP server IP",25))
#Receiving banner
    banner = s.recv(1024)
    print(banner)
#Send the command VRFY to have the server verify an email address
    s.send("VRFY" + sys.argv[1] + "\r\n")
    result = s.recv(1024)
    print(result)
#Close connection.
    s.close()
