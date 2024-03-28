package main

import (
	"fmt"

	"github.com/cch123/elasticsql"
)

var sql = `
SELECT * from product where name like '%楚瑶%' order by product_id desc
`

func main() {
	dsl, esType, _ := elasticsql.Convert(sql)
	fmt.Println(dsl)
	fmt.Println(esType)
}
