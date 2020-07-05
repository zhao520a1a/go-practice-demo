package main

func main() {

}

/*
字节数组的比较
*/
func Campare(a, b []byte) int {
	//比较长度
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	//比较内容
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	return 0
}
