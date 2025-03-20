#!/usr/bin/env python3

import json as j
import argparse
import xml.etree.ElementTree as ET

json = {}
ports = []

#def file_exists(file) -> bool:
#    
#    return False

def parse_port(xml_element):
    for child in xml_element:
        match child.tag:
            case 'state':
                json['port_state'] = child.attrib['state']
                json['reason'] = child.attrib['reason']
            case 'service':
                json['service'] = child.attrib['name']
                for attr in child.attrib:
                    # We do not always have these information
                    if attr == 'product':
                        json['product'] = child.attrib['product']
                    if attr == 'version':
                        json['product_version'] = child.attrib['version']

#############################################################

def parse_ports(xml_element):
    asd = {}
    for child in xml_element:
        match child.tag:
            case 'port':
                if child.attrib['portid'] not in ports:
                    asd['protocol'] = child.attrib['protocol']
                    asd['port'] = child.attrib['portid']
                    break
                #break
    return asd
                #parse_port(child)

def parse_hostnames(xml_element):
    asd = {}
    for child in xml_element:
        match child.tag:
            case 'hostname':
                asd['hostname'] = child.attrib['name']
    return asd

def parse_host(xml_element):
    asd = {}
    for child in xml_element:
        match child.tag:
            case 'address':
                asd['addr'] = child.attrib['addr']
            case 'hostnames':
                asd.update(parse_hostnames(child))
                #return asd
            case 'ports':
                asd.update(parse_ports(child))
                return asd
    

def parse(file):
    jsonres = {}
    with open(file, 'r') as xmlfile:
        tree = ET.parse(xmlfile)
        root = tree.getroot()
        # Grab the nmap command that has been executed
        #json['command_line'] = root.attrib['args']

        for child in root:
            match child.tag:
                case 'scaninfo':
                    jsonres['scan_type'] = child.attrib['type']
                case 'host':
                    #yield jsonres.update(parse_host(child))
                    yield parse_host(child)

#def parse(file):
#    with open(file, 'r') as xmlfile:
#        tree = ET.parse(xmlfile)
#        root = tree.getroot()
#        # Grab the nmap command that has been executed
#        #json['command_line'] = root.attrib['args']
#
#        for child in root:
#            if child.tag == 'scaninfo':
#                json['scan_type'] = child.attrib['type']
#                json['num_services'] = child.attrib['numservices']
#            elif child.tag == 'host':
#                host = child
#                child = None
#                for child in host:
#                    if child.tag == 'address':
#                        json['addr'] = child.attrib['addr']
#                    elif child.tag == 'hostnames':
#                        hostnames = child
#                        child = None
#                        for child in hostnames:
#                            json['hostname'] = child.attrib['name']
#                    elif child.tag == 'ports':
#                        ports = child
#                        child = None
#                        for child in ports:
#                            if child.tag == 'port':
#                                json['protocol'] = child.attrib['protocol']
#                                json['port'] = child.attrib['portid']
#                                port = child
#                                child = None
#                                for child in port:
#                                    if child.tag == 'state':
#                                        json['port_state'] = child.attrib['state']
#                                        json['reason'] = child.attrib['reason']
#                                    elif child.tag == 'service':
#                                        json['service'] = child.attrib['name']
#                                        #[(json[x] = child.attrib[x]) for x in child.attrib if (child.attrib == 'product' or child.attrib == 'version') ]
#                                        for x in child.attrib:
#                                            if x == 'product':
#                                                json['product'] = child.attrib['product']
#                                            elif x == 'version':
#                                                json['product_version'] = child.attrib['version']
#                                                                                    
#                                        yield json
#                                        #print(type(json))
#                                    # with open('/tmp/TEST.jsonl', '+a') as jsonlfile:
#                                    #     j.dump(json, jsonlfile)
 
if __name__ == '__main__':
    parser = argparse.ArgumentParser(
                    prog='nmap2jsonl',
                    description='Converts Nmap\'s XML output to JSONL')
    parser.add_argument('filename', help='XML file to convert')
    args = parser.parse_args()

    #print(parse(args.filename))
    print([x for x in parse(args.filename)])
    #print([j.dump(x) for x in parse(args.filename)])
    
    #with open('TEST.jsonl', '+a') as jsonfile:
    #    [j.dump(x, jsonfile) for x in parse(args.filename)]