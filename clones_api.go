/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.1.dev1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type ClonesApi struct {
	Configuration *Configuration
}

func NewClonesApi() *ClonesApi {
	configuration := NewConfiguration()
	return &ClonesApi{
		Configuration: configuration,
	}
}

func NewClonesApiWithBasePath(basePath string) *ClonesApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ClonesApi{
		Configuration: configuration,
	}
}

/**
 * Get clones
 * A list of the character&#39;s clones  ---  Alternate route: &#x60;/v2/characters/{character_id}/clones/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/clones/&#x60;   ---  This route is cached for up to 120 seconds
 *
 * @param characterId An EVE character ID
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return *GetCharactersCharacterIdClonesOk
 */
func (a ClonesApi) GetCharactersCharacterIdClones(characterId int32, datasource string, token string, userAgent string, xUserAgent string) (*GetCharactersCharacterIdClonesOk, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/characters/{character_id}/clones/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(evesso)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("datasource", a.Configuration.APIClient.ParameterToString(datasource, ""))
	localVarQueryParams.Add("token", a.Configuration.APIClient.ParameterToString(token, ""))
	localVarQueryParams.Add("user_agent", a.Configuration.APIClient.ParameterToString(userAgent, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "X-User-Agent"
	localVarHeaderParams["X-User-Agent"] = a.Configuration.APIClient.ParameterToString(xUserAgent, "")
	var successPayload = new(GetCharactersCharacterIdClonesOk)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCharactersCharacterIdClones", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

