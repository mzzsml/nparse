class Scan:
    """The Scan object.
    
    It stores the general information of a Nmap scan, such as scan type,
    protol and a list of hosts scanned.
    """

    def __init__(self, scan_type=None, protocol=None, hosts=[]):
        #self.cmdline = None # nmaprun
        self.scan_type = scan_type # scaninfo.type
        self.protocol = protocol # scaninfo.protocol
        self.hosts = hosts # List containing Host objects
    
    def asdict(self):
        return {
            'scan_type': self.scan_type,
            'protocol': self.protocol,
            'hosts': self.hosts
        }