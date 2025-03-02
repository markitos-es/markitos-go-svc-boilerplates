package domain

import (
	"crypto/rand"
	"fmt"
	"regexp"
)

func IsUUIDv4(uuid string) bool {
	uuidRegex := `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	matched, err := regexp.MatchString(uuidRegex, uuid)

	return err == nil && matched
}

func UUIDv4() string {
	var uuid [16]byte
	rand.Read(uuid[:])

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
