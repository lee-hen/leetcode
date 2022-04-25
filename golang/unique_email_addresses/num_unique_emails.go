package unique_email_addresses

import (
	"strings"
)

func NumUniqueEmails(emails []string) int {
	domains := make(map[string]map[string]struct{})
	for i := range emails {
		address := strings.Split(emails[i], "@")
		idx := strings.Index(address[0], "+")
		if idx == -1 {
			idx = len(address[0])
		}
		localName := strings.Replace(address[0][:idx], ".", "", -1)
		if domains[address[1]] == nil {
			domains[address[1]] = make(map[string]struct{})
		}
		domains[address[1]][localName] = struct{}{}
	}

	var count int
	for _, localNames := range domains{
		count += len(localNames)
	}
	return count
}
