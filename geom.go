package geom

import (
	"math"
)

const EPS = 1e-9

type Point struct{ x, y float64 }

// Model a linear equation ax + by + c = 0.
//  a is the slope.
// b = 0 if it is a vertical line; otherwise b = 1.
type Line struct{ a, b, c float64 }

// Given two points compute line: ax + by + c = 0.
func NewLine(p1, p2 Point) *Line {
	var l Line
	// l is vertical line
	if p1.x == p2.x {
		l.a = 1.0
		l.b = 0
		l.c = -p1.x
	} else {
		l.a = -(p1.y - p2.y) / (p1.x - p2.x)
		l.b = 1.0
		l.c = -(l.a * p1.x) - p1.y
	}
	return &l
}

func SameLines(l1, l2 *Line) bool {
	return fabs(l1.a-l2.a) < EPS &&
		fabs(l1.b-l2.b) < EPS &&
		fabs(l1.c-l2.c) < EPS
}

func fabs(fb float64) float64 {
	if fb < 0 {
		return fb * (-1)
	}
	return fb
}

func IntersectionPoint(l1, l2 *Line) *Point {
	// If l1 and l2 are parallel, return nil.
	if fabs(l1.a-l2.a) < EPS && fabs(l1.b-l2.b) < EPS {
		return nil // no intersection
	}

	var p Point
	// Solve the system of two linear algebraic equations.
	p.x = (l2.b*l1.c - l1.b*l2.c) / (l2.a*l1.b - l1.a*l2.b)

	if fabs(l1.b) > EPS {
		p.y = -(l1.a*p.x + l1.c)
	} else {
		p.y = -(l2.a*p.x + l2.c)
	}
	return &p
}

// Distance calculates the Euclidean distance between p1 and p2.
func Distance(p1, p2 Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

// Given p1, p2 and p3 are collinear, is p3 on the segment formed between p1 and p2
func IsOnSegment(p1, p2, p3 Point) bool {
	if Distance(p1, p3)-EPS <= Distance(p1, p2) &&
		Distance(p2, p3)-EPS <= Distance(p1, p2) {
		return true
	}
	return false
}

func max(x1, x2 float64) float64 {
	if x1 < x2 {
		return x2
	}
	return x1
}

func min(x1, x2 float64) float64 {
	if x1 < x2 {
		return x1
	}
	return x2
}

// Is p3 inside the rectangle formed by left-top p1 and right-bottom p2?
func IsInsideRectangle(p1, p2, p3 Point) bool {
	if p3.x <= max(p1.x, p2.x) &&
		p3.x >= min(p1.x, p2.x) &&
		p3.y <= max(p1.y, p2.y) &&
		p3.y >= min(p1.y, p2.y) {
		return true
	}
	return false
}

type Rect struct {
	p1, p2 Point
}

// Intersects is line drawn through two points intersects with rectangle.
func Intersects(pt1, pt2 Point, rect Rect) bool {
	corners := []Point{rect.p1, {0, 0}, rect.p2, {0, 0}}

	if IsInsideRectangle(corners[0], corners[2], pt1) ||
		IsInsideRectangle(corners[0], corners[2], pt2) {
		return true
	}

	corners[1] = corners[0]
	corners[3] = corners[2]
	corners[1].y, corners[3].y = corners[3].y, corners[1].y

	line := NewLine(pt1, pt2)
	intersect := false

	for i := 0; i < 4; i++ {
		edge := NewLine(corners[i], corners[(i+1)%4])

		if SameLines(line, edge) {
			if IsOnSegment(pt1, pt2, corners[i]) ||
				IsOnSegment(pt1, pt2, corners[(i+1)%4]) {
				intersect = true
				break
			}
			continue
		}

		if p := IntersectionPoint(line, edge); p != nil {
			if IsOnSegment(pt1, pt2, *p) &&
				IsOnSegment(corners[i], corners[(i+1)%4], *p) {
				intersect = true
				break
			}
		}
	}
	return intersect
}
