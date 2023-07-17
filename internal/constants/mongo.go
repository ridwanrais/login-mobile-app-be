package constants

import "os"

// database
var DATABASE_PRIMARY string

// collection
const COLLECTION_ACCOUNT = "accounts"

func InitConstants() {
	DATABASE_PRIMARY = os.Getenv("DATABASE_NAME")
}
