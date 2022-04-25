package license_key_formatting

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	s := "5F3Z-2e-9-w"
	k := 4

	require.Equal(t, "5F3Z-2E9W", LicenseKeyFormatting(s, k))
	require.Equal(t, "5F3Z-2E9W", licenseKeyFormatting(s, k))

	s = "2-5g-3-J"
	k = 2
	require.Equal(t, "2-5G-3J", LicenseKeyFormatting(s, k))
	require.Equal(t, "2-5G-3J", licenseKeyFormatting(s, k))
}
