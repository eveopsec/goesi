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
)

type UserInterfaceApi struct {
	Configuration *Configuration
}

func NewUserInterfaceApi() *UserInterfaceApi {
	configuration := NewConfiguration()
	return &UserInterfaceApi{
		Configuration: configuration,
	}
}

func NewUserInterfaceApiWithBasePath(basePath string) *UserInterfaceApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &UserInterfaceApi{
		Configuration: configuration,
	}
}

/**
 * Set Autopilot Waypoint
 * Set a solar system as autopilot waypoint  ---  Alternate route: &#x60;/v2/ui/autopilot/waypoint/&#x60;  Alternate route: &#x60;/dev/ui/autopilot/waypoint/&#x60; 
 *
 * @param addToBeginning Whether this solar system should be added to the beginning of all waypoints
 * @param clearOtherWaypoints Whether clean other waypoints beforing adding this one
 * @param destinationId The destination to travel to, can be solar system, station or structure&#39;s id
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return void
 */
func (a UserInterfaceApi) PostUiAutopilotWaypoint(addToBeginning bool, clearOtherWaypoints bool, destinationId int64, datasource string, token string, userAgent string, xUserAgent string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ui/autopilot/waypoint/"

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
	localVarQueryParams.Add("add_to_beginning", a.Configuration.APIClient.ParameterToString(addToBeginning, ""))
	localVarQueryParams.Add("clear_other_waypoints", a.Configuration.APIClient.ParameterToString(clearOtherWaypoints, ""))
	localVarQueryParams.Add("datasource", a.Configuration.APIClient.ParameterToString(datasource, ""))
	localVarQueryParams.Add("destination_id", a.Configuration.APIClient.ParameterToString(destinationId, ""))
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
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "PostUiAutopilotWaypoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * Open Contract Window
 * Open the contract window inside the client  ---  Alternate route: &#x60;/v1/ui/openwindow/contract/&#x60;  Alternate route: &#x60;/legacy/ui/openwindow/contract/&#x60;  Alternate route: &#x60;/dev/ui/openwindow/contract/&#x60; 
 *
 * @param contractId The contract to open
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return void
 */
func (a UserInterfaceApi) PostUiOpenwindowContract(contractId int32, datasource string, token string, userAgent string, xUserAgent string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ui/openwindow/contract/"

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
	localVarQueryParams.Add("contract_id", a.Configuration.APIClient.ParameterToString(contractId, ""))
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
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "PostUiOpenwindowContract", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * Open Information Window
 * Open the information window for a character, corporation or alliance inside the client  ---  Alternate route: &#x60;/v1/ui/openwindow/information/&#x60;  Alternate route: &#x60;/legacy/ui/openwindow/information/&#x60;  Alternate route: &#x60;/dev/ui/openwindow/information/&#x60; 
 *
 * @param targetId The target to open
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return void
 */
func (a UserInterfaceApi) PostUiOpenwindowInformation(targetId int32, datasource string, token string, userAgent string, xUserAgent string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ui/openwindow/information/"

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
	localVarQueryParams.Add("target_id", a.Configuration.APIClient.ParameterToString(targetId, ""))
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
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "PostUiOpenwindowInformation", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * Open Market Details
 * Open the market details window for a specific typeID inside the client  ---  Alternate route: &#x60;/v1/ui/openwindow/marketdetails/&#x60;  Alternate route: &#x60;/legacy/ui/openwindow/marketdetails/&#x60;  Alternate route: &#x60;/dev/ui/openwindow/marketdetails/&#x60; 
 *
 * @param typeId The item type to open in market window
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return void
 */
func (a UserInterfaceApi) PostUiOpenwindowMarketdetails(typeId int32, datasource string, token string, userAgent string, xUserAgent string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ui/openwindow/marketdetails/"

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
	localVarQueryParams.Add("type_id", a.Configuration.APIClient.ParameterToString(typeId, ""))
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
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "PostUiOpenwindowMarketdetails", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * Open New Mail Window
 * Open the New Mail window, according to settings from the request if applicable  ---  Alternate route: &#x60;/v1/ui/openwindow/newmail/&#x60;  Alternate route: &#x60;/legacy/ui/openwindow/newmail/&#x60;  Alternate route: &#x60;/dev/ui/openwindow/newmail/&#x60; 
 *
 * @param newMail The details of mail to create
 * @param datasource The server name you would like data from
 * @param token Access token to use, if preferred over a header
 * @param userAgent Client identifier, takes precedence over headers
 * @param xUserAgent Client identifier, takes precedence over User-Agent
 * @return void
 */
func (a UserInterfaceApi) PostUiOpenwindowNewmail(newMail PostUiOpenwindowNewmailNewMail, datasource string, token string, userAgent string, xUserAgent string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ui/openwindow/newmail/"

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
	localVarPostBody = &newMail
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "PostUiOpenwindowNewmail", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}
