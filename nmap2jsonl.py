#!/usr/bin/env python3

import os
import sys
import json as j
import argparse
import xml.etree.ElementTree as ET

def is_file_already_existing(file) -> bool:
    return True if os.path.exists(file) else False

def parse_port(xml_element):
    port_details = {
        "port_state": None,
        "reason": None,
        "service": None,
        "product": None,
        "version": None
    }
    for child in xml_element:
        match child.tag:
            case 'state':
                port_details['port_state'] = child.attrib['state']
                port_details['reason'] = child.attrib['reason']
            case 'service':
                for attr in child.attrib:
                    match attr:
                        case 'name':
                            port_details['service'] = child.attrib['name']
                        # We do not always have these information
                        case 'product':
                            port_details['product'] = child.attrib['product']
                        case 'version':
                            port_details['version'] = child.attrib['version']
                        # NOTE: Maybe it's actually better to have a separate column for better compatibility with endoflife.date apis.
                        case 'extrainfo':
                            if port_details['version'] == None:
                                port_details['version'] = ''
                            port_details['version'] += f" ({child.attrib['extrainfo']})"
                yield port_details

def parse_ports(xml_element):
    ports = {}
    for child in xml_element:
        match child.tag:
            case 'port':
                ports['protocol'] = child.attrib['protocol']
                ports['port'] = child.attrib['portid']
                for x in parse_port(child):
                    ports.update(x)
                    yield ports

def parse_hostnames(xml_element):
    asd = {}
    for child in xml_element:
        match child.tag:
            case 'hostname':
                asd['hostname'] = child.attrib['name']
    return asd

def parse_host(xml_element):
    hosts = {}
    for child in xml_element:
        match child.tag:
            case 'address':
                hosts['addr'] = child.attrib['addr']
            case 'hostnames':
                hosts.update(parse_hostnames(child))
            case 'ports':
                for x in parse_ports(child):
                    hosts.update(x)
                    yield hosts

def parse(file):
    host_details = {}
    with open(file, 'r') as xmlfile:
        tree = ET.parse(xmlfile)
        root = tree.getroot()
        # Grab the nmap command that has been executed
        #json['command_line'] = root.attrib['args']

        for child in root:
            match child.tag:
                case 'scaninfo':
                    host_details['scan_type'] = child.attrib['type']
                case 'host':
                    for x in parse_host(child):
                        host_details.update(x)
                        yield host_details

if __name__ == '__main__':
    parser = argparse.ArgumentParser(
                    prog='nmap2jsonl',
                    description='Converts Nmap\'s XML output to JSONL')
    parser.add_argument('filename', help='XML file to convert')
    parser.add_argument('-o', '--output',
                        #default='./test/TEST.jsonl',
                        default=sys.stdout,
                        required=False,
                        help='Output file.')
    args = parser.parse_args()

    #[print(x) for x in parse(args.filename)]
    if args.output == sys.stdout:
        [print(x) for x in parse(args.filename)]
    else:
        # If file exists and it's size is bigger than 0, delete it's content.
        if is_file_already_existing(args.output) and os.stat(args.output).st_size != 0:
            open(args.output, 'w').close()
        with open(args.output, '+a') as jsonfile:
            [j.dump(x, jsonfile) for x in parse(args.filename)]