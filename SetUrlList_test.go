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

func TestSetUrlListValid (t *testing.T) {
	var new_urls = []string{
		"http://example.com/",
		"http://example.net/",
	}
	old_urls := GetUrlList()

	err_list := SetUrlList(new_urls)

	// Reset url list (avoid side-effects on other tests)
	SetUrlList(old_urls)

	if err_list != nil {
		t.Error("Valid URL list had failures:", err_list)
	}
}

func TestSetUrlListInvalid (t *testing.T) {
	var new_urls = []string{
		"http://example.com/",
		"http://example.net/",
		"%xx%xx",
		"%zz%zz",
	}
	old_urls := GetUrlList()

	err_list := SetUrlList(new_urls)

	// Reset url list (avoid side-effects on other tests)
	SetUrlList(old_urls)

	if len(err_list) != 2 {
		t.Error("Invalid URLs not detected:", new_urls)
	}
}

func ExampleSetUrlList () {
	var new_urls = []string{
		"http://example.com/",
		"http://example.net/",
	}
	err_list := SetUrlList(new_urls)
	if err_list != nil {
		// Handle problem with URLs
	}
	// Output:
}
