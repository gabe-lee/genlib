package genlib

type Vec2[T Real] [2]T
type Vec3[T Real] [3]T
type Vec4[T Real] [4]T
type VecN[T Real] []T

type IVec2 = Vec2[int]
type UVec2 = Vec2[uint]
type I8Vec2 = Vec2[int8]
type U8Vec2 = Vec2[uint8]
type I16Vec2 = Vec2[int16]
type U16Vec2 = Vec2[uint16]
type I32Vec2 = Vec2[int32]
type U32Vec2 = Vec2[uint32]
type I64Vec2 = Vec2[int64]
type U64Vec2 = Vec2[uint64]
type F32Vec2 = Vec2[float32]
type F64Vec2 = Vec2[float64]

type IVec3 = Vec3[int]
type UVec3 = Vec3[uint]
type I8Vec3 = Vec3[int8]
type U8Vec3 = Vec3[uint8]
type I16Vec3 = Vec3[int16]
type U16Vec3 = Vec3[uint16]
type I32Vec3 = Vec3[int32]
type U32Vec3 = Vec3[uint32]
type I64Vec3 = Vec3[int64]
type U64Vec3 = Vec3[uint64]
type F32Vec3 = Vec3[float32]
type F64Vec3 = Vec3[float64]

type IVec4 = Vec4[int]
type UVec4 = Vec4[uint]
type I8Vec4 = Vec4[int8]
type U8Vec4 = Vec4[uint8]
type I16Vec4 = Vec4[int16]
type U16Vec4 = Vec4[uint16]
type I32Vec4 = Vec4[int32]
type U32Vec4 = Vec4[uint32]
type I64Vec4 = Vec4[int64]
type U64Vec4 = Vec4[uint64]
type F32Vec4 = Vec4[float32]
type F64Vec4 = Vec4[float64]

type IVecN = VecN[int]
type UVecN = VecN[uint]
type I8VecN = VecN[int8]
type U8VecN = VecN[uint8]
type I16VecN = VecN[int16]
type U16VecN = VecN[uint16]
type I32VecN = VecN[int32]
type U32VecN = VecN[uint32]
type I64VecN = VecN[int64]
type U64VecN = VecN[uint64]
type F32VecN = VecN[float32]
type F64VecN = VecN[float64]

type Line2[T Real] [2]Vec2[T]
type Tri2[T Real] [3]Vec2[T]
type Quad2[T Real] [4]Vec2[T]
type AABB2[T Real] [2]Vec2[T]
type Array2[T Real] []Vec2[T]

type Circ2[T Real] struct {
	Center Vec2[T]
	Radius T
}

type Line3[T Real] [2]Vec3[T]
type Tri3[T Real] [3]Vec3[T]
type Quad3[T Real] [4]Vec3[T]
type AABB3[T Real] [2]Vec3[T]
type Array3[T Real] []Vec3[T]

type Circ3[T Real] struct {
	Center Vec3[T]
	Radius T
}

type ILine2 = Line2[int]
type ULine2 = Line2[uint]
type I8Line2 = Line2[int8]
type U8Line2 = Line2[uint8]
type I16Line2 = Line2[int16]
type U16Line2 = Line2[uint16]
type I32Line2 = Line2[int32]
type U32Line2 = Line2[uint32]
type I64Line2 = Line2[int64]
type U64Line2 = Line2[uint64]
type F32Line2 = Line2[float32]
type F64Line2 = Line2[float64]

type ITri2 = Tri2[int]
type UTri2 = Tri2[uint]
type I8Tri2 = Tri2[int8]
type U8Tri2 = Tri2[uint8]
type I16Tri2 = Tri2[int16]
type U16Tri2 = Tri2[uint16]
type I32Tri2 = Tri2[int32]
type U32Tri2 = Tri2[uint32]
type I64Tri2 = Tri2[int64]
type U64Tri2 = Tri2[uint64]
type F32Tri2 = Tri2[float32]
type F64Tri2 = Tri2[float64]

type IQuad2 = Quad2[int]
type UQuad2 = Quad2[uint]
type I8Quad2 = Quad2[int8]
type U8Quad2 = Quad2[uint8]
type I16Quad2 = Quad2[int16]
type U16Quad2 = Quad2[uint16]
type I32Quad2 = Quad2[int32]
type U32Quad2 = Quad2[uint32]
type I64Quad2 = Quad2[int64]
type U64Quad2 = Quad2[uint64]
type F32Quad2 = Quad2[float32]
type F64Quad2 = Quad2[float64]

type IAABB2 = AABB2[int]
type UAABB2 = AABB2[uint]
type I8AABB2 = AABB2[int8]
type U8AABB2 = AABB2[uint8]
type I16AABB2 = AABB2[int16]
type U16AABB2 = AABB2[uint16]
type I32AABB2 = AABB2[int32]
type U32AABB2 = AABB2[uint32]
type I64AABB2 = AABB2[int64]
type U64AABB2 = AABB2[uint64]
type F32AABB2 = AABB2[float32]
type F64AABB2 = AABB2[float64]

type IRect2 = AABB2[int]
type URect2 = AABB2[uint]
type I8Rect2 = AABB2[int8]
type U8Rect2 = AABB2[uint8]
type I16Rect2 = AABB2[int16]
type U16Rect2 = AABB2[uint16]
type I32Rect2 = AABB2[int32]
type U32Rect2 = AABB2[uint32]
type I64Rect2 = AABB2[int64]
type U64Rect2 = AABB2[uint64]
type F32Rect2 = AABB2[float32]
type F64Rect2 = AABB2[float64]

type IArray2 = Array2[int]
type UArray2 = Array2[uint]
type I8Array2 = Array2[int8]
type U8Array2 = Array2[uint8]
type I16Array2 = Array2[int16]
type U16Array2 = Array2[uint16]
type I32Array2 = Array2[int32]
type U32Array2 = Array2[uint32]
type I64Array2 = Array2[int64]
type U64Array2 = Array2[uint64]
type F32Array2 = Array2[float32]
type F64Array2 = Array2[float64]

type ICirc2 = Circ2[int]
type UCirc2 = Circ2[uint]
type I8Circ2 = Circ2[int8]
type U8Circ2 = Circ2[uint8]
type I16Circ2 = Circ2[int16]
type U16Circ2 = Circ2[uint16]
type I32Circ2 = Circ2[int32]
type U32Circ2 = Circ2[uint32]
type I64Circ2 = Circ2[int64]
type U64Circ2 = Circ2[uint64]
type F32Circ2 = Circ2[float32]
type F64Circ2 = Circ2[float64]

type ILine3 = Line3[int]
type ULine3 = Line3[uint]
type I8Line3 = Line3[int8]
type U8Line3 = Line3[uint8]
type I16Line3 = Line3[int16]
type U16Line3 = Line3[uint16]
type I32Line3 = Line3[int32]
type U32Line3 = Line3[uint32]
type I64Line3 = Line3[int64]
type U64Line3 = Line3[uint64]
type F32Line3 = Line3[float32]
type F64Line3 = Line3[float64]

type ITri3 = Tri3[int]
type UTri3 = Tri3[uint]
type I8Tri3 = Tri3[int8]
type U8Tri3 = Tri3[uint8]
type I16Tri3 = Tri3[int16]
type U16Tri3 = Tri3[uint16]
type I32Tri3 = Tri3[int32]
type U32Tri3 = Tri3[uint32]
type I64Tri3 = Tri3[int64]
type U64Tri3 = Tri3[uint64]
type F32Tri3 = Tri3[float32]
type F64Tri3 = Tri3[float64]

type IQuad3 = Quad3[int]
type UQuad3 = Quad3[uint]
type I8Quad3 = Quad3[int8]
type U8Quad3 = Quad3[uint8]
type I16Quad3 = Quad3[int16]
type U16Quad3 = Quad3[uint16]
type I32Quad3 = Quad3[int32]
type U32Quad3 = Quad3[uint32]
type I64Quad3 = Quad3[int64]
type U64Quad3 = Quad3[uint64]
type F32Quad3 = Quad3[float32]
type F64Quad3 = Quad3[float64]

type IAABB3 = AABB3[int]
type UAABB3 = AABB3[uint]
type I8AABB3 = AABB3[int8]
type U8AABB3 = AABB3[uint8]
type I16AABB3 = AABB3[int16]
type U16AABB3 = AABB3[uint16]
type I32AABB3 = AABB3[int32]
type U32AABB3 = AABB3[uint32]
type I64AABB3 = AABB3[int64]
type U64AABB3 = AABB3[uint64]
type F32AABB3 = AABB3[float32]
type F64AABB3 = AABB3[float64]

type IRect3 = AABB3[int]
type URect3 = AABB3[uint]
type I8Rect3 = AABB3[int8]
type U8Rect3 = AABB3[uint8]
type I16Rect3 = AABB3[int16]
type U16Rect3 = AABB3[uint16]
type I32Rect3 = AABB3[int32]
type U32Rect3 = AABB3[uint32]
type I64Rect3 = AABB3[int64]
type U64Rect3 = AABB3[uint64]
type F32Rect3 = AABB3[float32]
type F64Rect3 = AABB3[float64]

type IArray3 = Array3[int]
type UArray3 = Array3[uint]
type I8Array3 = Array3[int8]
type U8Array3 = Array3[uint8]
type I16Array3 = Array3[int16]
type U16Array3 = Array3[uint16]
type I32Array3 = Array3[int32]
type U32Array3 = Array3[uint32]
type I64Array3 = Array3[int64]
type U64Array3 = Array3[uint64]
type F32Array3 = Array3[float32]
type F64Array3 = Array3[float64]

type ICirc3 = Circ3[int]
type UCirc3 = Circ3[uint]
type I8Circ3 = Circ3[int8]
type U8Circ3 = Circ3[uint8]
type I16Circ3 = Circ3[int16]
type U16Circ3 = Circ3[uint16]
type I32Circ3 = Circ3[int32]
type U32Circ3 = Circ3[uint32]
type I64Circ3 = Circ3[int64]
type U64Circ3 = Circ3[uint64]
type F32Circ3 = Circ3[float32]
type F64Circ3 = Circ3[float64]

func (self Vec2[T]) X() T {
	return self[0]
}
func (self Vec3[T]) X() T {
	return self[0]
}
func (self Vec4[T]) X() T {
	return self[0]
}
func (self Vec2[T]) R() T {
	return self[0]
}
func (self Vec3[T]) R() T {
	return self[0]
}
func (self Vec4[T]) R() T {
	return self[0]
}

func (self Vec2[T]) Y() T {
	return self[1]
}
func (self Vec3[T]) Y() T {
	return self[1]
}
func (self Vec4[T]) Y() T {
	return self[1]
}
func (self Vec2[T]) G() T {
	return self[1]
}
func (self Vec3[T]) G() T {
	return self[1]
}
func (self Vec4[T]) G() T {
	return self[1]
}

func (self Vec3[T]) Z() T {
	return self[2]
}
func (self Vec4[T]) Z() T {
	return self[2]
}
func (self Vec3[T]) B() T {
	return self[2]
}
func (self Vec4[T]) B() T {
	return self[2]
}

func (self Vec4[T]) W() T {
	return self[3]
}
func (self Vec4[T]) A() T {
	return self[3]
}

func (self Vec2[T]) WithX(val T) Vec2[T] {
	self[0] = val
	return self
}
func (self Vec3[T]) WithX(val T) Vec3[T] {
	self[0] = val
	return self
}
func (self Vec4[T]) WithX(val T) Vec4[T] {
	self[0] = val
	return self
}
func (self Vec2[T]) WithR(val T) Vec2[T] {
	self[0] = val
	return self
}
func (self Vec3[T]) WithR(val T) Vec3[T] {
	self[0] = val
	return self
}
func (self Vec4[T]) WithR(val T) Vec4[T] {
	self[0] = val
	return self
}

func (self Vec2[T]) WithY(val T) Vec2[T] {
	self[1] = val
	return self
}
func (self Vec3[T]) WithY(val T) Vec3[T] {
	self[1] = val
	return self
}
func (self Vec4[T]) WithY(val T) Vec4[T] {
	self[1] = val
	return self
}
func (self Vec2[T]) WithG(val T) Vec2[T] {
	self[1] = val
	return self
}
func (self Vec3[T]) WithG(val T) Vec3[T] {
	self[1] = val
	return self
}
func (self Vec4[T]) WithG(val T) Vec4[T] {
	self[1] = val
	return self
}

func (self Vec3[T]) WithZ(val T) Vec3[T] {
	self[2] = val
	return self
}
func (self Vec4[T]) WithZ(val T) Vec4[T] {
	self[2] = val
	return self
}
func (self Vec3[T]) WithB(val T) Vec3[T] {
	self[2] = val
	return self
}
func (self Vec4[T]) WithB(val T) Vec4[T] {
	self[2] = val
	return self
}

func (self Vec4[T]) WithW(val T) Vec4[T] {
	self[3] = val
	return self
}
func (self Vec4[T]) WithA(val T) Vec4[T] {
	self[3] = val
	return self
}

func MakeVec2[T Real](vals ...T) (result Vec2[T]) {
	return result.Swizzle(vals...)
}
func MakeVec3[T Real](vals ...T) (result Vec3[T]) {
	return result.Swizzle(vals...)
}
func MakeVec4[T Real](vals ...T) (result Vec4[T]) {
	return result.Swizzle(vals...)
}
func MakeVecN[T Real](vals ...T) (result VecN[T]) {
	return VecN[T](vals)
}

func (self Vec2[T]) Swizzle(vals ...T) (result Vec2[T]) {
	i := 0
	for i < len(result) && i < len(vals) {
		result[i] = vals[i]
		i += 1
	}
	for i < len(result) {
		result[i] = self[i]
		i += 1
	}
	return result
}
func (self Vec3[T]) Swizzle(vals ...T) (result Vec3[T]) {
	i := 0
	for i < len(result) && i < len(vals) {
		result[i] = vals[i]
		i += 1
	}
	for i < len(result) {
		result[i] = self[i]
		i += 1
	}
	return result
}
func (self Vec4[T]) Swizzle(vals ...T) (result Vec4[T]) {
	i := 0
	for i < len(result) && i < len(vals) {
		result[i] = vals[i]
		i += 1
	}
	for i < len(result) {
		result[i] = self[i]
		i += 1
	}
	return result
}

func (self *Vec2[T]) SetX(val T) {
	self[0] = val
}
func (self *Vec3[T]) SetX(val T) {
	self[0] = val
}
func (self *Vec4[T]) SetX(val T) {
	self[0] = val
}
func (self *Vec2[T]) SetR(val T) {
	self[0] = val
}
func (self *Vec3[T]) SetR(val T) {
	self[0] = val
}
func (self *Vec4[T]) SetR(val T) {
	self[0] = val
}

func (self *Vec2[T]) SetY(val T) {
	self[1] = val
}
func (self *Vec3[T]) SetY(val T) {
	self[1] = val
}
func (self *Vec4[T]) SetY(val T) {
	self[1] = val
}
func (self *Vec2[T]) SetG(val T) {
	self[1] = val
}
func (self *Vec3[T]) SetG(val T) {
	self[1] = val
}
func (self *Vec4[T]) SetG(val T) {
	self[1] = val
}

func (self *Vec3[T]) SetZ(val T) {
	self[2] = val
}
func (self *Vec4[T]) SetZ(val T) {
	self[2] = val
}
func (self *Vec3[T]) SetB(val T) {
	self[2] = val
}
func (self *Vec4[T]) SetB(val T) {
	self[2] = val
}

func (self *Vec4[T]) SetW(val T) {
	self[3] = val
}
func (self *Vec4[T]) SetA(val T) {
	self[3] = val
}

func (self Vec2[T]) Add(other Vec2[T]) (result Vec2[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	return result
}
func (self Vec2[T]) AddV3(other Vec3[T]) (result Vec2[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	return result
}
func (self Vec2[T]) AddV4(other Vec4[T]) (result Vec2[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	return result
}
func (self Vec3[T]) Add(other Vec3[T]) (result Vec3[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	result[2] = self[2] + other[2]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec3[T]) AddV2(other Vec2[T]) (result Vec3[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	result[2] = self[2]
	return result
}
func (self Vec3[T]) AddV4(other Vec4[T]) (result Vec3[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	result[2] = self[2] + other[2]
	return result
}
func (self Vec4[T]) Add(other Vec4[T]) (result Vec4[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	result[2] = self[2] + other[2]
	result[3] = self[3] + other[3]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) AddV2(other Vec2[T]) (result Vec4[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	result[2] = self[2]
	result[3] = self[3]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) AddV3(other Vec3[T]) (result Vec4[T]) {
	result[0] = self[0] + other[0]
	result[1] = self[1] + other[1]
	result[2] = self[2] + other[2]
	result[3] = self[3]
	return result
}

func (self Vec2[T]) Sub(other Vec2[T]) (result Vec2[T]) {
	return self.Add(other.Neg())
}
func (self Vec2[T]) SubV3(other Vec3[T]) (result Vec2[T]) {
	return self.AddV3(other.Neg())
}
func (self Vec2[T]) SubV4(other Vec4[T]) (result Vec2[T]) {
	return self.AddV4(other.Neg())
}
func (self Vec3[T]) Sub(other Vec3[T]) (result Vec3[T]) {
	return self.Add(other.Neg())
}

// Missing terms in 'other' are ignored
func (self Vec3[T]) SubV2(other Vec2[T]) (result Vec3[T]) {
	return self.AddV2(other.Neg())
}
func (self Vec3[T]) SubV4(other Vec4[T]) (result Vec3[T]) {
	return self.AddV4(other.Neg())
}
func (self Vec4[T]) Sub(other Vec4[T]) (result Vec4[T]) {
	return self.Add(other.Neg())
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) SubV2(other Vec2[T]) (result Vec4[T]) {
	return self.AddV2(other.Neg())
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) SubV3(other Vec3[T]) (result Vec4[T]) {
	return self.AddV3(other.Neg())
}

func (self Vec2[T]) Grow(add T) (result Vec2[T]) {
	result[0] = self[0] + add
	result[1] = self[1] + add
	return result
}
func (self Vec3[T]) Grow(add T) (result Vec3[T]) {
	result[0] = self[0] + add
	result[1] = self[1] + add
	result[2] = self[2] + add
	return result
}
func (self Vec4[T]) Grow(add T) (result Vec4[T]) {
	result[0] = self[0] + add
	result[1] = self[1] + add
	result[2] = self[2] + add
	result[3] = self[3] + add
	return result
}

func (self Vec2[T]) Mult(other Vec2[T]) (result Vec2[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	return result
}
func (self Vec2[T]) MultV3(other Vec3[T]) (result Vec2[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	return result
}
func (self Vec2[T]) MultV4(other Vec4[T]) (result Vec2[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	return result
}
func (self Vec3[T]) Mult(other Vec3[T]) (result Vec3[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	result[2] = self[2] * other[2]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec3[T]) MultV2(other Vec2[T]) (result Vec3[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	result[2] = self[2]
	return result
}
func (self Vec3[T]) MultV4(other Vec4[T]) (result Vec3[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	result[2] = self[2] * other[2]
	return result
}
func (self Vec4[T]) Mult(other Vec4[T]) (result Vec4[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	result[2] = self[2] * other[2]
	result[3] = self[3] * other[3]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) MultV2(other Vec2[T]) (result Vec4[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	result[2] = self[2]
	result[3] = self[3]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) MultV3(other Vec3[T]) (result Vec4[T]) {
	result[0] = self[0] * other[0]
	result[1] = self[1] * other[1]
	result[2] = self[2] * other[2]
	result[3] = self[3]
	return result
}

func (self Vec2[T]) Div(other Vec2[T]) (result Vec2[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	return result
}
func (self Vec2[T]) DivV3(other Vec3[T]) (result Vec2[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	return result
}
func (self Vec2[T]) DivV4(other Vec4[T]) (result Vec2[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	return result
}
func (self Vec3[T]) Div(other Vec3[T]) (result Vec3[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	result[2] = self[2] / other[2]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec3[T]) DivV2(other Vec2[T]) (result Vec3[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	result[2] = self[2]
	return result
}
func (self Vec3[T]) DivV4(other Vec4[T]) (result Vec3[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	result[2] = self[2] / other[2]
	return result
}
func (self Vec4[T]) Div(other Vec4[T]) (result Vec4[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	result[2] = self[2] / other[2]
	result[3] = self[3] / other[3]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) DivV2(other Vec2[T]) (result Vec4[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	result[2] = self[2]
	result[3] = self[3]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) DivV3(other Vec3[T]) (result Vec4[T]) {
	result[0] = self[0] / other[0]
	result[1] = self[1] / other[1]
	result[2] = self[2] / other[2]
	result[3] = self[3]
	return result
}

func (self Vec2[T]) Scale(val T) (result Vec2[T]) {
	result[0] = self[0] * val
	result[1] = self[1] * val
	return result
}
func (self Vec3[T]) Scale(val T) (result Vec3[T]) {
	result[0] = self[0] * val
	result[1] = self[1] * val
	result[2] = self[2] * val
	return result
}
func (self Vec4[T]) Scale(val T) (result Vec4[T]) {
	result[0] = self[0] * val
	result[1] = self[1] * val
	result[2] = self[2] * val
	result[3] = self[3] * val
	return result
}

func (self Vec2[T]) FMod(other Vec2[T]) (result Vec2[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	return result
}
func (self Vec2[T]) FModV3(other Vec3[T]) (result Vec2[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	return result
}
func (self Vec2[T]) FModV4(other Vec4[T]) (result Vec2[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	return result
}
func (self Vec3[T]) FMod(other Vec3[T]) (result Vec3[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	result[2] = FMod(self[2], other[2])
	return result
}

// Missing terms in 'other' are ignored
func (self Vec3[T]) FModV2(other Vec2[T]) (result Vec3[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	result[2] = self[2]
	return result
}
func (self Vec3[T]) FModV4(other Vec4[T]) (result Vec3[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	result[2] = FMod(self[2], other[2])
	return result
}
func (self Vec4[T]) FMod(other Vec4[T]) (result Vec4[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	result[2] = FMod(self[2], other[2])
	result[3] = FMod(self[3], other[3])
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) FModV2(other Vec2[T]) (result Vec4[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	result[2] = self[2]
	result[3] = self[3]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) FModV3(other Vec3[T]) (result Vec4[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = FMod(self[1], other[1])
	result[2] = FMod(self[2], other[2])
	result[3] = self[3]
	return result
}

func (self Vec2[T]) IMod(other Vec2[T]) (result Vec2[T]) {
	result[0] = FMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	return result
}
func (self Vec2[T]) IModV3(other Vec3[T]) (result Vec2[T]) {
	result[0] = IMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	return result
}
func (self Vec2[T]) IModV4(other Vec4[T]) (result Vec2[T]) {
	result[0] = IMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	return result
}
func (self Vec3[T]) IMod(other Vec3[T]) (result Vec3[T]) {
	result[0] = IMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	result[2] = IMod(self[2], other[2])
	return result
}

// Missing terms in 'other' are ignored
func (self Vec3[T]) IModV2(other Vec2[T]) (result Vec3[T]) {
	result[0] = IMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	result[2] = self[2]
	return result
}
func (self Vec3[T]) IModV4(other Vec4[T]) (result Vec3[T]) {
	result[0] = IMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	result[2] = IMod(self[2], other[2])
	return result
}
func (self Vec4[T]) IMod(other Vec4[T]) (result Vec4[T]) {
	result[0] = IMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	result[2] = IMod(self[2], other[2])
	result[3] = IMod(self[3], other[3])
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) IModV2(other Vec2[T]) (result Vec4[T]) {
	result[0] = IMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	result[2] = self[2]
	result[3] = self[3]
	return result
}

// Missing terms in 'other' are ignored
func (self Vec4[T]) IModV3(other Vec3[T]) (result Vec4[T]) {
	result[0] = IMod(self[0], other[0])
	result[1] = IMod(self[1], other[1])
	result[2] = IMod(self[2], other[2])
	result[3] = self[3]
	return result
}

func (self Vec2[T]) Abs() (result Vec2[T]) {
	result[0] = self[0]
	result[1] = self[1]
	if result[0] < 0 {
		result[0] = -result[0]
	}
	if result[1] < 0 {
		result[1] = -result[1]
	}
	return result
}
func (self Vec3[T]) Abs() (result Vec3[T]) {
	result[0] = self[0]
	result[1] = self[1]
	result[2] = self[2]
	if result[0] < 0 {
		result[0] = -result[0]
	}
	if result[1] < 0 {
		result[1] = -result[1]
	}
	if result[2] < 0 {
		result[2] = -result[2]
	}
	return result
}
func (self Vec4[T]) Abs() (result Vec4[T]) {
	result[0] = self[0]
	result[1] = self[1]
	result[2] = self[2]
	result[3] = self[3]
	if result[0] < 0 {
		result[0] = -result[0]
	}
	if result[1] < 0 {
		result[1] = -result[1]
	}
	if result[2] < 0 {
		result[2] = -result[2]
	}
	if result[3] < 0 {
		result[3] = -result[3]
	}
	return result
}

func (self Vec2[T]) Neg() (result Vec2[T]) {
	result[0] = -self[0]
	result[1] = -self[1]
	return result
}
func (self Vec3[T]) Neg() (result Vec3[T]) {
	result[0] = -self[0]
	result[1] = -self[1]
	result[2] = -self[2]
	return result
}
func (self Vec4[T]) Neg() (result Vec4[T]) {
	result[0] = -self[0]
	result[1] = -self[1]
	result[2] = -self[2]
	result[3] = -self[3]
	return result
}

func (self Vec2[T]) Inv() (result Vec2[T]) {
	result[0] = 1 / self[0]
	result[1] = 1 / self[1]
	return result
}
func (self Vec3[T]) Inv() (result Vec3[T]) {
	result[0] = 1 / self[0]
	result[1] = 1 / self[1]
	result[2] = 1 / self[2]
	return result
}
func (self Vec4[T]) Inv() (result Vec4[T]) {
	result[0] = 1 / self[0]
	result[1] = 1 / self[1]
	result[2] = 1 / self[2]
	result[3] = 1 / self[3]
	return result
}

func (self Vec2[T]) Dot(other Vec2[T]) T {
	return (self[0] * other[0]) + (self[1] * other[1])
}
func (self Vec3[T]) Dot(other Vec3[T]) T {
	return (self[0] * other[0]) + (self[1] * other[1]) + (self[2] * other[2])
}
func (self Vec4[T]) Dot(other Vec4[T]) T {
	return (self[0] * other[0]) + (self[1] * other[1]) + (self[2] * other[2]) + (self[3] * other[3])
}

func (self Vec2[T]) Cross(other Vec2[T]) T {
	return self.Detr(other)
}
func (self Vec3[T]) Cross(other Vec3[T]) (result Vec3[T]) {
	result[0] = (self[1] * other[2]) - (self[2] * other[1])
	result[1] = (self[2] * other[0]) - (self[0] * other[2])
	result[2] = (self[0] * other[1]) - (self[1] * other[0])
	return result
}

func (self Vec2[T]) Detr(other Vec2[T]) T {
	return (self[0] * other[1]) - (self[1] * other[0])
}
func (self Vec3[T]) Detr(other Vec3[T], third Vec3[T]) T {
	return (self[0] * other[1] * third[2]) + (self[1] * other[2] * third[0]) + (self[2] * other[0] * third[1]) -
		(self[2] * other[1] * third[0]) - (self[1] * other[0] * third[2]) - (self[0] * other[2] * third[1])
}

func (self Vec2[T]) Len() T {
	sq := (self[0] * self[0]) + (self[1] * self[1])
	return Root(sq, 2)
}
func (self Vec3[T]) Len() T {
	sq := (self[0] * self[0]) + (self[1] * self[1]) + (self[2] * self[2])
	return Root(sq, 2)
}
func (self Vec4[T]) Len() T {
	sq := (self[0] * self[0]) + (self[1] * self[1]) + (self[2] * self[2]) + (self[3] * self[3])
	return Root(sq, 2)
}

func (self Vec2[T]) Norm() (result Vec2[T]) {
	l := self.Len()
	result[0] = self[0] / l
	result[1] = self[1] / l
	return result
}
func (self Vec3[T]) Norm() (result Vec3[T]) {
	l := self.Len()
	result[0] = self[0] / l
	result[1] = self[1] / l
	result[2] = self[2] / l
	return result
}
func (self Vec4[T]) Norm() (result Vec4[T]) {
	l := self.Len()
	result[0] = self[0] / l
	result[1] = self[1] / l
	result[2] = self[2] / l
	result[3] = self[3] / l
	return result
}

func (self Vec2[T]) Perp() (resultCW Vec2[T], resultCCW Vec2[T]) {
	resultCCW[0] = -self[1]
	resultCCW[1] = self[0]
	resultCW[0] = self[1]
	resultCW[1] = -self[0]
	return resultCW, resultCCW
}
func (self Vec3[T]) Perp(other Vec3[T]) (resultRH Vec3[T], resultLH Vec3[T]) {
	resultRH = self.Cross(other)
	resultLH = resultRH.Neg()
	return resultRH, resultLH
}

func (self Vec2[T]) LerpSelf(amount float64) (result Vec2[T]) {
	result[0] = T(float64(self[0]) * amount)
	result[1] = T(float64(self[1]) * amount)
	return result
}
func (self Vec3[T]) LerpSelf(amount float64) (result Vec3[T]) {
	result[0] = T(float64(self[0]) * amount)
	result[1] = T(float64(self[1]) * amount)
	result[2] = T(float64(self[2]) * amount)
	return result
}
func (self Vec4[T]) LerpSelf(amount float64) (result Vec4[T]) {
	result[0] = T(float64(self[0]) * amount)
	result[1] = T(float64(self[1]) * amount)
	result[2] = T(float64(self[2]) * amount)
	result[3] = T(float64(self[3]) * amount)
	return result
}

func (self Vec2[T]) LerpOther(other Vec2[T], amount float64) (result Vec2[T]) {
	return other.Sub(self).LerpSelf(amount).Add(self)
}
func (self Vec3[T]) LerpOther(other Vec3[T], amount float64) (result Vec3[T]) {
	return other.Sub(self).LerpSelf(amount).Add(self)
}
func (self Vec4[T]) LerpOther(other Vec4[T], amount float64) (result Vec4[T]) {
	return other.Sub(self).LerpSelf(amount).Add(self)
}

func (self Vec2[T]) Rotate(degrees float64) (result Vec2[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	return self.rotInternal(sin, cos)
}

func (self Vec3[T]) RotateX(degrees float64) (result Vec3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	return self.rotInternalX(sin, cos)
}
func (self Vec3[T]) RotateY(degrees float64) (result Vec3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	return self.rotInternalY(sin, cos)
}
func (self Vec3[T]) RotateZ(degrees float64) (result Vec3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	return self.rotInternalZ(sin, cos)
}

func (self Vec3[T]) AsVec2() (result Vec2[T]) {
	result[0] = self[0]
	result[1] = self[1]
	return result
}
func (self Vec4[T]) AsVec2() (result Vec2[T]) {
	result[0] = self[0]
	result[1] = self[1]
	return result
}

func (self Vec2[T]) AsVec3() (result Vec3[T]) {
	result[0] = self[0]
	result[1] = self[1]
	return result
}
func (self Vec4[T]) AsVec3() (result Vec3[T]) {
	result[0] = self[0]
	result[1] = self[1]
	result[2] = self[2]
	return result
}

func (self Vec2[T]) AsVec4() (result Vec4[T]) {
	result[0] = self[0]
	result[1] = self[1]
	return result
}
func (self Vec3[T]) AsVec4() (result Vec4[T]) {
	result[0] = self[0]
	result[1] = self[1]
	result[2] = self[2]
	return result
}

func (self *Vec2[T]) Ptr() *T {
	return &self[0]
}
func (self *Vec3[T]) Ptr() *T {
	return &self[0]
}
func (self *Vec4[T]) Ptr() *T {
	return &self[0]
}

func (self Vec2[T]) Equals(other Vec2[T]) bool {
	return self[0] == other[0] && self[1] == other[1]
}
func (self Vec3[T]) Equals(other Vec3[T]) bool {
	return self[0] == other[0] && self[1] == other[1] && self[2] == other[2]
}
func (self Vec4[T]) Equals(other Vec4[T]) bool {
	return self[0] == other[0] && self[1] == other[1] && self[2] == other[2] && self[3] == other[3]
}

func (self Line2[T]) A() (result Vec2[T]) {
	return self[0]
}
func (self Line2[T]) B() (result Vec2[T]) {
	return self[1]
}
func (self Line2[T]) Delta() (result Vec2[T]) {
	return self[1].Sub(self[0])
}
func (self Line2[T]) Len() T {
	return self.Delta().Len()
}
func (self Line2[T]) Lerp(val float64) (result Vec2[T]) {
	delta := self.Delta()
	dx, dy := float64(delta.X())*val, float64(delta.Y())*val
	return self.A().Add(Vec2[T]{T(dx), T(dy)})
}
func (self Line2[T]) PerpLines(offset T) (left Line2[T], right Line2[T]) {
	mult := Vec2[T]{offset, offset}
	delta := self.Delta()
	norm := delta.Norm()
	perpCW, perpCCW := norm.Mult(mult).Perp()
	left = Line2[T]{self.A().Add(perpCCW), self.B().Add(perpCCW)}
	right = Line2[T]{self.A().Add(perpCW), self.B().Add(perpCW)}
	return left, right
}
func (self Line2[T]) IntersectionInfinite(other Line2[T]) (Vec2[T], bool) {
	xdiff := Vec2[T]{self[0][0] - self[1][0], other[0][0] - other[1][0]}
	ydiff := Vec2[T]{self[0][1] - self[1][1], other[0][1] - other[1][1]}
	div := xdiff.Detr(ydiff)
	if div == 0 {
		return Vec2[T]{}, false
	}
	d := Vec2[T]{self.A().Detr(self.B()), other.A().Detr(other.B())}
	x := d.Detr(xdiff) / div
	y := d.Detr(ydiff) / div
	return Vec2[T]{x, y}, true
}
func (self Line2[T]) IntersectionFinite(other Line2[T]) (Vec2[T], bool) {
	aStart := self.A()
	aDelta := self.Delta()
	bStart := other.A()
	bDelta := other.Delta()
	denom := aDelta.Cross(bDelta)
	numer := bStart.Sub(aStart).Cross(aDelta)
	if denom == 0 {
		if numer == 0 {
			var aLeft, aRight, bLeft, bRight Vec2[T]
			if self.A().X() < self.B().X() {
				aLeft = self.A()
				aRight = self.B()
			} else {
				aLeft = self.B()
				aRight = self.A()
			}
			if other.A().X() < other.B().X() {
				bLeft = other.A()
				bRight = other.B()
			} else {
				bLeft = other.B()
				bRight = other.A()
			}
			if aLeft.X() >= bLeft.X() && aLeft.X() <= bRight.X() {
				return aLeft, true
			} else if aRight.X() >= bLeft.X() && aRight.X() <= bRight.X() {
				return aRight, true
			} else {
				return Vec2[T]{}, false
			}
		} else {
			return Vec2[T]{}, false
		}
	}
	aLerp := bStart.Sub(aStart).Cross(bDelta) / denom
	bLerp := bStart.Sub(aStart).Cross(aDelta) / denom
	if aLerp >= 0 && aLerp <= 1 && bLerp >= 0 && bLerp <= 1 {
		intersection := self.A().LerpOther(self.B(), float64(aLerp))
		return intersection, true
	}
	return Vec2[T]{}, false
}
func (self Line2[T]) Translate(val Vec2[T]) (result Line2[T]) {
	result[0] = self[0].Add(val)
	result[1] = self[1].Add(val)
	return result
}
func (self Line2[T]) Scale(val T) (result Line2[T]) {
	result[0] = self[0].Scale(val)
	result[1] = self[1].Scale(val)
	return result
}
func (self Line2[T]) Mult(val Vec2[T]) (result Line2[T]) {
	result[0] = self[0].Mult(val)
	result[1] = self[1].Mult(val)
	return result
}
func (self Line2[T]) Rotate(degrees float64) (result Line2[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternal(sin, cos)
	result[1] = self[1].rotInternal(sin, cos)
	return result
}
func (self Line2[T]) Reverse() Line2[T] {
	return Line2[T]{self.B(), self.A()}
}

func (self Tri2[T]) A() Vec2[T] {
	return self[0]
}
func (self Tri2[T]) B() Vec2[T] {
	return self[1]
}
func (self Tri2[T]) C() Vec2[T] {
	return self[2]
}
func (self Tri2[T]) AB() Line2[T] {
	return Line2[T]{self.A(), self.B()}
}
func (self Tri2[T]) BA() Line2[T] {
	return Line2[T]{self.B(), self.A()}
}
func (self Tri2[T]) BC() Line2[T] {
	return Line2[T]{self.B(), self.C()}
}
func (self Tri2[T]) CB() Line2[T] {
	return Line2[T]{self.C(), self.B()}
}
func (self Tri2[T]) CA() Line2[T] {
	return Line2[T]{self.C(), self.A()}
}
func (self Tri2[T]) AC() Line2[T] {
	return Line2[T]{self.A(), self.C()}
}
func (self Tri2[T]) Translate(val Vec2[T]) (result Tri2[T]) {
	result[0] = self[0].Add(val)
	result[1] = self[1].Add(val)
	result[2] = self[2].Add(val)
	return result
}
func (self Tri2[T]) Scale(val T) (result Tri2[T]) {
	result[0] = self[0].Scale(val)
	result[1] = self[1].Scale(val)
	result[2] = self[2].Scale(val)
	return result
}
func (self Tri2[T]) Mult(val Vec2[T]) (result Tri2[T]) {
	result[0] = self[0].Mult(val)
	result[1] = self[1].Mult(val)
	result[2] = self[2].Mult(val)
	return result
}
func (self Tri2[T]) Rotate(degrees float64) (result Tri2[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternal(sin, cos)
	result[1] = self[1].rotInternal(sin, cos)
	result[2] = self[2].rotInternal(sin, cos)
	return result
}
func (self Tri2[T]) Area() T {
	return Abs(
		/**/ ((self[1].X()*self[0].Y())-(self[0].X()*self[1].Y()))+
			((self[2].X()*self[1].Y())-(self[1].X()*self[2].Y()))+
			((self[0].X()*self[2].Y())-(self[2].X()*self[0].Y()))) / 2
}
func (self Tri2[T]) Reverse() Tri2[T] {
	return Tri2[T]{self.C(), self.B(), self.A()}
}

func (self Quad2[T]) A() (result Vec2[T]) {
	return self[0]
}
func (self Quad2[T]) B() (result Vec2[T]) {
	return self[1]
}
func (self Quad2[T]) C() (result Vec2[T]) {
	return self[2]
}
func (self Quad2[T]) D() (result Vec2[T]) {
	return self[3]
}
func (self Quad2[T]) AB() Line2[T] {
	return Line2[T]{self.A(), self.B()}
}
func (self Quad2[T]) BA() Line2[T] {
	return Line2[T]{self.B(), self.A()}
}
func (self Quad2[T]) BC() Line2[T] {
	return Line2[T]{self.B(), self.C()}
}
func (self Quad2[T]) CB() Line2[T] {
	return Line2[T]{self.C(), self.B()}
}
func (self Quad2[T]) CD() Line2[T] {
	return Line2[T]{self.C(), self.D()}
}
func (self Quad2[T]) DC() Line2[T] {
	return Line2[T]{self.D(), self.C()}
}
func (self Quad2[T]) DA() Line2[T] {
	return Line2[T]{self.D(), self.A()}
}
func (self Quad2[T]) AD() Line2[T] {
	return Line2[T]{self.A(), self.D()}
}
func (self Quad2[T]) AC() Line2[T] {
	return Line2[T]{self.A(), self.C()}
}
func (self Quad2[T]) CA() Line2[T] {
	return Line2[T]{self.C(), self.A()}
}
func (self Quad2[T]) DB() Line2[T] {
	return Line2[T]{self.D(), self.B()}
}
func (self Quad2[T]) BD() Line2[T] {
	return Line2[T]{self.B(), self.D()}
}
func (self Quad2[T]) ABC() (result Tri2[T]) {
	return Tri2[T]{self.A(), self.B(), self.D()}
}
func (self Quad2[T]) BCD() (result Tri2[T]) {
	return Tri2[T]{self.B(), self.C(), self.D()}
}
func (self Quad2[T]) CDA() (result Tri2[T]) {
	return Tri2[T]{self.C(), self.D(), self.A()}
}
func (self Quad2[T]) DAB() (result Tri2[T]) {
	return Tri2[T]{self.D(), self.A(), self.B()}
}
func (self Quad2[T]) Area() T {
	return self.ABC().Area() + self.CDA().Area()
}
func (self Quad2[T]) MapLerp(xLerp float64, yLerp float32) Vec2[T] {
	h1 := self.AB().Lerp(xLerp)
	h2 := self.DC().Lerp(xLerp)
	v := h1.LerpOther(h2, float64(yLerp))
	return v
}
func (self Quad2[T]) Translate(val Vec2[T]) (result Quad2[T]) {
	result[0] = self[0].Add(val)
	result[1] = self[1].Add(val)
	result[2] = self[2].Add(val)
	result[3] = self[3].Add(val)
	return result
}
func (self Quad2[T]) Scale(val T) (result Quad2[T]) {
	result[0] = self[0].Scale(val)
	result[1] = self[1].Scale(val)
	result[2] = self[2].Scale(val)
	result[3] = self[3].Scale(val)
	return result
}
func (self Quad2[T]) Mult(val Vec2[T]) (result Quad2[T]) {
	result[0] = self[0].Mult(val)
	result[1] = self[1].Mult(val)
	result[2] = self[2].Mult(val)
	result[3] = self[3].Mult(val)
	return result
}
func (self Quad2[T]) Rotate(degrees float64) (result Quad2[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternal(sin, cos)
	result[1] = self[1].rotInternal(sin, cos)
	result[2] = self[2].rotInternal(sin, cos)
	result[3] = self[3].rotInternal(sin, cos)
	return result
}

func NewRect2[T Real](pos Vec2[T], size Vec2[T]) AABB2[T] {
	return AABB2[T]{pos, pos.Add(size)}
}
func (self AABB2[T]) X() T {
	return self[0].X()
}
func (self AABB2[T]) Y() T {
	return self[0].Y()
}
func (self AABB2[T]) W() T {
	return self[1].X() - self[0].X()
}
func (self AABB2[T]) H() T {
	return self[1].Y() - self[0].Y()
}
func (self AABB2[T]) Min() Vec2[T] {
	return self[0]
}
func (self AABB2[T]) Max() Vec2[T] {
	return self[1]
}
func (self AABB2[T]) Translate(val Vec2[T]) (result AABB2[T]) {
	result[0] = self[0].Add(val)
	result[1] = self[1].Add(val)
	return result
}
func (self AABB2[T]) Expand(val Vec2[T]) (result AABB2[T]) {
	result[0] = self[0]
	result[1] = self[1].Add(val)
	return result
}
func (self AABB2[T]) Scale(val T) (result AABB2[T]) {
	result[0] = self[0].Scale(val)
	result[1] = self[1].Scale(val)
	return result
}
func (self AABB2[T]) Mult(val Vec2[T]) (result AABB2[T]) {
	result[0] = self[0].Mult(val)
	result[1] = self[1].Mult(val)
	return result
}
func (self AABB2[T]) Area() T {
	return self.W() * self.H()
}
func (self AABB2[T]) Quad() Quad2[T] {
	return Quad2[T]{self.Min(), Vec2[T]{self.Max().X(), self.Min().Y()}, self.Max(), Vec2[T]{self.Min().X(), self.Max().Y()}}
}

func (self Array2[T]) Translate(val Vec2[T]) (result Array2[T]) {
	result = make([]Vec2[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].Add(val)
	}
	return result
}
func (self Array2[T]) Scale(val T) (result Array2[T]) {
	result = make([]Vec2[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].Scale(val)
	}
	return result
}
func (self Array2[T]) Mult(val Vec2[T]) (result Array2[T]) {
	result = make([]Vec2[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].Mult(val)
	}
	return result
}
func (self Array2[T]) Rotate(degrees float64) (result Array2[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result = make([]Vec2[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].rotInternal(sin, cos)
	}
	return result
}

func (self Line3[T]) Translate(val Vec3[T]) (result Line3[T]) {
	result[0] = self[0].Add(val)
	result[1] = self[1].Add(val)
	return result
}
func (self Line3[T]) Scale(val T) (result Line3[T]) {
	result[0] = self[0].Scale(val)
	result[1] = self[1].Scale(val)
	return result
}
func (self Line3[T]) Mult(val Vec3[T]) (result Line3[T]) {
	result[0] = self[0].Mult(val)
	result[1] = self[1].Mult(val)
	return result
}
func (self Line3[T]) RotateX(degrees float64) (result Line3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalX(sin, cos)
	result[1] = self[1].rotInternalX(sin, cos)
	return result
}
func (self Line3[T]) RotateY(degrees float64) (result Line3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalY(sin, cos)
	result[1] = self[1].rotInternalY(sin, cos)
	return result
}
func (self Line3[T]) RotateZ(degrees float64) (result Line3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalZ(sin, cos)
	result[1] = self[1].rotInternalZ(sin, cos)
	return result
}

func (self Tri3[T]) Translate(val Vec3[T]) (result Tri3[T]) {
	result[0] = self[0].Add(val)
	result[1] = self[1].Add(val)
	result[2] = self[2].Add(val)
	return result
}
func (self Tri3[T]) Scale(val T) (result Tri3[T]) {
	result[0] = self[0].Scale(val)
	result[1] = self[1].Scale(val)
	result[2] = self[2].Scale(val)
	return result
}
func (self Tri3[T]) Mult(val Vec3[T]) (result Tri3[T]) {
	result[0] = self[0].Mult(val)
	result[1] = self[1].Mult(val)
	result[2] = self[2].Mult(val)
	return result
}
func (self Tri3[T]) RotateX(degrees float64) (result Tri3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalX(sin, cos)
	result[1] = self[1].rotInternalX(sin, cos)
	result[2] = self[2].rotInternalX(sin, cos)
	return result
}
func (self Tri3[T]) RotateY(degrees float64) (result Tri3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalY(sin, cos)
	result[1] = self[1].rotInternalY(sin, cos)
	result[2] = self[2].rotInternalY(sin, cos)
	return result
}
func (self Tri3[T]) RotateZ(degrees float64) (result Tri3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalZ(sin, cos)
	result[1] = self[1].rotInternalZ(sin, cos)
	result[2] = self[2].rotInternalZ(sin, cos)
	return result
}
func (self Quad3[T]) Translate(val Vec3[T]) (result Quad3[T]) {
	result[0] = self[0].Add(val)
	result[1] = self[1].Add(val)
	result[2] = self[2].Add(val)
	result[3] = self[3].Add(val)
	return result
}
func (self Quad3[T]) Scale(val T) (result Quad3[T]) {
	result[0] = self[0].Scale(val)
	result[1] = self[1].Scale(val)
	result[2] = self[2].Scale(val)
	result[3] = self[3].Scale(val)
	return result
}
func (self Quad3[T]) Mult(val Vec3[T]) (result Quad3[T]) {
	result[0] = self[0].Mult(val)
	result[1] = self[1].Mult(val)
	result[2] = self[2].Mult(val)
	result[3] = self[3].Mult(val)
	return result
}
func (self Quad3[T]) RotateX(degrees float64) (result Quad3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalX(sin, cos)
	result[1] = self[1].rotInternalX(sin, cos)
	result[2] = self[2].rotInternalX(sin, cos)
	result[3] = self[3].rotInternalX(sin, cos)
	return result
}
func (self Quad3[T]) RotateY(degrees float64) (result Quad3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalY(sin, cos)
	result[1] = self[1].rotInternalY(sin, cos)
	result[2] = self[2].rotInternalY(sin, cos)
	result[3] = self[3].rotInternalY(sin, cos)
	return result
}
func (self Quad3[T]) RotateZ(degrees float64) (result Quad3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result[0] = self[0].rotInternalZ(sin, cos)
	result[1] = self[1].rotInternalZ(sin, cos)
	result[2] = self[2].rotInternalZ(sin, cos)
	result[3] = self[3].rotInternalZ(sin, cos)
	return result
}

func (self Array3[T]) Translate(val Vec3[T]) (result Array3[T]) {
	result = make([]Vec3[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].Add(val)
	}
	return result
}
func (self Array3[T]) Scale(val T) (result Array3[T]) {
	result = make([]Vec3[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].Scale(val)
	}
	return result
}
func (self Array3[T]) Mult(val Vec3[T]) (result Array3[T]) {
	result = make([]Vec3[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].Mult(val)
	}
	return result
}
func (self Array3[T]) RotateX(degrees float64) (result Array3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result = make([]Vec3[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].rotInternalX(sin, cos)
	}
	return result
}
func (self Array3[T]) RotateY(degrees float64) (result Array3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result = make([]Vec3[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].rotInternalY(sin, cos)
	}
	return result
}
func (self Array3[T]) RotateZ(degrees float64) (result Array3[T]) {
	cos := CosDeg(degrees)
	sin := SinDeg(degrees)
	result = make([]Vec3[T], len(self))
	for i := 0; i < len(self); i += 1 {
		result[i] = self[i].rotInternalZ(sin, cos)
	}
	return result
}

func (self Vec2[T]) rotInternal(sin float64, cos float64) (result Vec2[T]) {
	result[0] = T((float64(self[0]) * cos) - (float64(self[1]) * sin))
	result[1] = T((float64(self[0]) * sin) + (float64(self[1]) * cos))
	return self
}

func (self Vec3[T]) rotInternalX(sin float64, cos float64) (result Vec3[T]) {
	result[0] = self[0]
	result[1] = T((float64(self[1]) * cos) - (float64(self[2]) * sin))
	result[2] = T((float64(self[1]) * sin) + (float64(self[2]) * cos))
	return self
}
func (self Vec3[T]) rotInternalY(sin float64, cos float64) (result Vec3[T]) {
	result[0] = T((float64(self[0]) * cos) + (float64(self[2]) * sin))
	result[1] = self[1]
	result[2] = T((float64(-self[0]) * sin) + (float64(self[2]) * cos))
	return self
}
func (self Vec3[T]) rotInternalZ(sin float64, cos float64) (result Vec3[T]) {
	result[0] = T((float64(self[0]) * cos) - (float64(self[1]) * sin))
	result[1] = T((float64(self[0]) * sin) + (float64(self[1]) * cos))
	result[2] = self[2]
	return self
}

func (self VecN[T]) Vec2() Vec2[T] {
	return (Vec2[T]{0, 0}).Swizzle(self...)
}
func (self VecN[T]) Vec3() Vec3[T] {
	return (Vec3[T]{0, 0, 0}).Swizzle(self...)
}
func (self VecN[T]) Vec4() Vec4[T] {
	return (Vec4[T]{0, 0, 0, 0}).Swizzle(self...)
}
func (self Vec2[T]) VecN() VecN[T] {
	return VecN[T]{self[0], self[1]}
}
func (self Vec3[T]) VecN() VecN[T] {
	return VecN[T]{self[0], self[1], self[2]}
}
func (self Vec4[T]) VecN() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[3]}
}

func (self Vec2[T]) XX() VecN[T] {
	return VecN[T]{self[0], self[0]}
}
func (self Vec2[T]) XY() VecN[T] {
	return VecN[T]{self[0], self[1]}
}
func (self Vec2[T]) YY() VecN[T] {
	return VecN[T]{self[1], self[1]}
}
func (self Vec2[T]) YX() VecN[T] {
	return VecN[T]{self[1], self[0]}
}
func (self Vec2[T]) XXX() VecN[T] {
	return VecN[T]{self[0], self[0], self[0]}
}
func (self Vec2[T]) XXY() VecN[T] {
	return VecN[T]{self[0], self[0], self[1]}
}
func (self Vec2[T]) XYX() VecN[T] {
	return VecN[T]{self[0], self[1], self[0]}
}
func (self Vec2[T]) XYY() VecN[T] {
	return VecN[T]{self[0], self[1], self[1]}
}
func (self Vec2[T]) YXX() VecN[T] {
	return VecN[T]{self[1], self[0], self[0]}
}
func (self Vec2[T]) YXY() VecN[T] {
	return VecN[T]{self[1], self[0], self[1]}
}
func (self Vec2[T]) YYX() VecN[T] {
	return VecN[T]{self[1], self[1], self[0]}
}
func (self Vec2[T]) YYY() VecN[T] {
	return VecN[T]{self[1], self[1], self[1]}
}
func (self Vec2[T]) XXXX() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[0]}
}
func (self Vec2[T]) XXXY() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[1]}
}
func (self Vec2[T]) XXYX() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[0]}
}
func (self Vec2[T]) XXYY() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[1]}
}
func (self Vec2[T]) XYXX() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[0]}
}
func (self Vec2[T]) XYXY() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[1]}
}
func (self Vec2[T]) XYYX() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[0]}
}
func (self Vec2[T]) XYYY() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[1]}
}
func (self Vec2[T]) YXXX() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[0]}
}
func (self Vec2[T]) YXXY() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[1]}
}
func (self Vec2[T]) YXYX() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[1]}
}
func (self Vec2[T]) YXYY() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[1]}
}
func (self Vec2[T]) YYXX() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[0]}
}
func (self Vec2[T]) YYXY() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[1]}
}
func (self Vec2[T]) YYYX() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[0]}
}
func (self Vec2[T]) YYYY() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[1]}
}
func (self Vec3[T]) XX() VecN[T] {
	return VecN[T]{self[0], self[0]}
}
func (self Vec3[T]) XY() VecN[T] {
	return VecN[T]{self[0], self[1]}
}
func (self Vec3[T]) XZ() VecN[T] {
	return VecN[T]{self[0], self[2]}
}
func (self Vec3[T]) YX() VecN[T] {
	return VecN[T]{self[1], self[0]}
}
func (self Vec3[T]) YY() VecN[T] {
	return VecN[T]{self[1], self[1]}
}
func (self Vec3[T]) YZ() VecN[T] {
	return VecN[T]{self[1], self[2]}
}
func (self Vec3[T]) ZX() VecN[T] {
	return VecN[T]{self[2], self[0]}
}
func (self Vec3[T]) ZY() VecN[T] {
	return VecN[T]{self[2], self[1]}
}
func (self Vec3[T]) ZZ() VecN[T] {
	return VecN[T]{self[2], self[2]}
}
func (self Vec3[T]) XXX() VecN[T] {
	return VecN[T]{self[0], self[0], self[0]}
}
func (self Vec3[T]) XXY() VecN[T] {
	return VecN[T]{self[0], self[0], self[1]}
}
func (self Vec3[T]) XXZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[2]}
}
func (self Vec3[T]) XYX() VecN[T] {
	return VecN[T]{self[0], self[1], self[0]}
}
func (self Vec3[T]) XYY() VecN[T] {
	return VecN[T]{self[0], self[1], self[1]}
}
func (self Vec3[T]) XYZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[2]}
}
func (self Vec3[T]) XZX() VecN[T] {
	return VecN[T]{self[0], self[2], self[0]}
}
func (self Vec3[T]) XZY() VecN[T] {
	return VecN[T]{self[0], self[2], self[1]}
}
func (self Vec3[T]) XZZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[2]}
}
func (self Vec3[T]) YXX() VecN[T] {
	return VecN[T]{self[1], self[0], self[0]}
}
func (self Vec3[T]) YXY() VecN[T] {
	return VecN[T]{self[1], self[0], self[1]}
}
func (self Vec3[T]) YXZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[2]}
}
func (self Vec3[T]) YYX() VecN[T] {
	return VecN[T]{self[1], self[1], self[0]}
}
func (self Vec3[T]) YYY() VecN[T] {
	return VecN[T]{self[1], self[1], self[1]}
}
func (self Vec3[T]) YYZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[2]}
}
func (self Vec3[T]) YZX() VecN[T] {
	return VecN[T]{self[1], self[2], self[0]}
}
func (self Vec3[T]) YZY() VecN[T] {
	return VecN[T]{self[1], self[2], self[1]}
}
func (self Vec3[T]) YZZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[2]}
}
func (self Vec3[T]) ZXX() VecN[T] {
	return VecN[T]{self[2], self[0], self[0]}
}
func (self Vec3[T]) ZXY() VecN[T] {
	return VecN[T]{self[2], self[0], self[1]}
}
func (self Vec3[T]) ZXZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[2]}
}
func (self Vec3[T]) ZYX() VecN[T] {
	return VecN[T]{self[2], self[1], self[0]}
}
func (self Vec3[T]) ZYY() VecN[T] {
	return VecN[T]{self[2], self[1], self[1]}
}
func (self Vec3[T]) ZYZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[2]}
}
func (self Vec3[T]) ZZX() VecN[T] {
	return VecN[T]{self[2], self[2], self[0]}
}
func (self Vec3[T]) ZZY() VecN[T] {
	return VecN[T]{self[2], self[2], self[1]}
}
func (self Vec3[T]) ZZZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[2]}
}
func (self Vec3[T]) XXXX() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[0]}
}
func (self Vec3[T]) XXXY() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[1]}
}
func (self Vec3[T]) XXXZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[2]}
}
func (self Vec3[T]) XXYX() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[0]}
}
func (self Vec3[T]) XXYY() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[1]}
}
func (self Vec3[T]) XXYZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[2]}
}
func (self Vec3[T]) XXZX() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[0]}
}
func (self Vec3[T]) XXZY() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[1]}
}
func (self Vec3[T]) XXZZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[2]}
}
func (self Vec3[T]) XYXX() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[0]}
}
func (self Vec3[T]) XYXY() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[1]}
}
func (self Vec3[T]) XYXZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[2]}
}
func (self Vec3[T]) XYYX() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[0]}
}
func (self Vec3[T]) XYYY() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[1]}
}
func (self Vec3[T]) XYYZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[2]}
}
func (self Vec3[T]) XYZX() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[0]}
}
func (self Vec3[T]) XYZY() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[1]}
}
func (self Vec3[T]) XYZZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[2]}
}
func (self Vec3[T]) XZXX() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[0]}
}
func (self Vec3[T]) XZXY() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[1]}
}
func (self Vec3[T]) XZXZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[2]}
}
func (self Vec3[T]) XZYX() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[0]}
}
func (self Vec3[T]) XZYY() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[1]}
}
func (self Vec3[T]) XZYZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[2]}
}
func (self Vec3[T]) XZZX() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[0]}
}
func (self Vec3[T]) XZZY() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[1]}
}
func (self Vec3[T]) XZZZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[2]}
}
func (self Vec3[T]) YXXX() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[0]}
}
func (self Vec3[T]) YXXY() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[1]}
}
func (self Vec3[T]) YXXZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[2]}
}
func (self Vec3[T]) YXYX() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[0]}
}
func (self Vec3[T]) YXYY() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[1]}
}
func (self Vec3[T]) YXYZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[2]}
}
func (self Vec3[T]) YXZX() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[0]}
}
func (self Vec3[T]) YXZY() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[1]}
}
func (self Vec3[T]) YXZZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[2]}
}
func (self Vec3[T]) YYXX() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[0]}
}
func (self Vec3[T]) YYXY() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[1]}
}
func (self Vec3[T]) YYXZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[2]}
}
func (self Vec3[T]) YYYX() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[0]}
}
func (self Vec3[T]) YYYY() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[1]}
}
func (self Vec3[T]) YYYZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[2]}
}
func (self Vec3[T]) YYZX() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[0]}
}
func (self Vec3[T]) YYZY() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[1]}
}
func (self Vec3[T]) YYZZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[2]}
}
func (self Vec3[T]) YZXX() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[0]}
}
func (self Vec3[T]) YZXY() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[1]}
}
func (self Vec3[T]) YZXZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[2]}
}
func (self Vec3[T]) YZYX() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[0]}
}
func (self Vec3[T]) YZYY() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[1]}
}
func (self Vec3[T]) YZYZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[2]}
}
func (self Vec3[T]) YZZX() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[0]}
}
func (self Vec3[T]) YZZY() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[1]}
}
func (self Vec3[T]) YZZZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[2]}
}
func (self Vec3[T]) ZXXX() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[0]}
}
func (self Vec3[T]) ZXXY() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[1]}
}
func (self Vec3[T]) ZXXZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[2]}
}
func (self Vec3[T]) ZXYX() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[0]}
}
func (self Vec3[T]) ZXYY() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[1]}
}
func (self Vec3[T]) ZXYZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[2]}
}
func (self Vec3[T]) ZXZX() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[0]}
}
func (self Vec3[T]) ZXZY() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[1]}
}
func (self Vec3[T]) ZXZZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[2]}
}
func (self Vec3[T]) ZYXX() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[0]}
}
func (self Vec3[T]) ZYXY() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[1]}
}
func (self Vec3[T]) ZYXZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[2]}
}
func (self Vec3[T]) ZYYX() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[0]}
}
func (self Vec3[T]) ZYYY() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[1]}
}
func (self Vec3[T]) ZYYZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[2]}
}
func (self Vec3[T]) ZYZX() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[0]}
}
func (self Vec3[T]) ZYZY() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[1]}
}
func (self Vec3[T]) ZYZZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[2]}
}
func (self Vec3[T]) ZZXX() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[0]}
}
func (self Vec3[T]) ZZXY() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[1]}
}
func (self Vec3[T]) ZZXZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[2]}
}
func (self Vec3[T]) ZZYX() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[0]}
}
func (self Vec3[T]) ZZYY() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[1]}
}
func (self Vec3[T]) ZZYZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[2]}
}
func (self Vec3[T]) ZZZX() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[0]}
}
func (self Vec3[T]) ZZZY() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[1]}
}
func (self Vec3[T]) ZZZZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[2]}
}
func (self Vec4[T]) XX() VecN[T] {
	return VecN[T]{self[0], self[0]}
}
func (self Vec4[T]) XY() VecN[T] {
	return VecN[T]{self[0], self[1]}
}
func (self Vec4[T]) XZ() VecN[T] {
	return VecN[T]{self[0], self[2]}
}
func (self Vec4[T]) XW() VecN[T] {
	return VecN[T]{self[0], self[3]}
}
func (self Vec4[T]) YX() VecN[T] {
	return VecN[T]{self[0], self[0]}
}
func (self Vec4[T]) YY() VecN[T] {
	return VecN[T]{self[1], self[1]}
}
func (self Vec4[T]) YZ() VecN[T] {
	return VecN[T]{self[1], self[2]}
}
func (self Vec4[T]) YW() VecN[T] {
	return VecN[T]{self[1], self[3]}
}
func (self Vec4[T]) ZX() VecN[T] {
	return VecN[T]{self[2], self[0]}
}
func (self Vec4[T]) ZY() VecN[T] {
	return VecN[T]{self[2], self[1]}
}
func (self Vec4[T]) ZZ() VecN[T] {
	return VecN[T]{self[2], self[2]}
}
func (self Vec4[T]) ZW() VecN[T] {
	return VecN[T]{self[2], self[3]}
}
func (self Vec4[T]) WX() VecN[T] {
	return VecN[T]{self[3], self[0]}
}
func (self Vec4[T]) WY() VecN[T] {
	return VecN[T]{self[3], self[1]}
}
func (self Vec4[T]) WZ() VecN[T] {
	return VecN[T]{self[3], self[2]}
}
func (self Vec4[T]) WW() VecN[T] {
	return VecN[T]{self[3], self[3]}
}
func (self Vec4[T]) XXX() VecN[T] {
	return VecN[T]{self[0], self[0], self[0]}
}
func (self Vec4[T]) XXY() VecN[T] {
	return VecN[T]{self[0], self[0], self[1]}
}
func (self Vec4[T]) XXZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[2]}
}
func (self Vec4[T]) XXW() VecN[T] {
	return VecN[T]{self[0], self[0], self[3]}
}
func (self Vec4[T]) XYX() VecN[T] {
	return VecN[T]{self[0], self[1], self[0]}
}
func (self Vec4[T]) XYY() VecN[T] {
	return VecN[T]{self[0], self[1], self[1]}
}
func (self Vec4[T]) XYZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[2]}
}
func (self Vec4[T]) XYW() VecN[T] {
	return VecN[T]{self[0], self[1], self[3]}
}
func (self Vec4[T]) XZX() VecN[T] {
	return VecN[T]{self[0], self[2], self[0]}
}
func (self Vec4[T]) XZY() VecN[T] {
	return VecN[T]{self[0], self[2], self[1]}
}
func (self Vec4[T]) XZZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[2]}
}
func (self Vec4[T]) XZW() VecN[T] {
	return VecN[T]{self[0], self[2], self[3]}
}
func (self Vec4[T]) XWX() VecN[T] {
	return VecN[T]{self[0], self[3], self[0]}
}
func (self Vec4[T]) XWY() VecN[T] {
	return VecN[T]{self[0], self[3], self[1]}
}
func (self Vec4[T]) XWZ() VecN[T] {
	return VecN[T]{self[0], self[3], self[2]}
}
func (self Vec4[T]) XWW() VecN[T] {
	return VecN[T]{self[0], self[3], self[3]}
}
func (self Vec4[T]) YXX() VecN[T] {
	return VecN[T]{self[1], self[0], self[0]}
}
func (self Vec4[T]) YXY() VecN[T] {
	return VecN[T]{self[1], self[0], self[1]}
}
func (self Vec4[T]) YXZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[2]}
}
func (self Vec4[T]) YXW() VecN[T] {
	return VecN[T]{self[1], self[0], self[3]}
}
func (self Vec4[T]) YYX() VecN[T] {
	return VecN[T]{self[1], self[1], self[0]}
}
func (self Vec4[T]) YYY() VecN[T] {
	return VecN[T]{self[1], self[1], self[1]}
}
func (self Vec4[T]) YYZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[2]}
}
func (self Vec4[T]) YYW() VecN[T] {
	return VecN[T]{self[1], self[1], self[3]}
}
func (self Vec4[T]) YZX() VecN[T] {
	return VecN[T]{self[1], self[2], self[0]}
}
func (self Vec4[T]) YZY() VecN[T] {
	return VecN[T]{self[1], self[2], self[1]}
}
func (self Vec4[T]) YZZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[2]}
}
func (self Vec4[T]) YZW() VecN[T] {
	return VecN[T]{self[1], self[2], self[3]}
}
func (self Vec4[T]) YWX() VecN[T] {
	return VecN[T]{self[1], self[3], self[0]}
}
func (self Vec4[T]) YWY() VecN[T] {
	return VecN[T]{self[1], self[3], self[1]}
}
func (self Vec4[T]) YWZ() VecN[T] {
	return VecN[T]{self[1], self[3], self[2]}
}
func (self Vec4[T]) YWW() VecN[T] {
	return VecN[T]{self[1], self[3], self[3]}
}
func (self Vec4[T]) ZXX() VecN[T] {
	return VecN[T]{self[2], self[0], self[0]}
}
func (self Vec4[T]) ZXY() VecN[T] {
	return VecN[T]{self[2], self[0], self[1]}
}
func (self Vec4[T]) ZXZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[2]}
}
func (self Vec4[T]) ZXW() VecN[T] {
	return VecN[T]{self[2], self[0], self[3]}
}
func (self Vec4[T]) ZYX() VecN[T] {
	return VecN[T]{self[2], self[1], self[0]}
}
func (self Vec4[T]) ZYY() VecN[T] {
	return VecN[T]{self[2], self[1], self[1]}
}
func (self Vec4[T]) ZYZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[2]}
}
func (self Vec4[T]) ZYW() VecN[T] {
	return VecN[T]{self[2], self[1], self[3]}
}
func (self Vec4[T]) ZZX() VecN[T] {
	return VecN[T]{self[2], self[2], self[0]}
}
func (self Vec4[T]) ZZY() VecN[T] {
	return VecN[T]{self[2], self[2], self[1]}
}
func (self Vec4[T]) ZZZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[2]}
}
func (self Vec4[T]) ZZW() VecN[T] {
	return VecN[T]{self[2], self[2], self[3]}
}
func (self Vec4[T]) ZWX() VecN[T] {
	return VecN[T]{self[2], self[3], self[0]}
}
func (self Vec4[T]) ZWY() VecN[T] {
	return VecN[T]{self[2], self[3], self[1]}
}
func (self Vec4[T]) ZWZ() VecN[T] {
	return VecN[T]{self[2], self[3], self[2]}
}
func (self Vec4[T]) ZWW() VecN[T] {
	return VecN[T]{self[2], self[3], self[3]}
}
func (self Vec4[T]) WXX() VecN[T] {
	return VecN[T]{self[3], self[0], self[0]}
}
func (self Vec4[T]) WXY() VecN[T] {
	return VecN[T]{self[3], self[0], self[1]}
}
func (self Vec4[T]) WXZ() VecN[T] {
	return VecN[T]{self[3], self[0], self[2]}
}
func (self Vec4[T]) WXW() VecN[T] {
	return VecN[T]{self[3], self[0], self[3]}
}
func (self Vec4[T]) WYX() VecN[T] {
	return VecN[T]{self[3], self[1], self[0]}
}
func (self Vec4[T]) WYY() VecN[T] {
	return VecN[T]{self[3], self[1], self[1]}
}
func (self Vec4[T]) WYZ() VecN[T] {
	return VecN[T]{self[3], self[1], self[2]}
}
func (self Vec4[T]) WYW() VecN[T] {
	return VecN[T]{self[3], self[1], self[3]}
}
func (self Vec4[T]) WZX() VecN[T] {
	return VecN[T]{self[3], self[2], self[0]}
}
func (self Vec4[T]) WZY() VecN[T] {
	return VecN[T]{self[3], self[2], self[1]}
}
func (self Vec4[T]) WZZ() VecN[T] {
	return VecN[T]{self[3], self[2], self[2]}
}
func (self Vec4[T]) WZW() VecN[T] {
	return VecN[T]{self[3], self[2], self[3]}
}
func (self Vec4[T]) WWX() VecN[T] {
	return VecN[T]{self[3], self[3], self[0]}
}
func (self Vec4[T]) WWY() VecN[T] {
	return VecN[T]{self[3], self[3], self[1]}
}
func (self Vec4[T]) WWZ() VecN[T] {
	return VecN[T]{self[3], self[3], self[2]}
}
func (self Vec4[T]) WWW() VecN[T] {
	return VecN[T]{self[3], self[3], self[3]}
}
func (self Vec4[T]) XXXX() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[0]}
}
func (self Vec4[T]) XXXY() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[1]}
}
func (self Vec4[T]) XXXZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[2]}
}
func (self Vec4[T]) XXXW() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[3]}
}
func (self Vec4[T]) XXYX() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[0]}
}
func (self Vec4[T]) XXYY() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[1]}
}
func (self Vec4[T]) XXYZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[2]}
}
func (self Vec4[T]) XXYW() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[3]}
}
func (self Vec4[T]) XXZX() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[0]}
}
func (self Vec4[T]) XXZY() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[1]}
}
func (self Vec4[T]) XXZZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[2]}
}
func (self Vec4[T]) XXZW() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[3]}
}
func (self Vec4[T]) XXWX() VecN[T] {
	return VecN[T]{self[0], self[0], self[3], self[0]}
}
func (self Vec4[T]) XXWY() VecN[T] {
	return VecN[T]{self[0], self[0], self[3], self[1]}
}
func (self Vec4[T]) XXWZ() VecN[T] {
	return VecN[T]{self[0], self[0], self[3], self[2]}
}
func (self Vec4[T]) XXWW() VecN[T] {
	return VecN[T]{self[0], self[0], self[3], self[3]}
}
func (self Vec4[T]) XYXX() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[0]}
}
func (self Vec4[T]) XYXY() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[1]}
}
func (self Vec4[T]) XYXZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[2]}
}
func (self Vec4[T]) XYXW() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[3]}
}
func (self Vec4[T]) XYYX() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[0]}
}
func (self Vec4[T]) XYYY() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[1]}
}
func (self Vec4[T]) XYYZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[2]}
}
func (self Vec4[T]) XYYW() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[3]}
}
func (self Vec4[T]) XYZX() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[0]}
}
func (self Vec4[T]) XYZY() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[1]}
}
func (self Vec4[T]) XYZZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[2]}
}
func (self Vec4[T]) XYZW() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[3]}
}
func (self Vec4[T]) XYWX() VecN[T] {
	return VecN[T]{self[0], self[1], self[3], self[0]}
}
func (self Vec4[T]) XYWY() VecN[T] {
	return VecN[T]{self[0], self[1], self[3], self[1]}
}
func (self Vec4[T]) XYWZ() VecN[T] {
	return VecN[T]{self[0], self[1], self[3], self[2]}
}
func (self Vec4[T]) XYWW() VecN[T] {
	return VecN[T]{self[0], self[1], self[3], self[3]}
}
func (self Vec4[T]) XZXX() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[0]}
}
func (self Vec4[T]) XZXY() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[1]}
}
func (self Vec4[T]) XZXZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[2]}
}
func (self Vec4[T]) XZXW() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[3]}
}
func (self Vec4[T]) XZYX() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[0]}
}
func (self Vec4[T]) XZYY() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[1]}
}
func (self Vec4[T]) XZYZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[2]}
}
func (self Vec4[T]) XZYW() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[3]}
}
func (self Vec4[T]) XZZX() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[0]}
}
func (self Vec4[T]) XZZY() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[1]}
}
func (self Vec4[T]) XZZZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[2]}
}
func (self Vec4[T]) XZZW() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[3]}
}
func (self Vec4[T]) XZWX() VecN[T] {
	return VecN[T]{self[0], self[2], self[3], self[0]}
}
func (self Vec4[T]) XZWY() VecN[T] {
	return VecN[T]{self[0], self[2], self[3], self[1]}
}
func (self Vec4[T]) XZWZ() VecN[T] {
	return VecN[T]{self[0], self[2], self[3], self[2]}
}
func (self Vec4[T]) XZWW() VecN[T] {
	return VecN[T]{self[0], self[2], self[3], self[3]}
}
func (self Vec4[T]) XWXX() VecN[T] {
	return VecN[T]{self[0], self[3], self[0], self[0]}
}
func (self Vec4[T]) XWXY() VecN[T] {
	return VecN[T]{self[0], self[3], self[0], self[1]}
}
func (self Vec4[T]) XWXZ() VecN[T] {
	return VecN[T]{self[0], self[3], self[0], self[2]}
}
func (self Vec4[T]) XWXW() VecN[T] {
	return VecN[T]{self[0], self[3], self[0], self[3]}
}
func (self Vec4[T]) XWYX() VecN[T] {
	return VecN[T]{self[0], self[3], self[1], self[0]}
}
func (self Vec4[T]) XWYY() VecN[T] {
	return VecN[T]{self[0], self[3], self[1], self[1]}
}
func (self Vec4[T]) XWYZ() VecN[T] {
	return VecN[T]{self[0], self[3], self[1], self[2]}
}
func (self Vec4[T]) XWYW() VecN[T] {
	return VecN[T]{self[0], self[3], self[1], self[3]}
}
func (self Vec4[T]) XWZX() VecN[T] {
	return VecN[T]{self[0], self[3], self[2], self[0]}
}
func (self Vec4[T]) XWZY() VecN[T] {
	return VecN[T]{self[0], self[3], self[2], self[1]}
}
func (self Vec4[T]) XWZZ() VecN[T] {
	return VecN[T]{self[0], self[3], self[2], self[2]}
}
func (self Vec4[T]) XWZW() VecN[T] {
	return VecN[T]{self[0], self[3], self[2], self[3]}
}
func (self Vec4[T]) XWWX() VecN[T] {
	return VecN[T]{self[0], self[3], self[3], self[0]}
}
func (self Vec4[T]) XWWY() VecN[T] {
	return VecN[T]{self[0], self[3], self[3], self[1]}
}
func (self Vec4[T]) XWWZ() VecN[T] {
	return VecN[T]{self[0], self[3], self[3], self[2]}
}
func (self Vec4[T]) XWWW() VecN[T] {
	return VecN[T]{self[0], self[3], self[3], self[3]}
}
func (self Vec4[T]) YXXX() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[0]}
}
func (self Vec4[T]) YXXY() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[1]}
}
func (self Vec4[T]) YXXZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[2]}
}
func (self Vec4[T]) YXXW() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[3]}
}
func (self Vec4[T]) YXYX() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[0]}
}
func (self Vec4[T]) YXYY() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[1]}
}
func (self Vec4[T]) YXYZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[2]}
}
func (self Vec4[T]) YXYW() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[3]}
}
func (self Vec4[T]) YXZX() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[0]}
}
func (self Vec4[T]) YXZY() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[1]}
}
func (self Vec4[T]) YXZZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[2]}
}
func (self Vec4[T]) YXZW() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[3]}
}
func (self Vec4[T]) YXWX() VecN[T] {
	return VecN[T]{self[1], self[0], self[3], self[0]}
}
func (self Vec4[T]) YXWY() VecN[T] {
	return VecN[T]{self[1], self[0], self[3], self[1]}
}
func (self Vec4[T]) YXWZ() VecN[T] {
	return VecN[T]{self[1], self[0], self[3], self[2]}
}
func (self Vec4[T]) YXWW() VecN[T] {
	return VecN[T]{self[1], self[0], self[3], self[3]}
}
func (self Vec4[T]) YYXX() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[0]}
}
func (self Vec4[T]) YYXY() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[1]}
}
func (self Vec4[T]) YYXZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[2]}
}
func (self Vec4[T]) YYXW() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[3]}
}
func (self Vec4[T]) YYYX() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[0]}
}
func (self Vec4[T]) YYYY() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[1]}
}
func (self Vec4[T]) YYYZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[2]}
}
func (self Vec4[T]) YYYW() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[3]}
}
func (self Vec4[T]) YYZX() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[0]}
}
func (self Vec4[T]) YYZY() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[1]}
}
func (self Vec4[T]) YYZZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[2]}
}
func (self Vec4[T]) YYZW() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[3]}
}
func (self Vec4[T]) YYWX() VecN[T] {
	return VecN[T]{self[1], self[1], self[3], self[0]}
}
func (self Vec4[T]) YYWY() VecN[T] {
	return VecN[T]{self[1], self[1], self[3], self[1]}
}
func (self Vec4[T]) YYWZ() VecN[T] {
	return VecN[T]{self[1], self[1], self[3], self[2]}
}
func (self Vec4[T]) YYWW() VecN[T] {
	return VecN[T]{self[1], self[1], self[3], self[3]}
}
func (self Vec4[T]) YZXX() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[0]}
}
func (self Vec4[T]) YZXY() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[1]}
}
func (self Vec4[T]) YZXZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[2]}
}
func (self Vec4[T]) YZXW() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[3]}
}
func (self Vec4[T]) YZYX() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[0]}
}
func (self Vec4[T]) YZYY() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[1]}
}
func (self Vec4[T]) YZYZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[2]}
}
func (self Vec4[T]) YZYW() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[3]}
}
func (self Vec4[T]) YZZX() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[0]}
}
func (self Vec4[T]) YZZY() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[1]}
}
func (self Vec4[T]) YZZZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[2]}
}
func (self Vec4[T]) YZZW() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[3]}
}
func (self Vec4[T]) YZWX() VecN[T] {
	return VecN[T]{self[1], self[2], self[3], self[0]}
}
func (self Vec4[T]) YZWY() VecN[T] {
	return VecN[T]{self[1], self[2], self[3], self[1]}
}
func (self Vec4[T]) YZWZ() VecN[T] {
	return VecN[T]{self[1], self[2], self[3], self[2]}
}
func (self Vec4[T]) YZWW() VecN[T] {
	return VecN[T]{self[1], self[2], self[3], self[3]}
}
func (self Vec4[T]) YWXX() VecN[T] {
	return VecN[T]{self[1], self[3], self[0], self[0]}
}
func (self Vec4[T]) YWXY() VecN[T] {
	return VecN[T]{self[1], self[3], self[0], self[1]}
}
func (self Vec4[T]) YWXZ() VecN[T] {
	return VecN[T]{self[1], self[3], self[0], self[2]}
}
func (self Vec4[T]) YWXW() VecN[T] {
	return VecN[T]{self[1], self[3], self[0], self[3]}
}
func (self Vec4[T]) YWYX() VecN[T] {
	return VecN[T]{self[1], self[3], self[1], self[0]}
}
func (self Vec4[T]) YWYY() VecN[T] {
	return VecN[T]{self[1], self[3], self[1], self[1]}
}
func (self Vec4[T]) YWYZ() VecN[T] {
	return VecN[T]{self[1], self[3], self[1], self[2]}
}
func (self Vec4[T]) YWYW() VecN[T] {
	return VecN[T]{self[1], self[3], self[1], self[3]}
}
func (self Vec4[T]) YWZX() VecN[T] {
	return VecN[T]{self[1], self[3], self[2], self[0]}
}
func (self Vec4[T]) YWZY() VecN[T] {
	return VecN[T]{self[1], self[3], self[2], self[1]}
}
func (self Vec4[T]) YWZZ() VecN[T] {
	return VecN[T]{self[1], self[3], self[2], self[2]}
}
func (self Vec4[T]) YWZW() VecN[T] {
	return VecN[T]{self[1], self[3], self[2], self[3]}
}
func (self Vec4[T]) YWWX() VecN[T] {
	return VecN[T]{self[1], self[3], self[3], self[0]}
}
func (self Vec4[T]) YWWY() VecN[T] {
	return VecN[T]{self[1], self[3], self[3], self[1]}
}
func (self Vec4[T]) YWWZ() VecN[T] {
	return VecN[T]{self[1], self[3], self[3], self[2]}
}
func (self Vec4[T]) YWWW() VecN[T] {
	return VecN[T]{self[1], self[3], self[3], self[3]}
}
func (self Vec4[T]) ZXXX() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[0]}
}
func (self Vec4[T]) ZXXY() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[1]}
}
func (self Vec4[T]) ZXXZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[2]}
}
func (self Vec4[T]) ZXXW() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[3]}
}
func (self Vec4[T]) ZXYX() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[0]}
}
func (self Vec4[T]) ZXYY() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[1]}
}
func (self Vec4[T]) ZXYZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[2]}
}
func (self Vec4[T]) ZXYW() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[3]}
}
func (self Vec4[T]) ZXZX() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[0]}
}
func (self Vec4[T]) ZXZY() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[1]}
}
func (self Vec4[T]) ZXZZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[2]}
}
func (self Vec4[T]) ZXZW() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[3]}
}
func (self Vec4[T]) ZXWX() VecN[T] {
	return VecN[T]{self[2], self[0], self[3], self[0]}
}
func (self Vec4[T]) ZXWY() VecN[T] {
	return VecN[T]{self[2], self[0], self[3], self[1]}
}
func (self Vec4[T]) ZXWZ() VecN[T] {
	return VecN[T]{self[2], self[0], self[3], self[2]}
}
func (self Vec4[T]) ZXWW() VecN[T] {
	return VecN[T]{self[2], self[0], self[3], self[3]}
}
func (self Vec4[T]) ZYXX() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[0]}
}
func (self Vec4[T]) ZYXY() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[1]}
}
func (self Vec4[T]) ZYXZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[2]}
}
func (self Vec4[T]) ZYXW() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[3]}
}
func (self Vec4[T]) ZYYX() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[0]}
}
func (self Vec4[T]) ZYYY() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[1]}
}
func (self Vec4[T]) ZYYZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[2]}
}
func (self Vec4[T]) ZYYW() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[3]}
}
func (self Vec4[T]) ZYZX() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[0]}
}
func (self Vec4[T]) ZYZY() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[1]}
}
func (self Vec4[T]) ZYZZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[2]}
}
func (self Vec4[T]) ZYZW() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[3]}
}
func (self Vec4[T]) ZYWX() VecN[T] {
	return VecN[T]{self[2], self[1], self[3], self[0]}
}
func (self Vec4[T]) ZYWY() VecN[T] {
	return VecN[T]{self[2], self[1], self[3], self[1]}
}
func (self Vec4[T]) ZYWZ() VecN[T] {
	return VecN[T]{self[2], self[1], self[3], self[2]}
}
func (self Vec4[T]) ZYWW() VecN[T] {
	return VecN[T]{self[2], self[1], self[3], self[3]}
}
func (self Vec4[T]) ZZXX() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[0]}
}
func (self Vec4[T]) ZZXY() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[1]}
}
func (self Vec4[T]) ZZXZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[2]}
}
func (self Vec4[T]) ZZXW() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[3]}
}
func (self Vec4[T]) ZZYX() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[0]}
}
func (self Vec4[T]) ZZYY() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[1]}
}
func (self Vec4[T]) ZZYZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[2]}
}
func (self Vec4[T]) ZZYW() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[3]}
}
func (self Vec4[T]) ZZZX() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[0]}
}
func (self Vec4[T]) ZZZY() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[1]}
}
func (self Vec4[T]) ZZZZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[2]}
}
func (self Vec4[T]) ZZZW() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[3]}
}
func (self Vec4[T]) ZZWX() VecN[T] {
	return VecN[T]{self[2], self[2], self[3], self[0]}
}
func (self Vec4[T]) ZZWY() VecN[T] {
	return VecN[T]{self[2], self[2], self[3], self[1]}
}
func (self Vec4[T]) ZZWZ() VecN[T] {
	return VecN[T]{self[2], self[2], self[3], self[2]}
}
func (self Vec4[T]) ZZWW() VecN[T] {
	return VecN[T]{self[2], self[2], self[3], self[3]}
}
func (self Vec4[T]) ZWXX() VecN[T] {
	return VecN[T]{self[2], self[3], self[0], self[0]}
}
func (self Vec4[T]) ZWXY() VecN[T] {
	return VecN[T]{self[2], self[3], self[0], self[1]}
}
func (self Vec4[T]) ZWXZ() VecN[T] {
	return VecN[T]{self[2], self[3], self[0], self[2]}
}
func (self Vec4[T]) ZWXW() VecN[T] {
	return VecN[T]{self[2], self[3], self[0], self[3]}
}
func (self Vec4[T]) ZWYX() VecN[T] {
	return VecN[T]{self[2], self[3], self[1], self[0]}
}
func (self Vec4[T]) ZWYY() VecN[T] {
	return VecN[T]{self[2], self[3], self[1], self[1]}
}
func (self Vec4[T]) ZWYZ() VecN[T] {
	return VecN[T]{self[2], self[3], self[1], self[2]}
}
func (self Vec4[T]) ZWYW() VecN[T] {
	return VecN[T]{self[2], self[3], self[1], self[3]}
}
func (self Vec4[T]) ZWZX() VecN[T] {
	return VecN[T]{self[2], self[3], self[2], self[0]}
}
func (self Vec4[T]) ZWZY() VecN[T] {
	return VecN[T]{self[2], self[3], self[2], self[1]}
}
func (self Vec4[T]) ZWZZ() VecN[T] {
	return VecN[T]{self[2], self[3], self[2], self[2]}
}
func (self Vec4[T]) ZWZW() VecN[T] {
	return VecN[T]{self[2], self[3], self[2], self[3]}
}
func (self Vec4[T]) ZWWX() VecN[T] {
	return VecN[T]{self[2], self[3], self[3], self[0]}
}
func (self Vec4[T]) ZWWY() VecN[T] {
	return VecN[T]{self[2], self[3], self[3], self[1]}
}
func (self Vec4[T]) ZWWZ() VecN[T] {
	return VecN[T]{self[2], self[3], self[3], self[2]}
}
func (self Vec4[T]) ZWWW() VecN[T] {
	return VecN[T]{self[2], self[3], self[3], self[3]}
}
func (self Vec4[T]) WXXX() VecN[T] {
	return VecN[T]{self[3], self[0], self[0], self[0]}
}
func (self Vec4[T]) WXXY() VecN[T] {
	return VecN[T]{self[3], self[0], self[0], self[1]}
}
func (self Vec4[T]) WXXZ() VecN[T] {
	return VecN[T]{self[3], self[0], self[0], self[2]}
}
func (self Vec4[T]) WXXW() VecN[T] {
	return VecN[T]{self[3], self[0], self[0], self[3]}
}
func (self Vec4[T]) WXYX() VecN[T] {
	return VecN[T]{self[3], self[0], self[1], self[0]}
}
func (self Vec4[T]) WXYY() VecN[T] {
	return VecN[T]{self[3], self[0], self[1], self[1]}
}
func (self Vec4[T]) WXYZ() VecN[T] {
	return VecN[T]{self[3], self[0], self[1], self[2]}
}
func (self Vec4[T]) WXYW() VecN[T] {
	return VecN[T]{self[3], self[0], self[1], self[3]}
}
func (self Vec4[T]) WXZX() VecN[T] {
	return VecN[T]{self[3], self[0], self[2], self[0]}
}
func (self Vec4[T]) WXZY() VecN[T] {
	return VecN[T]{self[3], self[0], self[2], self[1]}
}
func (self Vec4[T]) WXZZ() VecN[T] {
	return VecN[T]{self[3], self[0], self[2], self[2]}
}
func (self Vec4[T]) WXZW() VecN[T] {
	return VecN[T]{self[3], self[0], self[2], self[3]}
}
func (self Vec4[T]) WXWX() VecN[T] {
	return VecN[T]{self[3], self[0], self[3], self[0]}
}
func (self Vec4[T]) WXWY() VecN[T] {
	return VecN[T]{self[3], self[0], self[3], self[1]}
}
func (self Vec4[T]) WXWZ() VecN[T] {
	return VecN[T]{self[3], self[0], self[3], self[2]}
}
func (self Vec4[T]) WXWW() VecN[T] {
	return VecN[T]{self[3], self[0], self[3], self[3]}
}
func (self Vec4[T]) WYXX() VecN[T] {
	return VecN[T]{self[3], self[1], self[0], self[0]}
}
func (self Vec4[T]) WYXY() VecN[T] {
	return VecN[T]{self[3], self[1], self[0], self[1]}
}
func (self Vec4[T]) WYXZ() VecN[T] {
	return VecN[T]{self[3], self[1], self[0], self[2]}
}
func (self Vec4[T]) WYXW() VecN[T] {
	return VecN[T]{self[3], self[1], self[0], self[3]}
}
func (self Vec4[T]) WYYX() VecN[T] {
	return VecN[T]{self[3], self[1], self[1], self[0]}
}
func (self Vec4[T]) WYYY() VecN[T] {
	return VecN[T]{self[3], self[1], self[1], self[1]}
}
func (self Vec4[T]) WYYZ() VecN[T] {
	return VecN[T]{self[3], self[1], self[1], self[2]}
}
func (self Vec4[T]) WYYW() VecN[T] {
	return VecN[T]{self[3], self[1], self[1], self[3]}
}
func (self Vec4[T]) WYZX() VecN[T] {
	return VecN[T]{self[3], self[1], self[2], self[0]}
}
func (self Vec4[T]) WYZY() VecN[T] {
	return VecN[T]{self[3], self[1], self[2], self[1]}
}
func (self Vec4[T]) WYZZ() VecN[T] {
	return VecN[T]{self[3], self[1], self[2], self[2]}
}
func (self Vec4[T]) WYZW() VecN[T] {
	return VecN[T]{self[3], self[1], self[2], self[3]}
}
func (self Vec4[T]) WYWX() VecN[T] {
	return VecN[T]{self[3], self[1], self[3], self[0]}
}
func (self Vec4[T]) WYWY() VecN[T] {
	return VecN[T]{self[3], self[1], self[3], self[1]}
}
func (self Vec4[T]) WYWZ() VecN[T] {
	return VecN[T]{self[3], self[1], self[3], self[2]}
}
func (self Vec4[T]) WYWW() VecN[T] {
	return VecN[T]{self[3], self[1], self[3], self[3]}
}
func (self Vec4[T]) WZXX() VecN[T] {
	return VecN[T]{self[3], self[2], self[0], self[0]}
}
func (self Vec4[T]) WZXY() VecN[T] {
	return VecN[T]{self[3], self[2], self[0], self[1]}
}
func (self Vec4[T]) WZXZ() VecN[T] {
	return VecN[T]{self[3], self[2], self[0], self[2]}
}
func (self Vec4[T]) WZXW() VecN[T] {
	return VecN[T]{self[3], self[2], self[0], self[3]}
}
func (self Vec4[T]) WZYX() VecN[T] {
	return VecN[T]{self[3], self[2], self[1], self[0]}
}
func (self Vec4[T]) WZYY() VecN[T] {
	return VecN[T]{self[3], self[2], self[1], self[1]}
}
func (self Vec4[T]) WZYZ() VecN[T] {
	return VecN[T]{self[3], self[2], self[1], self[2]}
}
func (self Vec4[T]) WZYW() VecN[T] {
	return VecN[T]{self[3], self[2], self[1], self[3]}
}
func (self Vec4[T]) WZZX() VecN[T] {
	return VecN[T]{self[3], self[2], self[2], self[0]}
}
func (self Vec4[T]) WZZY() VecN[T] {
	return VecN[T]{self[3], self[2], self[2], self[1]}
}
func (self Vec4[T]) WZZZ() VecN[T] {
	return VecN[T]{self[3], self[2], self[2], self[2]}
}
func (self Vec4[T]) WZZW() VecN[T] {
	return VecN[T]{self[3], self[2], self[2], self[3]}
}
func (self Vec4[T]) WZWX() VecN[T] {
	return VecN[T]{self[3], self[2], self[3], self[0]}
}
func (self Vec4[T]) WZWY() VecN[T] {
	return VecN[T]{self[3], self[2], self[3], self[1]}
}
func (self Vec4[T]) WZWZ() VecN[T] {
	return VecN[T]{self[3], self[2], self[3], self[2]}
}
func (self Vec4[T]) WZWW() VecN[T] {
	return VecN[T]{self[3], self[2], self[3], self[3]}
}
func (self Vec4[T]) WWXX() VecN[T] {
	return VecN[T]{self[3], self[3], self[0], self[0]}
}
func (self Vec4[T]) WWXY() VecN[T] {
	return VecN[T]{self[3], self[3], self[0], self[1]}
}
func (self Vec4[T]) WWXZ() VecN[T] {
	return VecN[T]{self[3], self[3], self[0], self[2]}
}
func (self Vec4[T]) WWXW() VecN[T] {
	return VecN[T]{self[3], self[3], self[0], self[3]}
}
func (self Vec4[T]) WWYX() VecN[T] {
	return VecN[T]{self[3], self[3], self[1], self[0]}
}
func (self Vec4[T]) WWYY() VecN[T] {
	return VecN[T]{self[3], self[3], self[1], self[1]}
}
func (self Vec4[T]) WWYZ() VecN[T] {
	return VecN[T]{self[3], self[3], self[1], self[2]}
}
func (self Vec4[T]) WWYW() VecN[T] {
	return VecN[T]{self[3], self[3], self[1], self[3]}
}
func (self Vec4[T]) WWZX() VecN[T] {
	return VecN[T]{self[3], self[3], self[2], self[0]}
}
func (self Vec4[T]) WWZY() VecN[T] {
	return VecN[T]{self[3], self[3], self[2], self[1]}
}
func (self Vec4[T]) WWZZ() VecN[T] {
	return VecN[T]{self[3], self[3], self[2], self[2]}
}
func (self Vec4[T]) WWZW() VecN[T] {
	return VecN[T]{self[3], self[3], self[2], self[3]}
}
func (self Vec4[T]) WWWX() VecN[T] {
	return VecN[T]{self[3], self[3], self[3], self[0]}
}
func (self Vec4[T]) WWWY() VecN[T] {
	return VecN[T]{self[3], self[3], self[3], self[1]}
}
func (self Vec4[T]) WWWZ() VecN[T] {
	return VecN[T]{self[3], self[3], self[3], self[2]}
}
func (self Vec4[T]) WWWW() VecN[T] {
	return VecN[T]{self[3], self[3], self[3], self[3]}
}
func (self Vec2[T]) RR() VecN[T] {
	return VecN[T]{self[0], self[0]}
}
func (self Vec2[T]) RG() VecN[T] {
	return VecN[T]{self[0], self[1]}
}
func (self Vec2[T]) GG() VecN[T] {
	return VecN[T]{self[1], self[1]}
}
func (self Vec2[T]) GR() VecN[T] {
	return VecN[T]{self[1], self[0]}
}
func (self Vec2[T]) RRR() VecN[T] {
	return VecN[T]{self[0], self[0], self[0]}
}
func (self Vec2[T]) RRG() VecN[T] {
	return VecN[T]{self[0], self[0], self[1]}
}
func (self Vec2[T]) RGR() VecN[T] {
	return VecN[T]{self[0], self[1], self[0]}
}
func (self Vec2[T]) RGG() VecN[T] {
	return VecN[T]{self[0], self[1], self[1]}
}
func (self Vec2[T]) GRR() VecN[T] {
	return VecN[T]{self[1], self[0], self[0]}
}
func (self Vec2[T]) GRG() VecN[T] {
	return VecN[T]{self[1], self[0], self[1]}
}
func (self Vec2[T]) GGR() VecN[T] {
	return VecN[T]{self[1], self[1], self[0]}
}
func (self Vec2[T]) GGG() VecN[T] {
	return VecN[T]{self[1], self[1], self[1]}
}
func (self Vec2[T]) RRRR() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[0]}
}
func (self Vec2[T]) RRRG() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[1]}
}
func (self Vec2[T]) RRGR() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[0]}
}
func (self Vec2[T]) RRGG() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[1]}
}
func (self Vec2[T]) RGRR() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[0]}
}
func (self Vec2[T]) RGRG() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[1]}
}
func (self Vec2[T]) RGGR() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[0]}
}
func (self Vec2[T]) RGGG() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[1]}
}
func (self Vec2[T]) GRRR() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[0]}
}
func (self Vec2[T]) GRRG() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[1]}
}
func (self Vec2[T]) GRGR() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[1]}
}
func (self Vec2[T]) GRGG() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[1]}
}
func (self Vec2[T]) GGRR() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[0]}
}
func (self Vec2[T]) GGRG() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[1]}
}
func (self Vec2[T]) GGGR() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[0]}
}
func (self Vec2[T]) GGGG() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[1]}
}
func (self Vec3[T]) RR() VecN[T] {
	return VecN[T]{self[0], self[0]}
}
func (self Vec3[T]) RG() VecN[T] {
	return VecN[T]{self[0], self[1]}
}
func (self Vec3[T]) RB() VecN[T] {
	return VecN[T]{self[0], self[2]}
}
func (self Vec3[T]) GR() VecN[T] {
	return VecN[T]{self[1], self[0]}
}
func (self Vec3[T]) GG() VecN[T] {
	return VecN[T]{self[1], self[1]}
}
func (self Vec3[T]) GB() VecN[T] {
	return VecN[T]{self[1], self[2]}
}
func (self Vec3[T]) BR() VecN[T] {
	return VecN[T]{self[2], self[0]}
}
func (self Vec3[T]) BG() VecN[T] {
	return VecN[T]{self[2], self[1]}
}
func (self Vec3[T]) BB() VecN[T] {
	return VecN[T]{self[2], self[2]}
}
func (self Vec3[T]) RRR() VecN[T] {
	return VecN[T]{self[0], self[0], self[0]}
}
func (self Vec3[T]) RRG() VecN[T] {
	return VecN[T]{self[0], self[0], self[1]}
}
func (self Vec3[T]) RRB() VecN[T] {
	return VecN[T]{self[0], self[0], self[2]}
}
func (self Vec3[T]) RGR() VecN[T] {
	return VecN[T]{self[0], self[1], self[0]}
}
func (self Vec3[T]) RGG() VecN[T] {
	return VecN[T]{self[0], self[1], self[1]}
}
func (self Vec3[T]) RGB() VecN[T] {
	return VecN[T]{self[0], self[1], self[2]}
}
func (self Vec3[T]) RBR() VecN[T] {
	return VecN[T]{self[0], self[2], self[0]}
}
func (self Vec3[T]) RBG() VecN[T] {
	return VecN[T]{self[0], self[2], self[1]}
}
func (self Vec3[T]) RBB() VecN[T] {
	return VecN[T]{self[0], self[2], self[2]}
}
func (self Vec3[T]) GRR() VecN[T] {
	return VecN[T]{self[1], self[0], self[0]}
}
func (self Vec3[T]) GRG() VecN[T] {
	return VecN[T]{self[1], self[0], self[1]}
}
func (self Vec3[T]) GRB() VecN[T] {
	return VecN[T]{self[1], self[0], self[2]}
}
func (self Vec3[T]) GGR() VecN[T] {
	return VecN[T]{self[1], self[1], self[0]}
}
func (self Vec3[T]) GGG() VecN[T] {
	return VecN[T]{self[1], self[1], self[1]}
}
func (self Vec3[T]) GGB() VecN[T] {
	return VecN[T]{self[1], self[1], self[2]}
}
func (self Vec3[T]) GBR() VecN[T] {
	return VecN[T]{self[1], self[2], self[0]}
}
func (self Vec3[T]) GBG() VecN[T] {
	return VecN[T]{self[1], self[2], self[1]}
}
func (self Vec3[T]) GBB() VecN[T] {
	return VecN[T]{self[1], self[2], self[2]}
}
func (self Vec3[T]) BRR() VecN[T] {
	return VecN[T]{self[2], self[0], self[0]}
}
func (self Vec3[T]) BRG() VecN[T] {
	return VecN[T]{self[2], self[0], self[1]}
}
func (self Vec3[T]) BRB() VecN[T] {
	return VecN[T]{self[2], self[0], self[2]}
}
func (self Vec3[T]) BGR() VecN[T] {
	return VecN[T]{self[2], self[1], self[0]}
}
func (self Vec3[T]) BGG() VecN[T] {
	return VecN[T]{self[2], self[1], self[1]}
}
func (self Vec3[T]) BGB() VecN[T] {
	return VecN[T]{self[2], self[1], self[2]}
}
func (self Vec3[T]) BBR() VecN[T] {
	return VecN[T]{self[2], self[2], self[0]}
}
func (self Vec3[T]) BBG() VecN[T] {
	return VecN[T]{self[2], self[2], self[1]}
}
func (self Vec3[T]) BBB() VecN[T] {
	return VecN[T]{self[2], self[2], self[2]}
}
func (self Vec3[T]) RRRR() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[0]}
}
func (self Vec3[T]) RRRG() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[1]}
}
func (self Vec3[T]) RRRB() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[2]}
}
func (self Vec3[T]) RRGR() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[0]}
}
func (self Vec3[T]) RRGG() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[1]}
}
func (self Vec3[T]) RRGB() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[2]}
}
func (self Vec3[T]) RRBR() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[0]}
}
func (self Vec3[T]) RRBG() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[1]}
}
func (self Vec3[T]) RRBB() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[2]}
}
func (self Vec3[T]) RGRR() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[0]}
}
func (self Vec3[T]) RGRG() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[1]}
}
func (self Vec3[T]) RGRB() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[2]}
}
func (self Vec3[T]) RGGR() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[0]}
}
func (self Vec3[T]) RGGG() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[1]}
}
func (self Vec3[T]) RGGB() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[2]}
}
func (self Vec3[T]) RGBR() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[0]}
}
func (self Vec3[T]) RGBG() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[1]}
}
func (self Vec3[T]) RGBB() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[2]}
}
func (self Vec3[T]) RBRR() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[0]}
}
func (self Vec3[T]) RBRG() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[1]}
}
func (self Vec3[T]) RBRB() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[2]}
}
func (self Vec3[T]) RBGR() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[0]}
}
func (self Vec3[T]) RBGG() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[1]}
}
func (self Vec3[T]) RBGB() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[2]}
}
func (self Vec3[T]) RBBR() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[0]}
}
func (self Vec3[T]) RBBG() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[1]}
}
func (self Vec3[T]) RBBB() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[2]}
}
func (self Vec3[T]) GRRR() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[0]}
}
func (self Vec3[T]) GRRG() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[1]}
}
func (self Vec3[T]) GRRB() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[2]}
}
func (self Vec3[T]) GRGR() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[0]}
}
func (self Vec3[T]) GRGG() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[1]}
}
func (self Vec3[T]) GRGB() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[2]}
}
func (self Vec3[T]) GRBR() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[0]}
}
func (self Vec3[T]) GRBG() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[1]}
}
func (self Vec3[T]) GRBB() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[2]}
}
func (self Vec3[T]) GGRR() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[0]}
}
func (self Vec3[T]) GGRG() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[1]}
}
func (self Vec3[T]) GGRB() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[2]}
}
func (self Vec3[T]) GGGR() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[0]}
}
func (self Vec3[T]) GGGG() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[1]}
}
func (self Vec3[T]) GGGB() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[2]}
}
func (self Vec3[T]) GGBR() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[0]}
}
func (self Vec3[T]) GGBG() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[1]}
}
func (self Vec3[T]) GGBB() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[2]}
}
func (self Vec3[T]) GBRR() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[0]}
}
func (self Vec3[T]) GBRG() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[1]}
}
func (self Vec3[T]) GBRB() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[2]}
}
func (self Vec3[T]) GBGR() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[0]}
}
func (self Vec3[T]) GBGG() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[1]}
}
func (self Vec3[T]) GBGB() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[2]}
}
func (self Vec3[T]) GBBR() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[0]}
}
func (self Vec3[T]) GBBG() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[1]}
}
func (self Vec3[T]) GBBB() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[2]}
}
func (self Vec3[T]) BRRR() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[0]}
}
func (self Vec3[T]) BRRG() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[1]}
}
func (self Vec3[T]) BRRB() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[2]}
}
func (self Vec3[T]) BRGR() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[0]}
}
func (self Vec3[T]) BRGG() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[1]}
}
func (self Vec3[T]) BRGB() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[2]}
}
func (self Vec3[T]) BRBR() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[0]}
}
func (self Vec3[T]) BRBG() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[1]}
}
func (self Vec3[T]) BRBB() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[2]}
}
func (self Vec3[T]) BGRR() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[0]}
}
func (self Vec3[T]) BGRG() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[1]}
}
func (self Vec3[T]) BGRB() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[2]}
}
func (self Vec3[T]) BGGR() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[0]}
}
func (self Vec3[T]) BGGG() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[1]}
}
func (self Vec3[T]) BGGB() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[2]}
}
func (self Vec3[T]) BGBR() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[0]}
}
func (self Vec3[T]) BGBG() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[1]}
}
func (self Vec3[T]) BGBB() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[2]}
}
func (self Vec3[T]) BBRR() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[0]}
}
func (self Vec3[T]) BBRG() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[1]}
}
func (self Vec3[T]) BBRB() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[2]}
}
func (self Vec3[T]) BBGR() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[0]}
}
func (self Vec3[T]) BBGG() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[1]}
}
func (self Vec3[T]) BBGB() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[2]}
}
func (self Vec3[T]) BBBR() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[0]}
}
func (self Vec3[T]) BBBG() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[1]}
}
func (self Vec3[T]) BBBB() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[2]}
}
func (self Vec4[T]) RR() VecN[T] {
	return VecN[T]{self[0], self[0]}
}
func (self Vec4[T]) RG() VecN[T] {
	return VecN[T]{self[0], self[1]}
}
func (self Vec4[T]) RB() VecN[T] {
	return VecN[T]{self[0], self[2]}
}
func (self Vec4[T]) RA() VecN[T] {
	return VecN[T]{self[0], self[3]}
}
func (self Vec4[T]) GR() VecN[T] {
	return VecN[T]{self[0], self[0]}
}
func (self Vec4[T]) GG() VecN[T] {
	return VecN[T]{self[1], self[1]}
}
func (self Vec4[T]) GB() VecN[T] {
	return VecN[T]{self[1], self[2]}
}
func (self Vec4[T]) GA() VecN[T] {
	return VecN[T]{self[1], self[3]}
}
func (self Vec4[T]) BR() VecN[T] {
	return VecN[T]{self[2], self[0]}
}
func (self Vec4[T]) BG() VecN[T] {
	return VecN[T]{self[2], self[1]}
}
func (self Vec4[T]) BB() VecN[T] {
	return VecN[T]{self[2], self[2]}
}
func (self Vec4[T]) BA() VecN[T] {
	return VecN[T]{self[2], self[3]}
}
func (self Vec4[T]) AR() VecN[T] {
	return VecN[T]{self[3], self[0]}
}
func (self Vec4[T]) AG() VecN[T] {
	return VecN[T]{self[3], self[1]}
}
func (self Vec4[T]) AB() VecN[T] {
	return VecN[T]{self[3], self[2]}
}
func (self Vec4[T]) AA() VecN[T] {
	return VecN[T]{self[3], self[3]}
}
func (self Vec4[T]) RRR() VecN[T] {
	return VecN[T]{self[0], self[0], self[0]}
}
func (self Vec4[T]) RRG() VecN[T] {
	return VecN[T]{self[0], self[0], self[1]}
}
func (self Vec4[T]) RRB() VecN[T] {
	return VecN[T]{self[0], self[0], self[2]}
}
func (self Vec4[T]) RRA() VecN[T] {
	return VecN[T]{self[0], self[0], self[3]}
}
func (self Vec4[T]) RGR() VecN[T] {
	return VecN[T]{self[0], self[1], self[0]}
}
func (self Vec4[T]) RGG() VecN[T] {
	return VecN[T]{self[0], self[1], self[1]}
}
func (self Vec4[T]) RGB() VecN[T] {
	return VecN[T]{self[0], self[1], self[2]}
}
func (self Vec4[T]) RGA() VecN[T] {
	return VecN[T]{self[0], self[1], self[3]}
}
func (self Vec4[T]) RBR() VecN[T] {
	return VecN[T]{self[0], self[2], self[0]}
}
func (self Vec4[T]) RBG() VecN[T] {
	return VecN[T]{self[0], self[2], self[1]}
}
func (self Vec4[T]) RBB() VecN[T] {
	return VecN[T]{self[0], self[2], self[2]}
}
func (self Vec4[T]) RBA() VecN[T] {
	return VecN[T]{self[0], self[2], self[3]}
}
func (self Vec4[T]) RAR() VecN[T] {
	return VecN[T]{self[0], self[3], self[0]}
}
func (self Vec4[T]) RAG() VecN[T] {
	return VecN[T]{self[0], self[3], self[1]}
}
func (self Vec4[T]) RAB() VecN[T] {
	return VecN[T]{self[0], self[3], self[2]}
}
func (self Vec4[T]) RAA() VecN[T] {
	return VecN[T]{self[0], self[3], self[3]}
}
func (self Vec4[T]) GRR() VecN[T] {
	return VecN[T]{self[1], self[0], self[0]}
}
func (self Vec4[T]) GRG() VecN[T] {
	return VecN[T]{self[1], self[0], self[1]}
}
func (self Vec4[T]) GRB() VecN[T] {
	return VecN[T]{self[1], self[0], self[2]}
}
func (self Vec4[T]) GRA() VecN[T] {
	return VecN[T]{self[1], self[0], self[3]}
}
func (self Vec4[T]) GGR() VecN[T] {
	return VecN[T]{self[1], self[1], self[0]}
}
func (self Vec4[T]) GGG() VecN[T] {
	return VecN[T]{self[1], self[1], self[1]}
}
func (self Vec4[T]) GGB() VecN[T] {
	return VecN[T]{self[1], self[1], self[2]}
}
func (self Vec4[T]) GGA() VecN[T] {
	return VecN[T]{self[1], self[1], self[3]}
}
func (self Vec4[T]) GBR() VecN[T] {
	return VecN[T]{self[1], self[2], self[0]}
}
func (self Vec4[T]) GBG() VecN[T] {
	return VecN[T]{self[1], self[2], self[1]}
}
func (self Vec4[T]) GBB() VecN[T] {
	return VecN[T]{self[1], self[2], self[2]}
}
func (self Vec4[T]) GBA() VecN[T] {
	return VecN[T]{self[1], self[2], self[3]}
}
func (self Vec4[T]) GAR() VecN[T] {
	return VecN[T]{self[1], self[3], self[0]}
}
func (self Vec4[T]) GAG() VecN[T] {
	return VecN[T]{self[1], self[3], self[1]}
}
func (self Vec4[T]) GAB() VecN[T] {
	return VecN[T]{self[1], self[3], self[2]}
}
func (self Vec4[T]) GAA() VecN[T] {
	return VecN[T]{self[1], self[3], self[3]}
}
func (self Vec4[T]) BRR() VecN[T] {
	return VecN[T]{self[2], self[0], self[0]}
}
func (self Vec4[T]) BRG() VecN[T] {
	return VecN[T]{self[2], self[0], self[1]}
}
func (self Vec4[T]) BRB() VecN[T] {
	return VecN[T]{self[2], self[0], self[2]}
}
func (self Vec4[T]) BRA() VecN[T] {
	return VecN[T]{self[2], self[0], self[3]}
}
func (self Vec4[T]) BGR() VecN[T] {
	return VecN[T]{self[2], self[1], self[0]}
}
func (self Vec4[T]) BGG() VecN[T] {
	return VecN[T]{self[2], self[1], self[1]}
}
func (self Vec4[T]) BGB() VecN[T] {
	return VecN[T]{self[2], self[1], self[2]}
}
func (self Vec4[T]) BGA() VecN[T] {
	return VecN[T]{self[2], self[1], self[3]}
}
func (self Vec4[T]) BBR() VecN[T] {
	return VecN[T]{self[2], self[2], self[0]}
}
func (self Vec4[T]) BBG() VecN[T] {
	return VecN[T]{self[2], self[2], self[1]}
}
func (self Vec4[T]) BBB() VecN[T] {
	return VecN[T]{self[2], self[2], self[2]}
}
func (self Vec4[T]) BBA() VecN[T] {
	return VecN[T]{self[2], self[2], self[3]}
}
func (self Vec4[T]) BAR() VecN[T] {
	return VecN[T]{self[2], self[3], self[0]}
}
func (self Vec4[T]) BAG() VecN[T] {
	return VecN[T]{self[2], self[3], self[1]}
}
func (self Vec4[T]) BAB() VecN[T] {
	return VecN[T]{self[2], self[3], self[2]}
}
func (self Vec4[T]) BAA() VecN[T] {
	return VecN[T]{self[2], self[3], self[3]}
}
func (self Vec4[T]) ARR() VecN[T] {
	return VecN[T]{self[3], self[0], self[0]}
}
func (self Vec4[T]) ARG() VecN[T] {
	return VecN[T]{self[3], self[0], self[1]}
}
func (self Vec4[T]) ARB() VecN[T] {
	return VecN[T]{self[3], self[0], self[2]}
}
func (self Vec4[T]) ARA() VecN[T] {
	return VecN[T]{self[3], self[0], self[3]}
}
func (self Vec4[T]) AGR() VecN[T] {
	return VecN[T]{self[3], self[1], self[0]}
}
func (self Vec4[T]) AGG() VecN[T] {
	return VecN[T]{self[3], self[1], self[1]}
}
func (self Vec4[T]) AGB() VecN[T] {
	return VecN[T]{self[3], self[1], self[2]}
}
func (self Vec4[T]) AGA() VecN[T] {
	return VecN[T]{self[3], self[1], self[3]}
}
func (self Vec4[T]) ABR() VecN[T] {
	return VecN[T]{self[3], self[2], self[0]}
}
func (self Vec4[T]) ABG() VecN[T] {
	return VecN[T]{self[3], self[2], self[1]}
}
func (self Vec4[T]) ABB() VecN[T] {
	return VecN[T]{self[3], self[2], self[2]}
}
func (self Vec4[T]) ABA() VecN[T] {
	return VecN[T]{self[3], self[2], self[3]}
}
func (self Vec4[T]) AAR() VecN[T] {
	return VecN[T]{self[3], self[3], self[0]}
}
func (self Vec4[T]) AAG() VecN[T] {
	return VecN[T]{self[3], self[3], self[1]}
}
func (self Vec4[T]) AAB() VecN[T] {
	return VecN[T]{self[3], self[3], self[2]}
}
func (self Vec4[T]) AAA() VecN[T] {
	return VecN[T]{self[3], self[3], self[3]}
}
func (self Vec4[T]) RRRR() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[0]}
}
func (self Vec4[T]) RRRG() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[1]}
}
func (self Vec4[T]) RRRB() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[2]}
}
func (self Vec4[T]) RRRA() VecN[T] {
	return VecN[T]{self[0], self[0], self[0], self[3]}
}
func (self Vec4[T]) RRGR() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[0]}
}
func (self Vec4[T]) RRGG() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[1]}
}
func (self Vec4[T]) RRGB() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[2]}
}
func (self Vec4[T]) RRGA() VecN[T] {
	return VecN[T]{self[0], self[0], self[1], self[3]}
}
func (self Vec4[T]) RRBR() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[0]}
}
func (self Vec4[T]) RRBG() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[1]}
}
func (self Vec4[T]) RRBB() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[2]}
}
func (self Vec4[T]) RRBA() VecN[T] {
	return VecN[T]{self[0], self[0], self[2], self[3]}
}
func (self Vec4[T]) RRAR() VecN[T] {
	return VecN[T]{self[0], self[0], self[3], self[0]}
}
func (self Vec4[T]) RRAG() VecN[T] {
	return VecN[T]{self[0], self[0], self[3], self[1]}
}
func (self Vec4[T]) RRAB() VecN[T] {
	return VecN[T]{self[0], self[0], self[3], self[2]}
}
func (self Vec4[T]) RRAA() VecN[T] {
	return VecN[T]{self[0], self[0], self[3], self[3]}
}
func (self Vec4[T]) RGRR() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[0]}
}
func (self Vec4[T]) RGRG() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[1]}
}
func (self Vec4[T]) RGRB() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[2]}
}
func (self Vec4[T]) RGRA() VecN[T] {
	return VecN[T]{self[0], self[1], self[0], self[3]}
}
func (self Vec4[T]) RGGR() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[0]}
}
func (self Vec4[T]) RGGG() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[1]}
}
func (self Vec4[T]) RGGB() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[2]}
}
func (self Vec4[T]) RGGA() VecN[T] {
	return VecN[T]{self[0], self[1], self[1], self[3]}
}
func (self Vec4[T]) RGBR() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[0]}
}
func (self Vec4[T]) RGBG() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[1]}
}
func (self Vec4[T]) RGBB() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[2]}
}
func (self Vec4[T]) RGBA() VecN[T] {
	return VecN[T]{self[0], self[1], self[2], self[3]}
}
func (self Vec4[T]) RGAR() VecN[T] {
	return VecN[T]{self[0], self[1], self[3], self[0]}
}
func (self Vec4[T]) RGAG() VecN[T] {
	return VecN[T]{self[0], self[1], self[3], self[1]}
}
func (self Vec4[T]) RGAB() VecN[T] {
	return VecN[T]{self[0], self[1], self[3], self[2]}
}
func (self Vec4[T]) RGAA() VecN[T] {
	return VecN[T]{self[0], self[1], self[3], self[3]}
}
func (self Vec4[T]) RBRR() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[0]}
}
func (self Vec4[T]) RBRG() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[1]}
}
func (self Vec4[T]) RBRB() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[2]}
}
func (self Vec4[T]) RBRA() VecN[T] {
	return VecN[T]{self[0], self[2], self[0], self[3]}
}
func (self Vec4[T]) RBGR() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[0]}
}
func (self Vec4[T]) RBGG() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[1]}
}
func (self Vec4[T]) RBGB() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[2]}
}
func (self Vec4[T]) RBGA() VecN[T] {
	return VecN[T]{self[0], self[2], self[1], self[3]}
}
func (self Vec4[T]) RBBR() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[0]}
}
func (self Vec4[T]) RBBG() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[1]}
}
func (self Vec4[T]) RBBB() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[2]}
}
func (self Vec4[T]) RBBA() VecN[T] {
	return VecN[T]{self[0], self[2], self[2], self[3]}
}
func (self Vec4[T]) RBAR() VecN[T] {
	return VecN[T]{self[0], self[2], self[3], self[0]}
}
func (self Vec4[T]) RBAG() VecN[T] {
	return VecN[T]{self[0], self[2], self[3], self[1]}
}
func (self Vec4[T]) RBAB() VecN[T] {
	return VecN[T]{self[0], self[2], self[3], self[2]}
}
func (self Vec4[T]) RBAA() VecN[T] {
	return VecN[T]{self[0], self[2], self[3], self[3]}
}
func (self Vec4[T]) RARR() VecN[T] {
	return VecN[T]{self[0], self[3], self[0], self[0]}
}
func (self Vec4[T]) RARG() VecN[T] {
	return VecN[T]{self[0], self[3], self[0], self[1]}
}
func (self Vec4[T]) RARB() VecN[T] {
	return VecN[T]{self[0], self[3], self[0], self[2]}
}
func (self Vec4[T]) RARA() VecN[T] {
	return VecN[T]{self[0], self[3], self[0], self[3]}
}
func (self Vec4[T]) RAGR() VecN[T] {
	return VecN[T]{self[0], self[3], self[1], self[0]}
}
func (self Vec4[T]) RAGG() VecN[T] {
	return VecN[T]{self[0], self[3], self[1], self[1]}
}
func (self Vec4[T]) RAGB() VecN[T] {
	return VecN[T]{self[0], self[3], self[1], self[2]}
}
func (self Vec4[T]) RAGA() VecN[T] {
	return VecN[T]{self[0], self[3], self[1], self[3]}
}
func (self Vec4[T]) RABR() VecN[T] {
	return VecN[T]{self[0], self[3], self[2], self[0]}
}
func (self Vec4[T]) RABG() VecN[T] {
	return VecN[T]{self[0], self[3], self[2], self[1]}
}
func (self Vec4[T]) RABB() VecN[T] {
	return VecN[T]{self[0], self[3], self[2], self[2]}
}
func (self Vec4[T]) RABA() VecN[T] {
	return VecN[T]{self[0], self[3], self[2], self[3]}
}
func (self Vec4[T]) RAAR() VecN[T] {
	return VecN[T]{self[0], self[3], self[3], self[0]}
}
func (self Vec4[T]) RAAG() VecN[T] {
	return VecN[T]{self[0], self[3], self[3], self[1]}
}
func (self Vec4[T]) RAAB() VecN[T] {
	return VecN[T]{self[0], self[3], self[3], self[2]}
}
func (self Vec4[T]) RAAA() VecN[T] {
	return VecN[T]{self[0], self[3], self[3], self[3]}
}
func (self Vec4[T]) GRRR() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[0]}
}
func (self Vec4[T]) GRRG() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[1]}
}
func (self Vec4[T]) GRRB() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[2]}
}
func (self Vec4[T]) GRRA() VecN[T] {
	return VecN[T]{self[1], self[0], self[0], self[3]}
}
func (self Vec4[T]) GRGR() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[0]}
}
func (self Vec4[T]) GRGG() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[1]}
}
func (self Vec4[T]) GRGB() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[2]}
}
func (self Vec4[T]) GRGA() VecN[T] {
	return VecN[T]{self[1], self[0], self[1], self[3]}
}
func (self Vec4[T]) GRBR() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[0]}
}
func (self Vec4[T]) GRBG() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[1]}
}
func (self Vec4[T]) GRBB() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[2]}
}
func (self Vec4[T]) GRBA() VecN[T] {
	return VecN[T]{self[1], self[0], self[2], self[3]}
}
func (self Vec4[T]) GRAR() VecN[T] {
	return VecN[T]{self[1], self[0], self[3], self[0]}
}
func (self Vec4[T]) GRAG() VecN[T] {
	return VecN[T]{self[1], self[0], self[3], self[1]}
}
func (self Vec4[T]) GRAB() VecN[T] {
	return VecN[T]{self[1], self[0], self[3], self[2]}
}
func (self Vec4[T]) GRAA() VecN[T] {
	return VecN[T]{self[1], self[0], self[3], self[3]}
}
func (self Vec4[T]) GGRR() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[0]}
}
func (self Vec4[T]) GGRG() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[1]}
}
func (self Vec4[T]) GGRB() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[2]}
}
func (self Vec4[T]) GGRA() VecN[T] {
	return VecN[T]{self[1], self[1], self[0], self[3]}
}
func (self Vec4[T]) GGGR() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[0]}
}
func (self Vec4[T]) GGGG() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[1]}
}
func (self Vec4[T]) GGGB() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[2]}
}
func (self Vec4[T]) GGGA() VecN[T] {
	return VecN[T]{self[1], self[1], self[1], self[3]}
}
func (self Vec4[T]) GGBR() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[0]}
}
func (self Vec4[T]) GGBG() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[1]}
}
func (self Vec4[T]) GGBB() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[2]}
}
func (self Vec4[T]) GGBA() VecN[T] {
	return VecN[T]{self[1], self[1], self[2], self[3]}
}
func (self Vec4[T]) GGAR() VecN[T] {
	return VecN[T]{self[1], self[1], self[3], self[0]}
}
func (self Vec4[T]) GGAG() VecN[T] {
	return VecN[T]{self[1], self[1], self[3], self[1]}
}
func (self Vec4[T]) GGAB() VecN[T] {
	return VecN[T]{self[1], self[1], self[3], self[2]}
}
func (self Vec4[T]) GGAA() VecN[T] {
	return VecN[T]{self[1], self[1], self[3], self[3]}
}
func (self Vec4[T]) GBRR() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[0]}
}
func (self Vec4[T]) GBRG() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[1]}
}
func (self Vec4[T]) GBRB() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[2]}
}
func (self Vec4[T]) GBRA() VecN[T] {
	return VecN[T]{self[1], self[2], self[0], self[3]}
}
func (self Vec4[T]) GBGR() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[0]}
}
func (self Vec4[T]) GBGG() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[1]}
}
func (self Vec4[T]) GBGB() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[2]}
}
func (self Vec4[T]) GBGA() VecN[T] {
	return VecN[T]{self[1], self[2], self[1], self[3]}
}
func (self Vec4[T]) GBBR() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[0]}
}
func (self Vec4[T]) GBBG() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[1]}
}
func (self Vec4[T]) GBBB() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[2]}
}
func (self Vec4[T]) GBBA() VecN[T] {
	return VecN[T]{self[1], self[2], self[2], self[3]}
}
func (self Vec4[T]) GBAR() VecN[T] {
	return VecN[T]{self[1], self[2], self[3], self[0]}
}
func (self Vec4[T]) GBAG() VecN[T] {
	return VecN[T]{self[1], self[2], self[3], self[1]}
}
func (self Vec4[T]) GBAB() VecN[T] {
	return VecN[T]{self[1], self[2], self[3], self[2]}
}
func (self Vec4[T]) GBAA() VecN[T] {
	return VecN[T]{self[1], self[2], self[3], self[3]}
}
func (self Vec4[T]) GARR() VecN[T] {
	return VecN[T]{self[1], self[3], self[0], self[0]}
}
func (self Vec4[T]) GARG() VecN[T] {
	return VecN[T]{self[1], self[3], self[0], self[1]}
}
func (self Vec4[T]) GARB() VecN[T] {
	return VecN[T]{self[1], self[3], self[0], self[2]}
}
func (self Vec4[T]) GARA() VecN[T] {
	return VecN[T]{self[1], self[3], self[0], self[3]}
}
func (self Vec4[T]) GAGR() VecN[T] {
	return VecN[T]{self[1], self[3], self[1], self[0]}
}
func (self Vec4[T]) GAGG() VecN[T] {
	return VecN[T]{self[1], self[3], self[1], self[1]}
}
func (self Vec4[T]) GAGB() VecN[T] {
	return VecN[T]{self[1], self[3], self[1], self[2]}
}
func (self Vec4[T]) GAGA() VecN[T] {
	return VecN[T]{self[1], self[3], self[1], self[3]}
}
func (self Vec4[T]) GABR() VecN[T] {
	return VecN[T]{self[1], self[3], self[2], self[0]}
}
func (self Vec4[T]) GABG() VecN[T] {
	return VecN[T]{self[1], self[3], self[2], self[1]}
}
func (self Vec4[T]) GABB() VecN[T] {
	return VecN[T]{self[1], self[3], self[2], self[2]}
}
func (self Vec4[T]) GABA() VecN[T] {
	return VecN[T]{self[1], self[3], self[2], self[3]}
}
func (self Vec4[T]) GAAR() VecN[T] {
	return VecN[T]{self[1], self[3], self[3], self[0]}
}
func (self Vec4[T]) GAAG() VecN[T] {
	return VecN[T]{self[1], self[3], self[3], self[1]}
}
func (self Vec4[T]) GAAB() VecN[T] {
	return VecN[T]{self[1], self[3], self[3], self[2]}
}
func (self Vec4[T]) GAAA() VecN[T] {
	return VecN[T]{self[1], self[3], self[3], self[3]}
}
func (self Vec4[T]) BRRR() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[0]}
}
func (self Vec4[T]) BRRG() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[1]}
}
func (self Vec4[T]) BRRB() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[2]}
}
func (self Vec4[T]) BRRA() VecN[T] {
	return VecN[T]{self[2], self[0], self[0], self[3]}
}
func (self Vec4[T]) BRGR() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[0]}
}
func (self Vec4[T]) BRGG() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[1]}
}
func (self Vec4[T]) BRGB() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[2]}
}
func (self Vec4[T]) BRGA() VecN[T] {
	return VecN[T]{self[2], self[0], self[1], self[3]}
}
func (self Vec4[T]) BRBR() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[0]}
}
func (self Vec4[T]) BRBG() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[1]}
}
func (self Vec4[T]) BRBB() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[2]}
}
func (self Vec4[T]) BRBA() VecN[T] {
	return VecN[T]{self[2], self[0], self[2], self[3]}
}
func (self Vec4[T]) BRAR() VecN[T] {
	return VecN[T]{self[2], self[0], self[3], self[0]}
}
func (self Vec4[T]) BRAG() VecN[T] {
	return VecN[T]{self[2], self[0], self[3], self[1]}
}
func (self Vec4[T]) BRAB() VecN[T] {
	return VecN[T]{self[2], self[0], self[3], self[2]}
}
func (self Vec4[T]) BRAA() VecN[T] {
	return VecN[T]{self[2], self[0], self[3], self[3]}
}
func (self Vec4[T]) BGRR() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[0]}
}
func (self Vec4[T]) BGRG() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[1]}
}
func (self Vec4[T]) BGRB() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[2]}
}
func (self Vec4[T]) BGRA() VecN[T] {
	return VecN[T]{self[2], self[1], self[0], self[3]}
}
func (self Vec4[T]) BGGR() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[0]}
}
func (self Vec4[T]) BGGG() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[1]}
}
func (self Vec4[T]) BGGB() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[2]}
}
func (self Vec4[T]) BGGA() VecN[T] {
	return VecN[T]{self[2], self[1], self[1], self[3]}
}
func (self Vec4[T]) BGBR() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[0]}
}
func (self Vec4[T]) BGBG() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[1]}
}
func (self Vec4[T]) BGBB() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[2]}
}
func (self Vec4[T]) BGBA() VecN[T] {
	return VecN[T]{self[2], self[1], self[2], self[3]}
}
func (self Vec4[T]) BGAR() VecN[T] {
	return VecN[T]{self[2], self[1], self[3], self[0]}
}
func (self Vec4[T]) BGAG() VecN[T] {
	return VecN[T]{self[2], self[1], self[3], self[1]}
}
func (self Vec4[T]) BGAB() VecN[T] {
	return VecN[T]{self[2], self[1], self[3], self[2]}
}
func (self Vec4[T]) BGAA() VecN[T] {
	return VecN[T]{self[2], self[1], self[3], self[3]}
}
func (self Vec4[T]) BBRR() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[0]}
}
func (self Vec4[T]) BBRG() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[1]}
}
func (self Vec4[T]) BBRB() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[2]}
}
func (self Vec4[T]) BBRA() VecN[T] {
	return VecN[T]{self[2], self[2], self[0], self[3]}
}
func (self Vec4[T]) BBGR() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[0]}
}
func (self Vec4[T]) BBGG() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[1]}
}
func (self Vec4[T]) BBGB() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[2]}
}
func (self Vec4[T]) BBGA() VecN[T] {
	return VecN[T]{self[2], self[2], self[1], self[3]}
}
func (self Vec4[T]) BBBR() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[0]}
}
func (self Vec4[T]) BBBG() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[1]}
}
func (self Vec4[T]) BBBB() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[2]}
}
func (self Vec4[T]) BBBA() VecN[T] {
	return VecN[T]{self[2], self[2], self[2], self[3]}
}
func (self Vec4[T]) BBAR() VecN[T] {
	return VecN[T]{self[2], self[2], self[3], self[0]}
}
func (self Vec4[T]) BBAG() VecN[T] {
	return VecN[T]{self[2], self[2], self[3], self[1]}
}
func (self Vec4[T]) BBAB() VecN[T] {
	return VecN[T]{self[2], self[2], self[3], self[2]}
}
func (self Vec4[T]) BBAA() VecN[T] {
	return VecN[T]{self[2], self[2], self[3], self[3]}
}
func (self Vec4[T]) BARR() VecN[T] {
	return VecN[T]{self[2], self[3], self[0], self[0]}
}
func (self Vec4[T]) BARG() VecN[T] {
	return VecN[T]{self[2], self[3], self[0], self[1]}
}
func (self Vec4[T]) BARB() VecN[T] {
	return VecN[T]{self[2], self[3], self[0], self[2]}
}
func (self Vec4[T]) BARA() VecN[T] {
	return VecN[T]{self[2], self[3], self[0], self[3]}
}
func (self Vec4[T]) BAGR() VecN[T] {
	return VecN[T]{self[2], self[3], self[1], self[0]}
}
func (self Vec4[T]) BAGG() VecN[T] {
	return VecN[T]{self[2], self[3], self[1], self[1]}
}
func (self Vec4[T]) BAGB() VecN[T] {
	return VecN[T]{self[2], self[3], self[1], self[2]}
}
func (self Vec4[T]) BAGA() VecN[T] {
	return VecN[T]{self[2], self[3], self[1], self[3]}
}
func (self Vec4[T]) BABR() VecN[T] {
	return VecN[T]{self[2], self[3], self[2], self[0]}
}
func (self Vec4[T]) BABG() VecN[T] {
	return VecN[T]{self[2], self[3], self[2], self[1]}
}
func (self Vec4[T]) BABB() VecN[T] {
	return VecN[T]{self[2], self[3], self[2], self[2]}
}
func (self Vec4[T]) BABA() VecN[T] {
	return VecN[T]{self[2], self[3], self[2], self[3]}
}
func (self Vec4[T]) BAAR() VecN[T] {
	return VecN[T]{self[2], self[3], self[3], self[0]}
}
func (self Vec4[T]) BAAG() VecN[T] {
	return VecN[T]{self[2], self[3], self[3], self[1]}
}
func (self Vec4[T]) BAAB() VecN[T] {
	return VecN[T]{self[2], self[3], self[3], self[2]}
}
func (self Vec4[T]) BAAA() VecN[T] {
	return VecN[T]{self[2], self[3], self[3], self[3]}
}
func (self Vec4[T]) ARRR() VecN[T] {
	return VecN[T]{self[3], self[0], self[0], self[0]}
}
func (self Vec4[T]) ARRG() VecN[T] {
	return VecN[T]{self[3], self[0], self[0], self[1]}
}
func (self Vec4[T]) ARRB() VecN[T] {
	return VecN[T]{self[3], self[0], self[0], self[2]}
}
func (self Vec4[T]) ARRA() VecN[T] {
	return VecN[T]{self[3], self[0], self[0], self[3]}
}
func (self Vec4[T]) ARGR() VecN[T] {
	return VecN[T]{self[3], self[0], self[1], self[0]}
}
func (self Vec4[T]) ARGG() VecN[T] {
	return VecN[T]{self[3], self[0], self[1], self[1]}
}
func (self Vec4[T]) ARGB() VecN[T] {
	return VecN[T]{self[3], self[0], self[1], self[2]}
}
func (self Vec4[T]) ARGA() VecN[T] {
	return VecN[T]{self[3], self[0], self[1], self[3]}
}
func (self Vec4[T]) ARBR() VecN[T] {
	return VecN[T]{self[3], self[0], self[2], self[0]}
}
func (self Vec4[T]) ARBG() VecN[T] {
	return VecN[T]{self[3], self[0], self[2], self[1]}
}
func (self Vec4[T]) ARBB() VecN[T] {
	return VecN[T]{self[3], self[0], self[2], self[2]}
}
func (self Vec4[T]) ARBA() VecN[T] {
	return VecN[T]{self[3], self[0], self[2], self[3]}
}
func (self Vec4[T]) ARAR() VecN[T] {
	return VecN[T]{self[3], self[0], self[3], self[0]}
}
func (self Vec4[T]) ARAG() VecN[T] {
	return VecN[T]{self[3], self[0], self[3], self[1]}
}
func (self Vec4[T]) ARAB() VecN[T] {
	return VecN[T]{self[3], self[0], self[3], self[2]}
}
func (self Vec4[T]) ARAA() VecN[T] {
	return VecN[T]{self[3], self[0], self[3], self[3]}
}
func (self Vec4[T]) AGRR() VecN[T] {
	return VecN[T]{self[3], self[1], self[0], self[0]}
}
func (self Vec4[T]) AGRG() VecN[T] {
	return VecN[T]{self[3], self[1], self[0], self[1]}
}
func (self Vec4[T]) AGRB() VecN[T] {
	return VecN[T]{self[3], self[1], self[0], self[2]}
}
func (self Vec4[T]) AGRA() VecN[T] {
	return VecN[T]{self[3], self[1], self[0], self[3]}
}
func (self Vec4[T]) AGGR() VecN[T] {
	return VecN[T]{self[3], self[1], self[1], self[0]}
}
func (self Vec4[T]) AGGG() VecN[T] {
	return VecN[T]{self[3], self[1], self[1], self[1]}
}
func (self Vec4[T]) AGGB() VecN[T] {
	return VecN[T]{self[3], self[1], self[1], self[2]}
}
func (self Vec4[T]) AGGA() VecN[T] {
	return VecN[T]{self[3], self[1], self[1], self[3]}
}
func (self Vec4[T]) AGBR() VecN[T] {
	return VecN[T]{self[3], self[1], self[2], self[0]}
}
func (self Vec4[T]) AGBG() VecN[T] {
	return VecN[T]{self[3], self[1], self[2], self[1]}
}
func (self Vec4[T]) AGBB() VecN[T] {
	return VecN[T]{self[3], self[1], self[2], self[2]}
}
func (self Vec4[T]) AGBA() VecN[T] {
	return VecN[T]{self[3], self[1], self[2], self[3]}
}
func (self Vec4[T]) AGAR() VecN[T] {
	return VecN[T]{self[3], self[1], self[3], self[0]}
}
func (self Vec4[T]) AGAG() VecN[T] {
	return VecN[T]{self[3], self[1], self[3], self[1]}
}
func (self Vec4[T]) AGAB() VecN[T] {
	return VecN[T]{self[3], self[1], self[3], self[2]}
}
func (self Vec4[T]) AGAA() VecN[T] {
	return VecN[T]{self[3], self[1], self[3], self[3]}
}
func (self Vec4[T]) ABRR() VecN[T] {
	return VecN[T]{self[3], self[2], self[0], self[0]}
}
func (self Vec4[T]) ABRG() VecN[T] {
	return VecN[T]{self[3], self[2], self[0], self[1]}
}
func (self Vec4[T]) ABRB() VecN[T] {
	return VecN[T]{self[3], self[2], self[0], self[2]}
}
func (self Vec4[T]) ABRA() VecN[T] {
	return VecN[T]{self[3], self[2], self[0], self[3]}
}
func (self Vec4[T]) ABGR() VecN[T] {
	return VecN[T]{self[3], self[2], self[1], self[0]}
}
func (self Vec4[T]) ABGG() VecN[T] {
	return VecN[T]{self[3], self[2], self[1], self[1]}
}
func (self Vec4[T]) ABGB() VecN[T] {
	return VecN[T]{self[3], self[2], self[1], self[2]}
}
func (self Vec4[T]) ABGA() VecN[T] {
	return VecN[T]{self[3], self[2], self[1], self[3]}
}
func (self Vec4[T]) ABBR() VecN[T] {
	return VecN[T]{self[3], self[2], self[2], self[0]}
}
func (self Vec4[T]) ABBG() VecN[T] {
	return VecN[T]{self[3], self[2], self[2], self[1]}
}
func (self Vec4[T]) ABBB() VecN[T] {
	return VecN[T]{self[3], self[2], self[2], self[2]}
}
func (self Vec4[T]) ABBA() VecN[T] {
	return VecN[T]{self[3], self[2], self[2], self[3]}
}
func (self Vec4[T]) ABAR() VecN[T] {
	return VecN[T]{self[3], self[2], self[3], self[0]}
}
func (self Vec4[T]) ABAG() VecN[T] {
	return VecN[T]{self[3], self[2], self[3], self[1]}
}
func (self Vec4[T]) ABAB() VecN[T] {
	return VecN[T]{self[3], self[2], self[3], self[2]}
}
func (self Vec4[T]) ABAA() VecN[T] {
	return VecN[T]{self[3], self[2], self[3], self[3]}
}
func (self Vec4[T]) AARR() VecN[T] {
	return VecN[T]{self[3], self[3], self[0], self[0]}
}
func (self Vec4[T]) AARG() VecN[T] {
	return VecN[T]{self[3], self[3], self[0], self[1]}
}
func (self Vec4[T]) AARB() VecN[T] {
	return VecN[T]{self[3], self[3], self[0], self[2]}
}
func (self Vec4[T]) AARA() VecN[T] {
	return VecN[T]{self[3], self[3], self[0], self[3]}
}
func (self Vec4[T]) AAGR() VecN[T] {
	return VecN[T]{self[3], self[3], self[1], self[0]}
}
func (self Vec4[T]) AAGG() VecN[T] {
	return VecN[T]{self[3], self[3], self[1], self[1]}
}
func (self Vec4[T]) AAGB() VecN[T] {
	return VecN[T]{self[3], self[3], self[1], self[2]}
}
func (self Vec4[T]) AAGA() VecN[T] {
	return VecN[T]{self[3], self[3], self[1], self[3]}
}
func (self Vec4[T]) AABR() VecN[T] {
	return VecN[T]{self[3], self[3], self[2], self[0]}
}
func (self Vec4[T]) AABG() VecN[T] {
	return VecN[T]{self[3], self[3], self[2], self[1]}
}
func (self Vec4[T]) AABB() VecN[T] {
	return VecN[T]{self[3], self[3], self[2], self[2]}
}
func (self Vec4[T]) AABA() VecN[T] {
	return VecN[T]{self[3], self[3], self[2], self[3]}
}
func (self Vec4[T]) AAAR() VecN[T] {
	return VecN[T]{self[3], self[3], self[3], self[0]}
}
func (self Vec4[T]) AAAG() VecN[T] {
	return VecN[T]{self[3], self[3], self[3], self[1]}
}
func (self Vec4[T]) AAAB() VecN[T] {
	return VecN[T]{self[3], self[3], self[3], self[2]}
}
func (self Vec4[T]) AAAA() VecN[T] {
	return VecN[T]{self[3], self[3], self[3], self[3]}
}
