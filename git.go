package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func hasGitRepo() (bool, error) {
	_, err := os.Stat(path.Join(wikiDir, ".git"))
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func copier(dir string) error {
	err := filepath.Walk(dir, func(curPath string, info os.FileInfo, err error) error {
		if !info.IsDir() {

			from, err := os.Open(curPath)
			if err != nil {
				return err
			}
			defer from.Close()

			destPath := path.Join(wikiDir, strings.TrimPrefix(curPath, dir))
			destDir, _ := filepath.Split(destPath)
			if destDir != "" {
				if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
					return err
				}
			}

			to, err := os.OpenFile(destPath, os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				return err
			}
			defer to.Close()

			_, err = io.Copy(to, from)
			if err != nil {
				return err
			}

		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func gitInit() error {
	hasGitRepo, err := hasGitRepo()
	if err != nil {
		return err
	}
	if hasGitRepo {
		return nil
	}

	_, err = gitInvoke("init")
	_, err = gitInvoke("config", "user.name", "Kele")
	_, err = gitInvoke("config", "user.email", "app@kele.wiki")
	if err != nil {
		return err
	}

	err = copier(dataDir)
	if err != nil {
		fmt.Printf("err.Error() = %+v\n", err.Error())
		return err
	}

	_, err = gitInvoke("add", ".")
	_, err = gitInvoke("commit", "-a", "-m", "initial commit by Kele!")
	if err != nil {
		return err
	}
	return nil
}

func InitHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		_ = gitInit()
	}
}
