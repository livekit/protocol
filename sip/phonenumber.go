package sip

import (
	"strings"

	"github.com/livekit/protocol/livekit"
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
func DetermineNumberType(phoneNumber string) livekit.PhoneNumberType {
	// Parse the phone number without defaulting to any country
	num, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		// If parsing fails, fall back to unknown
		return livekit.PhoneNumberType_PHONE_NUMBER_TYPE_UNKNOWN
	}

	numberType := phonenumbers.GetNumberType(num)

	// We are excluding a bunch of number types for now
	switch numberType {
	case phonenumbers.MOBILE:
		return livekit.PhoneNumberType_PHONE_NUMBER_TYPE_MOBILE
	case phonenumbers.FIXED_LINE:
		return livekit.PhoneNumberType_PHONE_NUMBER_TYPE_LOCAL
	case phonenumbers.FIXED_LINE_OR_MOBILE:
		return livekit.PhoneNumberType_PHONE_NUMBER_TYPE_LOCAL
	case phonenumbers.TOLL_FREE:
		return livekit.PhoneNumberType_PHONE_NUMBER_TYPE_TOLL_FREE
	default:
		return livekit.PhoneNumberType_PHONE_NUMBER_TYPE_UNKNOWN
	}
}

// IsEmergencyNumber reports whether a dialed number reaches the 911 emergency
// service. It accepts the bare number and the +1 or 1 prefixed forms.
func IsEmergencyNumber(number string) bool {
	digits := strings.Map(func(r rune) rune {
		if strings.ContainsRune("+ -()", r) {
			return -1
		}
		return r
	}, number)
	return digits == "911" || digits == "1911"
}

// URIUser returns the user part of a sip, sips or tel URI, with the display
// name, angle brackets, host and parameters removed. A URI with no user part
// and anything that is not a URI returns an empty string.
func URIUser(uri string) string {
	if i := strings.Index(uri, "<"); i >= 0 {
		uri = uri[i+1:]
		if j := strings.Index(uri, ">"); j >= 0 {
			uri = uri[:j]
		}
	}
	var user string
	switch {
	case strings.HasPrefix(uri, "sip:"), strings.HasPrefix(uri, "sips:"):
		uri = uri[strings.Index(uri, ":")+1:]
		at := strings.Index(uri, "@")
		if at < 0 {
			return ""
		}
		user = uri[:at]
	case strings.HasPrefix(uri, "tel:"):
		user = uri[len("tel:"):]
	default:
		return ""
	}
	if i := strings.IndexAny(user, ";?"); i >= 0 {
		user = user[:i]
	}
	return user
}
