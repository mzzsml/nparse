# Nparse

Nmap Parser.

`nparse` converts Nmap's XML output to JSON and JSONL.

## Usage

```txt
$ nparse nmap-out.xml
```

By default Nparse will print it's output in JSON format.<br>
To have a JSON formatted output, specify the `-jsonl` flag.

The JSONL output will be in the followin format:

```json
{"addr": "192.168.1.33", "hostname": "host.local", "port": "138", "protocol": "tcp", "port_state": "open", "reason": "syn-ack", "service": "netbios-ssn", "product": "Samba smbd", "version": "3", "extrainfo": null, "scan_type": "syn"}
```
