package gonetcheck

/*********************************************************************
 * Testing: gonetcheck - Go package to check general network health
 *
 * func: checkTcp
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
	"time"
)

func TestChecktcp(t *testing.T) {
	outputQueue := make(chan error)
	go checkTcp("www.google.com:80", outputQueue)

	select {
	case err := <-outputQueue:
		if err != nil {
			t.Error("expected TCP connection to www.google.com:80 to succeed:", err)
		}
	case <-time.After(20 * time.Second):
		t.Error("No response received after 20s")
	}
}
