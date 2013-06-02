package gonetcheck

/*********************************************************************
 * gonetcheck - Go package to check general network health
 *
 * Debugging
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
	"log"
)

// Debug flag - higher for more output
const (
	DBG_OFF = iota
	DBG_QUIET
	DBG_MEDIUM
	DBG_VERBOSE
)

var DEBUG int = DBG_OFF
var DEBUG_DEFAULT int = DEBUG

// Debug logging
func debug_log(debug_level int, a ...interface{}) {
	if DEBUG >= debug_level {
		log.Println(a...)
	}
}
