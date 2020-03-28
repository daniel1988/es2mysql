package Common

import (
    "io/ioutil"
    "os"
    "path"
)

func FilePutContents(file string, content string) error {
    f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

    defer f.Close()
    f.WriteString(content)
    return err
}

func FileGetContents(file string) (content string) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return ""
    }
    return string(data)
}

func Touch(file string) error {
    _, err := os.Stat(file)
    if err == nil {
        return nil
    }
    _, err = os.Create(file)
    return err
}

func DirName(file string) (base string) {
    return path.Base(file)
}

