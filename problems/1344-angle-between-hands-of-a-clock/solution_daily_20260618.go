package main

import "math"

// angleClock computes the smaller angle (in degrees) between the hour and
// minute hands of a clock.
//
// Approach (pure math, O(1)):
//   - The minute hand moves 360/60 = 6 degrees per minute.
//   - The hour hand moves 360/12 = 30 degrees per hour, plus 0.5 degrees per
//     minute to account for its continuous drift within the hour.
//   - We reduce the hour modulo 12 so that 12 o'clock maps to 0 degrees.
//   - The two hands divide the circle into two arcs; we return the smaller,
//     which is min(diff, 360 - diff).
func angleClock(hour int, minutes int) float64 {
	minuteAngle := float64(minutes) * 6.0
	hourAngle := float64(hour%12)*30.0 + float64(minutes)*0.5

	diff := math.Abs(hourAngle - minuteAngle)
	return math.Min(diff, 360.0-diff)
}
