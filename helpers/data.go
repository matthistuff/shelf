package helpers
import (
	"github.com/matthistuff/shelf/data"
	"fmt"
)

func ValidId(maybeId string) string {
	objectId, exists := data.AssertGuid(maybeId)
	ErrExit(objectId == "", "No object ID given!")
	ErrExit(!exists, fmt.Sprintf("No cached entry %s exists!", maybeId))

	return objectId
}