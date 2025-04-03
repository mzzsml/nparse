class Host:
    def __init__(self, addr=None, hostname=None, ports=[]):
        self.addr = addr # host > address
        self.hostname = hostname # hostnames > hostname.name
        self.ports = ports # List of Port objects.
    
    def asdict(self):
        return {
            'addr': self.addr,
            'hostname': self.hostname,
            'ports': self.ports
        }