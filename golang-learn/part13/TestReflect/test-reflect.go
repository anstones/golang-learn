package TestReflect

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 解析如下参数

/*
	#this is comment
	;this a comment
	;[]表示一个section
	[server]
	ip = 10.238.2.2
	port = 8080

	[mysql]
	username = root
	passwd = admin
	database = test
	host = 192.168.10.10
	port = 8000
	timeout = 1.2
 */

// 读取文件，做反序列化
func UnMarshalFile(path string, result interface{}) (err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil{
		panic(err)
	}
	// 反序列化
	return UnMarshal(file, result)
}

// 反序列化  []byte --->  struct
func UnMarshal(data []byte, result interface{}) (err error) {
	// 先判断是否是指针
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr{
		return
	}
	// 判断是否是结构体
	if typeInfo.Kind() != reflect.Struct{
		return
	}
	// 转类型，按行切割
	lineArr := strings.Split(string(data), "\n")
	// 定义全局标签名   也就是server和mysql
	var myFiledName string

	for _, line := range lineArr{
		// 各种严谨判断
		line = strings.TrimSpace(line)
		fmt.Println(line)
		// 处理文档中有注释的情况
		if len(line) == 0|| line[0] == ';'||line[0] == '#'{
			continue
		}

		if line[0] =='['{
			// 按照达标钱去处理
			myFiledName, err = myLabel(line, typeInfo.Elem())
			if err != nil{
				return
			}
			continue
		}

		// 正常处理
		err = myFiled(myFiledName, line, result)
		if err != nil{
			return
		}
	}
	return
}

// 处理大标签
func myLabel(line string, typeInfo reflect.Type) (filedName string, err error) {
	// 去字符串头和尾
	labelName := line[1:len(line)-1]
	// 循环去结构体找tag，对应上才能解析
	for i := 0;i< typeInfo.NumField(); i++{
		filed := typeInfo.Field(i)
		tagValue := filed.Tag.Get("ini")

		// 判断tag
		if labelName == tagValue{
			filedName = filed.Name
			break
		}
	}
	return
}

// 解析属性
// 参数：大标签名，行数据，对象
func myFiled(filedName string, line string, result interface{})(err error)  {
	fmt.Println(line)
	key := strings.TrimSpace(line[0:strings.Index(line, "=")])
	value := strings.TrimSpace(line[strings.Index(line, "=")+1:])

	// 解析结构体
	reultValue := reflect.ValueOf(result)

	// 拿到字段值，这里直接设置不知道类型
	labelValue := reultValue.Elem().FieldByName(filedName)
	fmt.Println(labelValue)

	labelType := labelValue.Type()

	var keyName string

	for i := 0; i< labelType.NumField(); i++{
		// 获取结构体字段
		filed := labelType.Field(i)
		tagValue := filed.Tag.Get("ini")
		if tagValue == key{
			keyName = filed.Name
			break
		}
	}

	// 给字段赋值
	// 取字段值
	filedValue := labelValue.FieldByName(keyName)

	switch filedValue.Type().Kind() {
	case reflect.String:
		filedValue.SetString(value)
	case reflect.Int:
		i, err2 := strconv.ParseInt(value, 10, 64)
		if err2 != nil{
			fmt.Println(err2)
			return
		}
		filedValue.SetInt(i)
	case reflect.Uint:
		i, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			fmt.Println(err)
			return err
		}
		filedValue.SetUint(i)
	case reflect.Float32:
		f , _ := strconv.ParseFloat(value,64)
		filedValue.SetFloat(f)
	}
	return
}


// 序列化
// 传入的结构体 ——> []byte
// 基本思路：反射解析出传入的数据，转字节切片
func Marshal(data interface{})(result []byte, err error)  {
	// 获取类型
	typeInfo := reflect.TypeOf(data)
	fmt.Println("======================",typeInfo)
	valueInfo := reflect.ValueOf(data)
	if typeInfo.Kind() != reflect.Struct{
		return
	}
	var conf []string
	for i:=0;i<typeInfo.NumField();i++{
		// 取字段
		labelFiled := typeInfo.Field(i)
		fmt.Println("==============", labelFiled)
		// 取值
		labelValue := valueInfo.Field(i)
		filedType := labelFiled.Type

		// 判断字段类型
		if filedType.Kind() != reflect.Struct{
			continue
		}
		// 获取个tag
		tagValue := labelFiled.Tag.Get("ini")
		if len(tagValue) == 0{
			tagValue = labelFiled.Name
		}
		label := fmt.Sprintf("\n[%s]\n", tagValue)
		conf = append(conf, label)

		// 拼k=v
		for j := 0; j < filedType.NumField();j++{
			// 取大写
			keyFiled := filedType.Field(j)
			//取tag
			filedTagvalue := keyFiled.Tag.Get("ini")
			if len(filedTagvalue) == 0{
				filedTagvalue = keyFiled.Name
			}
			// 取值
			valFiled := labelValue.Field(j)
			// Interface()取真正对应的值
			item := fmt.Sprintf("%s=%v\n", filedTagvalue, valFiled.Interface())
			conf = append(conf, item)
		}
	}
	// 遍历切片类型
	for _, val := range conf{
		byteValue := []byte(val)
		result = append(result, byteValue...)
	}
	return
}

func MarshalFile(fileNme string, data interface{}) (err error)  {
	// 数据的系列化
	result , err := Marshal(data)
	if err != nil{
		return
	}
	// 将序列化好的数据写到fileName
	return ioutil.WriteFile(fileNme, result, 0666)
}

