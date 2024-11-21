package uuids

import "github.com/google/uuid"

var (
	V4UUID    bool
	V5UUID    bool
	NilUUID   bool
	namespace = uuid.NameSpaceURL
	url       = []byte("https://jtprog.ru")
)

func GetUUID() string {
	switch {
	case V4UUID:
		return uuid.New().String()
	case V5UUID:
		return uuid.NewMD5(namespace, url).String()
	case NilUUID:
		return "00000000-0000-0000-0000-000000000000"
	default:
		return "00000000-0000-0000-0000-000000000000"
	}
}
