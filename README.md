github.com/bjdean/gonetcheck
============================

Go package for checking general network health

The purpose of this package is to provide a simple yes/no check as to
whether the general internet is directly accessible from the calling
system.

This could be used as a pre- or post-check when trying to connect to
a different system (helping to determine that system is down, or if
it is a local networking problem).

See also: http://bjdean.id.au/wiki/Programming_Notes/GoLang

A simple wrapper program that uses this package can be found
here: https://github.com/bjdean/can_access_internet

Synopsis
--------

	import (
		"github.com/bjdean/gonetcheck"
		"time"
	)
	
	canAccessInternet, err := gonetcheck.CheckInternetAccess(
		time.Duration(10 * time.Second),
		string[]{ "http://example.com", "http://example.net", "http://example.org/" },
		string[]{ "example.org:80", "example.net:443", "example.org:22" } )
	
	if canAccessInternet {
		// Do something requiring internet access
	}


License
-------

Copyright 2013 Bradley Dean

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http:www.gnu.org/licenses/>.

