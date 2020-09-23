import requests
import re

regex = r"(\w{32,}?)\s+\d+?\s+(.*)$"


def read_checksums(url):
    req = requests.get(url)
    matches = re.finditer(regex, req.text, re.MULTILINE)
    for matchNum, match in enumerate(matches):
        print "Checksum: {checksum}\tFilename: {filename}".format(checksum=match.groups()[0], filename=match.groups()[1])


read_checksums('http://ftp.debian.org/debian/dists/jessie/Release')
