# Worker

Status: WIP

- [ ] handle load parameter
- [ ] handle memory parameter
- [ ] add metrics endpoint

## Overview
This program is used to generate a base load per cpu core and a base memory usage for the entire system.

## Build locally
```bash
$ make amd64
$ ./build/bin/worker-linux-amd64
Usage of ./build/bin/worker-linux-amd64:
  -avg-load uint
        Average load target in % (default 10)
  -mem-usage uint
        Memory allocation target in MB (default 512)
```
