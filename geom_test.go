package geom

import "testing"

func TestIntersect(t *testing.T) {

	pt1 := Point{4, 9}
	pt2 := Point{11, 2}
	rect := Rect{Point{1, 5}, Point{7, 1}}
	if Intersects(pt1, pt2, rect) {
		t.Error("got: intersects=true, want: intersects=false")
	}

	pt1 = Point{2, -5}
	pt2 = Point{-10, 3}
	rect = Rect{Point{-1, 0}, Point{2, -3}}
	if !Intersects(pt1, pt2, rect) {
		t.Error("got: intersects=false, want: intersects=true")
	}

}
