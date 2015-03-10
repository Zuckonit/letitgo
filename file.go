// Package main provides ...
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "errors"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func read(filename string) string{
    dat, err := ioutil.ReadFile(filename)
    check(err)
    return string(dat)
}

func readBytes(filename string, b int) (string, error){
    fp, err := os.Open(filename)
    defer fp.Close()

    if err == nil {
        bytes := make([]byte, b)
        _, err := fp.Read(bytes)
        if err == nil {
            return string(bytes), nil
        }
        return "", err
    }
    return "", errors.New("error")
}

func peekBytes(filename string, b int) (string, error) {
    fp, err := os.Open(filename)
    defer fp.Close()

    if err == nil {
        nb := bufio.NewReader(fp)
        bytes, err := nb.Peek(b)
        if err == nil {
            return string(bytes), nil
        }
        return "", err
    }
    return "", errors.New("error")
}

func writeString(filename string, content string) {
    err := ioutil.WriteFile(filename, []byte(content), 0644)
    check(err)
}

func writeString2(filename string, content string) (int, error) {
    fp, err := os.Create(filename)
    defer fp.Close()

    ret, err := fp.Write([]byte(content))
    if err == nil {
        return ret, nil
    } 
    return 0, err
}

func writeBuffer(filename string, content string) (int, error) {
    fp, err := os.Create(filename)
    defer fp.Close()
    if err == nil {
        buf := bufio.NewWriter(fp)
        ret, err := buf.WriteString(content)
        buf.Flush()
        if err == nil {
            return ret, nil
        }
        return 0, err
    }
    return 0, err
}

func main() {
    filename := "/tmp/dat"
    readStr  := read(filename)
    readByte, _ := readBytes(filename, 2)
    peekByte, _ := peekBytes(filename, 2)

    fmt.Println(readStr)
    fmt.Println(readByte)
    fmt.Println(peekByte)

    writeString(filename, "hello")
    writeString2(filename, "world")
    writeBuffer(filename, "golang")
    rs := read(filename)
    fmt.Println(rs)
}
