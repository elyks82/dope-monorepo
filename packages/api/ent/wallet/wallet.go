// Code generated by entc, DO NOT EDIT.

package wallet

const (
	// Label holds the string label denoting the wallet type in the database.
	Label = "wallet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPaper holds the string denoting the paper field in the database.
	FieldPaper = "paper"
	// EdgeDopes holds the string denoting the dopes edge name in mutations.
	EdgeDopes = "dopes"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeHustlers holds the string denoting the hustlers edge name in mutations.
	EdgeHustlers = "hustlers"
	// Table holds the table name of the wallet in the database.
	Table = "wallets"
	// DopesTable is the table that holds the dopes relation/edge.
	DopesTable = "dopes"
	// DopesInverseTable is the table name for the Dope entity.
	// It exists in this package in order to avoid circular dependency with the "dope" package.
	DopesInverseTable = "dopes"
	// DopesColumn is the table column denoting the dopes relation/edge.
	DopesColumn = "wallet_dopes"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "wallet_items"
	// HustlersTable is the table that holds the hustlers relation/edge.
	HustlersTable = "hustlers"
	// HustlersInverseTable is the table name for the Hustler entity.
	// It exists in this package in order to avoid circular dependency with the "hustler" package.
	HustlersInverseTable = "hustlers"
	// HustlersColumn is the table column denoting the hustlers relation/edge.
	HustlersColumn = "wallet_hustlers"
)

// Columns holds all SQL columns for wallet fields.
var Columns = []string{
	FieldID,
	FieldPaper,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
