package now

import (
	"testing"
	"time"
)

var format string

func init() {
	format = "2006-01-02 15:04:05.999999999"
}

func TestBeginningOf(t *testing.T) {
	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.Now().Location())

	if New(n).BeginningOfMinute().Format(format) != "2013-11-18 17:51:00" {
		t.Errorf("BeginningOfMinute")
	}

	if New(n).BeginningOfHour().Format(format) != "2013-11-18 17:00:00" {
		t.Errorf("BeginningOfHour")
	}

	if New(n).BeginningOfDay().Format(format) != "2013-11-18 00:00:00" {
		t.Errorf("BeginningOfDay")
	}

	FirstDayMonday = true
	if New(n).BeginningOfWeek().Format(format) != "2013-11-18 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDayMonday")
	}

	FirstDayMonday = false
	if New(n).BeginningOfWeek().Format(format) != "2013-11-17 00:00:00" {
		t.Errorf("BeginningOfWeek")
	}

	if New(n).BeginningOfMonth().Format(format) != "2013-11-01 00:00:00" {
		t.Errorf("BeginningOfMonth")
	}

	if New(n).BeginningOfYear().Format(format) != "2013-01-01 00:00:00" {
		t.Errorf("BeginningOfYear")
	}
}

func TestEndOf(t *testing.T) {
	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.Now().Location())

	if New(n).EndOfMinute().Format(format) != "2013-11-18 17:51:59.999999999" {
		t.Errorf("EndOfMinute")
	}

	if New(n).EndOfHour().Format(format) != "2013-11-18 17:59:59.999999999" {
		t.Errorf("EndOfHour")
	}

	if New(n).EndOfDay().Format(format) != "2013-11-18 23:59:59.999999999" {
		t.Errorf("EndOfDay")
	}

	FirstDayMonday = true
	if New(n).EndOfWeek().Format(format) != "2013-11-24 23:59:59.999999999" {
		t.Errorf("EndOfWeek, FirstDayMonday")
	}

	FirstDayMonday = false
	if New(n).EndOfWeek().Format(format) != "2013-11-23 23:59:59.999999999" {
		t.Errorf("EndOfWeek")
	}

	if New(n).EndOfMonth().Format(format) != "2013-11-30 23:59:59.999999999" {
		t.Errorf("EndOfMonth")
	}

	if New(n).EndOfYear().Format(format) != "2013-12-31 23:59:59.999999999" {
		t.Errorf("EndOfYear")
	}

	n1 := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.Now().Location())
	if New(n1).EndOfMonth().Format(format) != "2013-02-28 23:59:59.999999999" {
		t.Errorf("EndOfMonth for 2013/02")
	}

	n2 := time.Date(1900, 02, 18, 17, 51, 49, 123456789, time.Now().Location())
	if New(n2).EndOfMonth().Format(format) != "1900-02-28 23:59:59.999999999" {
		t.Errorf("EndOfMonth")
	}
}
