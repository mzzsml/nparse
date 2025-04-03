from models.scan import Scan

class Target:
    """The Target object.
    
    It stores the information of the target host in a standalone, complete JSON.
    One for every host and port."""

    def __init__(self, addr=None, hostname=None, port=None, protocol=None,
                 port_state=None, reason=None, service=None, product=None,
                 version=None, scan_type=None, extrainfo=None):
        self.addr = addr
        self.hostname = hostname
        self.port = port
        self.protocol = protocol
        self.port_state = port_state
        self.reason = reason
        self.service = service
        self.product = product
        self.version = version
        self.scan_type = scan_type
        self.extrainfo = extrainfo
    
    def asdict(self):
        return {
            'addr': self.addr,
            'hostname': self.hostname,
            'port': self.port,
            'protocol': self.protocol,
            'port_state': self.port_state,
            'reason': self.reason,
            'service': self.service,
            'product': self.product,
            'version': self.version,
            'extrainfo': self.extrainfo,
            'scan_type': self.scan_type
        }