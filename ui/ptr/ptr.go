package ptr

func Int(num int) *int {
	return &num
}

func String(str string) *string {
	return &str
}

func Bool(b bool) *bool {
	return &b
}
