
package ft

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

func (client *Client) RpcHttpApi2(request *RpcHttpApi2Request) (response *RpcHttpApi2Response, err error) {
response = CreateRpcHttpApi2Response()
err = client.DoAction(request, response)
return
}

func (client *Client) RpcHttpApi2WithChan(request *RpcHttpApi2Request) (<-chan *RpcHttpApi2Response, <-chan error) {
responseChan := make(chan *RpcHttpApi2Response, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.RpcHttpApi2(request)
responseChan <- response
errChan <- err
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

func (client *Client) RpcHttpApi2WithCallback(request *RpcHttpApi2Request, callback func(response *RpcHttpApi2Response, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *RpcHttpApi2Response
var err error
defer close(result)
response, err = client.RpcHttpApi2(request)
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

type RpcHttpApi2Request struct {
*requests.RpcRequest
            Sleep  string `position:"Query" name:"Sleep"`
            P1  string `position:"Query" name:"P1"`
}


type RpcHttpApi2Response struct {
*responses.BaseResponse
            Params struct {
            RequestId     string `json:"RequestId"`
            CallerUid     string `json:"CallerUid"`
            Sleep     string `json:"Sleep"`
            P1     string `json:"P1"`
            }  `json:"Params"`
}

func CreateRpcHttpApi2Request() (request *RpcHttpApi2Request) {
request = &RpcHttpApi2Request{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ft", "2015-01-01", "RpcHttpApi2", "", "")
return
}

func CreateRpcHttpApi2Response() (response *RpcHttpApi2Response) {
response = &RpcHttpApi2Response{
BaseResponse: &responses.BaseResponse{},
}
return
}
