package license_key_formatting

import (
	"strings"
)

func LicenseKeyFormatting(s string, k int) string {
	sb := new(strings.Builder)
	for j := len(s)-1; j >= 0; j-- {
		if s[j] != '-' {
			if sb.Len() % (k + 1)  == k {
				sb.WriteString("-")
			}

			sb.WriteByte(s[j])
		}
	}

	return strings.ToUpper(ReverseString([]byte(sb.String())))
}

func ReverseString(bytes []byte) string {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func licenseKeyFormatting(s string, k int) string {
	str := strings.Join(strings.Split(s, "-"), "")
	strGroup := make([]string, 0)
	eachStr := make([]byte, k)
	var cnt int
	for j := len(str)-1; j >= 0; j-- {
		if cnt < k {
			if str[j] >= 'a' && str[j] <= 'z' {
				eachStr[k-cnt-1] = str[j] - ('a' - 'A')
			} else {
				eachStr[k-cnt-1] = str[j]
			}

			cnt++
		}

		if cnt == k || j == 0 {
			strGroup = append(strGroup, string(eachStr[k-cnt:]))
			eachStr = make([]byte, k)
			cnt = 0
		}
	}

	formatStr := new(strings.Builder)
	for j := len(strGroup)-1; j >= 0; j-- {
		formatStr.WriteString(strGroup[j])
		if j != 0 {
			formatStr.WriteString("-")
		}
	}

	return formatStr.String()
}
