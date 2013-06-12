github.com/bjdean/gonetcheck TODO
=================================

* Extend CheckInternetAccess to do some other types of network checks (eg. DNS lookups, pings, other protcols)
  Wed May 29 14:13:24 EST 2013


DONE
====

Sat Jun  1 00:49:49 EST 2013
----------------------------
* Add ability to set the test url list (package variable, or consider adding a struct to allow multiple different test urls to run within a single program.

Wed Jun 12 23:56:55 EST 2013
----------------------------
* Add proxy support (for URL tests) by using http.ProxyFromEnvironment

*Actually* I didn't make any changes for this, because the net/http DefaultClient uses the DefaultTransport, which in turn checks for the proxy - as per the docs:

	var DefaultTransport RoundTripper = &Transport{Proxy: ProxyFromEnvironment}
