package injector

import (
	mongodb "crud/adapters/outbound/repository"
	"crud/domain/repository"
	"crud/infra/persistence"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/wire"
)

func ProvideProductRepository(db *mongo.Database) repository.IProductDomainRepo {
	return mongodb.NewProductRepository(db)
}

var RepositoryProviderSet = wire.NewSet(
	persistence.SetupMongoDB,
	ProvideProductRepository,
)
