import sys
#Simply prints hex charaters which can be used to find bad characters when developing your exploit.
for x in range(0, 256):
  sys.stdout.write("\\x" + '{:02x}'.format(x))
