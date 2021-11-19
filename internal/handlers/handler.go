package handlers

import (
	"fmt"

	"github.com/latonaio/salesforce-data-models"
	"github.com/latonaio/aion-core/pkg/log"
)

func HandlePriceType(metadata map[string]interface{}) error {
	pt, err := models.MetadataToPriceTypes(metadata)
	if err != nil {
		return fmt.Errorf("failed to convert price_type: %v", err)
	}
	if err := models.RegisterPriceTypesAndCacheClear(pt); err != nil {
		return fmt.Errorf("failed to register price_type: %v", err)
	}
	log.Print("Successful price_type registration.")
	return nil
}
