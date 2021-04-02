#Oneliner for foward dns lookup. Enter common hostnames in any text file. ex: mail, owa, vpn, ftp
for ip in $(cat <file.txt>); do host $ip.<domain>;done
