package genlib

func PointOnCircle[T_ROT Real, T_VEC Real](radians T_ROT, radius T_VEC, center Vec2[T_VEC]) Vec2[T_VEC] {
	sin := Sin(float64(radians))
	cos := Cos(float64(radians))
	x, y := T_VEC(cos*float64(radius)), T_VEC(sin*float64(radius))
	return Vec2[T_VEC]{x, y}.Add(center)
}
func PointOnCircleDeg[T_ROT Real, T_VEC Real](degrees T_ROT, radius T_VEC, center Vec2[T_VEC]) Vec2[T_VEC] {
	return PointOnCircle(float64(degrees)*RAD_TO_DEG, radius, center)
}

func PointsOnCircle[T_ROT Real, T_VEC Real, I_CNT Integer](rotationRadians T_ROT, radius T_VEC, center Vec2[T_VEC], count I_CNT) []Vec2[T_VEC] {
	count = Floor(count)
	thetaStep := TAU / float64(count)
	points := make([]Vec2[T_VEC], count)
	for i, t := I_CNT(0), float64(rotationRadians); i < count; i, t = i+1, t+thetaStep {
		points[i] = PointOnCircle(t, radius, center)
	}
	return points
}
func PointsOnCircleDeg[T_ROT Real, T_VEC Real, I_CNT Integer](degrees T_ROT, radius T_VEC, center Vec2[T_VEC], count I_CNT) []Vec2[T_VEC] {
	return PointsOnCircle(float64(degrees)*RAD_TO_DEG, radius, center, count)
}

func PointsOnRing[T_ROT Real, T_VEC Real, I_CNT Integer](rotationRadians T_ROT, radiusInner T_VEC, radiusOuter T_VEC, center Vec2[T_VEC], count I_CNT) []Vec2[T_VEC] {
	count = Floor(count)
	thetaStep := TAU / float64(count)
	points := make([]Vec2[T_VEC], count*2)
	for i, t := I_CNT(0), float64(rotationRadians); i < count*2; i, t = i+2, t+thetaStep {
		points[i] = PointOnCircle(t, radiusInner, center)
		points[i+1] = PointOnCircle(t, radiusOuter, center)
	}
	return points
}
func PointsOnRingDeg[T_ROT Real, T_VEC Real, I_CNT Integer](degrees T_ROT, radiusInner T_VEC, radiusOuter T_VEC, center Vec2[T_VEC], count I_CNT) []Vec2[T_VEC] {
	return PointsOnRing(float64(degrees)*RAD_TO_DEG, radiusInner, radiusOuter, center, count)
}

func Circumference[T Real](radius T) T {
	return T(TAU * float64(radius))
}

func AreaOfTriangle[T Real](v Tri2[T]) T {
	return Abs(((v[1].X()*v[0].Y())-(v[0].X()*v[1].Y()))+((v[2].X()*v[1].Y())-(v[1].X()*v[2].Y()))+((v[0].X()*v[2].Y())-(v[2].X()*v[0].Y()))) / 2
}

func AreaOfCircle[T Real](radius T) T {
	return T(PI * Pow(float64(radius), 2))
}
