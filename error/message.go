package error

import "errors"

var ErrUnauthorized = NewUnBundledErrorMessages(401, errors.New("E-1-CMD-SRV-006"), nil)
var ErrReservedValueString = NewUnBundledErrorMessages(400, errors.New("E-4-CMD-DTO-006"), errFieldNameConverter)
var ErrEmptyField = NewUnBundledErrorMessages(400, errors.New("E-4-CMD-DTO-001"), errFieldNameConverter)
var ErrUnknownData = NewUnBundledErrorMessages(400, errors.New("E-4-CMD-DTO-004"), errFieldNameConverter)
var ErrFormatFieldRule = NewUnBundledErrorMessages(400, errors.New("E-4-CMD-DTO-003"), errFieldRuleConverter)
var ErrFormatField = NewUnBundledErrorMessages(400, errors.New("E-4-CMD-DTO-002"), errFieldNameConverter)

var errFieldNameConverter = func(value ...interface{}) map[string]ErrorParam {
	result := make(map[string]ErrorParam)
	result["FieldName"] = ErrorParam{value[0], true}
	return result
}

var errFieldRuleConverter = func(value ...interface{}) map[string]ErrorParam {
	result := make(map[string]ErrorParam)
	result["FieldName"] = ErrorParam{value[0], true}
	result["RuleName"] = ErrorParam{value[1], true}
	result["Other"] = ErrorParam{value[2], false}
	return result
}
