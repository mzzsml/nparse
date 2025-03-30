class Target:
    def __init__(self):
        self.addr = None
        self.hostname = None
        self.port = None
        self.protocol = None
        self.port_state = None
        self.reason = None
        self.service = None
        self.product = None
        self.version = None
        self.scan_type = None
        self.extrainfo = None
    
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