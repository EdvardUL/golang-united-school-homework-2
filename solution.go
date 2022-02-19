package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type intCustomType int

const SidesTriangle intCustomType = 3
const SidesSquare intCustomType = 4
const SidesCircle intCustomType = 0

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	if sidesNum == 3 {
		return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	}
	if sidesNum == 4 {
		return math.Pow(sideLen, 2)
	}
	if sidesNum == 0 {
		return math.Pow(sideLen, 2) * math.Pi
	}
	return 0
}
