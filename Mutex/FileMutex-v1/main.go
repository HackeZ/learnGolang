package main

import (
    "os"
    "errors"
    "io"
    "sync"
)

/*
 * rsn : 最后被读取的数据块序号
 * wsn : 最后被写入的数据块序号
 * woffset : 写入偏移量
 * roffset : 读取偏移量
 */

// Data is []byte
type Data []byte

// DataFile 数据文件的接口类型
type DataFile interface {
    // read a data bolck
    Read() (rsn int64, d Data, err error)
    // write a data bolck
    Write(d Data) (wsn int64, err error)
    // Get the value of rsn
    Rsn() int64
    // Get the value of wsn
    Wsn() int64
    // Get the length if data bolck
    DataLen() uint32
}


type myDataFile struct {
    f       *os.File        // File
    fmutex  sync.RWMutex    // RWMutex for File (control the f)
    woffset int64           // write offset
    roffset int64           // read offset
    wmutex  sync.Mutex      // write mutex (control the woffset)
    rmutex  sync.Mutex      // read mutex (control the roffset)
    dataLen uint32          // length of data block (uint32 is easy to trans int or int64)
}

// NewDataFile init myDataFile
func NewDataFile(path string, dataLen uint32) (DataFile, error)  {
    f, err := os.Create(path)
    if err != nil {
        return nil, err
    }
    if dataLen == 0 {
        return nil, errors.New("Invalid data length!")
    }
    // Other variable will be initialized to the default value
    df := &myDataFile{f: f, dataLen: dataLen}
    return df, nil
}

func (df *myDataFile) DataLen() uint32 {
    return df.dataLen
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
    // read and update read offset
    var offset int64
    df.rmutex.Lock()
    offset = df.roffset
    df.roffset += int64(df.dataLen)
    df.rmutex.Unlock()
    
    // Read a Data Bolck
    rsn = offset / int64(df.dataLen)
    bytes := make([]byte, df.dataLen)
    for {
        df.fmutex.Lock()
        _, err = df.f.ReadAt(bytes, offset)
        if err != nil {
            if err == io.EOF {
                df.fmutex.Unlock()
                continue
            }
            df.fmutex.Unlock()
            return
        }
        // Read the Data Successful
        d = bytes
        df.fmutex.Unlock()
        return
    }
}

func (df *myDataFile) Write(d Data) (wsn int64, err error)  {
    // read and update write offset
    var offset int64
    df.wmutex.Lock()
    offset = df.woffset
    df.woffset += int64(df.dataLen)
    df.wmutex.Unlock()
    
    // Write a Data Bolck
    wsn = offset / int64(df.dataLen)
    var bytes []byte
    if len(d) > int(df.dataLen) {
        bytes = d[0:df.dataLen]
    }  else {
        bytes = d
    }
    df.fmutex.Lock()
    defer df.fmutex.Unlock()
    _, err = df.f.Write(bytes)
    return
}

func (df *myDataFile) Rsn() int64 {
    df.rmutex.Lock()
    defer df.rmutex.Unlock()
    return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) Wsn() int64 {
    df.wmutex.Lock()
    defer df.wmutex.Unlock()
    return df.woffset / int64(df.dataLen)
}