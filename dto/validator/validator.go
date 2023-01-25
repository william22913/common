package validator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/william22913/common/constanta"
	errors "github.com/william22913/common/error"
	rgx "github.com/william22913/common/regex"
	"github.com/william22913/common/util"
)

func NewValidator() Validator {
	validator := &validator{}
	validator.build()
	return validator
}

type validator struct {
	enum       map[string][]string
	dateFormat map[string]string
	regex      map[string]regexValue
}

func (v *validator) AddEnum(key string, value ...string) *validator {
	v.enum[key] = value
	return v
}

func (v *validator) AddDateFormat(key string, dateFormat string) *validator {
	v.dateFormat[key] = dateFormat
	return v
}

func (v *validator) AddRegex(key string, regex string, name string) *validator {
	v.regex[key] = regexValue{
		regex:    regex,
		ruleName: name,
	}
	return v
}

func (v *validator) build() {
	v.enum = make(map[string][]string)
	v.enum["record_status"] = []string{"A", "N", "P"}
	v.enum["sharing_permission"] = []string{"edit", "view"}
	v.enum["boolean_permission"] = []string{"Y", "N"}

	v.dateFormat = make(map[string]string)
	v.dateFormat["default"] = constanta.DefaultTimeFormat
	v.dateFormat["date_only"] = constanta.DateOnlyTimeFormat

	v.regex = make(map[string]regexValue)
	v.regex["profile_name"] = regexValue{
		regex:    rgx.PROFILE_NAME,
		ruleName: "PROFILE_NAME_REGEX_MESSAGE",
	}
	v.regex["directory_name"] = regexValue{
		regex:    rgx.DIRECTORY_NAME,
		ruleName: "DIRECTORY_NAME_REGEX_MESSAGE",
	}
	v.regex["name"] = regexValue{
		regex:    rgx.NAME_STANDARD,
		ruleName: "NAME_REGEX_MESSAGE",
	}
	v.regex["text_only"] = regexValue{
		regex:    rgx.TEXT_ONLY,
		ruleName: "DESCRIPTION_REGEX_MESSAGE",
	}
	v.regex["alphanumeric"] = regexValue{
		regex:    rgx.ALPHANUMERIC,
		ruleName: "ALPHANUMERIC_REGEX",
	}
	v.regex["country_code"] = regexValue{
		regex:    rgx.COUNTRY_CODE,
		ruleName: "COUNTRY_CODE_REGEX",
	}
	v.regex["email"] = regexValue{
		regex:    rgx.EMAIL_REGEX,
		ruleName: "EMAIL_REGEX",
	}
	v.regex["phone"] = regexValue{
		regex:    rgx.PHONE_NUMBER_WITH_COUNTRY_CODE,
		ruleName: "PHONE_NUMBER_REGEX",
	}
	v.regex["additional_info"] = regexValue{
		regex:    rgx.ADDITIONAL_INFO,
		ruleName: "ADDITIONAL_INFO_REGEX",
	}
	v.regex["username"] = regexValue{
		regex:    rgx.USERNAME,
		ruleName: "USERNAME_REGEX",
	}
	v.regex["lowercase"] = regexValue{
		regex:    rgx.LOWERCASE,
		ruleName: "LOWERCASE_REGEX",
	}
	v.regex["upercase"] = regexValue{
		regex:    rgx.UPERCASE,
		ruleName: "UPERCASE_REGEX",
	}
	v.regex["npwp"] = regexValue{
		regex:    rgx.NPWP,
		ruleName: "NPWP_REGEX",
	}
	v.regex["nik"] = regexValue{
		regex:    rgx.NIK,
		ruleName: "NIK_REGEX",
	}
	v.regex["fax"] = regexValue{
		regex:    rgx.FAX,
		ruleName: "FAX_REGEX",
	}
	v.regex["lowercase_number"] = regexValue{
		regex:    rgx.LOWERCASE_AND_NUMBER,
		ruleName: "LOWERCASE_AND_NUMBER_REGEX",
	}
	v.regex["numeric"] = regexValue{
		regex:    rgx.LONG_NUMERIC,
		ruleName: "LONG_NUMERIC_REGEX",
	}
	v.regex["permission"] = regexValue{
		regex:    rgx.PERMISSION,
		ruleName: "PERMISSION_REGEX",
	}
	v.regex["profile_name"] = regexValue{
		regex:    rgx.PROFILE_NAME,
		ruleName: "PROFILE_NAME_REGEX",
	}
	v.regex["data_scope"] = regexValue{
		regex:    rgx.DATA_SCOPE,
		ruleName: "DATA_SCOPE_REGEX",
	}
	v.regex["directory_name"] = regexValue{
		regex:    rgx.DIRECTORY_NAME,
		ruleName: "DIRECTORY_NAME_REGEX",
	}
}

var (
	intType   = []string{reflect.Int64.String(), reflect.Int.String(), reflect.Int32.String()}
	floatType = []string{reflect.Float32.String(), reflect.Float64.String()}
)

func (v *validator) BasicValidatorByTag(
	dto interface{},
	language string,
	menu string,
) (
	err error,
) {
	reflectType := reflect.TypeOf(dto).Elem()
	reflectValue := reflect.ValueOf(dto).Elem()

	max := 0
	min := 0
	isMinFound := false
	isMaxFound := false

	for i := 0; i < reflectType.NumField(); i++ {
		currentField := reflectType.Field(i)
		currentValue := reflectValue.FieldByName(currentField.Name)

		if currentField.Name == "AbstractDTO" {
			continue
		}

		if currentField.Type.Kind() == reflect.Struct {
			newDTO := currentValue.Addr().Interface()
			err = v.BasicValidatorByTag(newDTO, menu, language)
			if err != nil {
				return
			}
		}

		required := currentField.Tag.Get("required")
		jsonField := strings.ToUpper(currentField.Tag.Get("json"))
		requiredArray := strings.Split(required, ",")
		reservedValues := currentField.Tag.Get("reserved")
		if reservedValues != "" {
			reservedValue := strings.Split(reservedValues, ",")
			if util.ValidateStringContainInStringArray(reservedValue, currentValue.String()) {
				err = errors.ErrReservedValueString.Language(language).Param(jsonField)
				return
			}
		}

		if util.ValidateStringContainInStringArray(requiredArray, menu) {
			defaultValue := currentField.Tag.Get("default")
			min, isMinFound, max, isMaxFound = v.getMinMaxValue(currentField)
			if util.ValidateStringContainInStringArray(intType, currentField.Type.String()) {
				if currentValue.IsZero() {
					valueIn, _ := strconv.Atoi(defaultValue)
					currentValue.SetInt(int64(valueIn))
				}

				value := currentValue.Int()
				if isMinFound {
					if min != 0 && int(value) == 0 {
						err = errors.ErrEmptyField.Language(language).Param(jsonField)
						return
					}
					if int(value) < min {
						return errors.ErrFormatFieldRule.Language(language).Param(jsonField, "NEED_MORE_THAN", strconv.Itoa(min))
					}
				}
				if isMaxFound {
					if int(value) > max {
						return errors.ErrFormatFieldRule.Language(language).Param(jsonField, "NEED_LESS_THAN", strconv.Itoa(max))
					}
				}
			} else if util.ValidateStringContainInStringArray(floatType, currentField.Type.String()) {
				if currentValue.IsZero() {
					valueIn, _ := strconv.ParseFloat(defaultValue, 64)
					currentValue.SetFloat(valueIn)
				}
				value := currentValue.Float()
				if isMinFound {
					if value < float64(min) {
						return errors.ErrFormatFieldRule.Language(language).Param(jsonField, "NEED_MORE_THAN", strconv.Itoa(min))
					}
				}
				if isMaxFound {
					if value > float64(max) {
						return errors.ErrFormatFieldRule.Language(language).Param(jsonField, "NEED_LESS_THAN", strconv.Itoa(max))
					}
				}
			} else if reflect.String.String() == currentField.Type.String() {
				currentValue.SetString(strings.Trim(currentValue.String(), " "))
				if currentValue.IsZero() {
					currentValue.SetString(defaultValue)
				}

				value := currentValue.String()
				err = v.ValidateMinMaxString(language, value, jsonField, min, max)
				if err != nil {
					return
				}

				enumField := currentField.Tag.Get("enum")
				if enumField != "" {
					if !util.ValidateStringContainInStringArray(v.enum[enumField], currentValue.String()) {
						err = errors.ErrUnknownData.Language(language).Param(jsonField)
						return
					}
				}

				// autoFix := currentField.Tag.Get("auto_fix")
				// if autoFix != "" {
				// 	if defaultFields.autoFix[autoFix] != nil {
				// 		defaultFields.autoFix[autoFix](currentValue)
				// 	}
				// }

				regexField := currentField.Tag.Get("regex")
				if regexField != "" {
					if v.regex[regexField].regex != "" {
						if len(value) > 0 || min != 0 {
							if !regexp.MustCompile(v.regex[regexField].regex).MatchString(currentValue.String()) {
								return errors.ErrFormatFieldRule.Language(language).Param(jsonField, v.regex[regexField].ruleName, "")
							}
						}

					}
				}
			} else if "time.Time" == currentField.Type.String() {
				var timeObject time.Time
				dateFormatTag := currentField.Tag.Get("dateFormat")
				strField := currentField.Name + "Str"

				timeFormatUsed := v.dateFormat["default"]
				if v.dateFormat[dateFormatTag] != "" {
					timeFormatUsed = v.dateFormat[dateFormatTag]
				}

				timeObject, err = v.TimeStrToTime(language, reflectValue.FieldByName(strField).String(), jsonField, timeFormatUsed)
				if err != nil {
					return
				}
				currentValue.Set(reflect.ValueOf(timeObject))
			} else if currentValue.Kind() == reflect.Slice {
				if isMinFound {
					if currentValue.Len() == 0 {
						return errors.ErrEmptyField.Language(language).Param(jsonField)
					}
					if currentValue.Len() < min {
						return errors.ErrFormatFieldRule.Language(language).Param(jsonField, "NEED_MORE_THAN", strconv.Itoa(min))
					}
				}
				if isMaxFound {
					if currentValue.Len() > max {
						return errors.ErrFormatFieldRule.Language(language).Param(jsonField, "NEED_LESS_THAN", strconv.Itoa(max))
					}
				}

				for i := 0; i < currentValue.Len(); i++ {
					temp := currentValue.Index(i)
					if temp.Type().String() == reflect.Struct.String() || temp.Type().Kind().String() == reflect.Struct.String() {
						newDTO := currentValue.Index(i).Addr().Interface()
						err = v.BasicValidatorByTag(newDTO, menu, language)
						if err != nil {
							return
						}
					}
				}
			}
		}
	}
	return
}

func (v *validator) getMinMaxValue(field reflect.StructField) (min int, isMinFound bool, max int, isMaxFound bool) {
	maxStr, isMaxFound := field.Tag.Lookup("max")
	minStr, isMinFound := field.Tag.Lookup("min")

	min, _ = strconv.Atoi(minStr)
	max, _ = strconv.Atoi(maxStr)

	return
}

func (v *validator) TimeStrToTime(
	language string,
	timeStr string,
	fieldName string,
	format string,
) (
	output time.Time,
	err error,
) {
	output, err = time.Parse(format, timeStr)
	if err != nil {
		err = errors.ErrFormatField.Language(language).Param(fieldName)
		return
	}

	return output, nil
}

func (v *validator) ValidateMinMaxString(
	language string,
	inputStr string,
	fieldName string,
	min int,
	max int,
) error {
	if min != 0 {
		if len(inputStr) == 0 {
			return errors.ErrEmptyField.Language(language).Param(fieldName)
		}
		if len(inputStr) < min {
			if min == 1 {
				return errors.ErrEmptyField.Language(language).Param(fieldName)
			} else {
				return errors.ErrFormatFieldRule.Language(language).Param(fieldName, "NEED_MORE_THAN", strconv.Itoa(min))
			}
		}
	}
	if max != 0 {
		if len(inputStr) > max {
			return errors.ErrFormatFieldRule.Language(language).Param(fieldName, "NEED_LESS_THAN", strconv.Itoa(max))
		}
	}

	return nil
}
