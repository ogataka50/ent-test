// Code generated by entc, DO NOT EDIT.

package item

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeItemDescription holds the string denoting the item_description edge name in mutations.
	EdgeItemDescription = "item_description"
	// EdgeItemVariation holds the string denoting the item_variation edge name in mutations.
	EdgeItemVariation = "item_variation"
	// EdgeItemGroup holds the string denoting the item_group edge name in mutations.
	EdgeItemGroup = "item_group"

	// Table holds the table name of the item in the database.
	Table = "items"
	// ItemDescriptionTable is the table the holds the item_description relation/edge.
	ItemDescriptionTable = "item_descriptions"
	// ItemDescriptionInverseTable is the table name for the ItemDescription entity.
	// It exists in this package in order to avoid circular dependency with the "itemdescription" package.
	ItemDescriptionInverseTable = "item_descriptions"
	// ItemDescriptionColumn is the table column denoting the item_description relation/edge.
	ItemDescriptionColumn = "item_item_description"
	// ItemVariationTable is the table the holds the item_variation relation/edge.
	ItemVariationTable = "item_variations"
	// ItemVariationInverseTable is the table name for the ItemVariation entity.
	// It exists in this package in order to avoid circular dependency with the "itemvariation" package.
	ItemVariationInverseTable = "item_variations"
	// ItemVariationColumn is the table column denoting the item_variation relation/edge.
	ItemVariationColumn = "item_item_variation"
	// ItemGroupTable is the table the holds the item_group relation/edge. The primary key declared below.
	ItemGroupTable = "item_item_group"
	// ItemGroupInverseTable is the table name for the ItemGroup entity.
	// It exists in this package in order to avoid circular dependency with the "itemgroup" package.
	ItemGroupInverseTable = "item_groups"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ItemGroupPrimaryKey and ItemGroupColumn2 are the table columns denoting the
	// primary key for the item_group relation (M2M).
	ItemGroupPrimaryKey = []string{"item_id", "item_group_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}