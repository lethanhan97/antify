package fileserver

import "net/http"

func GetMediaFileServerHandler() http.Handler {
	fs := http.FileServer(http.Dir("./assets/media"))
	return http.StripPrefix("/video/", fs)
}

func GetHtmlFileServerHandler() http.Handler {
	fs := http.FileServer(http.Dir("./assets/html"))
	return http.StripPrefix("/", fs)
}
