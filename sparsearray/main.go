package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 稀疏数组算法案例，稀疏数组可以减少内存的占用
func main(){
	Initial()
	SparseArray()
	ReadSparseArray()
}

// Initial 输出原始值，这个时候存的值是 11*11
func Initial(){
	// 实现 11 * 11 的原始值
	var chessMap [11][11]int
	chessMap[1][2] =1
	chessMap[2][3]=2

	// 输出原始值
	for _, v := range chessMap {
		for _, item := range v {
			fmt.Printf("%d\t",item)
		}
		fmt.Printf("\n")
	}
}

// SparseArray 使用稀疏数组的方式存入文件
func SparseArray(){
	// 实现 11 * 11 的原始值
	var chessMap [11][11]int
	chessMap[1][2] =1
	chessMap[2][3]=2
	chessMap[2][4]=2

	// 输出稀疏数组
	out:=[]RowColVal{}
	// 标准的稀疏数组还有一个表示 数组的规模 和一个默认值在第一个元素
	scale:=RowColVal{
		Row:11,
		Col:11,
		Val:0,
	}
	out= append(out, scale)
	for m, v := range chessMap {
		for n, item := range v {
			// 如果m*n的值大于0，则把该值存在一个结构体切片内
			if item>0{
				rcv:=RowColVal{
					Row:m,
					Col:n,
					Val:item,
				}
				out= append(out, rcv)
			}
		}
	}
	for _, v := range out {
		fmt.Println(v)
	}
	// 写入文件
	content,err:=json.Marshal(out)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	WriteFile(content)
}

type RowColVal struct {
	Row int
	Col int
	Val int
}

func WriteFile (content []byte){
	// 获取当前目录
	dir, err := os.Getwd()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	file := dir + "/sparsearray/data.txt"
	var fh *os.File
	fi, _ := os.Stat(file)
	if fi == nil {
		fh, _ = os.Create(file) // 文件不存在就创建
	} else {
		fh, _ = os.OpenFile(file, os.O_RDWR, 0666) // 文件存在就打开
	}
	defer fh.Close()
	w := []byte(content)
	_, err = fh.Write(w)
	if err!=nil{
		fmt.Println( err.Error())
	}
}

// ReadSparseArray 从文件读数稀疏数组并还原原始值
func ReadSparseArray(){
	fmt.Println("====================================")
	// 获取当前目录
	dir, err := os.Getwd()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	file := dir + "/sparsearray/data.txt"

	f,err:=os.Open(file)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	b,err:=ioutil.ReadAll(f)
	out:=[]RowColVal{}
	err = json.Unmarshal(b,&out)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	for _, v := range out {
		fmt.Println(v)
	}

	//scale:=out[0]
	//fmt.Println(scale)
	//row:=scale.Row
	//col:=scale.Col
	var initial [11][11]int
	//
	for _, v := range out[1:] {
		fmt.Println(v)
		initial[v.Row][v.Col]=v.Val
	}

	for _, v := range initial {
		for _, item := range v {
			fmt.Printf("%d\t",item)
		}
		fmt.Printf("\n")
	}
}

