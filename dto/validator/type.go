package validator

type Validator interface {
	AddEnum(
		key string,
		value ...string,
	) *validator

	AddDateFormat(
		key string,
		dateFormat string,
	) *validator

	AddRegex(
		key string,
		regex string,
		name string,
	) *validator

	BasicValidatorByTag(
		dto interface{},
		language string,
		menu string,
	) (
		err error,
	)
}

type regexValue struct {
	regex    string
	ruleName string
}
