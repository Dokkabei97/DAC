# DAC (Data Auto Cleaner)

## Description
This is realTime monitoring Linux Server `/data` directory and clean up old files.

## policy

```yaml
config:
  monitoring:
    alert-usage: 80
    platform: telegram
    api-key: api-key
  
  clean-policy:
    usage: 85
    path-fileformat:
      - "/data/example/logs:*.log"
      - "/data/dump/*:*"
    time: "now-5"
    auto: true
```