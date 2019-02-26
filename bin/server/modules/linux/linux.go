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

//CreateLinuxUser --
func (l *Linux) CreateLinuxUser(req str.LinuxUSE) error {
	if !isValidUsername(req.UserName) {

	}
	return nil
}

//OpenLinuxPort -
func (l *Linux) OpenLinuxPort(req str.LinuxUSE) error {

	return nil
}

//CloseLinuxPort -
func (l *Linux) CloseLinuxPort(req str.LinuxUSE) error {

	return nil
}

//ChangeLinuxPassword -
func (l *Linux) ChangeLinuxPassword(req str.LinuxUSE) error {

	return nil
}

//ListLinuxUser -
func (l *Linux) ListLinuxUser(req str.LinuxUSE) error {

	return nil
}
