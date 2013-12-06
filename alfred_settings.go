package Alfred

import (
    // "errors"
    "fmt"
    "github.com/mkrautz/plist"
    "io/ioutil"
    "os"
    "strings"
)

type AlfredSettings map[string]interface{}

func (ga *GoAlfred) Set(key, value string) (err error) {
    settings, err := ga.loadSettings()
    if err != nil {
        fmt.Printf("err type %T, %v\n", err, err)
        return err
    }
    // fmt.Println(settings, key, value)
    settings[key] = value
    b, err := plist.Marshal(settings)
    if err != nil {
        return err
    }
    err = ga.saveSettings(b)
    if err != nil {
        return err
    }
    return
}

func (ga *GoAlfred) Get(key string) (string, error) {
    settings, err := ga.loadSettings()
    if err != nil {
        return "", err
    }
    v, ok := settings[key]
    if !ok {
        return "", err
    } else {
        return v.(string), err
    }
}

func (ga *GoAlfred) loadSettings() (settings map[string]interface{}, err error) {
    settings = make(map[string]interface{})
    buf, err := ioutil.ReadFile(ga.SettingsFN)
    if err != nil && strings.Contains(err.Error(), "no such file") { // first time?
        file, err := os.Create(ga.SettingsFN) // Create the file
        defer file.Close()
        if err != nil {
            return settings, err
        }
        file.WriteString(emptyPlist)
        return settings, err
    } else if err != nil {
        return nil, err
    }
    err = plist.Unmarshal(buf, &settings)
    if err != nil {
        return nil, err
    }
    return
}

func (ga *GoAlfred) saveSettings(b []byte) (err error) {
    file, err := os.OpenFile(ga.SettingsFN, os.O_CREATE|os.O_WRONLY, 0666)
    defer file.Close()
    if err != nil {
        return err
    }
    _, err = file.WriteString(string(b))
    if err != nil {
        return err
    }
    return
}

var emptyPlist = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0"></plist>`
