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

type CorporationApi struct {
	Configuration *Configuration
}

func NewCorporationApi() *CorporationApi {
	configuration := NewConfiguration()
	return &CorporationApi{
		Configuration: configuration,
	}
}

func NewCorporationApiWithBasePath(basePath string) *CorporationApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &CorporationApi{
		Configuration: configuration,
	}
}

/**
 * Get corporation information
 * Public information about a corporation  ---  Alternate route: &#x60;/v3/corporations/{corporation_id}/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId An Eve corporation ID
 * @param datasource The server name you would like data from
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return *GetCorporationsCorporationIdOk
 */
func (a CorporationApi) GetCorporationsCorporationId(corporationId int32, datasource string, userAgent string, xUserAgent string) (*GetCorporationsCorporationIdOk, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/{corporation_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("datasource", a.Configuration.APIClient.ParameterToString(datasource, ""))
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
	var successPayload = new(GetCorporationsCorporationIdOk)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCorporationsCorporationId", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

/**
 * Get alliance history
 * Get a list of all the alliances a corporation has been a member of  ---  Alternate route: &#x60;/v1/corporations/{corporation_id}/alliancehistory/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/alliancehistory/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/alliancehistory/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId An EVE corporation ID
 * @param datasource The server name you would like data from
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return []GetCorporationsCorporationIdAlliancehistory200Ok
 */
func (a CorporationApi) GetCorporationsCorporationIdAlliancehistory(corporationId int32, datasource string, userAgent string, xUserAgent string) ([]GetCorporationsCorporationIdAlliancehistory200Ok, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/{corporation_id}/alliancehistory/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("datasource", a.Configuration.APIClient.ParameterToString(datasource, ""))
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
	var successPayload = new([]GetCorporationsCorporationIdAlliancehistory200Ok)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCorporationsCorporationIdAlliancehistory", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Get corporation icon
 * Get the icon urls for a corporation  ---  Alternate route: &#x60;/v1/corporations/{corporation_id}/icons/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/icons/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/icons/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId An EVE corporation ID
 * @param datasource The server name you would like data from
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return *GetCorporationsCorporationIdIconsOk
 */
func (a CorporationApi) GetCorporationsCorporationIdIcons(corporationId int32, datasource string, userAgent string, xUserAgent string) (*GetCorporationsCorporationIdIconsOk, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/{corporation_id}/icons/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("datasource", a.Configuration.APIClient.ParameterToString(datasource, ""))
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
	var successPayload = new(GetCorporationsCorporationIdIconsOk)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCorporationsCorporationIdIcons", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

/**
 * Get corporation members
 * Read the current list of members if the calling character is a member.  ---  Alternate route: &#x60;/v2/corporations/{corporation_id}/members/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/members/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/members/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId A corporation ID
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return []GetCorporationsCorporationIdMembers200Ok
 */
func (a CorporationApi) GetCorporationsCorporationIdMembers(corporationId int32, datasource string, token string, userAgent string, xUserAgent string) ([]GetCorporationsCorporationIdMembers200Ok, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/{corporation_id}/members/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

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
	var successPayload = new([]GetCorporationsCorporationIdMembers200Ok)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCorporationsCorporationIdMembers", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Get corporation member roles
 * Return the roles of all members if the character has the personnel manager role or any grantable role.  ---  Alternate route: &#x60;/v1/corporations/{corporation_id}/roles/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/roles/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/roles/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId A corporation ID
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return []GetCorporationsCorporationIdRoles200Ok
 */
func (a CorporationApi) GetCorporationsCorporationIdRoles(corporationId int32, datasource string, token string, userAgent string, xUserAgent string) ([]GetCorporationsCorporationIdRoles200Ok, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/{corporation_id}/roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

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
	var successPayload = new([]GetCorporationsCorporationIdRoles200Ok)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCorporationsCorporationIdRoles", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Get corporation structures
 * Get a list of corporation structures  ---  Alternate route: &#x60;/v1/corporations/{corporation_id}/structures/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/structures/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/structures/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId A corporation ID
 * @param datasource The server name you would like data from
 * @param language Language to use in the response
 * @param page Which page to query, 250 structures per page
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return []GetCorporationsCorporationIdStructures200Ok
 */
func (a CorporationApi) GetCorporationsCorporationIdStructures(corporationId int32, datasource string, language string, page int32, token string, userAgent string, xUserAgent string) ([]GetCorporationsCorporationIdStructures200Ok, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/{corporation_id}/structures/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

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
	localVarQueryParams.Add("language", a.Configuration.APIClient.ParameterToString(language, ""))
	localVarQueryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
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
	var successPayload = new([]GetCorporationsCorporationIdStructures200Ok)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCorporationsCorporationIdStructures", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Get corporation names
 * Resolve a set of corporation IDs to corporation names  ---  Alternate route: &#x60;/v1/corporations/names/&#x60;  Alternate route: &#x60;/legacy/corporations/names/&#x60;  Alternate route: &#x60;/dev/corporations/names/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationIds A comma separated list of corporation IDs
 * @param datasource The server name you would like data from
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return []GetCorporationsNames200Ok
 */
func (a CorporationApi) GetCorporationsNames(corporationIds []int64, datasource string, userAgent string, xUserAgent string) ([]GetCorporationsNames200Ok, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/names/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var collectionFormat = "multi"
	localVarQueryParams.Add("corporation_ids", a.Configuration.APIClient.ParameterToString(corporationIds, collectionFormat))

	localVarQueryParams.Add("datasource", a.Configuration.APIClient.ParameterToString(datasource, ""))
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
	var successPayload = new([]GetCorporationsNames200Ok)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCorporationsNames", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Get npc corporations
 * Get a list of npc corporations  ---  Alternate route: &#x60;/v1/corporations/npccorps/&#x60;  Alternate route: &#x60;/legacy/corporations/npccorps/&#x60;  Alternate route: &#x60;/dev/corporations/npccorps/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param datasource The server name you would like data from
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return []int32
 */
func (a CorporationApi) GetCorporationsNpccorps(datasource string, userAgent string, xUserAgent string) ([]int32, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/npccorps/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("datasource", a.Configuration.APIClient.ParameterToString(datasource, ""))
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
	var successPayload = new([]int32)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCorporationsNpccorps", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Update structure vulnerability schedule
 * Update the vulnerability window schedule of a corporation structure  ---  Alternate route: &#x60;/v1/corporations/{corporation_id}/structures/{structure_id}/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/structures/{structure_id}/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/structures/{structure_id}/&#x60; 
 *
 * @param corporationId A corporation ID
 * @param newSchedule New vulnerability window schedule for the structure
 * @param structureId A structure ID
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return void
 */
func (a CorporationApi) PutCorporationsCorporationIdStructuresStructureId(corporationId int32, newSchedule []PutCorporationsCorporationIdStructuresStructureIdNewSchedule, structureId int64, datasource string, token string, userAgent string, xUserAgent string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/corporations/{corporation_id}/structures/{structure_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"structure_id"+"}", fmt.Sprintf("%v", structureId), -1)

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
	// body params
	localVarPostBody = &newSchedule
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "PutCorporationsCorporationIdStructuresStructureId", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

