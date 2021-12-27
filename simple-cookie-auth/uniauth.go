package scauth

import (
	"fmt"
	"sync"
)

var cm CookieManager

var accessControlEnabled = false

func init() {
	cm.Mutex = &sync.RWMutex{}
	cm.Mutex.Lock()
	cm.Map = make(map[string]Cookie)
	cm.Mutex.Unlock()
}

func EnableAccessControl() {
	fmt.Println("Warning: you are enabling access control. This will result in huge performance issues")
	accessControlEnabled = true
}

func NewCookie(id string) string {
	if accessControlEnabled {
		return cm.newCookieAC(id)
	}
	return cm.newCookie(id)
}

func CheckExists(cookie string) bool {
	if accessControlEnabled {
		return cm.checkExistsAC(cookie)
	}
	return cm.checkExists(cookie)
}

func GetID(cookie string) (bool, string) {
	if accessControlEnabled {
		return cm.getIDAC(cookie)
	}
	return cm.getID(cookie)
}

func DeleteCookie(cookie string) {
	cm.delete(cookie)
}
