package common

func OptionalInt(i ...int) *int {
	if len(i) == 0 {
		return nil
	}
	result := i[0]
	return &result
}

func OptionalIntWithDefault(def int, i ...int) int {
	if result := OptionalInt(i...); result != nil {
		return *result
	}
	// default
	return def
}

func OptionalString(i ...string) *string {
	if len(i) == 0 {
		return nil
	}
	result := i[0]
	return &result
}

func OptionalStringWithDefault(def string, i ...string) string {
	if result := OptionalString(i...); result != nil {
		return *result
	}
	// default
	return def
}
