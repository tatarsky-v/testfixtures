package testfixtures

type clickhouse struct {
	baseHelper

	useAlterConstraint bool

	tables                   []string
	sequences                []string
	nonDeferrableConstraints []pgConstraint
	tablesChecksum           map[string]string
}
