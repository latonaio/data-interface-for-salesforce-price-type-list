package resources

import (
	"errors"
	"fmt"
)

// PriceType struct
type PriceType struct {
	method   string
	metadata map[string]interface{}
}

func (p *PriceType) objectName() string {
	const obName = "PriceTypeList"
	return obName
}

func NewPriceType(metadata map[string]interface{}) (*PriceType, error) {
	rawMethod, ok := metadata["method"]
	if !ok {
		return nil, errors.New("missing required parameters: method")
	}
	method, ok := rawMethod.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	return &PriceType{
		method:   method,
		metadata: metadata,
	}, nil
}

// getMetadata mold PriceType get metadata
func (p *PriceType) getMetadata() (map[string]interface{}, error) {
	idIF, ok := p.metadata["price_type_id"]
	if !ok {
		return buildMetadata(p.method, p.objectName(), "", nil, "", "price_type_get"), nil
	}
	queryParam, ok := idIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	return buildMetadata(p.method, p.objectName(), "", map[string]string{"priceTypeId": queryParam}, "", "price_type_get"), nil
}

// BuildMetadata
func (p *PriceType) BuildMetadata() (map[string]interface{}, error) {
	switch p.method {
	case "get":
		return p.getMetadata()
	}
	return nil, fmt.Errorf("invalid method: %s", p.method)
}

func buildMetadata(method, object, pathParam string, queryParams map[string]string, body string, connectionKey string) map[string]interface{} {
	metadata := map[string]interface{}{
		"method":         method,
		"object":         object,
		"connection_key": connectionKey,
	}
	if len(pathParam) > 0 {
		metadata["path_param"] = pathParam
	}
	if queryParams != nil {
		metadata["query_params"] = queryParams
	}
	if body != "" {
		metadata["body"] = body
	}
	return metadata
}
