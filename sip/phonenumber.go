package sip

import (
	"github.com/nyaruka/phonenumbers"
)

// ExtractAreaCode extracts the area code from a phone number using the phonenumbers library
func ExtractAreaCode(phoneNumber string) string {
	// Parse the phone number without defaulting to any country
	num, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		// If parsing fails, fall back to empty string
		return ""
	}

	// Get the country code
	countryCode := phonenumbers.GetRegionCodeForNumber(num)

	// Only handle US numbers for now
	if countryCode != "US" {
		return ""
	}

	// Get the national number and extract first 3 digits (area code for US)
	nationalNumber := phonenumbers.GetNationalSignificantNumber(num)
	if len(nationalNumber) < 3 {
		return ""
	}
	return nationalNumber[:3]
}

// DetermineNumberType determines the phone number type using the phonenumbers library
func DetermineNumberType(phoneNumber string) string {
	// Parse the phone number without defaulting to any country
	num, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		// If parsing fails, fall back to "local"
		return "unknown"
	}

	numberType := phonenumbers.GetNumberType(num)

	// We are excluding a bunch of number types for now
	switch numberType {
	case phonenumbers.MOBILE:
		return "mobile"
	case phonenumbers.FIXED_LINE:
		return "local"
	case phonenumbers.TOLL_FREE:
		return "toll-free"
	default:
		return "unknown"
	}
}
