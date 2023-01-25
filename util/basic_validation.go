package util

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/william22913/common/regex"
)

func IsMacAddress(input string) (output bool) {
	temp := strings.Split(input, ":")
	if len(temp) == 6 {
		output = true
		for i := 0; i < len(temp); i++ {
			if len(temp[i]) != 2 {
				output = false
			}
		}
		return
	}
	return false
}

func IsEmailAddress(input string) (output bool) {
	emailRegexp := regexp.MustCompile(regex.EMAIL_REGEX)
	return emailRegexp.MatchString(input)
}

func IsPhoneNumber(input string) (number int, isValid bool) {
	number, isValid = IsNumeric(input)
	return number, len(input) <= 13 && isValid
}
func IsPhoneNumberWithCountryCode(input string) bool {
	phoneNumberRegexp := regexp.MustCompile(regex.PHONE_NUMBER_WITH_COUNTRY_CODE)
	return phoneNumberRegexp.MatchString(input)
}
func IsCountryCode(input string) bool {
	countryCodeRegexp := regexp.MustCompile(regex.COUNTRY_CODE)
	return countryCodeRegexp.MatchString(input)
}
func IsIPPrivate(input string) (output bool) {
	ipPrivateRegexp := regexp.MustCompile(regex.IP_PRIVATE_LOCALHOST)
	if ipPrivateRegexp.MatchString(input) {
		return true
	}
	ipPrivateRegexp = regexp.MustCompile(regex.IP_PRIVATE_192)
	if ipPrivateRegexp.MatchString(input) {
		return true
	}

	ipPrivateRegexp = regexp.MustCompile(regex.IP_PRIVATE_OTHER)
	return ipPrivateRegexp.MatchString(input)
}

func IsNexsoftPasswordStandardValid(password string) (bool, string, string) {
	if len(password) < 8 {
		return false, "NEED_MORE_THAN", "8"
	} else if len(password) >= 8 && len(password) <= 50 {
	next:
		for name, classes := range map[string][]*unicode.RangeTable{
			"UPPERCASE": {unicode.Upper, unicode.Title},
			"LOWERCASE": {unicode.Lower},
			"NUMERIC":   {unicode.Number, unicode.Digit},
			"SPECIAL":   {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
		} {
			for _, r := range password {
				if unicode.IsOneOf(classes, r) {
					continue next
				}
			}
			return false, name, ""
		}
	} else {
		return false, "NEED_LESS_THAN", "50"
	}
	return true, "", ""
}

func IsNexsoftUsernameStandardValid(username string) (bool, string, string) {
	if len(username) < 6 {
		return false, "NEED_MORE_THAN", "6"
	} else if len(username) > 20 {
		return false, "NEED_LESS_THAN", "20"
	} else {
		usernameRegex := regexp.MustCompile(regex.USERNAME)
		return usernameRegex.MatchString(username), "USERNAME_REGEX_MESSAGE", ""
	}
}

func IsNexsoftNameStandardValid(username string) (bool, string, string) {
	usernameRegex := regexp.MustCompile(regex.NAME_STANDARD)
	return usernameRegex.MatchString(username), "NAME_REGEX_MESSAGE", ""
}

func IsNexsoftAdditionalInformationKeyStandardValid(username string) (bool, string) {
	usernameRegex := regexp.MustCompile(regex.ADDITIONAL_INFO)
	return usernameRegex.MatchString(username), "ADDITIONAL_INFO_REGEX"
}

func IsOnlyContainLowerCase(username string) (bool, string) {
	usernameRegex := regexp.MustCompile(regex.LOWERCASE)
	return usernameRegex.MatchString(username), "LOWERCASE_REGEX"
}

func IsOnlyContainLowerCaseAndNumber(username string) (bool, string) {
	usernameRegex := regexp.MustCompile(regex.LOWERCASE_AND_NUMBER)
	return usernameRegex.MatchString(username), "LOWERCASE_AND_NUMBER_REGEX"
}

func IsStringEmpty(input string) bool {
	return input == ""
}

func IsTimestampValid(input string) (bool, string) {
	format := "2006-01-02T15:04:05.999999999"
	timestamp, err := time.Parse(format, input)

	if err != nil {
		return false, ""
	} else {
		return true, strings.Replace(timestamp.UTC().Format(time.RFC3339Nano), "Z", "", -1)
	}
}

func IsNumeric(input string) (int, bool) {
	result, err := strconv.Atoi(input)
	if err != nil {
		return -1, false
	} else {
		return result, true
	}
}

func IsNexsoftPermissionStandardValid(permission string) (bool, string) {
	permissionRegex := regexp.MustCompile(regex.PERMISSION)
	return permissionRegex.MatchString(permission), "PERMISSION_REGEX_MESSAGE"
}

func IsNPWPValid(npwp string) bool {
	npwpRegex := regexp.MustCompile(regex.NPWP)
	return npwpRegex.MatchString(npwp)
}

func IsNIKValid(nik string) bool {
	nikRegex := regexp.MustCompile(regex.NIK)
	return nikRegex.MatchString(nik)
}

func IsFacsimileValid(fax string) bool {
	facsimileRegex := regexp.MustCompile(regex.FAX)
	return facsimileRegex.MatchString(fax)
}

func IsNexsoftProfileNameStandardValid(profileName string) (bool, string) {
	NameOrTitle := regexp.MustCompile(regex.PROFILE_NAME)
	return NameOrTitle.MatchString(profileName), "PROFILE_NAME_REGEX_MESSAGE"
}

func IsNameWithUppercaseValid(input string) bool {
	UppercaseRegex := regexp.MustCompile(regex.UPERCASE)
	return UppercaseRegex.MatchString(input)
}

func IsLongNumeric(input string) bool {
	longNumericRegex := regexp.MustCompile(regex.LONG_NUMERIC)
	return longNumericRegex.MatchString(input)
}

func IsDataScopeValid(input string) (bool, string) {
	dataScopeRegexp := regexp.MustCompile(regex.DATA_SCOPE)
	return dataScopeRegexp.MatchString(input), "DATA_SCOPE"
}

func IsNexsoftDirectoryNameStandardValid(profileName string) (bool, string) {
	NameOrTitle := regexp.MustCompile(regex.DIRECTORY_NAME)
	return NameOrTitle.MatchString(profileName), "DIRECTORY_NAME_REGEX_MESSAGE"
}
