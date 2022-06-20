package genlib

type LengthImperial float64

func (l LengthImperial) ToMetric() LengthMetric {
	return LengthMetric(l * MIL_INCH_TO_MILLIMETER)
}

type LengthMetric float64

func (l LengthMetric) ToImperial() LengthImperial {
	return LengthImperial(l * MILLIMETER_TO_MILL_INCH)
}

type AngleDegrees float64

func (a AngleDegrees) ToRadians() AngleRadians {
	return AngleRadians(a * DEG_TO_RAD)
}

type AngleRadians float64

func (a AngleRadians) ToDegrees() AngleDegrees {
	return AngleDegrees(a * RAD_TO_DEG)
}

const (
	RADIAN = AngleRadians(1)
	DEGREE = AngleDegrees(1)

	DEG_TO_RAD = 0.01745329251994329576923690768488612713442871888541725456097191 // Tau / 360
	RAD_TO_DEG = 57.2957795130823208767981548141051703324054724665643215491602439 // 360 / Tau

	MIL_INCH = LengthImperial(1)
	INCH     = MIL_INCH * 1000
	FOOT     = INCH * 12
	YARD     = FOOT * 3
	MILE     = FOOT * 5280

	MILLIMETER = LengthMetric(1)
	CENTIMETER = MILLIMETER * 10
	METER      = CENTIMETER * 100

	MIL_INCH_TO_MILLIMETER  = 0.0254
	MILLIMETER_TO_MILL_INCH = 39.37007874015748031496062992125984
)
