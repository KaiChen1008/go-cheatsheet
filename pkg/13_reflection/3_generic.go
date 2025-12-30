package reflection

import "reflect"

// watch the video: https://youtu.be/pA724Jkh9Lw?si=-z_CoXcNbNC_gTAJ&t=205
type Number interface {
	int | string
}

func Square[T Number](n T) T {
	v := reflect.ValueOf(n)
	switch v.Kind() {
	case reflect.Int:
		return any(v.Int() * v.Int()).(T)
	case reflect.String:
		return any(v.String() + v.String()).(T)
	default: // <- no need
		var r T
		return r
	}
}
