from models.scan import Scan
from models.target import Target

def tojsonl(scan) -> list:
    """Parse the `Scan` object into JSONL format.
    
    Returns a list of single-line JSON strings, a line for each port.
    """

    _ = []
    for x in scan.hosts:
        for j in x.ports:
            t = Target()
            t.scan_type = scan.scan_type
            t.protocol = scan.protocol
            t.addr = x.addr
            t.hostname = x.hostname
            t.port = j.port
            t.port_state = j.port_state
            t.reason = j.reason
            t.service = j.service
            t.product = j.product
            t.version = j.version
            t.extrainfo = j.extrainfo
            _.append(t.asdict())
    return _