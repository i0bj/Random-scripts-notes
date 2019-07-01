#! /usr/bin/python
import urllib2
from collections import Counter

cnt = Counter()
webSite = urllib2.open(<Enter Site to scrape>).read() #Enter the website you want to harvest, this opens the site and reads the words found on the page.
cnt.update(webSite.split()) #This will split all of the words grabbed from the site and turn it into a dictionary
print(cnt.most_common(n)) #n is the number of common words you want to scrape from the site.
for x in cnt.keys(): #For loop will iterate through harvested words on Site
    if "<" in x: del cnt[k]; continue
    if ">" in x: del cnt[k]; continue
    if "=" in x: del cnt[k]; continue
print(cnt.most_common(n)) #n is the common number of words you want to harvest.
