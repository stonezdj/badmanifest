package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	location := flag.String("location", "/data/registry", "location of the registry")
	fmt.Printf("Going to scan location:%s\n", *location)
	flag.Parse()
	found := false
	vistor := func(path string, info os.DirEntry, err error) error {
		if strings.Contains(path, "repositories") &&
			strings.Contains(path, "_manifests/revisions/sha256/") &&
			strings.HasSuffix(path, "link") &&
			!info.IsDir() &&
			isBadLink(path) {
			fmt.Printf("Bad link found, path:%v\n", path)
			found = true
		}

		// fmt.Printf("%s\n", path)
		return nil
	}
	filepath.WalkDir(*location, vistor)
	if !found {
		fmt.Println("bad link not found ")
	}
}

func isBadLink(path string) bool {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return true
	}
	digest := string(content)
	if !strings.Contains(digest, "sha256:") {
		return true
	}
	return !checkFileExist(getBlobPath(getBasePath(filepath.Dir(path)), digest))
}

func checkFileExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir() && fileInfo.Size() > 0
}
func getBasePath(path string) string {
	return path[0:strings.Index(path, "/repositories/")]
}
func getBlobPath(basePath, digest string) string {
	if strings.HasPrefix(digest, "sha256:") {
		return basePath + "/blobs/sha256/" + digest[7:][0:2] + "/" + digest[7:] + "/data"
	}
	return basePath + "/blobs/sha256/" + digest[0:2] + "/" + digest + "/data"
}
