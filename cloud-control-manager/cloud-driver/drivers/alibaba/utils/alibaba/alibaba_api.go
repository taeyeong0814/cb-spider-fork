package alibaba

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func CreateCluster(access_key string, access_secret string, region_id string, body string) (string, error) {
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()

	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters"
	request.Headers["Content-Type"] = "application/json"

	request.Content = []byte(body)

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}

func GetClusters(access_key string, access_secret string, region_id string) (string, error) {

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()

	request.Method = "GET"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/api/v1/clusters"
	request.Headers["Content-Type"] = "application/json"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}

func GetCluster(access_key string, access_secret string, region_id string, cluster_id string) (string, error) {

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		return "", err
	}

	request := requests.NewCommonRequest()

	request.Method = "GET"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters/c622a22eab740403cb3e6e675c61a4e00"
	request.Headers["Content-Type"] = "application/json"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}

func DeleteCluster(access_key string, access_secret string, region_id string, cluster_id string) (string, error) {

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()

	request.Method = "DELETE"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters/" + cluster_id
	request.Headers["Content-Type"] = "application/json"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}

func CreateNodeGroup(access_key string, access_secret string, region_id string, cluster_id string, body string) (string, error) {

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		return "", err
	}

	request := requests.NewCommonRequest()

	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters/" + cluster_id + "/nodepools"
	request.Headers["Content-Type"] = "application/json"

	request.Content = []byte(body)

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}

func ListNodeGroup(access_key string, access_secret string, region_id string, cluster_id string) (string, error) {

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		return "", err
	}

	request := requests.NewCommonRequest()

	request.Method = "GET"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters/" + cluster_id + "/nodepools"
	request.Headers["Content-Type"] = "application/json"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}

func GetNodeGroup(access_key string, access_secret string, region_id string, cluster_id string, nodepool_id string) (string, error) {
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		return "", err
	}

	request := requests.NewCommonRequest()

	request.Method = "GET"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters/" + cluster_id + "/nodepools/" + nodepool_id
	request.Headers["Content-Type"] = "application/json"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}

func DeleteNodeGroup(access_key string, access_secret string, region_id string, cluster_id string, nodepool_id string) (string, error) {
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		return "", err
	}

	request := requests.NewCommonRequest()

	request.Method = "DELETE"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters/" + cluster_id + "/nodepools/" + nodepool_id
	request.Headers["Content-Type"] = "application/json"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}

func UpgradeCluster(access_key string, access_secret string, region_id string, cluster_id string, body string) (string, error) {

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(access_key, access_secret)
	client, err := sdk.NewClientWithOptions(region_id, config, credential)
	if err != nil {
		return "", err
	}

	request := requests.NewCommonRequest()

	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "cs." + region_id + ".aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/api/v2/clusters/" + cluster_id + "/upgrade"
	request.Headers["Content-Type"] = "application/json"

	// {
	//   "next_version" : "1.22.3-aliyun.1"
	// }
	request.Content = []byte(body)

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}

	return response.GetHttpContentString(), nil
}
