package radical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadical(t *testing.T) {
	t.Logf("Adjust to maximum by: %d", Max32)
	const orig = 90.000001 // test precisions too
	d2r := Degrees2Radians(orig)
	t.Logf("%f degrees ==> %f radians", orig, d2r)
	back := Radians2Degrees(d2r)
	t.Logf("%f radians ==> %f degrees ", d2r, back)
	adjusted := Degrees2Adjusted(orig)
	t.Logf("%f degrees == %d adjusted", orig, adjusted)
	deg := adjusted.ToDegrees()
	t.Logf("%f == %f", deg, orig)
	radical()
}

func TestKeypoint(t *testing.T) {
	const (
		lat = 45.0
		lon = 110.0
	)
	kp := MakeKeyPoint(lat, lon)
	kLat, kLon := kp.ToDegrees()
	assert.Equal(t, lat, kLat)
	assert.Equal(t, lon, kLon)
}
