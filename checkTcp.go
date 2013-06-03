package gonetcheck

/*********************************************************************
 * gonetcheck - Go package to check general network health
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
	"net"
)

// Attempt TCP connection and put the error status onto
// the outQueue channel
func checkTcp(addr string, outQueue chan error) {
	debugLog(DBG_MEDIUM, "checkTcp: ", addr)

	tcpAddr, tcpAddrErr := net.ResolveTCPAddr("tcp", addr)
	if tcpAddrErr != nil {
		debugLog(DBG_VERBOSE, "net.ResolveTCPAddr failed:", tcpAddrErr)
		outQueue <- tcpAddrErr
	}
	conn, connErr := net.DialTCP("tcp", nil, tcpAddr)
	if connErr != nil {
		debugLog(DBG_VERBOSE, "net.DialTCP failed:", connErr)
		outQueue <- connErr
	}
	defer conn.Close()

	// Otherwise it was a success - no error
	debugLog(DBG_VERBOSE, "checkTcp SUCCESS:", addr)
	outQueue <- nil
}
