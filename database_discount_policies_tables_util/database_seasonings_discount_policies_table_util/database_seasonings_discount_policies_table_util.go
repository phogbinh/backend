package database_seasonings_discount_policies_table_util

import (
	"database/sql"

	discountPoliciesTablesConst "backend/database_discount_policies_tables_util/database_discount_policies_tables_const"
	"backend/database_util"
	"backend/util"
)

const (
	tableName           = discountPoliciesTablesConst.SeasoningsDiscountPoliciesTableName
	codeColumnName      = discountPoliciesTablesConst.SeasoningsDiscountPoliciesCodeColumnName
	rateColumnName      = discountPoliciesTablesConst.SeasoningsDiscountPoliciesRateColumnName
	beginDateColumnName = discountPoliciesTablesConst.SeasoningsDiscountPoliciesBeginDateColumnName
	endDateColumnName   = discountPoliciesTablesConst.SeasoningsDiscountPoliciesEndDateColumnName
)

const (
	queryCreateTable = "CREATE TABLE IF NOT EXISTS " + tableName + util.EndOfLine +
		"(" + util.EndOfLine +
		codeColumnName + "		CHAR(9)			NOT NULL," + util.EndOfLine +
		rateColumnName + "		DECIMAL(3, 2)	NOT NULL," + util.EndOfLine +
		beginDateColumnName + "	DATE            NOT NULL," + util.EndOfLine +
		endDateColumnName + "	DATE            NOT NULL," + util.EndOfLine +
		"PRIMARY KEY(" + codeColumnName + ")," + util.EndOfLine +
		"FOREIGN KEY(" + codeColumnName + ") REFERENCES " + discountPoliciesTablesConst.DiscountPoliciesTableName + "(" + discountPoliciesTablesConst.DiscountPoliciesCodeColumnName + ")" + util.EndOfLine +
		"	ON DELETE CASCADE" + util.EndOfLine +
		")"
)

// CreateTableIfNotExists creates table `seasonings_discount_policies`.
func CreateTableIfNotExists(databasePtr *sql.DB) error {
	return database_util.CreateTableIfNotExists(databasePtr, queryCreateTable)
}
