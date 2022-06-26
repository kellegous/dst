#!/usr/bin/env python3

import argparse
from ast import arg
import json
import os
import sys


def ReadAll(src):
    with open(src, 'r') as r:
        return json.load(r)


def FindZone(data, id):
    for feature in data.get('features'):
        properties = feature.get('properties')
        if properties['tzid'] == id:
            return feature
    return None


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--id',
                        default='America/New_York',
                        help="the zone to extract")
    parser.add_argument('--geojson-file',
                        default='combined-with-oceans.json',
                        help="the geojson file read")
    args = parser.parse_args()
    data = ReadAll(args.geojson_file)

    # for feature in data.get('features'):
    #     properties = feature.get('properties')
    #     print(properties['tzid'])
    feature = FindZone(data, args.id)
    print(json.dumps(feature, indent=2))


if __name__ == '__main__':
    sys.exit(main())
