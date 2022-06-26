#!/usr/bin/env python3

import math
import os
import sys


def lat_to_y(lat, h):
    lat_rad = lat * math.pi/180
    merc_n = math.log(math.tan((math.pi/4)+(lat_rad/2)))
    return (h/2) - (h*merc_n)/(2*math.pi)


def main():
    for lat in [-89.99, 0, 89.99]:
        print("lat = {}, y = {}".format(lat, lat_to_y(lat, 1)))


if __name__ == '__main__':
    sys.exit(main())
