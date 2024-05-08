package util

func StrToNilIfEmpty(str string) any {
	if str == "" {
		return nil
	}
	return str
}

func StrToPointer(str string) *string {
	return &str
}

func BoolToPointer(bo bool) *bool {
	return &bo
}
