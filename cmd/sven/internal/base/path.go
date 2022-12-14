package base

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func SvenHome() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	home := path.Join(dir, ".sven")
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}
	return home
}

func SvenHomeWithDir(dir string) string {
	home := path.Join(SvenHome(), dir)
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}
	return home
}

func copyFile(src, dst string, replaces []string) error {
	var err error
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	buf, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	var old string
	for i, next := range replaces {
		if i%2 == 0 {
			old = next
			continue
		}

		//fmt.Println("old",old)
		//fmt.Println("next",next)
		buf = bytes.ReplaceAll(buf, []byte(old), []byte(next))
		//fmt.Printf("%s\n", buf)
	}
	return os.WriteFile(dst, buf, srcinfo.Mode())
}

func copyDir(src, dst string, replaces, ignores []string) error {
	//fmt.Println("src:",src)
	//fmt.Println("dst:",dst)

	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	//Stat返回描述文件f的FileInfo类型值。如果出错，错误底层类型是*PathError。
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		if hasSets(fd.Name(), ignores) {
			continue
		}
		srcfp := path.Join(src, fd.Name())
		//fmt.Println("fd.name",fd.Name())
		//fmt.Println("srcfp:",srcfp)
		dstfp := path.Join(dst, fd.Name())
		//fmt.Println("dstfp:",dstfp)

		var e error
		if fd.IsDir() {
			e = copyDir(srcfp, dstfp, replaces, ignores)
		} else {
			e = copyFile(srcfp, dstfp, replaces)
		}
		if e != nil {
			return e
		}
	}
	return nil
}


func hasSets(name string, sets []string) bool {
	for _, ig := range sets {
		if ig == name {
			return true
		}
	}
	return false
}

func Tree(path string, dir string) {
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() {
			fmt.Printf("%s %s (%v bytes)\n", color.GreenString("CREATED"), strings.Replace(path, dir+"/", "", -1), info.Size())
		}
		return nil
	})
}
