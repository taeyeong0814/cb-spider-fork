package resources

import (
	"encoding/json"
	"fmt"
	"strings"

	idrv "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces"
	irs "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces/resources"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/pricing"
)

type AwsPriceInfoHandler struct {
	Region idrv.RegionInfo
	Client *pricing.Pricing
}

// AWS에서는 Region이 Product list에 영향을 주지 않습니다.
// 3개 Region Endpoint에서만 Product 정보를 리턴합니다.
// getPricingClient에 Client *pricing.Pricing 정의
func (priceInfoHandler *AwsPriceInfoHandler) ListProductFamily(regionName string) ([]string, error) {
	var result []string
	result = append(result, "AmazonEC2")
	// input := &pricing.GetAttributeValuesInput{
	// 	AttributeName: aws.String("productfamily"),
	// 	MaxResults:    aws.Int64(32), // 2024.01 기준 32개
	// 	ServiceCode:   aws.String("AmazonEC2"),
	// }

	// cblogger.Info("input 1321312434242341312312", input)
	// for {
	// 	attributeValues, err := priceInfoHandler.Client.GetAttributeValues(input)
	// 	if err != nil {
	// 		if aerr, ok := err.(awserr.Error); ok {
	// 			switch aerr.Code() {
	// 			case pricing.ErrCodeInternalErrorException:
	// 				cblogger.Error(pricing.ErrCodeInternalErrorException, aerr.Error())
	// 			case pricing.ErrCodeInvalidParameterException:
	// 				cblogger.Error(pricing.ErrCodeInvalidParameterException, aerr.Error())
	// 			case pricing.ErrCodeNotFoundException:
	// 				cblogger.Error(pricing.ErrCodeNotFoundException, aerr.Error())
	// 			case pricing.ErrCodeInvalidNextTokenException:
	// 				cblogger.Error(pricing.ErrCodeInvalidNextTokenException, aerr.Error())
	// 			case pricing.ErrCodeExpiredNextTokenException:
	// 				cblogger.Error(pricing.ErrCodeExpiredNextTokenException, aerr.Error())
	// 			default:
	// 				cblogger.Error(aerr.Error())
	// 			}
	// 		} else {
	// 			// Prnit the error, cast err to awserr.Error to get the Code and
	// 			// Message from an error.
	// 			cblogger.Error(err.Error())
	// 		}
	// 	}

	// 	for _, attributeValue := range attributeValues.AttributeValues {

	// 		//result = append(result, *attributeValue.Value)
	// 		result = append(result, *attributeValue.Value)
	// 	}

	// 	for i := range attributeValues.AttributeValues {
	// 		result[i] = removeSpaces(result[i])
	// 	}

	// 	for _, attributeValue := range attributeValues.AttributeValues {
	// 		attributeValue.Value = aws.String(strings.ReplaceAll(*attributeValue.Value, " ", ""))
	// 	}

	// 	// 결과 출력
	// 	cblogger.Info("rkskekfkekfkekfkekf", attributeValues)
	// 	fmt.Printf("%+v\n", attributeValues)

	// 	cblogger.Info("attributeValue0000000000000000000000000000", attributeValues.AttributeValues)

	// 	cblogger.Info("attributeValue===============================", result)
	// 	if attributeValues.NextToken != nil {
	// 		input = &pricing.GetAttributeValuesInput{
	// 			NextToken: attributeValues.NextToken,
	// 		}
	// 	} else {
	// 		break
	// 	}
	// }

	return result, nil
}
func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// AWS에서는 ListProductFamily를 통해 ProductFamily와 AttributeName을 수집하고,
// GetAttributeValues를 통해 AttributeValue를 수집하여 필터로 사용합니다.
// GetPriceInfo는 DescribeServices를 통해 올바른 productFamily 인자만 검사합니다. -> AttributeName에 오류가 있을경우 빈값을 리턴

func (priceInfoHandler *AwsPriceInfoHandler) GetPriceInfo(productFamily string, regionName string, filterList []irs.KeyValue) (string, error) {
	priceMap := make(map[string]irs.Price)
	cblogger.Infof("productFamily======", productFamily)

	cblogger.Infof("filter value : %+v", filterList)
	describeServicesinput := &pricing.DescribeServicesInput{
		ServiceCode: aws.String(productFamily),
		MaxResults:  aws.Int64(1),
	}
	// for the test
	// cblogger.Info("describeServicesinput", describeServicesinput)

	services, err := priceInfoHandler.Client.DescribeServices(describeServicesinput)
	if services == nil {
		cblogger.Error("No services in given productFamily. CHECK productFamily!")
		return "", err
	}

	if err != nil {
		cblogger.Error(err)
		return "", err
	}
	getProductsinputfilters := []*pricing.Filter{}

	if filterList != nil {
		for _, filter := range filterList {
			if filter.Key == "instanceType" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("instanceType"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}

			if filter.Key == "operatingSystem" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("operatingSystem"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}
			if filter.Key == "vcpu" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("vcpu"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}
			if filter.Key == "productId" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("sku"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}
			if filter.Key == "memory" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("memory"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}
			if filter.Key == "storage" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("storage"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}
			if filter.Key == "gpu" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("gpu"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}
			if filter.Key == "gpuMemory" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("gpuMemory"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}
			if filter.Key == "preInstalledSw" {
				getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
					Field: aws.String("preInstalledSw"),
					Type:  aws.String("TERM_MATCH"),
					Value: aws.String(filter.Value),
				})
			}

		}
	}
	if regionName != "" {
		getProductsinputfilters = append(getProductsinputfilters, &pricing.Filter{
			Field: aws.String("regionCode"),
			Type:  aws.String("EQUALS"),
			Value: aws.String(regionName),
		})
	}

	getProductsinput := &pricing.GetProductsInput{
		Filters:     getProductsinputfilters,
		ServiceCode: aws.String(productFamily),
	}

	priceinfos, err := priceInfoHandler.Client.GetProducts(getProductsinput)
	if err != nil {
		cblogger.Error(err)
		return "", err
	}

	result := &irs.CloudPriceData{}
	result.Meta.Version = "v0.1"
	result.Meta.Description = "Multi-Cloud Price Info"
	// for the test
	//  "productInfo", priceinfos)
	for _, price := range priceinfos.PriceList {
		jsonString, err := json.MarshalIndent(price["product"].(map[string]interface{})["attributes"], "", "    ")
		if err != nil {
			cblogger.Error(err)
		}

		var productInfo irs.ProductInfo
		ReplaceEmptyWithNA(&productInfo)
		err = json.Unmarshal(jsonString, &productInfo)
		if err != nil {
			cblogger.Error(err)
		}

		productId := fmt.Sprintf("%s", price["product"].(map[string]interface{})["sku"])
		productInfo.ProductId = fmt.Sprintf("%s", price["product"].(map[string]interface{})["sku"])
		productInfo.RegionName = fmt.Sprintf("%s", price["product"].(map[string]interface{})["attributes"].(map[string]interface{})["regionCode"])
		productInfo.Description = fmt.Sprintf("productFamily= %s, version= %s", price["product"].(map[string]interface{})["productFamily"], price["version"])
		productInfo.CSPProductInfo = price["product"]
		productInfo.ZoneName = "NA" // AWS zone is Not Applicable - 202401

		var priceInfo irs.PriceInfo
		priceInfo.CSPPriceInfo = price["terms"]
		for termsKey, termsValue := range price["terms"].(map[string]interface{}) {

			hasTerm := false
			termVal := ""
			hasPriceDimension := false
			priceDemensionVal := ""
			hasunit := false
			unitVal := ""
			if filterList != nil {

				for _, filter := range filterList {
					// find filter conditions
					if filter.Key == "pricingPolicy" {
						hasTerm = true
						termVal = filter.Value
						continue
					}

					if filter.Key == "pricingId" {
						hasPriceDimension = true
						priceDemensionVal = filter.Value
						continue
					}
					if filter.Key == "unit" {
						hasunit = true
						unitVal = filter.Value
						continue
					}
				}
				// check filters
				if hasTerm && termVal != termsKey {
					continue
				}
			}

			for _, policyvalue := range termsValue.(map[string]interface{}) {
				var pricingPolicy irs.PricingPolicies
				for innerpolicyKey, innerpolicyValue := range policyvalue.(map[string]interface{}) {
					if innerpolicyKey == "priceDimensions" {
						for priceDimensionsKey, priceDimensionsValue := range innerpolicyValue.(map[string]interface{}) {
							if filterList != nil {
								// check filters
								if hasPriceDimension && priceDemensionVal != priceDimensionsKey {

									continue

								}
								//pricingId의 unit값이 필터 값으로 들어오면 unit 값을 받은 값으로 설정
								foundSku := false
								for _, skukey := range priceDimensionsValue.(map[string]interface{}) {
									// check filters
									if hasunit && unitVal == skukey {
										foundSku = true
										break
									}
								}
								if hasunit && !foundSku { // sku를 못 찾았으면 skip.
									continue
								}
							}

							pricingPolicy.PricingId = priceDimensionsKey
							pricingPolicy.PricingPolicy = termsKey
							pricingPolicy.Description = fmt.Sprintf("%s", priceDimensionsValue.(map[string]interface{})["description"])
							for key, val := range priceDimensionsValue.(map[string]interface{})["pricePerUnit"].(map[string]interface{}) {
								pricingPolicy.Currency = key
								pricingPolicy.Price = fmt.Sprintf("%s", val)
								// USD is Default.
								// if NO USD data, accept other currency.
								if key == "USD" {
									break
								}
							}
							pricingPolicy.Unit = fmt.Sprintf("%s", priceDimensionsValue.(map[string]interface{})["unit"])

							var cspPriceInfo []string

							// Convert the []interface{} to []string before appending
							for _, item := range price["terms"].(map[string]interface{}) {
								jsonString, err := json.Marshal(item)
								if err != nil {
									cblogger.Error(err)
									continue
								}
								cspPriceInfo = append(cspPriceInfo, string(jsonString))
							}

							priceInfo.PricingPolicies = append(priceInfo.PricingPolicies, pricingPolicy)
							aPrice, ok := priceMap[productId]

							if ok { // product가 존재하면 policy 추가
								aPrice.PriceInfo.PricingPolicies = append(aPrice.PriceInfo.PricingPolicies, pricingPolicy)
								aPrice.PriceInfo.CSPPriceInfo = append(aPrice.PriceInfo.CSPPriceInfo.([][]string), cspPriceInfo)
								// var priceInfo irs.PriceInfo
								// priceInfo.CSPPriceInfo = price["terms"]
								priceMap[productId] = aPrice

							} else { // product가 없으면 price 추가

								newPriceInfo := irs.PriceInfo{}
								newPolicies := []irs.PricingPolicies{}
								newPolicies = append(newPolicies, pricingPolicy)

								newPriceInfo.PricingPolicies = newPolicies

								newCSPPriceInfo := [][]string{}
								newCSPPriceInfo = append(newCSPPriceInfo, cspPriceInfo)

								newPrice := irs.Price{}
								newPrice.PriceInfo = newPriceInfo
								newPrice.ProductInfo = productInfo
								newPrice.PriceInfo.CSPPriceInfo = newCSPPriceInfo

								priceMap[productId] = newPrice
							}
						}
					}
				}
			}
		}

		// price info
		// var priceListone irs.Price
		// priceListone.ProductInfo = productInfo
		// priceListone.PriceInfo = priceInfo
		priceList := []irs.Price{}
		for _, value := range priceMap {
			priceList = append(priceList, value)
		}

		priceone := irs.CloudPrice{
			CloudName: "AWS",
		}
		// priceone.PriceList = append(priceone.PriceList, priceList...)
		priceone.PriceList = priceList
		result.CloudPriceList = append(result.CloudPriceList, priceone)
	}

	resultString, err := json.Marshal(result)
	if err != nil {
		cblogger.Error(err)
		return "", err
	}

	return string(resultString), nil
}
