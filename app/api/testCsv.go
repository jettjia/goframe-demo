package api

import (
	"encoding/csv"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"io"
	"log"
	"my-app/library/response"
	"os"
	"path"
)

// 注册控制器
var TestCsv = testCsvApi{}

type testCsvApi struct{}

type UserTest struct {
	ID   string
	Name string
	Tel  string
	Addr string
}

// csv down
func (*testCsvApi) DownCsv(r *ghttp.Request) {
	Users1 := []UserTest{
		{"1", "贾直接", "adminTel", "武汉硚口1号"},
		{"2", "test", "testTel", "testAddr"},
	}
	StructToCsv("tmp/user.csv", Users1)

	response.JsonExit(r, 0, "ok")
}

// csv up
func (*testCsvApi) UpCsv(r *ghttp.Request) {
	file := r.GetUploadFile("file")

	savePath := "tmp/"
	saveFileName, err := file.Save(savePath, true)

	if err != nil {
		response.JsonExit(r, 1, "err", err.Error())
	}

	filename := file.Filename

	fmt.Println("上传文件名称是：", filename)
	fmt.Println("========================")
	fmt.Println("上传后缀是：", path.Ext(filename))
	fmt.Println("========================")
	fmt.Println("上传文件大小是：", file.Size)

	fmt.Println("============")
	fmt.Println("保存后的文件是：：", saveFileName)

	CsvToUp(savePath + saveFileName)

	response.JsonExit(r, 0, "ok")
}

// 追加写入
func StructToCsv(filename string, UsersDb []UserTest) {
	newFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		newFile.Close()
	}()
	// 写入UTF-8
	newFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	// 写数据到csv文件
	w := csv.NewWriter(newFile)
	header := []string{"ID", "Name", "Tel", "Addr"} //标题
	w.Write(header)
	for _, v := range UsersDb {
		context := []string{
			v.ID,
			v.Name,
			v.Tel,
			v.Addr,
		}
		// data = append(data, context)
		w.Write(context)
	}
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	w.Flush()
}

var Users []*UserTest

// 解析上传csv
func CsvToUp(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	reader := csv.NewReader(f)
	result1 := make([][]string, 0)
	result2 := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		result1 = append(result1, record)
	}
	for k, _ := range result1 {
		if result1[k][1] == "Name" && result1[k][2] == "Tel" && result1[k][3] == "Addr" {
			result2 = append(result1[:k], result1[k+1:]...)
		}
	}
	for _, v := range result2 {
		user := &UserTest{
			ID:   v[0],
			Name: v[1],
			Addr: v[2],
			Tel:  v[3],
		}
		Users = append(Users, user)
	}
	for _, v := range Users {
		fmt.Println(v)
	}
}
