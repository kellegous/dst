#!/usr/bin/env python3
import datetime
import dateutil.tz
import sys

def to_rfc3339(dt):
    s = dt.isoformat()
    if s.endswith('+00:00'):
        return s[:-6] + 'Z'
    return s

def main():
    tz = dateutil.tz.gettz('America/New_York')

    t = datetime.datetime(2021, 3, 14, 2, 0, 0, 0, tz)
    print(to_rfc3339(t.astimezone(datetime.timezone.utc)))
    print("{}\t{}".format(
        to_rfc3339(t),
        to_rfc3339(t.astimezone(datetime.timezone.utc))))

    t = datetime.datetime(2021, 11, 7, 1, 0, 0, 0, tz)
    print("{}\t{}".format(
        to_rfc3339(t),
        to_rfc3339(t.astimezone(datetime.timezone.utc))))


if __name__ == '__main__':
    sys.exit(main())
