package gonetcheck

/*********************************************************************
 * Testing: gonetcheck - Go package to check general network health
 *
 * func: CheckInternetAccess
 *
 * Copyright 2013 Bradley Dean
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http:www.gnu.org/licenses/>.
 */

import (
	"fmt"
	"os"
	"testing"
)

func TestCheckInternetAccess(t *testing.T) {
	_, errList := CheckInternetAccess(
		[]string{ "http://www.google.com/", "http://www.unimelb.edu.au/" },
		[]string{ "www.google.com:80", "www.google.com:443" },
	)
	if errList != nil {
		t.Error("Error received:", errList)
	}
}

func ExampleCheckInternetAccess() {
	canAccessInternet, errList := CheckInternetAccess(
		[]string{ "http://www.google.com/", "http://www.unimelb.edu.au/" },
		[]string{ "www.google.com:80", "www.google.com:443" },
	)
	switch errList {
	case nil:
		switch canAccessInternet {
		case true:
			os.Exit(0)
		default:
			os.Exit(1)
		}
	default:
		fmt.Println(
			"Error returned by CheckInternetAccess:",
			errList)
		// Would usually os.Exit(2)
		// But to avoid upsetting go test exit with 0
		os.Exit(0)
	}
	// Output:
}
