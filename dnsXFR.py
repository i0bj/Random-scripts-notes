#!/usr/bin/python
#zone tranfer in python
import dns.query
import dns.zone
z = dns.zone.from_xfr(dns.query.xfr('IP of dns server', 'domain'))
names = z.nodes.keys()
names.sort()
    for n in names:
        print(z[n].to_text(n)
