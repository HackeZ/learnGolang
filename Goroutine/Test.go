package main

import (
    "fmt"
    "time"
)

// Addr city is city district is district
type Addr struct {
    city     string
    district string
}

// Person is The Best
type Person struct {
    Name    string
    Age     uint8
    Address Addr
}

type PersonHandler interface {
    Batch(origs <-chan Person) <-chan Person
    Handle(origs *Person)
}

type PersonHandlerImpl struct{}

func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
    dests := make(chan Person, 100)
    go func ()  {
        for p := range origs {
            handler.Handle(&p)
            dests <- p
        }
        fmt.Println("All the information has been handled.")
        close(dests)
    }()
    return dests
}

func (handler PersonHandlerImpl) Handle(orig *Person) {
    if orig.Address.district == "Haidian" {
        orig.Address.district = "Shijingshan"
    }
}

// 需要处理的人员总数
var personTotal = 200

// 初始化人员切片
var persons = make([]Person, personTotal)

// 当前处理的人员id
var personCount int

// 初始化所有人
func init() {
    for i := 0; i < personTotal; i++ {
        name := fmt.Sprintf("%s%d", "P", i)
        p := Person{name, 21, Addr{"Beijing", "Haidian"}}
        persons[i] = p
    }
}


func main() {
    handler := getPersonHandler()
    origs := make(chan Person, 100)
    dests := handler.Batch(origs)
    fetchPerson(origs)
    sign := savePerson(dests)
    <-sign
}

func getPersonHandler() PersonHandler {
    return PersonHandlerImpl{}
}

// 获取人员信息，只允许插入origs通道，不允许取出
func fetchPerson(origs chan<- Person) {
    origsCap := cap(origs)
    // 为防止origs是一个非缓存通道的必要措施
    buffered := origsCap > 0
    goTicketTotal := origsCap / 2
    goTicket := initGoTicket(goTicketTotal)
    
    go func() {
        for {
            // 依次返回 id 为 i 的 persons 切片数据
            p, ok := fetchPersonSlice()
            if !ok {
               // 数据读取完毕
               for {
                   // 阻塞在这里，直到辅助通道数据读取完毕或者origs是一个非缓存通道（意味着数据已经结束）
                   if !buffered || len(goTicket) == goTicketTotal {
                       break
                   }
                   time.Sleep(time.Nanosecond)
               }
               close(origs)
               break
            }
            // 将数据插入origs通道
            if buffered {
                // 非缓存通道
                <-goTicket      // 先取出一个
                go func() {
                    origs <- p
                    goTicket <- 1   // 再写入一个
                }()
            } else {
                // 缓存通道
                origs <- p
            }
        }
    }()
}

// 初始化一个辅助缓存通道
func initGoTicket(total int) chan byte {
    var goTicket chan byte
    if total == 0 {
        return goTicket
    }
    goTicket = make(chan byte, total)
    // 填满一个空 goTicket 通道
    for i := 0; i < total; i++ {
        goTicket <- 1
    }
    return goTicket
}


func fetchPersonSlice() (Person, bool) {
    if personCount < personTotal {
        p := persons[personCount]
        personCount++
        return p, true
    }
    return Person{}, false
}


func savePerson(dest <-chan Person) <-chan byte {
    // sign用于阻塞，直至数据保存完毕
    sign := make(chan byte, 1)
    go func() {
       for {
           p, ok := <-dest
           if !ok {
               // 读取完毕，往sign中插入一个byte，返回，结束程序
               fmt.Println("All the information has been saved.")
               sign <- 0
               break
           }
           savePersonSlice(p)
       } 
    }()
    return sign
}


func savePersonSlice(p Person) bool {
    time.Sleep(1*time.Nanosecond)
    fmt.Println("Person district -->" ,p.Address.district)
    return true
}