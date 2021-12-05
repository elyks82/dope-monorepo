// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DopesColumns holds the columns for the "dopes" table.
	DopesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "claimed", Type: field.TypeBool, Default: false},
		{Name: "opened", Type: field.TypeBool, Default: false},
		{Name: "wallet_dopes", Type: field.TypeString, Nullable: true},
	}
	// DopesTable holds the schema information for the "dopes" table.
	DopesTable = &schema.Table{
		Name:       "dopes",
		Columns:    DopesColumns,
		PrimaryKey: []*schema.Column{DopesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dopes_wallets_dopes",
				Columns:    []*schema.Column{DopesColumns[3]},
				RefColumns: []*schema.Column{WalletsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"weapon", "clothes", "vehicle", "waist", "foot", "hand", "drugs", "neck", "ring", "accessory"}},
		{Name: "name_prefix", Type: field.TypeString, Nullable: true},
		{Name: "name_suffix", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "suffix", Type: field.TypeString, Nullable: true},
		{Name: "augmented", Type: field.TypeBool, Nullable: true},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
	}
	// WalletsColumns holds the columns for the "wallets" table.
	WalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "paper", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "numeric"}},
	}
	// WalletsTable holds the schema information for the "wallets" table.
	WalletsTable = &schema.Table{
		Name:       "wallets",
		Columns:    WalletsColumns,
		PrimaryKey: []*schema.Column{WalletsColumns[0]},
	}
	// DopeItemsColumns holds the columns for the "dope_items" table.
	DopeItemsColumns = []*schema.Column{
		{Name: "dope_id", Type: field.TypeString},
		{Name: "item_id", Type: field.TypeString},
	}
	// DopeItemsTable holds the schema information for the "dope_items" table.
	DopeItemsTable = &schema.Table{
		Name:       "dope_items",
		Columns:    DopeItemsColumns,
		PrimaryKey: []*schema.Column{DopeItemsColumns[0], DopeItemsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dope_items_dope_id",
				Columns:    []*schema.Column{DopeItemsColumns[0]},
				RefColumns: []*schema.Column{DopesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "dope_items_item_id",
				Columns:    []*schema.Column{DopeItemsColumns[1]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DopesTable,
		ItemsTable,
		WalletsTable,
		DopeItemsTable,
	}
)

func init() {
	DopesTable.ForeignKeys[0].RefTable = WalletsTable
	DopeItemsTable.ForeignKeys[0].RefTable = DopesTable
	DopeItemsTable.ForeignKeys[1].RefTable = ItemsTable
}
