package gonetcheck

/*********************************************************************
 * Testing: gonetcheck - Go package to check general network health
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
	"testing"
)

// XXX Make sure to reset DEBUG_DEFAULT after each test

func TestDebugLogOff(t *testing.T) {
	DEBUG = DBG_OFF
	debugLog(DBG_QUIET, "hello", "there")
	DEBUG = DEBUG_DEFAULT
}

func TestDebugLogQuiet(t *testing.T) {
	DEBUG = DBG_QUIET
	debugLog(DBG_QUIET, "hello", "there")
	DEBUG = DEBUG_DEFAULT
}

func TestDebugLogMedium(t *testing.T) {
	DEBUG = DBG_MEDIUM
	debugLog(DBG_QUIET, "hello", "there")
	DEBUG = DEBUG_DEFAULT
}

func TestDebugLogVerbose(t *testing.T) {
	DEBUG = DBG_VERBOSE
	debugLog(DBG_QUIET, "hello", "there")
	DEBUG = DEBUG_DEFAULT
}

func ExampleTestDebugLogOff() {
	DEBUG = DBG_OFF
	debugLog(DBG_QUIET, "hello", "there")
	DEBUG = DEBUG_DEFAULT
	// Output:
}

func ExampleDebugLogQuiet() {
	DEBUG = DBG_QUIET
	debugLog(DBG_QUIET, "hello", "there")
	DEBUG = DEBUG_DEFAULT
	// Output:
}

func ExampleDebugLogMedium() {
	DEBUG = DBG_MEDIUM
	debugLog(DBG_QUIET, "hello", "there")
	DEBUG = DEBUG_DEFAULT
	// Output:
}

func ExampleDebugLogVerbose() {
	DEBUG = DBG_VERBOSE
	debugLog(DBG_QUIET, "hello", "there")
	DEBUG = DEBUG_DEFAULT
	// Output:
}
