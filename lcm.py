#!/usr/bin/python3
import math
import sys

limit = 100

for a in sys.argv:
    x = a.split('=')
    if len(x) == 2 and x[0] == 'count':
        limit = int(x[1])

def lcm(i, j):
    if i == 0 and j == 0:
        return 0
    return int(abs(i * j) / math.gcd(i, j))

for i in range(limit):
    for j in range(limit):
        print('lcm({},{})={}'.format(i, j, lcm(i, j)))
