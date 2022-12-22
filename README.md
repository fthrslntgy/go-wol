# go-wol

Wake on LAN magic packet generator for golang. Forked from https://github.com/sabhiram/go-wol and edited to be used as a library.

## Usage

To import:
```go
import(
    ...
    "github.com/fthrslntgy/go-wol/"
    ...
)
```

To use:
```go
macAddr := "a1:b2:c3:d4:e5:f6"
bcastAddr := "192.168.0.255"
bcastPort := 7
err := wol.Wake(macAddr, bcastAddr, bcastPort)
if err != nil {
    log.Println(err.Error())
}
```

Broadcast port can be send as `nil`. It is default `7`.
```go
macAddr := "aa:bb:1c:2d:ee:33"
bcastAddr := "192.168.0.255"
err := wol.Wake(macAddr, bcastAddr, nil)
if err != nil {
    log.Println(err.Error())
}
```
## Supported MAC address formats

MAC addresses' delimiter can be `:` or `-`. It can not be dot (`.`) or space. Letters are case insensitive. 

According to these rules, following MAC addresses are valid:
```
ab:cd:e0:b1:22
A1-2C-6E-4B-23-DD
aa-cc-23-d5-f1-e1
```

Following MAC addresses are not valid:
```
1-2-3-4-5-6
AA BB CC 11 22 33
1a.2b.3c.4d.5e.6f
```