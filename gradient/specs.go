package gradient

import (
	"image"
	"math"

	"github.com/shelepuginivan/color/internal/degrees"
)

// angleSpec represents configuration for gradient angle.
//
// As a standard, value is 0 at top and increases clockwise. Values returned
// by [NormalizedRadians] and [NormalizedRadiansWithCenter] are the values
// adjusted to match the standard that are ready to use in formulas.
type angleSpec struct {
	angle     *int
	direction *Direction
}

func (a *angleSpec) NormalizedRadiansWithCenter(rect image.Rectangle, center point) float64 {
	if a.angle != nil {
		return degrees.ToRadians(90 - *a.angle)
	}

	if a.direction == nil {
		return math.Pi / 2
	}

	switch *a.direction {
	case TopLeft:
		dx := center.x - rect.Min.X
		dy := center.y - rect.Min.Y
		return math.Atan2(float64(dy), -float64(dx))
	case Top:
		return math.Pi / 2
	case TopRight:
		dx := rect.Max.X - center.x
		dy := center.y - rect.Min.Y
		return math.Atan2(float64(dy), float64(dx))
	case Right:
		return 0
	case BottomRight:
		dx := rect.Max.X - center.x
		dy := rect.Max.Y - center.y
		return math.Atan2(-float64(dy), float64(dx))
	case Bottom:
		return 3 * math.Pi / 2
	case BottomLeft:
		dx := center.x - rect.Min.X
		dy := rect.Max.Y - center.y
		return math.Atan2(-float64(dy), -float64(dx))
	case Left:
		return math.Pi
	default:
		return 0
	}
}

func (a *angleSpec) NormalizedRadians(rect image.Rectangle) float64 {
	if a.angle != nil {
		return degrees.ToRadians(90 - *a.angle)
	}

	if a.direction == nil {
		return math.Pi / 2
	}

	dx := rect.Dx()
	dy := rect.Dy()

	switch *a.direction {
	case TopLeft:
		return math.Atan2(float64(dy), -float64(dx))
	case Top:
		return math.Pi / 2
	case TopRight:
		return math.Atan2(float64(dy), float64(dx))
	case Right:
		return 0
	case BottomRight:
		return math.Atan2(-float64(dy), float64(dx))
	case Bottom:
		return 3 * math.Pi / 2
	case BottomLeft:
		return math.Atan2(-float64(dy), -float64(dx))
	case Left:
		return math.Pi
	default:
		return 0
	}
}

// point represents a point on the image.
type point struct {
	x, y int
}

// pointSpec represents configuration for point.
//
// Position can be configured with either absolute or relative coordinates,
// absolute take precedence.
type pointSpec struct {
	x, y       *int
	relX, relY *float64
}

func (p pointSpec) Position(rect image.Rectangle) point {
	if p.x != nil && p.y != nil {
		return point{*p.x, *p.y}
	}

	if p.relX == nil {
		d := 0.5
		p.relX = &d
	}

	if p.relY == nil {
		d := 0.5
		p.relY = &d
	}

	absX := *p.relX * float64(rect.Dx())
	absY := *p.relY * float64(rect.Dy())

	x := rect.Min.X + int(math.Round(absX))
	y := rect.Min.Y + int(math.Round(absY))

	return point{x, y}
}
