package main

import (
	"reflect"
	"testing"
)

func TestPlaceSpace(t *testing.T) {
	// var msg []SearchPlacesResponse
	msg := searchPlaces(" ")
	if len(msg) != 0 {
		t.Errorf("Expected empty list for empty search string, got msg %v", msg)
	}
}
func TestPlaceValid(t *testing.T) {
	// var msg []SearchPlacesResponse
	msg := searchPlaces("Gainesville")
	if !(len(msg) >= 1) {
		t.Errorf("Expected list of nearby places, got msg %v", msg)
	}
}

func TestPlaceClosestMatch(t *testing.T) {
	// var msg []SearchPlacesResponse
	msg1 := searchPlaces("New Yrk")
	msg2 := searchPlaces("New York")
	if !(reflect.DeepEqual(msg1, msg2)) {
		t.Errorf("Expected search string to return place list for closest match, got msg %v", msg1)
	}
}
func TestPlaceNoMatch(t *testing.T) {
	// var msg []SearchPlacesResponse
	msg1 := searchPlaces("New Yr")
	if len(msg1) != 0 {
		t.Errorf("Expected empty list for invalid search string, got msg %v", msg1)
	}
}
func TestPlaceEmpty(t *testing.T) {
	// var msg []SearchPlacesResponse
	msg1 := searchPlaces("")
	var msg2 []SearchPlacesResponse
	if !(reflect.DeepEqual(msg1, msg2)) {
		t.Errorf("Expected empty result for empty search string, got %v", msg1)
	}
}
