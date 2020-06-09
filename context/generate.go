package context

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateReqID() string {
	reqID := uuid.New().String()
	return strings.ReplaceAll(reqID, "-", "")
}
