// Code generated by entc, DO NOT EDIT.

package itemvariation

const (
	// Label holds the string label denoting the itemvariation type in the database.
	Label = "item_variation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVariantName holds the string denoting the variant_name field in the database.
	FieldVariantName = "variant_name"

	// EdgeOriginItem holds the string denoting the origin_item edge name in mutations.
	EdgeOriginItem = "origin_item"

	// Table holds the table name of the itemvariation in the database.
	Table = "item_variations"
	// OriginItemTable is the table the holds the origin_item relation/edge.
	OriginItemTable = "item_variations"
	// OriginItemInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	OriginItemInverseTable = "items"
	// OriginItemColumn is the table column denoting the origin_item relation/edge.
	OriginItemColumn = "item_item_variation"
)

// Columns holds all SQL columns for itemvariation fields.
var Columns = []string{
	FieldID,
	FieldVariantName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the ItemVariation type.
var ForeignKeys = []string{
	"item_item_variation",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
