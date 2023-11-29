package use_case

import (
	"PostHubApp/domain/entity"
)

type ServiceProducer interface {
	Produce(message entity.Migrations) error
}
