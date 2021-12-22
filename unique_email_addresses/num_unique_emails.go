package unique_email_addresses

import (
	"strings"
)

func NumUniqueEmails(emails []string) int {
	domains := make(map[string][]string)
	for i := range emails {
		address := strings.Split(emails[i], "@")
		idx := strings.Index(address[0], "+")
		if idx == -1 {
			idx = len(address[0])
		}
		localName := strings.Replace(address[0][:idx], ".", "", -1)
		if localNames, ok := domains[address[1]]; ok && localNames[len(localNames)-1] == localName {
			continue
		}
		domains[address[1]] = append(domains[address[1]], localName)
	}

	var count int
	for _, localNames := range domains{
		count += len(localNames)
	}
	return count
}
