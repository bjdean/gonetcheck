package gonetcheck

/*********************************************************************
 * Testing: gonetcheck - Go package to check general network health
 *
 * func: SetUrlList
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
	"testing"
)

func TestSetUrlListValid(t *testing.T) {
	var newUrls = []string{
		"http://example.com/",
		"http://example.net/",
	}
	oldUrls := GetUrlList()

	errList := SetUrlList(newUrls)

	// Reset url list (avoid side-effects on other tests)
	SetUrlList(oldUrls)

	if errList != nil {
		t.Error("Valid URL list had failures:", errList)
	}
}

func TestSetUrlListInvalid(t *testing.T) {
	var newUrls = []string{
		"http://example.com/",
		"http://example.net/",
		"%xx%xx",
		"%zz%zz",
	}
	oldUrls := GetUrlList()

	errList := SetUrlList(newUrls)

	// Reset url list (avoid side-effects on other tests)
	SetUrlList(oldUrls)

	if len(errList) != 2 {
		t.Error("Invalid URLs not detected:", newUrls)
	}
}

func ExampleSetUrlList() {
	var newUrls = []string{
		"http://example.com/",
		"http://example.net/",
	}
	errList := SetUrlList(newUrls)
	if errList != nil {
		// Handle problem with URLs
	}
	// Output:
}
