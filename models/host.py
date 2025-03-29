class Host:
    def __init__(self):
        self.addr = None # host > address
        self.hostname = None # hostnames > hostname.name
        self.ports = []
    
    def asdict(self):
        return {
            'addr': self.addr,
            'hostname': self.hostname,
            'ports': self.ports
        }