import socket
from datetime import datetime
net = input("Enter Subnet: ")
net1 = net.split('.')
a = '.'

net2 = net1[0] + a + net1[1] + a + net1[2] + a
st1 = int(input("First Number: "))
en1 = int(input("Last Number: "))
en1 = en1 + 1
t1 = datetime.now()

def scan(addr):
   s = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
   socket.setdefaulttimeout(1)
   result = s.connect_ex((addr,135))
   if result == 0:
      return 1
   else :
      return 0

def run1():
   for i in range(st1,en1):
      addr = net2 + str(i)
      if (scan(addr)):
         print (addr , "is up")
         
run1()
t2 = datetime.now()
total = t2 - t1
print ("Completed in: " , total)
