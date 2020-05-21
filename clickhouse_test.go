// +build sqlserver

package testfixtures

import (
	"os"
	"testing"
)

func TestClickhouse(t *testing.T) {
	testLoader(
		t,
		"clickhouse",
		os.Getenv("CLICKHOUSE_CONN_STRING"),
		"testdata/schema/clickhouse.sql",
		DangerousSkipTestDatabaseCheck(),
	)
}
