package linux

import (
	"strings"

	str "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
)

//Linux structure contains structures for manipulating internal liux systems
type Linux struct {
	root string
}

func isValidUsername(some string) bool {
	if strings.Contains(some, "root") {
		return false
	}
	for _, element := range strings.Split(some, "") {
		if element == "%" || element == "*" || element == "?" || element == ";" {
			return false
		}
	}
	return true
}

func (l *Linux) CreateLinuxUser(req str.LinuxUSE) error {

	return nil
}

func (l *Linux) OpenLinuxPort() error {

	return nil
}

func (l *Linux) CloseLinuxPort() error {

	return nil
}

func (l *Linux) ChangeLinuxPassword() error {

	return nil
}

func (l *Linux) ListLinuxUser() error {

	return nil
}
