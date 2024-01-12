#!/usr/bin/python3
import math
import sys

limit = 100

for a in sys.argv:
    x = a.split('=')
    if len(x) == 2 and x[0] == 'count':
        limit = int(x[1])

for i in range(limit):
    for j in range(limit):
        print('gcd({},{})={}'.format(i, j, math.gcd(i, j)))
