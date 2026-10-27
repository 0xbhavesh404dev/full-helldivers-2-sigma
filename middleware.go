package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.9" )

func inVpSlLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSmLGRaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTTHZfEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TOV8KHKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tr5n1Lq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvktufzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmvGGT5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjkUhFXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGpZfa5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOSnWfKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWNwrDB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gseqDjmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SrkxFZ4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ew60JLQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jz0Y1N5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnNszr9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1SrHGOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oor6mbCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvvGlj4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIMen1ceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNWMNZkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKVsRIxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWcYaVIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMhn3GiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cBojX4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wj6dLP2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoZhXyqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OczJBTklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4fOG3ZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EguoEyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sW87NoO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWXmzulUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHUEY6wjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgyM3Nd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AR6wsMQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gO277XpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0c153clSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RS2V2iySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxGMNiKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPPDEv0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnLUpEJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsEThPRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyKe5oPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIw8GAnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UAvdmQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzUFat9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CDGe63mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzHLaCCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v78LjCikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZGwRNgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIgSboXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fr6Kwl7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8Lun9BsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bb3YgqfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKgkWkSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQfbV5wGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bO7ba0gZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PQe95JjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGXlYtfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21mCmnMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YErFAydJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnzPXVODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjdCq9cjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0iWmtD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KIz786ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHgI8iqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuUDDsvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OXFT7jtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFhXQxdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AfJe5krWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EejXYQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUZpHt7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4ECIkRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfZSNpPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rwMQRTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkBiW3n3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuRNFlnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUHByh4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E56rchE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsnffPJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xMwdAMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bIJDAF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMakAmmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynKx3XBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWa1YUEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XdMsaKw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LL3Hy8epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8Ng5w7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8XEUll7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOuBjO36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsSYuwKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDneJflgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5V0ZVrxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func po4CfdxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFTal2lvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50O9WBokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOtwI4eEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jr6pwwV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mb5QrMS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZG9GxHjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MXuhS1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qSjI8wfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wZsQyzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtnyilu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aihuw0NNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctnWTgTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esslEINOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4MBGtkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vN2bmZssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8S2WsRdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBj3uUdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ar2g9RyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HiIBKUgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zdlvi9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func doTCHnEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zO4vDrF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34QPa41vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcwtInbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgs7YMXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6GPF6YnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JdB3NJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBBCawUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfET4RolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYF7dsLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4d7xAiwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OO6afiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRiI2JB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkQzohaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDNCTM7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rugOCfVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIoOM1TpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LGhDqW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAeVrr8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdUyrklzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i50O6NxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6pGvfOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4t96DU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eyJONBjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjkpKRW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVhmn6ebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GX1XscL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1vURhq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTrm7Y6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCqsmUZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5s2JZVC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJyzwBeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hexq3mWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYWVbuNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ir4GnozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cux6j3KcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uBRZ1dRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqplxHN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0M7Za6QlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyr6bs0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cASG5UfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eiZ4hMtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvJZ9lURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJ0YAEGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htAHh0ZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnfjy5t8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8jKW9AVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fJ9pVguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52W2DPG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSKzTDMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpGekAyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YpzWEyghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwR0ycpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcijrG3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6f7AW66OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HuLLHn9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nO6EsVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYIL6P7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iuAg9yMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HImiIlsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUua5lwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFws0ARIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0JSEQGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJxcLBZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08LrmBdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwJIWC8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaVELhJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDk1xXTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOjksffZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qa21x22iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WeNVryL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YH9xD801Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uArZ9rkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WC2S0lPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JZumZCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBtHUfKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDueQzedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ji4bKBGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEzm0ZRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXACEfMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6RBwpO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9aQpi3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9Gqnbc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeC2l0D3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfaRIvVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCdCFZFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYpYnkF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0ZM3Z97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5pJrnnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCUN8rYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKlwwoAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqj5Sln4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gCaCIIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EPeopbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2EEhv1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqjDnyHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfME31bXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryHuqpclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiqjJzwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Olsn4cj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZtP7dmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFYddZcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8CfJmo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0EFc5L6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJmGRRzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOyjgV1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkPmbrFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUXcHZLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODJ3B7pbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKwrrFA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImMotXUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flqFmSW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWBsKb4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tU4qG4cGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHeNmQR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKyZt6CLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7gOSI8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41ynutg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEDTEnIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xcDfX9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a31EwyfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vD2Za6AOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvWNyod8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnBPYVQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwzsgpy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0IS0YR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 346m1xHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0y15jiWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9n03Op0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqWlXGyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjMHTQCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSaWwVMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbQZWvDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjcdSCxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyxhofQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVeLpZTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J136R7ZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2LzDdgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92xwU0CGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOkda1wDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gs2wZNAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddMeAoguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sz7OSiwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1Hzw5LJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asaYaPicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucX5yVdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xlh2jlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFMuadOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8e5oCs5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CBltOlrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0gkbyPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTFEn5eIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPwB92MQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZ5iKlPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0aPhEsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6tNCM3TzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8cwAzvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VP1lFZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIv8tPVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfDTJg42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Njx0SGw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdRCPTCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVhG5jMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8llxcPRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSckZqInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLiDy2ZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caWbkfvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6rqUJu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 888NL5evWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I57TqTPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9ehEHxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqpwZsObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UOZMjbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8E7MHmkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPk2VF01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxO6OddMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3NUD4C8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hekVuJ7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuGgcmLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PS8YloMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woXhdKBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2V4lqEbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKNAsGcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSNNtXbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1A1MZrz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4MLwu0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVUdGGBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSzPFfTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPbH3BhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RN2YTuyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEIMgKU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvakSSHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6eSsF9CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sbj45ivdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22e0ciGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aucSH0oXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvOgFRPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98j7zhdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzAZukzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlXwnpubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQOZP2SmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ibiveiWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Am3LSJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZE7vH1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqTHR0TkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sS8rSHZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkmPF32JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7MZ2udbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9yJ66CDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbjOa3qvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9YqWEbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEQxrwT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hc7bszFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FhuA05AnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gez4zhTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOupDqpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXFPbriAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqtS1BpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Zq8CyLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDxFBntJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kU1MZRCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4L4pkaaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YL4H390qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyL9ecL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46A7x293Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQp0U63rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHIVdx5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqc80kqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkulrRFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5YghGjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1k9SIc8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7BMGYSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axXpT5G2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcwqA2KkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n32t9JDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vL28Ie6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQNPS3UGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Ct29fkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpY6cxiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dka0FXhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8yLvxQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WkKRSGSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8v1BbjxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWzCG1ZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9yoRvbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwJK0TiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxkp9A7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQFSvQBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0EyP0N7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEHCtox2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rC1dc3KmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qA1VtJDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHYbxWtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kldnqo7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOTD73wwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjzJLqAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxHh5tTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ja3iU3Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6mmlCN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBR6SCwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELhhWZgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxkxB3J5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bu4YmzleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAzeJEBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYoHfBiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cug5B3CoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSgPD2mxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vv4nLmtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwvNlzpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nPe5zkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpjBfsGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NriWyIwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWgcGO6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMin3PYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ObOna8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kF0DWZBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VovGEFpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nH1xXrYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F0yTg517Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrQ78UgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dpTZn81cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ks5OeisWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YRiAlDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdazzFdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYycvPp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdA9ELlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pm75rFcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyKijbNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ret1iwLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGvNH0UhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlBrvTAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PllLiYvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1PMBLV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjHOBmubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2T5nexIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFQwmcvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zmk7iHosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URMXMCD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fH5W3GshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGZxQnrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pm57v5vkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoKhZ8ttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66sA3n9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAq2HBJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAyyZKFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkFcUjyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UftKXRrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpNTSfrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3xkTWOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsTAAuxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdZU2VqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJVqYuYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81lmrfeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlgzORYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lTTL9EdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDALyzPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQqvNLvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5YJpJ9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXrZzKubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q6eZay9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eE1bcjaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GGx42pTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEGBag3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQuPGI0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8XSh14hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFJk5WjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8OypEg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1ZnYn0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ieav04wAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLKIb7aOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yx7iynGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zj1gqAxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RELGV0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nHvTHSb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRxq8OQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ArXPgXp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CTYbkP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNDSYWqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zTgqV4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aToLk1GKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0X5L7I2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymlpomToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1sNGtsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iR4U2CLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuJPSRZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUNSw5RXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKbCRSNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSd8vH9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fl7IVoV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OiYkasdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLtWlmiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNkIYzcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OlGisMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BR6qgdsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZY5KxwPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgoeHm5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeKC6Kp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7truU9wmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Bh8NLjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYBd0J6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ht36cFr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKxj2l2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWIJwAkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAn6ib4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOchyVDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2RRZBAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67v3vpuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fj6hSGH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLTLjWIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvYCuHaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYjdHWyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVqYuFnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8RlyUS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKlkLZKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfQQchfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHNFB0ruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INnmjyy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcjgKDsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6Fl0oBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfVDHbpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmMTRFLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M71DfMrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfawACybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1To9hCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OH7CVyEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuKwOXmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdcdGNxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksGwwSY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8w9xtzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jL7s5OMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJ08x7TzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFg61JRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LX5Ims4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jh9ojT8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chqB3Eh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGNMPD3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlJLaZiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9FT9Pps8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1W01n28yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wk8DrU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAIs8VkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VrbZdkwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRZF5MugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99dsaFUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8osKVWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oH713gBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCtvyjfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEwEgTTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uATMB2omWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFixQlMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyCJV4VWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVt2kqr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfRnmckCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUyr2sTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSFVgTlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onoXEr91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36LSBjViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5eclbFz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18HICrdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OB122Y0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1sJQa6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLcPfmBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1yQmevuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UYaOmGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6tkI9NfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjWXy75FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRs64R9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djr2Y48PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOGzT6xfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2xPXUfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IVSImjWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNkZuUBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0P1UxUI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYRvverKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7VlbTQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gzuSmNlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Hkif73SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxKMMgT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZqqRS3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwbAJNQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6NUNLfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vrS5Rs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yiJtCkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTCeoFBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZQVcdG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSm86jxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZqpl5vvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4Seov5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RH0ikdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZVxYLSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVAkRbnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6UlUQyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h19U3RzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2T6I27wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v4obR8aRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwaQSXLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FcqOZcq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsSZDHfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gG7yJ7IaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uquYtSx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqkaDIqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMI99etWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpyxIl2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IndVOGScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v74V435NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESupztdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qm8Sr7XWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZN4Qx7TLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jb6e90SxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGvv881kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAThWHjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGnyuGh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKVkOr4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvaXuRneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4adMLDKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hZrz1VXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JonqNnxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYAcTuqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twRXFIUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bo6uYly4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oT5TypVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jwj5diJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIeOS72YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiHBzNXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8A16MuyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmLFBXMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0WSpT73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21g27NkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBMb4tHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axBFWroRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aF4zR0GIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEbo6aczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpSIrt7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XOdnjddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZKObzLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AnAM4h4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vX6z30LKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xf5hQRN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymJ3E2URWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJTknnCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHf0bDR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cY1l8EKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YZo0S1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2LCkBbZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhMZuy7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOl8gQiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccZ2TjjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HZ9qaFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RE3OuURaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vw1ZGYmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVqwRcdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSOA1D1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RQ3QSx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDK1Gg9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfxJEA0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acy5kR1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5W2uEfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnPioOZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lYaDLG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JJRtkpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcln2MZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func emb3qWXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcGIFElMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QyKRjhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chfVkH2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmMYyXLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vk9iMpj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func po5AQYJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xnIR2gF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCc0dl0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TNsLcoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wD6I3QBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eo7uTiChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQbQFRiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgqeUuUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2k0AUokkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDWLN0xZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSW0ueJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0Io6DRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GD6frYACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3hjwdAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OynSzRlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dp1JdxfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ry2NbVYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3A73P6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zW6R3FWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaiQ53zJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8VGT0OhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PafjKS2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LH1Ul35PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xj1QVTBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l50TudYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImQ0iEZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xnYOTFFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yf3dNTbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykTpeGKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qy4RswrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVSbvV5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTIarinYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wMTgs6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rdq5KkPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OOIF1RLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DonR7Vs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vM9WERuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1GeikOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxZLcr2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8ByN6PZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXu9vVOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmVhzMgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZj5YQ0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knSddAatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbwseWQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94NIhuF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUXoCmEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cP9PtYVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5v6cL5PZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvyXsxlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsAoNxC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AufUkdhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQUlVUmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Phnjc3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQ284r8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVbh9ZWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6u23I10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOq9VhMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5xElFhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpkG7AcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWEgW1MwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzMsyxGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U2fz5Z8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lCp3zqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzi8jmt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1MMdXVlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ilsh7LnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3q9yZdwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sr6yKmQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOCHAJEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1gVQvKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeafxZlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQJ0RkAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 606bZCExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOnc3SNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gA53EZ6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5y65hCt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RU2MDQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Fy8ayLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func doY4c1uoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqjHya54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgFezJ4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7bGAfPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCRtnV9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kc368VIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7h8zK2PHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GasmTkHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KEBrkq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEtKX0lbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xle1GqNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Si6WNymkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xc6jG5sqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mU2NzDkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnizYnwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrNJEP7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ragDhEjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6tNsqPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K31htkqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PjQ1rifDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpN793wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZF8b0QI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYBJkpOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zj4jejuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THitmZx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BImcXhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guxP3u2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCpjWgPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DG20QOIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vP3Qgl0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qgno5PRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rhNXARLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qEy9MKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOTdWxyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ThvlfJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtqsDe7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjWRsgDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBcFOPbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82h9GiJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owaxY0bBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWSDZwKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjYcOP6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWW8BeTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BE82nPKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6NnLDQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUubNnBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGDmdR6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7gY52yDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFyHhbzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIXjz5cwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfKaFrxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVLSIaLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVhewlRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBTHedUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cA4HTEpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NDuq0pvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kjmb9fsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbeF6c3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpFohJ0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1889U3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvzDmuyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGVvjNq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x541jhhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I13Z28kSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIx5yYgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J17hox0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIY6R05JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xeSN6H0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yf9j2u9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qg2AFAO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UTIMPuEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qnln62dKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rr6CBHz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZjTUouGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlxfqPhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMiGIc7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nl2oBvWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9Ry35jXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zivi94M8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0SnOBAgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mo6ThOqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpA57nIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24tLogFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0aQxta6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJVhdcMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0JbG5nOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzzWI4SsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func quknanE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIPt5e4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ga3zUEhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzrNOjNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9koGzOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkEjUa3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSHIrOzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2gNS9n3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFSkUj1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8LzRKlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjPkPDNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LukuOGekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgiOdAVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0NLfT5rJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76GCL5ftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmrJZbJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XV6y2mQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P41oAURpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHEFkCKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syzlyQEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12lHQXauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbvkgZYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAOkMS8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sF1bGphaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmJLZOJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJKXvGSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5bDWAg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPckhee0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HIF2z2AGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyFmtTDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V70rbgwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mdKOx7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2L31F6zoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hn8ZODB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8m8FrPyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpRocNS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lfTveR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zX3bbqt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOBmkeovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehcschqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqNsqZy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOEmaD9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxyyzV4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNbJAEEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y08trzsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2nlTO0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SG17dZeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeHCy9HmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fKAVoHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UugTztgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4uDNjxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSk5JXA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JFFjSNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqSS88BrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18EHQkKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoBBWYHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2MrnmkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pX9qqursWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nf1c4EMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTZ1ybv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojAvMqhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJtJ1AQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9dBuGGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txSzf773Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WO6ZrlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EojCOkPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3RoktZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0d9gROTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oN1V6XwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jTMVCObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6fRCPceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tq1wpV7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szTLKrYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhl3otwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4Sz6tlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynoA3QJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOzib8HNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFCAGhIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U6d1it4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuS4OUqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8XNAxdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glcTrxU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zog26gsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDN8MuIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OgYK1KJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpUAxSrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7nwYn42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vh6bTSsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTYiENCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtDLO4QbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0x1VqBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQ9sL8zXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsdGed09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulDsnnDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOvlf6VrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiIVswfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9dnmqlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FowKh4qHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MV8vhGsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4umTiF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTQc6bILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOPHkXrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tzQgrXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XIg4pe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahEu47DFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvybfZY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPvVsSbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgQM66DbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxuNGH9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZSJTbttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nsgz84EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ecy2KE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fM3rfSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTRdAg0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQu0HukoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epNIwGpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvwmbsskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func puW0KiEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mH0j7uoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyU7p7HRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIBaPzG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbWsWSwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mjSglESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56qszAuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EB90xUf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvbmRr9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xU8hC9TIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QGAlBYfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhEnb1uIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Q9ovrnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2c2kvDI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yyBWO19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V7nIbogZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yrlrk68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lp6x2zKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJY8uihLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gWJLdxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxmd5w6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rELJ6J3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ejj8UxgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zYiDQAhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52gDp777Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZZDDABMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5v3xRMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAof6c8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYXoFlfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbgznEk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXd9rGaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3DFxBEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0OGxqglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5k5xEtgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwsbADRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzOZgM1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOFTXIviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXgZy9PnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSPxcys9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEoe9ormWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGeD5clNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fm4aKGxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KLVKOZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2YlQMiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVA29bkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boqH7Bl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfE17CtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxaYhp82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZHjygDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMEMuqULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlgyhSxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbY5tYXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 02kkQQ6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jm3nrw9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrLp0BnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJr0DxshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vs7EeurdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsVd5NmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFuJNHMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8SQUzWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxgKxcSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A99fzCX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuu4eVBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuHt4aBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5YAEFxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbvpRIYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9R7PzuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1q42fmWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TPITIP4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQKuqxjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z88AVoQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXxlbzFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ek5QtfP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NpC8oZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEuOdEOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttLBln4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOYDviAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlgWDiq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEBO2DakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ui4zfKLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZqDzQvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqRavtYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXi7LUOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLuh7TuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yrAU0Z3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fY4qSq6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNFoO3fXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4v5tlz1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74AQr9l8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAHW2KYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xtvwtls6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buMtzVEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peJWG8B6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIbBdzAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4r5jW4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JtjEPJPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWEhb6LBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5qolnE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6jLMByzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iRSM9i9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdM1XINaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxctVv1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lh3kkUHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZJx0aMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaK33akfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUqqlfpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bItF4PbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odMh2KPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44Jl1CDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxfZW6z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTbHvUHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LXCerWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQQKdifdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7U6fkDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qOp0sE6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkZSRlCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zh1JZ5QPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5eJcBNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1deuOAPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kifiiZIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0PK7qFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHxqY16vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbXWN1M4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjtmxODpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6M2hLdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iboG4pymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func to5oLmOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulttBwgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1pJxM6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8A6IUVjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fm4BsujFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0McxvPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7z5cZMKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdlQiM1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poxWwcNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KARSH3tIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DdTUxT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WgO1MJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCDSsjlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cXQLVcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bW3pBs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c16z2ifYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cHM0duFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CupckNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 880Z0tn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVjN6PnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aM7lBemIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywbjSKX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ve6EXBoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUcGP9rPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNa8ipHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPiGCucUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5qyg36yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFMamcDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITsBU40UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1l865sifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAPkbotNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xaAXe2N0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvXuyT56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUyY5XU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySnj5lV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YR6R8JsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHPK4BYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mQOkVUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlUNFGz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MldOQNMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKkLioH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8wW4mqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGGTJ2ooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3imNqXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcJnDW6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgiXDhcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8jBIpfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fxn2llOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FM0uCm5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFcKwyYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eI75LiFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XI2BOGVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlU32CnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7t2SNvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWTmLknFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qY2tJDmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

