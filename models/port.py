class Port:
    """The Port object.
    
    Stores information about the state of a port and the
    service running on it.
    """

    def __init__(self, port=None, port_state=None, reason=None, service=None,
                 product=None, version=None, extrainfo=None):
        self.port = port # ports > port.portid
        self.port_state = port_state
        self.reason = reason
        self.service = service # port > service.name
        self.product = product # port > service.product
        self.version = version # port > service.version
        self.extrainfo = extrainfo
    
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