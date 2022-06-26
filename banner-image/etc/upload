#!/usr/bin/env python3

import argparse
import os
import subprocess
import sys


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        'image',
        help="the file to upload")
    parser.add_argument(
        '--tag',
        default='00f1',
        help='the tag to use at the destination')
    args = parser.parse_args()
    return subprocess.call([
        'gsutil',
        'cp',
        args.image,
        'gs://fs.kellegous.com/s/b/{}.png'.format(args.tag),
    ])


if __name__ == '__main__':
    sys.exit(main())
