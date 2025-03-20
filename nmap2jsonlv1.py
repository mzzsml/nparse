#!/usr/bin/env python3

import json as j
import argparse
import xml.etree.ElementTree as ET

json = {}

def parse(file):
    with open(file, 'r') as xmlfile:
        tree = ET.parse(xmlfile)
        root = tree.getroot()
        # Grab the nmap command that has been executed
        #json['command_line'] = root.attrib['args']

        for child in root:
            if child.tag == 'scaninfo':
                json['scan_type'] = child.attrib['type']
                json['num_services'] = child.attrib['numservices']
            elif child.tag == 'host':
                host = child
                child = None
                for child in host:
                    if child.tag == 'address':
                        json['addr'] = child.attrib['addr']
                    elif child.tag == 'hostnames':
                        hostnames = child
                        child = None
                        for child in hostnames:
                            json['hostname'] = child.attrib['name']
                    elif child.tag == 'ports':
                        ports = child
                        child = None
                        for child in ports:
                            if child.tag == 'port':
                                json['protocol'] = child.attrib['protocol']
                                json['port'] = child.attrib['portid']
                                port = child
                                child = None
                                for child in port:
                                    if child.tag == 'state':
                                        json['port_state'] = child.attrib['state']
                                        json['reason'] = child.attrib['reason']
                                    elif child.tag == 'service':
                                        json['service'] = child.attrib['name']
                                        #[(json[x] = child.attrib[x]) for x in child.attrib if (child.attrib == 'product' or child.attrib == 'version') ]
                                        for x in child.attrib:
                                            if x == 'product':
                                                json['product'] = child.attrib['product']
                                            elif x == 'version':
                                                json['product_version'] = child.attrib['version']
                                                                                    
                                        yield json
                                        #print(type(json))
                                    # with open('/tmp/TEST.jsonl', '+a') as jsonlfile:
                                    #     j.dump(json, jsonlfile)
 
if __name__ == '__main__':
    parser = argparse.ArgumentParser(
                    prog='nmap2jsonl',
                    description='Converts Nmap\'s XML output to JSONL')
    parser.add_argument('filename', help='XML file to convert')
    args = parser.parse_args()

    #print(parse(args.filename))