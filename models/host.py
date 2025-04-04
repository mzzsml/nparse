class Host:
    """The Host object.
    
    Stores information about a host, such as its IP address,
    hostname and a list of its ports.
    """

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