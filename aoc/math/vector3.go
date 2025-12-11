package math

type Vector3[T Numbers] struct {
	X T
	Y T
	Z T
}

func NewVector3[T Numbers](x, y, z T) Vector3[T] {
	return Vector3[T]{X: x, Y: y, Z: z}
}

func (v Vector3[T]) Add(av Vector3[T]) Vector3[T] {
	return NewVector3(v.X+av.X, v.Y+av.Y, v.Z+av.Z)
}

func (v Vector3[T]) Sub(av Vector3[T]) Vector3[T] {
	return NewVector3(v.X-av.X, v.Y-av.Y, v.Z-av.Z)
}

func (v Vector3[T]) Dst2(av Vector3[T]) T {
	diff := v.Sub(av)

	return diff.X*diff.X + diff.Y*diff.Y + diff.Z*diff.Z
}
