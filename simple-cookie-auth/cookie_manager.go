package scauth

import (
	"github.com/foxclore/go-utils/crand"
	"github.com/foxclore/go-utils/hash"
	"sync"
	"time"
)

type CookieManager struct {
	Mutex *sync.RWMutex
	Map   map[string]Cookie
}

func (cm *CookieManager) newCookie(id string) string {
	cm.Mutex.Lock()
	cookie := hash.Hash(crand.Random(10), time.Now().Unix())
	cm.Map[cookie] = Cookie{
		Created: time.Now().Unix(),
		ID:      id,
	}
	cm.Mutex.Unlock()
	return cookie
}

func (cm *CookieManager) checkExists(cookie string) bool {
	cm.Mutex.RLock()
	_, exists := cm.Map[cookie]
	cm.Mutex.RUnlock()
	return exists
}

func (cm *CookieManager) getID(cookie string) (bool, string) {
	cm.Mutex.RLock()
	c, exists := cm.Map[cookie]
	cm.Mutex.RUnlock()
	return exists, c.ID
}

func (cm *CookieManager) delete(cookie string) {
	cm.Mutex.Lock()
	delete(cm.Map, cookie)
	cm.Mutex.Unlock()
}

func (cm *CookieManager) newCookieAC(id string) string {
	cm.Mutex.Lock()
	cookie := hash.Hash(crand.Random(10), time.Now().Unix())
	cm.Map[cookie] = Cookie{
		Created:      time.Now().Unix(),
		ID:           id,
		LastAccessed: time.Now().Unix(),
	}
	cm.Mutex.Unlock()
	return cookie
}

func (cm *CookieManager) checkExistsAC(cookie string) bool {
	cm.Mutex.Lock()
	c, exists := cm.Map[cookie]
	if exists {
		cm.Map[cookie] = Cookie{
			Created:      c.Created,
			LastAccessed: time.Now().Unix(),
			ID:           c.ID,
		}
	}
	cm.Mutex.Unlock()
	return exists
}

func (cm *CookieManager) getIDAC(cookie string) (bool, string) {
	cm.Mutex.RLock()
	c, exists := cm.Map[cookie]
	if exists {
		cm.Map[cookie] = Cookie{
			Created:      c.Created,
			LastAccessed: time.Now().Unix(),
			ID:           c.ID,
		}
	}
	cm.Mutex.Unlock()
	return exists, c.ID
}
