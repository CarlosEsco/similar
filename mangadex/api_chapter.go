/*
 * MangaDex API
 *
 * MangaDex is an ad-free manga reader offering high-quality images!  This document details our API as it is right now. It is in no way a promise to never change it, although we will endeavour to publicly notify any major change.  # Acceptable use policy  Usage of our services implies acceptance of the following: - You **MUST** credit us - You **MUST** credit scanlation groups if you offer the ability to read chapters - You **CANNOT** run ads or paid services on your website and/or apps  These may change at any time for any and no reason and it is up to you check for updates from time to time.  # Security issues  If you believe you found a security issue in our API, please check our [security.txt](/security.txt) to get in touch privately.
 *
 * API version: 5.9.0
 * Contact: support@mangadex.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mangadex

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ChapterApiService service

/*
ChapterApiService Delete Chapter
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id Chapter ID

@return Response
*/
func (a *ChapterApiService) DeleteChapterId(ctx context.Context, id string) (Response, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/chapter/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Response
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChapterApiService Chapter list
Chapter list. If you want the Chapters of a given Manga, please check the feed endpoints.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ChapterApiGetChapterOpts - Optional Parameters:
     * @param "Limit" (optional.Int32) -
     * @param "Offset" (optional.Int32) -
     * @param "Ids" (optional.Interface of []string) -  Chapter ids (limited to 100 per request)
     * @param "Title" (optional.String) -
     * @param "Groups" (optional.Interface of []string) -
     * @param "Uploader" (optional.Interface of Uploader) -
     * @param "Manga" (optional.Interface of string) -
     * @param "Volume" (optional.Interface of Volume) -
     * @param "Chapter" (optional.Interface of Chapter) -
     * @param "TranslatedLanguage" (optional.Interface of []string) -
     * @param "OriginalLanguage" (optional.Interface of []string) -
     * @param "ExcludedOriginalLanguage" (optional.Interface of []string) -
     * @param "ContentRating" (optional.Interface of []string) -
     * @param "ExcludedGroups" (optional.Interface of []string) -
     * @param "ExcludedUploaders" (optional.Interface of []string) -
     * @param "IncludeFutureUpdates" (optional.String) -
     * @param "IncludeEmptyPages" (optional.Int32) -
     * @param "IncludeFuturePublishAt" (optional.Int32) -
     * @param "IncludeExternalUrl" (optional.Int32) -
     * @param "CreatedAtSince" (optional.String) -
     * @param "UpdatedAtSince" (optional.String) -
     * @param "PublishAtSince" (optional.String) -
     * @param "OrderCreatedAt" (optional.String) -
     * @param "OrderUpdatedAt" (optional.String) -
     * @param "OrderPublishAt" (optional.String) -
     * @param "OrderReadableAt" (optional.String) -
     * @param "OrderVolume" (optional.String) -
     * @param "OrderChapter" (optional.String) -
     * @param "Includes" (optional.Interface of []string) -
@return ChapterList
*/

type ChapterApiGetChapterOpts struct {
	Limit                    optional.Int32
	Offset                   optional.Int32
	Ids                      optional.Interface
	Title                    optional.String
	Groups                   optional.Interface
	Uploader                 optional.Interface
	Manga                    optional.Interface
	Volume                   optional.Interface
	Chapter                  optional.Interface
	TranslatedLanguage       optional.Interface
	OriginalLanguage         optional.Interface
	ExcludedOriginalLanguage optional.Interface
	ContentRating            optional.Interface
	ExcludedGroups           optional.Interface
	ExcludedUploaders        optional.Interface
	IncludeFutureUpdates     optional.String
	IncludeEmptyPages        optional.Int32
	IncludeFuturePublishAt   optional.Int32
	IncludeExternalUrl       optional.Int32
	CreatedAtSince           optional.String
	UpdatedAtSince           optional.String
	PublishAtSince           optional.String
	OrderCreatedAt           optional.String
	OrderUpdatedAt           optional.String
	OrderPublishAt           optional.String
	OrderReadableAt          optional.String
	OrderVolume              optional.String
	OrderChapter             optional.String
	Includes                 optional.Interface
}

func (a *ChapterApiService) GetChapter(ctx context.Context, localVarOptionals *ChapterApiGetChapterOpts) (ChapterList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ChapterList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/chapter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ids.IsSet() {
		localVarQueryParams.Add("ids[]", parameterToString(localVarOptionals.Ids.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Title.IsSet() {
		localVarQueryParams.Add("title", parameterToString(localVarOptionals.Title.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Groups.IsSet() {
		localVarQueryParams.Add("groups[]", parameterToString(localVarOptionals.Groups.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Uploader.IsSet() {
		localVarQueryParams.Add("uploader", parameterToString(localVarOptionals.Uploader.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Manga.IsSet() {
		localVarQueryParams.Add("manga", parameterToString(localVarOptionals.Manga.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Volume.IsSet() {
		localVarQueryParams.Add("volume[]", parameterToString(localVarOptionals.Volume.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Chapter.IsSet() {
		localVarQueryParams.Add("chapter", parameterToString(localVarOptionals.Chapter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TranslatedLanguage.IsSet() {
		localVarQueryParams.Add("translatedLanguage[]", parameterToString(localVarOptionals.TranslatedLanguage.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.OriginalLanguage.IsSet() {
		localVarQueryParams.Add("originalLanguage[]", parameterToString(localVarOptionals.OriginalLanguage.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludedOriginalLanguage.IsSet() {
		localVarQueryParams.Add("excludedOriginalLanguage[]", parameterToString(localVarOptionals.ExcludedOriginalLanguage.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ContentRating.IsSet() {
		localVarQueryParams.Add("contentRating[]", parameterToString(localVarOptionals.ContentRating.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludedGroups.IsSet() {
		localVarQueryParams.Add("excludedGroups[]", parameterToString(localVarOptionals.ExcludedGroups.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludedUploaders.IsSet() {
		localVarQueryParams.Add("excludedUploaders[]", parameterToString(localVarOptionals.ExcludedUploaders.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeFutureUpdates.IsSet() {
		localVarQueryParams.Add("includeFutureUpdates", parameterToString(localVarOptionals.IncludeFutureUpdates.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeEmptyPages.IsSet() {
		localVarQueryParams.Add("includeEmptyPages", parameterToString(localVarOptionals.IncludeEmptyPages.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeFuturePublishAt.IsSet() {
		localVarQueryParams.Add("includeFuturePublishAt", parameterToString(localVarOptionals.IncludeFuturePublishAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeExternalUrl.IsSet() {
		localVarQueryParams.Add("includeExternalUrl", parameterToString(localVarOptionals.IncludeExternalUrl.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreatedAtSince.IsSet() {
		localVarQueryParams.Add("createdAtSince", parameterToString(localVarOptionals.CreatedAtSince.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdatedAtSince.IsSet() {
		localVarQueryParams.Add("updatedAtSince", parameterToString(localVarOptionals.UpdatedAtSince.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PublishAtSince.IsSet() {
		localVarQueryParams.Add("publishAtSince", parameterToString(localVarOptionals.PublishAtSince.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderCreatedAt.IsSet() {
		localVarQueryParams.Add("order[createdAt]", parameterToString(localVarOptionals.OrderCreatedAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderUpdatedAt.IsSet() {
		localVarQueryParams.Add("order[updatedAt]", parameterToString(localVarOptionals.OrderUpdatedAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderPublishAt.IsSet() {
		localVarQueryParams.Add("order[publishAt]", parameterToString(localVarOptionals.OrderPublishAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderReadableAt.IsSet() {
		localVarQueryParams.Add("order[readableAt]", parameterToString(localVarOptionals.OrderReadableAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderVolume.IsSet() {
		localVarQueryParams.Add("order[volume]", parameterToString(localVarOptionals.OrderVolume.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderChapter.IsSet() {
		localVarQueryParams.Add("order[chapter]", parameterToString(localVarOptionals.OrderChapter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Includes.IsSet() {
		localVarQueryParams.Add("includes", parameterToString(localVarOptionals.Includes.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ChapterList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChapterApiService Get Chapter
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Chapter ID
 * @param optional nil or *ChapterApiGetChapterIdOpts - Optional Parameters:
     * @param "Includes" (optional.Interface of []string) -
@return ChapterResponse
*/

type ChapterApiGetChapterIdOpts struct {
	Includes optional.Interface
}

func (a *ChapterApiService) GetChapterId(ctx context.Context, id string, localVarOptionals *ChapterApiGetChapterIdOpts) (ChapterResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ChapterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/chapter/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Includes.IsSet() {
		localVarQueryParams.Add("includes[]", parameterToString(localVarOptionals.Includes.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ChapterResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChapterApiService Update Chapter
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param contentType
 * @param id Chapter ID
 * @param optional nil or *ChapterApiPutChapterIdOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ChapterEdit) -  The size of the body is limited to 32KB.
@return ChapterResponse
*/

type ChapterApiPutChapterIdOpts struct {
	Body optional.Interface
}

func (a *ChapterApiService) PutChapterId(ctx context.Context, contentType string, id string, localVarOptionals *ChapterApiPutChapterIdOpts) (ChapterResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ChapterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/chapter/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ChapterResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
