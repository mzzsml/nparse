# Nmap2jsonl

`Nmap2jsonl` converts Nmap's XML output to JSONL.

No external libraries required.

## Usage

`nmap2jsonl` takes the Nmap XML output:

```txt
$ python3 ./nmap2jsonl.py nmap-out.xml
```

Use the `-o` to write the content to a file:

```txt
$ python3 ./nmap2jsonl.py nmap-out.xml -o nmap-out.jsonl
```

The output will be in the followin format:

```json
{"addr": "192.168.1.33", "hostname": "host.local", "port": "138", "protocol": "tcp", "port_state": "open", "reason": "syn-ack", "service": "netbios-ssn", "product": "Samba smbd", "version": "3", "extrainfo": null, "scan_type": "syn"}
```

### Converting the output in regular JSON

To convert the output in regular JSON the following command can be used:

```txt
$ cat nmap-out.jsonl | jq -scr '[.[]]' > nmap-out.json
```

This will create an array of JSONs.