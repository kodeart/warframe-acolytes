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
	value := int64(1581282911976)
	expect := int64(1581282911)
	ms := IntMillisToTime(value)

	if ms.Unix() != expect {
		t.Errorf("Converted milliseconds are not correct, expected %d, got %d", expect, ms.Unix())
	}
}
