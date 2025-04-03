import xml.etree.ElementTree as ET

from models.scan import Scan
from models.host import Host
from models.port import Port

def parse_ports(xml_port) -> Port:
    """Parse the `ports` node of the Nmap XML output."""

    port = None
    port_state = None
    reason = None
    service = None
    product = None
    version = None
    extrainfo = None

    _ = []
    for child in xml_port:
        p = Port()
        if child.tag == 'port':
            port = child.attrib['portid']
            # Other attributes are a lever lower.
            for x in child:
                if x.tag == 'state':
                    port_state = x.attrib['state']
                    reason = x.attrib['reason']
                elif x.tag == 'service':
                    service = x.attrib['name']
                    # We do not always have `product` and `version`,
                    # so we need to check manually.
                    for j in x.attrib:
                        if j == 'product':
                            product = x.attrib['product']
                        if j == 'version':
                            version = x.attrib['version']
                        if j == 'extrainfo':
                            extrainfo = x.attrib['extrainfo']
                    _.append(Port(port=port, port_state=port_state,
                                  reason=reason, service=service,
                                  product=product, version=version,
                                  extrainfo=extrainfo))
    return _

def parse_host(xml_host) -> Host:
    """Parse the `host` node of the Nmap XML output.
    
    Returns a Host object."""

    addr = ""
    hostname = ""
    ports = []

    for child in xml_host:
        if child.tag == 'address':
            addr = child.attrib['addr']
        elif child.tag == 'hostnames':
            # The hostname is a level deeper, so we need to iterate.
            for x in child:
               if x.tag == 'hostname':
                hostname = x.attrib['name'] 
        elif child.tag == 'ports':
            ports = parse_ports(child)
    return Host(addr=addr, hostname=hostname, ports=ports)

def parse(xmlfile) -> Scan:
    """Parse the Nmap XML output.
    
    Returns a Scan object."""

    tree = ET.parse(xmlfile)
    root = tree.getroot()

    scan_type = ""
    protocol = ""
    hosts = []

    for child in root:
        #scan.cmdline = root.attrib['args']
        if child.tag == 'scaninfo':
            scan_type = child.attrib['type']
            protocol = child.attrib['protocol']
        elif child.tag == 'host':
            hosts.append(parse_host(child))
    return Scan(scan_type=scan_type, protocol=protocol, hosts=hosts)