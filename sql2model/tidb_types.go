package sql2model

/********** copy from github.com/pingcap/tidb/types end ******************/
// Flag information.
const (
	NotNullFlag        uint = 1 << 0  /* Field can't be NULL */
	PriKeyFlag         uint = 1 << 1  /* Field is part of a primary key */
	UniqueKeyFlag      uint = 1 << 2  /* Field is part of a unique key */
	MultipleKeyFlag    uint = 1 << 3  /* Field is part of a key */
	BlobFlag           uint = 1 << 4  /* Field is a blob */
	UnsignedFlag       uint = 1 << 5  /* Field is unsigned */
	ZerofillFlag       uint = 1 << 6  /* Field is zerofill */
	BinaryFlag         uint = 1 << 7  /* Field is binary   */
	EnumFlag           uint = 1 << 8  /* Field is an enum */
	AutoIncrementFlag  uint = 1 << 9  /* Field is an auto increment field */
	TimestampFlag      uint = 1 << 10 /* Field is a timestamp */
	SetFlag            uint = 1 << 11 /* Field is a set */
	NoDefaultValueFlag uint = 1 << 12 /* Field doesn't have a default value */
	OnUpdateNowFlag    uint = 1 << 13 /* Field is set to NOW on UPDATE */
	PartKeyFlag        uint = 1 << 14 /* Intern: Part of some keys */
	NumFlag            uint = 1 << 15 /* Field is a num (for clients) */

	GroupFlag             uint = 1 << 15 /* Internal: Group field */
	UniqueFlag            uint = 1 << 16 /* Internal: Used by sql_yacc */
	BinCmpFlag            uint = 1 << 17 /* Internal: Used by sql_yacc */
	ParseToJSONFlag       uint = 1 << 18 /* Internal: Used when we want to parse string to JSON in CAST */
	IsBooleanFlag         uint = 1 << 19 /* Internal: Used for telling boolean literal from integer */
	PreventNullInsertFlag uint = 1 << 20 /* Prevent this Field from inserting NULL values */
	EnumSetAsIntFlag      uint = 1 << 21 /* Internal: Used for inferring enum eval type. */
	DropColumnIndexFlag   uint = 1 << 22 /* Internal: Used for indicate the column is being dropped with index */
)

// MySQL type information.
const (
	TypeUnspecified byte = 0
	TypeTiny        byte = 1 // TINYINT
	TypeShort       byte = 2 // SMALLINT
	TypeLong        byte = 3 // INT
	TypeFloat       byte = 4
	TypeDouble      byte = 5
	TypeNull        byte = 6
	TypeTimestamp   byte = 7
	TypeLonglong    byte = 8 // BIGINT
	TypeInt24       byte = 9 // MEDIUMINT
	TypeDate        byte = 10
	/* TypeDuration original name was TypeTime, renamed to TypeDuration to resolve the conflict with Go type Time.*/
	TypeDuration byte = 11
	TypeDatetime byte = 12
	TypeYear     byte = 13
	TypeNewDate  byte = 14
	TypeVarchar  byte = 15
	TypeBit      byte = 16

	TypeJSON       byte = 0xf5
	TypeNewDecimal byte = 0xf6
	TypeEnum       byte = 0xf7
	TypeSet        byte = 0xf8
	TypeTinyBlob   byte = 0xf9
	TypeMediumBlob byte = 0xfa
	TypeLongBlob   byte = 0xfb
	TypeBlob       byte = 0xfc
	TypeVarString  byte = 0xfd
	TypeString     byte = 0xfe
	TypeGeometry   byte = 0xff
)

// HasUnsignedFlag checks if UnsignedFlag is set.
func HasUnsignedFlag(flag uint) bool {
	return (flag & UnsignedFlag) > 0
}

/********** copy from github.com/pingcap/tidb/types end ******************/
