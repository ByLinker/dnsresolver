# DNSResolver
This repo contains a pretty simple recursive resolver made using [meikg/dns](github.com/miekg/dns) library.

**NOTE: This is made for educational purposes, don't use it for production at all**


## How to run it:
- Install Go
- Clone this repo, and from your terminal `cd` into the directory
- Then run the following:
```bash
go run *.go -host github.com
```

The output should be like this:
```
We found answers github.com
[github.com.	60	IN	A	140.82.121.4]
```


