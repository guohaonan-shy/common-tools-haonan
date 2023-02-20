package generic

// generic try

func Contains[T comparable](target T, quene []T) bool {
	for _, element := range quene {
		if target == element {
			return true
		}
	}
	return false
}

func ContainsByFunc[T comparable](fn func(value T) bool, quene []T) bool {
	for _, element := range quene {
		if fn(element) {
			return true
		}
	}
	return false
}

func SortByFunc[T comparable](quene []T, lessFn func(a, b T) bool) []T {
	length := len(quene)

	for i := 0; i < length-1; i++ {
		if lessFn(quene[i], quene[i+1]) {
			continue
		} else {
			tmp := quene[i]
			quene[i] = quene[i+1]
			quene[i+1] = tmp
		}
	}
	return quene
}

// FilterByFunc self-define value函数
func FilterByFunc[Key comparable, Value any](kv map[Key]Value, fn func(value any) bool) map[Key]Value {
	for k, v := range kv {
		if fn(v) {
			delete(kv, k)
		}
	}
	return kv
}

func FilterDuplicate[T comparable](list []T) []T {
	res, set := make([]T, 0), make(map[T]struct{}, 0)
	for _, element := range list {
		if _, ok := set[element]; !ok {
			res = append(res, element)
			set[element] = struct{}{}
		}
	}
	return res
}
