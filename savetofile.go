package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"fmt"
	"crypto/sha256"
	"path/filepath"
)

func main() {
	err := realMain()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func realMain() error {
	if len(os.Args) != 7 {
		return nil
	}

	// savetofile <mode> <filepath> <uid> <gid> <perm> <data>
	mode := os.Args[1]
	path := os.Args[2]
	uid_arg := os.Args[3]
	gid_arg := os.Args[4]
	perm_arg := os.Args[5]
	data := os.Args[6]

  os.MkdirAll(filepath.Dir(path), os.ModePerm)

	uid, err := strconv.Atoi(uid_arg)
	if err != nil {
		return err
	}

	gid, err := strconv.Atoi(gid_arg)
	if err != nil {
		return err
	}

	perm_u, err := strconv.ParseUint(perm_arg, 8, 32)
	if err != nil {
		return err
	}

	perm := os.FileMode(perm_u)

	switch mode {
	case "append":
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, perm)
		if err != nil {
			return err
		}

		defer f.Close()

		if _, err = f.WriteString(data); err != nil {
			return err
		}
	case "append-nl":
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, perm)
		if err != nil {
			return err
		}

		defer f.Close()

		if _, err = f.WriteString(data); err != nil {
			return err
		}

		if _, err = f.WriteString("\n"); err != nil {
			return err
		}
	case "create-nl":
		err := ioutil.WriteFile(path, append([]byte(data), []byte("\n")...), perm)
		if err != nil {
			return err
		}
	default: // "create"
		err := ioutil.WriteFile(path, []byte(data), perm)
		if err != nil {
			return err
		}
	}

	err = os.Chown(path, uid, gid)
	if err != nil {
		return err
	}

	err = os.Chmod(path, perm)
	if err != nil {
		return err
	}

	fmt.Printf("%x", sha256.Sum256([]byte(data)))

	return nil
}
