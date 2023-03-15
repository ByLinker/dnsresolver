## First let's a lookup on madrasati.sa

```
dig ANY madrasati.sa @8.8.8.8
```
- SOA
- NS
- A


How did 8.8.8.8 find the answer? Let's do it manually...



### Trying to reach absher.sa

First the resolver has to figure out what's the IP for absher.sa (A record).

Where do we start?
The way it works the dns is operated as a herirechal tree. Root zone is operated by ican

```
dig NS .
```

Take one of the root servers do an A record query on it. It works why? because the IP for each root server is hardcoded in what's called a hints file maintained by iana. https://www.iana.org/domains/root/files

This is how resolvers bootstrap themselves.


```
dig NS sa @a.root-servers.net.
```

it'll give us a list of nameservers operating .sa TLD and additionally their IPs. Why? to avoid circular dependencies, because imagine if the nameservers are operated in the same zone. Which they're in this case c1.dns.sa. is operating in .sa, but we don't know how to find an ip for a .sa domain!
The easy fix is to allow for DNS servers to respond with additional data, called glue-records!
This way it removes the circullar dependency.
