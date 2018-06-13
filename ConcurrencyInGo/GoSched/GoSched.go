package main
import(
	"runtime"
	"fmt"
	"time"
	"strconv"
)
func showNumber(num int) {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num)
	fmt.Println(num, tstamp)
	time.Sleep(time.Millisecond * 10)
}
func main() {
	runtime.GOMAXPROCS(0)
	iterations := 10
	for i := 0; i<=iterations; i++ {

		go showNumber(i)

	}
	fmt.Println("Goodbye!")
	runtime.Gosched()
}