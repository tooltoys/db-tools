package tools

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type MulticolumnIndex struct {
	conn *sqlx.DB
}

func NewMulticolumnIndex(conn *sqlx.DB) *MulticolumnIndex {
	return &MulticolumnIndex{
		conn: conn,
	}
}

func (t *MulticolumnIndex) AnalysisOrder(table string, columns ...string) []byte {
	buildCountQueries := func(columns ...string) string {
		var countQueries = make([]string, 0, len(columns)+1)
		for _, col := range columns {
			countQueries = append(countQueries, fmt.Sprintf("COUNT(DISTINCT %s) / COUNT(*) AS %s_selectivity", col, col))
		}
		countQueries = append(countQueries, "COUNT(*) AS countall")

		return strings.Join(countQueries, ", ")
	}

	countQueries := buildCountQueries(columns...)
	query := fmt.Sprintf("SELECT %s FROM %s", countQueries, table)

	rows, _ := t.conn.Query(query) // Note: Ignoring errors for brevity
	cols, _ := rows.Columns()

	fmt.Println(cols)
	for rows.Next() {
		columns := make([]float64, len(cols))
		columnPointers := make([]any, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			log.Fatal(err)
		}

		fmt.Println(columns)
	}

	return nil
}

type tmp struct {
	StaffIdSelectivity    float64
	CustomerIdSelectivity float64
	COUNTALL              int
}
