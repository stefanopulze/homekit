# Homekit Server

### config.yaml
```yaml
homekit:
  name: 'bridge_name'
  pin: '00102003'
  storagePath: './db'

server:
  port: 4000

entities:
  - name: 'Sensor 01'
    serialNumber: '123'
    firmware: '0.0.1b'
    type: 'relay'
    options:
      mode: 'http'
      ip: '10.0.0.2'
      realyId: 0
      reloadTimeout: 1000

  - name: 'Shelly1'
    type: 'relay'
    options:
      mode: 'http'
      ip: '10.0.0.3'
      realyId: 0
```