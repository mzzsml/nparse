class Scan:
    def __init__(self, scan_type=None, protocol=None, hosts=[]):
        #self.cmdline = None # nmaprun
        self.scan_type = scan_type # scaninfo.type
        self.protocol = protocol # scaninfo.protocol
        self.hosts = hosts # List containing Port objects
    
    def asdict(self):
        return {
            'scan_type': self.scan_type,
            'protocol': self.protocol,
            'hosts': self.hosts
        }