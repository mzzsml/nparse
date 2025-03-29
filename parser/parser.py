import xml.etree.ElementTree as ET

from models.scan import Scan
from models.host import Host
from models.port import Port

def parse_ports(xml_port):
    _ = []
    for child in xml_port:
        p = Port()
        if child.tag == 'port':
            p.port = child.attrib['portid']
            # Other attributes are a lever lower.
            for x in child:
                if x.tag == 'state':
                    p.port_state = x.attrib['state']
                    p.reason = x.attrib['reason']
                elif x.tag == 'service':
                    p.service = x.attrib['name']
                    # We do not always have `product` and `version`,
                    # so we need to check manually.
                    for j in x.attrib:
                        if j == 'product':
                            p.product = x.attrib['product']
                        if j == 'version':
                            p.version = x.attrib['version']
                        if j == 'extrainfo':
                            p.extrainfo = x.attrib['extrainfo']
                    _.append(p.asdict())
    return _

def parse_host(xml_host):
    h = Host()
    for child in xml_host:
        if child.tag == 'address':
            h.addr = child.attrib['addr']
        elif child.tag == 'hostnames':
            # The hostname is one level deeper, so we need to iterate.
            for x in child:
               if x.tag == 'hostname':
                h.hostname = x.attrib['name'] 
        elif child.tag == 'ports':
            h.ports = parse_ports(child)
    return h.asdict()

def parse(xmlfile):
    tree = ET.parse(xmlfile)
    root = tree.getroot()

    scan = Scan()
    for child in root:
        #scan.cmdline = root.attrib['args']
        if child.tag == 'scaninfo':
            scan.scan_type = child.attrib['type']
            scan.protocol = child.attrib['protocol']
        elif child.tag == 'host':
            scan.hosts.append(parse_host(child))
    return scan.asdict()