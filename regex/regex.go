package regex

const EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
const PHONE_NUMBER_WITH_COUNTRY_CODE = "[+][0-9]+[-][1-9][0-9]{9,12}"
const COUNTRY_CODE = "^[+][0-9]{1,3}$"
const IP_PRIVATE_LOCALHOST = "^https?://localhost.*"
const IP_PRIVATE_192 = "^https?://192.168..*"
const IP_PRIVATE_OTHER = "^https?://fd[0-9a-z]+:.*"
const NAME_STANDARD = "^[A-Z][a-z]+(([ ][A-Z][a-z])?[a-z]*)*$"
const ADDITIONAL_INFO = "^[a-z][a-z0-9_.-]+$"
const USERNAME = "^[a-z][a-z0-9_.]+$"
const LOWERCASE = "^[a-z]+$"
const UPERCASE = "^[A-Z]+$"
const LOWERCASE_AND_NUMBER = "^[a-z][a-z0-9]+$"
const LONG_NUMERIC = "^[0-9]+$"
const ALPHANUMERIC = "^[A-Za-z0-9]+$"
const PERMISSION = "^[a-z]+[a-z._-]+[a-z1-9]+[:][a-z]+[a-z-][a-z]+$"
const NPWP = "^[0-9]{2}[.][0-9]{3}[.][0-9]{3}[.][0-9][-][0-9]{3}[.][0-9]{3}$"
const NIK = "^[0-9]{16,20}$"
const FAX = "^[0-9]{5,25}$"
const PROFILE_NAME = "^[A-Z0-9](?:|(?:[a-z0-9]+|(?:[a-z0-9]|[a-z0-9])(?:([_-]|)[a-z0-9])+)|[ ]([A-Z0-9](?:|(?:[a-z0-9]+|(?:[a-z0-9]|[a-z0-9])(?:([_-]|)[a-z0-9])+))+)+)+$"
const DATA_SCOPE = "^(nexsoft[.]([a-z]+[a-z._-]+[a-z1-9]))[:]([1-9][0-9]*|all)$"
const DIRECTORY_NAME = "^(([a-z0-9](?:([_]|)[a-z0-9])+))$"
const TEXT_ONLY = "^[A-Za-z ]+$"
