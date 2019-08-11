package main
import "encoding/json"
import "fmt"
import "os"
type response1 struct {
    Id   int
    Fname []string
}
type response2 struct {
    Id   int      `json:"id"`
    Fname []string `json:"fname"`
}
func main() {

}