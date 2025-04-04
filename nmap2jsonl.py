#!/usr/bin/env python3

import sys
import argparse

from models.scan import Scan
from models.file import File

from lib.parser import parse
from lib.tojsonl import tojsonl

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
        xmlfile = File(args.filename)
        s = Scan()
        s = parse(xmlfile=xmlfile.path)
        [print(x) for x in tojsonl(s)]
    else:
        f = File(path=args.output)
        # If file exists and it's size is bigger than 0, delete it's content.
        if f.exists() and f.size != 0:
            f.deletecontent()
        f.writecontent(parse(xmlfile=args.filename))