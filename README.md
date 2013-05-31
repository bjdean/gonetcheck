github.com/bjdean/gonetcheck
============================

go package for checking general network health

Synopsis
--------

	import "github.com/bjdean/gonetcheck"
	
	can_access_internet, err := gonetcheck.CheckInternetAccess()
	
	if can_access_internet {
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

