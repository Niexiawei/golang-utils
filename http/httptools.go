package http

import "net/http"

func CookiesMerge(old, new []*http.Cookie) []*http.Cookie {
	newMap := map[string]*http.Cookie{}
	oldMap := map[string]*http.Cookie{}
	var resCookies []*http.Cookie
	for idx, v := range new {
		newMap[v.Name] = new[idx]
	}
	for idx, v := range old {
		oldMap[v.Name] = old[idx]
	}

	for k, v := range newMap {
		oldMap[k] = v
	}

	for _, v := range oldMap {
		resCookies = append(resCookies, v)
	}

	return resCookies
}
