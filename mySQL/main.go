package main
import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)
type Tag struct {
	ID int 
	Name string 
}
func main() {
fmt.Println("MySql!")

	db, err := sql.Open("mysql","username:password@tcp(192.168.0.6)/test")

	if (err!=nil) {
		panic(err.Error())
	}

	defer db.Close()
//	insert,err := db.Query("INSERT INTO test VALUES(2. 'test')")
results, err := db.Query("SELECT id,name FROM tags")
	if (err!=nil) {
		panic(err.Error())
	}
	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.ID , &tag.Name)
		if (err!=nil) {
			panic(err.Error())
		}
		log.Printf(tag.Name)
	}
}