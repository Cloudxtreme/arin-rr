# Unofficial ARIN Route Registry Tool (arin-rr)
Tool to update the ARIN Route Registry. Templates used are based on:

https://www.arin.net/resources/routing/templates.html

## Installation

### From Source
```shell
go get -u github.com/kkirsche/arin-rr
```

### Example

```shell
~/g/s/g/k/arin-rr git:master ❯❯❯ arin-rr route -r "1.2.3.4/24" -a 1234 -d "This is a description" -n "notifyNetworking@company.com" -g "changedByMe@company.com 20160406" -m "maintainer@company.com"
To: rr@arin.net
From:
Subject: route

route: 1.2.3.4/24
descr: This is a description
origin: AS1234
notify: notifyNetworking@company.com
mnt-by: maintainer@company.com
changed: changedByMe@company.com 20160406
source: ARIN
```
