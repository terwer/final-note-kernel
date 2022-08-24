// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// SessionData represents the session.
type SessionData struct {
	UID     uint64 // user ID
	UName   string // username
	UB3Key  string // user B3 key
	URole   int    // user role
	UAvatar string // user avatar URL
	BID     uint64 // blog ID
	BURL    string // blog url
}

// AvatarURLWithSize returns avatar URL with the specified size.
func (sd *SessionData) AvatarURLWithSize(size int) string {
	return ImageSize(sd.UAvatar, size, size)
}

// Save saves the current session of the specified context.
func Save(w http.ResponseWriter, k string, v string) {
	value := map[string]string{
		k: v,
	}
	if encoded, err := cookieHandler.Encode(k, value); err == nil {
		cookie := &http.Cookie{
			Name:   k,
			Value:  encoded,
			Path:   "/",
			MaxAge: 3600,
		}
		http.SetCookie(w, cookie)
	}
}

// GetSession returns session of the specified context.
func GetSession(r *http.Request, key string) *string {
	var val string
	if cookie, err := r.Cookie(key); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode(key, cookie.Value, &cookieValue); err == nil {
			val = cookieValue["your-name"]
		}
	}
	return &val
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "your-name",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
