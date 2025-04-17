package readRepo

import (
	"github.com/CallumLewisGH/Boondock-Service-Base/database"
	"github.com/google/uuid"
)

func Get[T any](Guid uuid.UUID) {
	database.ConnectDatabase()

}
