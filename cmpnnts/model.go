package mdls

/**
 * main model for database
 */
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

type test struct{
	db *sql.DB
	err error
}

func Generate(usr, pss, db string) {
	var(
		id int
		nombre string
	)
	obj := new(test)
	obj.db,obj.err = sql.Open("mysql",usr+":@"+pss+"/"+db)
	rows, err := obj.db.Query("select id, nombre from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &nombre)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, nombre)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}