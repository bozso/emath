package rand

import (
    "math/rand"
)

// type alias for the rand package's rand.Source type
type Source rand.Source

// wrapper for the rand package NewSource function
func NewSource(seed int64) (s Source) {
    s = Source(rand.NewSource(seed))
    return
}

type Scaler interface {
    Int(int) int
    Float32(float32) float32
    Float64(float64) float64
}

/*
 * Configuration for shifting the mean value and scaling the standard
 * deviation of random values.
 */
type Scale struct {
    Mean float64
    Std float64
}

// scale an integer number
func (s Scale) Int(i int) (ri int) {
    return int(s.Std) * i + int(s.Mean)
}

// scale a float32 number
func (s Scale) Float32(f float32) (rf float32) {
    return float32(s.Std) * f + float32(s.Mean)
}

// scale a float64 number
func (s Scale) Float64(f float64) (rf float64) {
    return s.Std * f + s.Mean
}

/*
 * Wrapper for the rand.Rand struct. Applies scaling defined by a Config
 * struct.
 */
type Rand struct {
    scaler Scaler
    *rand.Rand
}

// Create a new Rand struct with configuration defined by Config.
func (sc Scale) New(s rand.Source) (r Rand) {
    r.scaler, r.Rand = sc, rand.New(s)
    return
}

// Get a random Float32.
func (r Rand) Float32() (f float32) {
    return r.scaler.Float32(r.Rand.Float32())
}

// Get a random Float64.
func (r Rand) Float64() (f float64) {
    return r.scaler.Float64(r.Rand.Float64())
}

// Get a random Int.
func (r Rand) Int() (i int) {
    return r.scaler.Int(r.Rand.Int())
}

// Returns a copy of the default configuration with zero mean and 1.0
// standard deviation.
func DefaultScale() (s Scale) {
    return defaultScale
}

var defaultScale = Scale{
    Mean: 0.0,
    Std: 1.0,
}

// Create a new random number generator with the default configuration.
func NoScale(s rand.Source) (r Rand) {
    return defaultScale.New(s)
}
