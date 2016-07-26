package environment

import (
	"github.com/DataDrake/unpackage/config"
	"os"
	"errors"
	"strings"
	"io/ioutil"
	"path/filepath"
)

func ModuleFileWalk(path, version string) (location string, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		err = errors.New("Could not read Module directory '" + path + "', reason: " + err.Error())
		return
	}
	if len(version) == 0 && len(files) > 0{
		location = filepath.Join(path,files[0].Name())
		return
	}
	for _,f := range files {
		if !f.IsDir() && f.Name() == version {
			location = filepath.Join(path,version)
			return
		}
	}
	return
}

func ModuleWalk(path, name, version string) (location string, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		err = errors.New("Could not read Module directory '" + path + "', reason: " + err.Error())
		return
	}
	for _,f := range files {
		if f.IsDir() {
			if f.Name() == name {
				location, err = ModuleFileWalk(filepath.Join(path,f.Name()),version)
			} else {
				location, err = ModuleWalk(filepath.Join(path, f.Name()), name, version)
			}
			if err != nil {
				return
			}
			if len(location) > 0 {
				return
			}
		}
	}
	return
}

func FindModule(name string) (c *config.Config, err error) {
	module := strings.Split(name,"/")
	version := ""
	if len(module) > 1 {
		name = module[0]
		version = module[1]
	}
	mp := os.Getenv("MODULE_PATH")
	if len(mp) == 0 {
		err = errors.New("MODULE_PATH empty or not set")
		return
	}
	mps := strings.Split(mp,":")
	var location string
	for _,p := range mps {
		location, err = ModuleWalk(p,name,version)
		if err != nil {
			return
		}
		if len(location) != 0 {
			return
		}
	}
	err = errors.New("Module '" + name + "' not found")
	return
}
