package smartag

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateFlowLog invokes the smartag.CreateFlowLog API synchronously
func (client *Client) CreateFlowLog(request *CreateFlowLogRequest) (response *CreateFlowLogResponse, err error) {
	response = CreateCreateFlowLogResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFlowLogWithChan invokes the smartag.CreateFlowLog API asynchronously
func (client *Client) CreateFlowLogWithChan(request *CreateFlowLogRequest) (<-chan *CreateFlowLogResponse, <-chan error) {
	responseChan := make(chan *CreateFlowLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFlowLog(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateFlowLogWithCallback invokes the smartag.CreateFlowLog API asynchronously
func (client *Client) CreateFlowLogWithCallback(request *CreateFlowLogRequest, callback func(response *CreateFlowLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFlowLogResponse
		var err error
		defer close(result)
		response, err = client.CreateFlowLog(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateFlowLogRequest is the request struct for api CreateFlowLog
type CreateFlowLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	NetflowVersion       string           `position:"Query"`
	Description          string           `position:"Query"`
	InactiveAging        requests.Integer `position:"Query"`
	SlsRegionId          string           `position:"Query"`
	ActiveAging          requests.Integer `position:"Query"`
	OutputType           string           `position:"Query"`
	ProjectName          string           `position:"Query"`
	LogstoreName         string           `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	NetflowServerPort    requests.Integer `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	NetflowServerIp      string           `position:"Query"`
	Name                 string           `position:"Query"`
}

// CreateFlowLogResponse is the response struct for api CreateFlowLog
type CreateFlowLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	FlowLogId string `json:"FlowLogId" xml:"FlowLogId"`
}

// CreateCreateFlowLogRequest creates a request to invoke CreateFlowLog API
func CreateCreateFlowLogRequest() (request *CreateFlowLogRequest) {
	request = &CreateFlowLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "CreateFlowLog", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFlowLogResponse creates a response to parse from CreateFlowLog response
func CreateCreateFlowLogResponse() (response *CreateFlowLogResponse) {
	response = &CreateFlowLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}