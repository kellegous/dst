#!/usr/bin/env python3

import datetime
import dateutil.tz
import sys


def main():
    tz = dateutil.tz.gettz('America/New_York')

    t = datetime.datetime(2021, 3, 14, 2, 0, 0, 0, tz)
    print("{}\t{}".format(
        t.isoformat(),
        t.astimezone(datetime.timezone.utc).isoformat()))

    t = datetime.datetime(2021, 11, 7, 1, 0, 0, 0, tz)
    print("{}\t{}".format(
        t.isoformat(),
        t.astimezone(datetime.timezone.utc).isoformat()))


if __name__ == '__main__':
    sys.exit(main())
