package util

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// 返回当前时间
func GetDate() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}

// 获取当前系统环境
func GetRunTime() string {
	//获取系统环境变量
	RUN_TIME := os.Getenv("RUN_TIME")
	if RUN_TIME == "" {
		fmt.Println("No RUN_TIME Can't start")
	}
	return RUN_TIME
}

// MD5 加密字符串
func GetMD5(plainText string) string {
	h := md5.New()
	h.Write([]byte(plainText))
	return hex.EncodeToString(h.Sum(nil))
}

//计算文件的md5，适用于本地文件计算
func GetMd5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(md5hash.Sum(nil)), nil
}

//从流中直接读取数据计算md5 并返回流的副本，不能用于计算大文件流否则内存占用很大
//@return io.Reader @params file的副本
func GetMd52(file io.Reader) (io.Reader, string, error) {
	var b bytes.Buffer
	md5hash := md5.New()
	if _, err := io.Copy(&b, io.TeeReader(file, md5hash)); err != nil {
		return nil, "", err
	}
	return &b, hex.EncodeToString(md5hash.Sum(nil)), nil
}

//解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func getDir(path string) string {
	return SubString(path, 0, strings.LastIndex(path, "/"))
}

func SubString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func UploadFile(file *multipart.FileHeader, path string) (string, error) {
	if reflect.ValueOf(file).IsNil() || !reflect.ValueOf(file).IsValid() {
		return "", errors.New("invalid memory address or nil pointer dereference")
	}
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return "", err
	}
	err = MkDir(path)
	if err != nil {
		return "", err
	}
	// Destination
	// 去除空格
	filename := strings.Replace(file.Filename, " ", "", -1)
	// 去除换行符
	filename = strings.Replace(filename, "\n", "", -1)

	dst, err := os.Create(path + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return filename, nil
}

func GetFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	//文件大小
	fsize := fileInfo.Size()
	return fsize, nil
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func MaxIntArr(vals...int64) (int,int64) {
	maxVal := vals[0]
	maxValIndex := 0
	for key, val := range vals {
		if val > maxVal {
			maxVal = val
			maxValIndex = key
		}
	}
	return maxValIndex,maxVal
}


func MinIntArr(vals...int64) (int,int64) {
	minVal  := vals[0]
	minValIndex := 0
	for key, val := range vals {
		if minVal == 0 || val <= minVal {
			minVal = val
			minValIndex = key
		}
	}
	return minValIndex,minVal
}

// 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// 判断是不是真实手机号码
func IsMobile(mobile string) bool {
	result, _ := regexp.MatchString(`^(1\d{10})$`, mobile)
	if result {
		return true
	} else {
		return false
	}
}

// 合并字符串
func StrCombine(str ...string) string {
	var bt bytes.Buffer
	for _, arg := range str {
		bt.WriteString(arg)
	}
	//获得拼接后的字符串
	return bt.String()
}