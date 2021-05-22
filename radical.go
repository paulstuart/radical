package radical

import (
	"fmt"
	"log"
	"math"
)

//var RadAdjust =
/*
	const lat = 180.0
	rad := ra(lat)
	max := math.MaxInt32
	xlt := float32(max) / math.Pi
	adj := int32(rad * xlt)

*/

type RadInt int32
type RadFloat float32
type KeyPoint uint64

var (
	//rad = ra(180.0)
	RadIntMaxF = float64(math.MaxInt32) / math.Pi
	Max32F     = RadFloat(math.MaxInt32) / math.Pi
	Max32      = int32(float32(math.MaxInt32) / math.Pi)
)

func (kp *KeyPoint) FromDegrees(lat, lon float64) {
	if kp == nil {
		log.Println("nil keypoint")
	} else {
		iLat := int32(Degrees2Adjusted(lat))
		iLon := int32(Degrees2Adjusted(lon))
		*kp = KeyPoint(To(iLat, iLon))
	}
}

func (kp *KeyPoint) ToDegrees() (lat, lon float64) {
	if kp == nil {
		return
	}
	iLat, iLon := From(uint64(*kp))
	return Adjusted2Degrees(RadInt(iLat)), Adjusted2Degrees(RadInt(iLon))
}

func MakeKeyPoint(lat, lon float64) KeyPoint {
	var kp KeyPoint
	(&kp).FromDegrees(lat, lon)
	return kp
}

func (r *RadInt) FromDegrees(d float64) {
	if r != nil {
		*r = Degrees2Adjusted(d) //RadInt(Degrees2Radians(d) * Max32F)
	}
}

func (r *RadInt) ToDegrees() float64 {
	if r != nil {
		return Adjusted2Degrees(*r) //float64(*r) * 180.0 / RadIntMaxF
		//*r = Degrees2Radians(d)
	}
	return 0.0
}

func Degrees2Adjusted(d float64) RadInt {
	return RadInt(Degrees2Radians(d) * Max32F)
	//return (d * math.Pi) / 180.0
}

func Adjusted2Degrees(a RadInt) float64 {
	return Radians2Degrees(RadFloat(a) / Max32F)
	//return (d * math.Pi) / 180.0
}

func Degrees2Radians(d float64) RadFloat {
	return (RadFloat(d) * math.Pi) / 180.0
}

func Radians2Degrees(d RadFloat) float64 {
	return float64((RadFloat(d) * 180.0) / math.Pi)
}

func ra(f float32) float32 {
	return (f * math.Pi) / 180
}

func deg(f float32) float32 {
	return (f * 180) / math.Pi
}

func radical() {
	//const lat = 24.79634
	//const lat = -179.87654
	const lat = 180.0
	//const lat = 270.0
	rad := ra(lat)
	max := math.MaxInt32
	xlt := float32(max) / math.Pi
	adj := int32(rad * xlt)
	fmt.Printf("xlt by: %f -- %d\n", xlt, int(xlt))
	fmt.Printf("From: %.5f to:%d\n", lat, adj)
	fmt.Printf("Hexd: %.5f to:%X\n", lat, adj)
	bak := deg(float32(adj) / xlt)
	//fmt.Println("max:",max)
	fmt.Printf("bak:%.5f\n", bak)
}

/*
xlt by: 683565248.000000 -- 683565248
From: -179.87654 to:-2146010624
bak:-179.87654
*/
