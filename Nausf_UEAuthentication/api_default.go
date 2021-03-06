/*
 * AUSF API
 *
 * OpenAPI specification for AUSF
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package Nausf_UEAuthentication

 import (
	 "free5gc/lib/openapi"
	 "free5gc/lib/openapi/models"
	 "free5gc/lib/openapi/logger"

	 "context"
	 "fmt"
	 "io/ioutil"
	 "net/http"
	 "net/url"
	 "strings"

	 "github.com/antihax/optional"
 )

 // Linger please
 var (
	 _ context.Context
 )

 type DefaultApiService service

 /*
 DefaultApiService
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param authCtxId
  * @param optional nil or *EapAuthMethodParamOpts - Optional Parameters:
  * @param "EapSession" (optional.Interface of models.EapSession) -
 @return models.EapSession
 */

 type EapAuthMethodParamOpts struct {
	 EapSession optional.Interface
 }

 func (a *DefaultApiService) EapAuthMethod(ctx context.Context, authCtxId string, localVarOptionals *EapAuthMethodParamOpts) (models.EapSession, *http.Response, error) {
	 var (
		 localVarHTTPMethod   = strings.ToUpper("Post")
		 localVarPostBody     interface{}
		 localVarFormFileName string
		 localVarFileName     string
		 localVarFileBytes    []byte
		 localVarReturnValue  models.EapSession
	 )

	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath() + "/ue-authentications/{authCtxId}/eap-session"
	 localVarPath = strings.Replace(localVarPath, "{"+"authCtxId"+"}", fmt.Sprintf("%v", authCtxId), -1)

	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}

	 localVarHTTPContentTypes := []string{"application/json"}

	 localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	 // to determine the Accept header
	 localVarHTTPHeaderAccepts := []string{"application/json", "application/3gppHal+json", "application/problem+json"}

	 // set Accept header
	 localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	 if localVarHTTPHeaderAccept != "" {
		 localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	 }

	 // body params
	 if localVarOptionals != nil && localVarOptionals.EapSession.IsSet() {
		 localVarOptionalEapSession, localVarOptionalEapSessionok := localVarOptionals.EapSession.Value().(models.EapSession)
		 if !localVarOptionalEapSessionok {
			 return localVarReturnValue, nil, openapi.ReportError("eapSession should be models.EapSession")
		 }
		 localVarPostBody = &localVarOptionalEapSession
	 }

	 r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	 if err != nil {
		 return localVarReturnValue, nil, err
	 }

	 localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	 if err != nil || localVarHTTPResponse == nil {
		 return localVarReturnValue, localVarHTTPResponse, err
	 }

	 localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	 localVarHTTPResponse.Body.Close()
	 if err != nil {
		 return localVarReturnValue, localVarHTTPResponse, err
	 }

	 apiError := openapi.GenericOpenAPIError{
		 RawBody:     localVarBody,
		 ErrorStatus: localVarHTTPResponse.Status,
	 }

	 switch localVarHTTPResponse.StatusCode {
	 case 200:
		 err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
		 }
		 return localVarReturnValue, localVarHTTPResponse, nil
	 case 400:
		 var v models.ProblemDetails
		 err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
			 return localVarReturnValue, localVarHTTPResponse, apiError
		 }
		 apiError.ErrorModel = v
		 return localVarReturnValue, localVarHTTPResponse, apiError
	 case 500:
		 var v models.ProblemDetails
		 err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
			 return localVarReturnValue, localVarHTTPResponse, apiError
		 }
		 apiError.ErrorModel = v
		 return localVarReturnValue, localVarHTTPResponse, apiError
	 default:
		 return localVarReturnValue, localVarHTTPResponse, openapi.ReportError("%d is not a valid status code in EapAuthMethod", localVarHTTPResponse.StatusCode)
	 }
 }

 /*
 DefaultApiService
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param authCtxId
  * @param optional nil or *UeAuthenticationsAuthCtxId5gAkaConfirmationPutParamOpts - Optional Parameters:
  * @param "ConfirmationData" (optional.Interface of models.ConfirmationData) -
 @return models.ConfirmationDataResponse
 */

 type UeAuthenticationsAuthCtxId5gAkaConfirmationPutParamOpts struct {
	 ConfirmationData optional.Interface
 }

 func (a *DefaultApiService) UeAuthenticationsAuthCtxId5gAkaConfirmationPut(ctx context.Context, authCtxId string, localVarOptionals *UeAuthenticationsAuthCtxId5gAkaConfirmationPutParamOpts) (models.ConfirmationDataResponse, *http.Response, error) {
	 var (
		 localVarHTTPMethod   = strings.ToUpper("Put")
		 localVarPostBody     interface{}
		 localVarFormFileName string
		 localVarFileName     string
		 localVarFileBytes    []byte
		 localVarReturnValue  models.ConfirmationDataResponse
	 )

	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath() + "/ue-authentications/{authCtxId}/5g-aka-confirmation"
	 localVarPath = strings.Replace(localVarPath, "{"+"authCtxId"+"}", fmt.Sprintf("%v", authCtxId), -1)

	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}

	 localVarHTTPContentTypes := []string{"application/json"}

	 localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	 // to determine the Accept header
	 localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	 // set Accept header
	 localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	 if localVarHTTPHeaderAccept != "" {
		 localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	 }

	 // body params
	 if localVarOptionals != nil && localVarOptionals.ConfirmationData.IsSet() {
		 localVarOptionalConfirmationData, localVarOptionalConfirmationDataok := localVarOptionals.ConfirmationData.Value().(models.ConfirmationData)
		 if !localVarOptionalConfirmationDataok {
			 return localVarReturnValue, nil, openapi.ReportError("confirmationData should be models.ConfirmationData")
		 }
		 localVarPostBody = &localVarOptionalConfirmationData
	 }

	 r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	 if err != nil {
		 return localVarReturnValue, nil, err
	 }

	 localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	 if err != nil || localVarHTTPResponse == nil {
		 return localVarReturnValue, localVarHTTPResponse, err
	 }

	 localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	 localVarHTTPResponse.Body.Close()
	 if err != nil {
		 return localVarReturnValue, localVarHTTPResponse, err
	 }

	 apiError := openapi.GenericOpenAPIError{
		 RawBody:     localVarBody,
		 ErrorStatus: localVarHTTPResponse.Status,
	 }

	 switch localVarHTTPResponse.StatusCode {
	 case 200:
		 err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
		 }
		 return localVarReturnValue, localVarHTTPResponse, nil
	 case 400:
		 var v models.ProblemDetails
		 err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
			 return localVarReturnValue, localVarHTTPResponse, apiError
		 }
		 apiError.ErrorModel = v
		 return localVarReturnValue, localVarHTTPResponse, apiError
	 case 500:
		 var v models.ProblemDetails
		 err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
			 return localVarReturnValue, localVarHTTPResponse, apiError
		 }
		 apiError.ErrorModel = v
		 return localVarReturnValue, localVarHTTPResponse, apiError
	 default:
		 return localVarReturnValue, localVarHTTPResponse, openapi.ReportError("%d is not a valid status code in UeAuthenticationsAuthCtxId5gAkaConfirmationPut", localVarHTTPResponse.StatusCode)
	 }
 }

 /*
 DefaultApiService
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param authenticationInfo
 @return models.UeAuthenticationCtx
 */

 func (a *DefaultApiService) UeAuthenticationsPost(ctx context.Context, authenticationInfo models.AuthenticationInfo) (models.UeAuthenticationCtx, *http.Response, error) {
	 var (
		 localVarHTTPMethod   = strings.ToUpper("Post")
		 localVarPostBody     interface{}
		 localVarFormFileName string
		 localVarFileName     string
		 localVarFileBytes    []byte
		 localVarReturnValue  models.UeAuthenticationCtx
	 )

	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath() + "/ue-authentications"

	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}

	 localVarHTTPContentTypes := []string{"application/json"}

	 localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	 // to determine the Accept header
	 localVarHTTPHeaderAccepts := []string{"application/3gppHal+json", "application/problem+json"}

	 // set Accept header
	 localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	 if localVarHTTPHeaderAccept != "" {
		 localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	 }
	 // suci convert to ipv6 address and insert it to kwaf header
	 // suci-0(SUPI type)-mcc-mnc-routingIndentifier-protectionScheme-homeNetworkPublicKeyIdentifier-schemeOutput.
	 logger.OpenApiLog.Info("[Fima] Start Header insertion Change")
	 const supiTypePlace = 1 //SUPI type
	 const mccPlace = 2
	 const mncPlace = 3
	 const routingIndentifier = 4
	 const schemePlace = 5
	 const homeNetworkPublicKeyIdentifier = 6
	 const schemeOutput = 7

	 suciPart := strings.Split(authenticationInfo.SupiOrSuci, "-")
	 suciPrefix := suciPart[0]
	 if suciPrefix == "imsi" || suciPrefix == "nai" {
		 logger.OpenApiLog.Info("Got supi\n")
	 } else if suciPrefix == "suci" {
		 if len(suciPart) < 6 {
			 logger.OpenApiLog.Info("Suci with wrong format\n")
		 }
	 }
	 //mccMnc := suciPart[mccPlace] + suciPart[mncPlace]
	 msin := suciPart[schemeOutput]
	 msinArray := strings.Split(msin,"")
	 //in case  of msin less than 10 digits [0000 7487] prepand will add more zeros
	 if (len(msinArray)) < 10 {
		 m := []string{}
		 for i := 0; i < 10-(len(msinArray)); i++ {
			 m = append(m,"0")
		 }
		 msinArray = append(m,msinArray...)
	 }
	 logger.OpenApiLog.Info("MSIN Array:",msinArray)
	 chunkSize := 4
	 var msinArrayDevided []string
	 for i := len(msinArray); i > 0 ; i -= chunkSize {
		 start:= i - chunkSize
		 if start < 0 {
			 start = 0
		 }
		 msinArrayDevided = append(msinArrayDevided , strings.Join(msinArray[start:i],""))
	 }
	 logger.OpenApiLog.Info("MSIN Array Devided is: ",msinArrayDevided)
	 kwafHeader := suciPart[mccPlace]+":"+suciPart[mncPlace]+":"+suciPart[routingIndentifier]+":"+suciPart[homeNetworkPublicKeyIdentifier]+":"+msinArrayDevided[2]+":"+msinArrayDevided[1]+":"+msinArrayDevided[0]
	 kwafHeaderFinal :=  "FE80:"+kwafHeader
	 logger.OpenApiLog.Info("This is final ipv6 header: ", kwafHeaderFinal)
	 localVarHeaderParams["kwaf-suci"] = kwafHeaderFinal

	 // body params
	 localVarPostBody = &authenticationInfo

	 r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	 if err != nil {
		 return localVarReturnValue, nil, err
	 }

	 localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	 if err != nil || localVarHTTPResponse == nil {
		 return localVarReturnValue, localVarHTTPResponse, err
	 }

	 localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	 localVarHTTPResponse.Body.Close()
	 if err != nil {
		 return localVarReturnValue, localVarHTTPResponse, err
	 }

	 apiError := openapi.GenericOpenAPIError{
		 RawBody:     localVarBody,
		 ErrorStatus: localVarHTTPResponse.Status,
	 }

	 switch localVarHTTPResponse.StatusCode {
	 case 201:
		 err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
		 }
		 return localVarReturnValue, localVarHTTPResponse, nil
	 case 400:
		 var v models.ProblemDetails
		 err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
			 return localVarReturnValue, localVarHTTPResponse, apiError
		 }
		 apiError.ErrorModel = v
		 return localVarReturnValue, localVarHTTPResponse, apiError
	 case 403:
		 var v models.ProblemDetails
		 err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
			 return localVarReturnValue, localVarHTTPResponse, apiError
		 }
		 apiError.ErrorModel = v
		 return localVarReturnValue, localVarHTTPResponse, apiError
	 case 500:
		 var v models.ProblemDetails
		 err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		 if err != nil {
			 apiError.ErrorStatus = err.Error()
			 return localVarReturnValue, localVarHTTPResponse, apiError
		 }
		 apiError.ErrorModel = v
		 return localVarReturnValue, localVarHTTPResponse, apiError
	 default:
		 return localVarReturnValue, localVarHTTPResponse, openapi.ReportError("%d is not a valid status code in UeAuthenticationsPost", localVarHTTPResponse.StatusCode)
	 }
 }