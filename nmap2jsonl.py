#!/usr/bin/env python3

import os
import sys
import json
import argparse

from models.scan import Scan

from parser.parser import parse

def is_file_already_existing(file) -> bool:
    return True if os.path.exists(file) else False

if __name__ == '__main__':
    parser = argparse.ArgumentParser(
                    prog='nmap2jsonl',
                    description='Converts Nmap\'s XML output to JSONL')
    parser.add_argument('filename', help='XML file to convert')
    parser.add_argument('-o', '--output',
                        default=sys.stdout,
                        required=False,
                        help='Output file.')
    args = parser.parse_args()

    if args.output == sys.stdout:
        with open(args.filename, mode='r') as f:
            s = Scan()
            s = parse(xmlfile=f)
            print(json.dumps(s))
    else:
        # If file exists and it's size is bigger than 0, delete it's content.
        if is_file_already_existing(args.output) and os.stat(args.output).st_size != 0:
            open(args.output, 'w').close()
        with open(args.output, '+a') as f:
            json.dump(parse(args.filename, f))