package vec2d

import "math"

// Vector type defines vector using exported float64 values: X and Y
type Vector struct {
	X, Y float64
}

// Init initializes already created vector
func (v *Vector) Init(x, y float64) *Vector {
	v.X = x
	v.Y = y
	return v
}

// New returns a new vector
func New(x, y float64) *Vector {
	return new(Vector).Init(x, y)
}

// IsEqual compares а Vector with another and returns true if they're equal
func (v *Vector) IsEqual(other *Vector) bool {
	return v.X == other.X && v.Y == other.Y
}

// Angle returns the Vector's angle in float64
func (v *Vector) Angle() float64 {
	return math.Atan2(v.Y, v.X) / (math.Pi / 180)
}

// SetAngle changes Vector's angle using vector rotation
func (v *Vector) SetAngle(angle_degrees float64) {
	v.X = v.Length()
	v.Y = 0.0
	v.Rotate(angle_degrees)
}

// Length returns... well the Vector's length
func (v *Vector) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// SetLength changes Vector's length, which obviously changes
// the values of Vector.X and Vector.Y
func (v *Vector) SetLength(value float64) {
	length := v.Length()
	v.X *= value / length
	v.Y *= value / length
}

// Rotate Vector by given angle degrees in float64
func (v *Vector) Rotate(angle_degrees float64) {
	radians := (math.Pi / 180) * angle_degrees
	sin := math.Sin(radians)
	cos := math.Cos(radians)

	x := v.X*cos - v.Y*sin
	y := v.X*sin + v.Y*cos
	v.X = x
	v.Y = y
}

// Collect changes Vector's X and Y by collecting them with other's
func (v *Vector) Collect(other *Vector) {
	v.X += other.X
	v.Y += other.Y
}

// CollectToFloat64 changes Vector's X and Y by collecting them with value
func (v *Vector) CollectToFloat64(value float64) {
	v.X += value
	v.Y += value
}

// Sub changes Vector's X and Y by substracting them with other's
func (v *Vector) Sub(other *Vector) {
	v.X -= other.X
	v.Y -= other.Y
}

// SubToFloat64 changes Vector's X and Y by substracting them with value
func (v *Vector) SubToFloat64(value float64) {
	v.X -= value
	v.Y -= value
}

// Mul changes Vector's X and Y by multiplying them with other's
func (v *Vector) Mul(other *Vector) {
	v.X *= other.X
	v.Y *= other.Y
}

// MulToFloat64 changes Vector's X and Y by multiplying them with value
func (v *Vector) MulToFloat64(value float64) {
	v.X *= value
	v.Y *= value
}

// Div changes Vector's X and Y by dividing them with other's
func (v *Vector) Div(other *Vector) {
	v.X /= other.X
	v.Y /= other.Y
}

// DivToFloat64 changes Vector's X and Y by dividing them with value
func (v *Vector) DivToFloat64(value float64) {
	v.X /= value
	v.Y /= value
}
