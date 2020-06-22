package main

import "os"

// go 判斷文文件或文件夾是否存在的方法為使用 os.Stat() 函數返回的錯誤值進行判斷
// 1. 如果返回瘩錯誤是 nil 說明文件或文件夾存在
// 2. 如果返回的錯誤類型使用 os.IsNotExist() 判斷為 true 說明文件或資料夾不存在
// 3. 如果返回的錯誤為期他類型 則不確定是否存在

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {

}
