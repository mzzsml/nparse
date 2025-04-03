from models.target import Target

def tojsonl(scan) -> list:
    """Parse the `Scan` object into JSONL format.
    
    Returns a list of single-line JSON strings, a line for each port.
    """

    _ = []
    for x in scan.hosts:
        for j in x.ports:
            t = Target(x.addr,
                       x.hostname,
                       j.port,
                       scan.protocol,
                       j.port_state,
                       j.reason,
                       j.service,
                       j.product,
                       j.version,
                       scan.scan_type,
                       j.extrainfo)
            _.append(t.asdict())
    return _