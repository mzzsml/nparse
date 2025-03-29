class Port:
    def __init__(self):
        self.port = None # ports > port.portid
        self.port_state = None
        self.reason = None
        self.service = None # port > service.name
        self.product = None # port > service.product
        self.version = None # port > service.version
        self.extrainfo = None
    
    def asdict(self):
        return {
            'port': self.port,
            'port_state': self.port_state,
            'reason': self.reason,
            'service': self.service,
            'product': self.product,
            'version': self.version,
            'extrainfo': self.extrainfo
        }