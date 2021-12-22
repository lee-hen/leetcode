package unique_email_addresses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumUniqueEmails(t *testing.T) {
	emails := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	require.Equal(t, 2, NumUniqueEmails(emails))

	emails = []string{"a@leetcode.com","b@leetcode.com","c@leetcode.com"}
	require.Equal(t, 3, NumUniqueEmails(emails))
}
