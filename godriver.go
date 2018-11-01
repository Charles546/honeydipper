package main

import (
	"os/exec"
	"strings"
)

// GoDriver : a driver type that runs a golang program process
type GoDriver struct {
	Driver
	Package string
}

// NewGoDriver : create a driver object to run the golang program process
func NewGoDriver(data map[string]interface{}) GoDriver {
	driver := NewDriver(data)

	pack, ok := data["Package"].(string)
	if !ok {
		log.Panic("Package is not sepcified in driver")
	}

	if driver.Executable == "" {
		packParts := strings.Split(pack, "/")
		driver.Executable = packParts[len(packParts)-1]
	}

	godriver := GoDriver{
		Driver:  driver,
		Package: pack,
	}
	godriver.Type = "go"
	godriver.PreStart = godriver.preStart
	return godriver
}

func (g *GoDriver) preStart(service string, runtime *DriverRuntime) {
	log.Infof("[%s] pre-start dirver %s", service, runtime.meta.Name)
	check := execCommand("go", "list", g.Package)
	outp, err := check.CombinedOutput()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			install := execCommand("go", "get", g.Package)
			if outp, err := install.CombinedOutput(); err != nil {
				log.Panicf("[%s] Unable to install the go package for driver [%s] %+v", service, runtime.meta.Name, string(outp))
			}
		} else {
			log.Panicf("[%s] driver [%s] prestart failed %+v %+v", service, runtime.meta.Name, err, string(outp))
		}
	}
}
