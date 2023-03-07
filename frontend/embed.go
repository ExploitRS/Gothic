package frontend

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

// test
// go:generate npm i
// go:generate npm run build
// go:ambed all:build
var files embed.FS

func PrintFsys() {
	fsys, err := fs.Sub(files, "build")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
	fmt.Println(fsys)
}

// func SvelteKitHandler() []string {
// 	fsys, err := fs.Sub(files, "build")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	filesystem := http.FS(fsys)

// 		path := strings.TrimPrefix(r.URL.Path, path)
// 		// try if file exists at path, if not append .html (SvelteKit adapter-static specific)
// 		_, err := filesystem.Open(path)
// 		if errors.Is(err, os.ErrNotExist) {
// 			path = fmt.Sprintf("%s.html", path)
// 		}

// 		r.URL.Path = path
// 		http.FileServer(filesystem).ServeHTTP(w, r)
// 	})
// }
