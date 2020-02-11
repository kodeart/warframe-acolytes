package internal

import (
	"testing"
	"time"
)

func TestStrMillisToTime(t *testing.T) {
	_, err := StrMillisToTime("1581282911976")

	if err != nil {
		t.Error("Failed to convert string milliseconds to time.TIme")
	}

	ms, _ := StrMillisToTime("")

	if !ms.Equal(time.Time{}) {
		t.Error("Invalid milliseconds string should resolve to time.Time{}")
	}
}

func TestIntMillisToTime(t *testing.T) {
	ms := IntMillisToTime(1581282911976)

	if ms.Unix() != 1581282911 {
		t.Errorf("Converted milliseconds are not correct, the result was %v", ms.Unix())
	}

	expect := "2020-02-09T22:15:11+01:00"
	dt := ms.Format(time.RFC3339)
	if dt != expect {
		t.Errorf("Converted milliseconds are not correct, expected result is %s", expect)
	}
}
