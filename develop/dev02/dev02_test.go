package main

import "testing"

func TestCorrectString(t *testing.T) {
	arr, err := unpack("a4bc2d5e")
	if err != nil {
		t.Error(err.Error())
	}
	if arr != "aaaabccddddde" {
		t.Error("not correct string")
	}

	arr, err = unpack("av2c")
	if err != nil {
		t.Error(err.Error())
	}
	if arr != "avvc" {
		t.Error("not correct string")
	}

	arr, err = unpack("abcd")
	if err != nil {
		t.Error(err.Error())
	}
	if arr != "abcd" {
		t.Error("not correct string")
	}

	arr, err = unpack("")
	if err != nil {
		t.Error(err.Error())
	}
	if arr != "" {
		t.Error("not correct string")
	}
}

func TestNotCorrectStringWithZero(t *testing.T) {
	_, err := unpack("a4bc2d5e0")
	if err == nil {
		t.Error("this is not correct string")
	}
	_, err = unpack("a4bc0d5e")
	if err == nil {
		t.Error("this is not correct string")
	}
}

func TestNotCorrectStringWithNumberOnStart(t *testing.T) {
	_, err := unpack("3a4bc2d5e0")
	if err == nil {
		t.Error("this is not correct string")
	}
	_, err = unpack("5")
	if err == nil {
		t.Error("this is not correct string")
	}
	_, err = unpack("5a")
	if err == nil {
		t.Error("this is not correct string")
	}
}
