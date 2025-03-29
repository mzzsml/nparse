class Scan:
    def __init__(self):
        #self.cmdline = None # nmaprun
        self.scan_type = None # scaninfo.type
        self.protocol = None # scaninfo.protocol
        self.hosts = [] # SARA LISTA DI OGGETTI DI TIPO HOST
    
    def asdict(self):
        return {
            'scan_type': self.scan_type,
            'protocol': self.protocol,
            'hosts': self.hosts
        }