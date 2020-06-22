與文件相關的動作分成
1. 開檔
2. 寫檔
3. 讀檔


1. 開檔: 需要手動關檔
    file, err := os.Open(filePath)
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
    file.Close()


2. 寫檔:
分成 全部寫入 / 寫入緩衝後才寫入檔案

全部寫入:
    err = ioutil.WriteFile(desFile, data, 0666)
寫入緩衝後才寫入檔案:
    writer := bufio.NewWriter(file)
    writer.WriteString(str + "\r\n")
    err = writer.Flush() 因為是 bufio 所以實際上 WriteString 是寫入到緩衝區 沒有真正寫到檔案內


3. 讀檔:
分成 全部讀入 / 寫入緩衝後讀入
全部讀入:
    file, err := ioutil.ReadFile(filePath)
寫入緩衝:
    reader := bufio.NewReader(file)


