package urlshort

import "net/http"

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	
	// HandlerFunc - adapter to allow the use of ordinary functions as HTTP handlers
	return func(w http.ResponseWriter, r *http.Request) {
		// get path from request
		path := r.URL.Path
		
		// valid path - redirect 
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		
		// invalid path - fallback
		fallback.ServeHTTP(w, r)
	}
}

// func 

