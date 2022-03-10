package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)
func main(){
	file , err := os.OpenFile("d:/abc.txt",os.O_WRONLY,0666)
	if err != nil {
		fmt.Println("open file err =",err )
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
}