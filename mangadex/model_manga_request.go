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

type MangaRequest struct {
	Title                  *map[string]string  `json:"title,omitempty"`
	AltTitles              []map[string]string `json:"altTitles,omitempty"`
	Description            *map[string]string  `json:"description,omitempty"`
	Authors                []string            `json:"authors,omitempty"`
	Artists                []string            `json:"artists,omitempty"`
	Links                  map[string]string   `json:"links,omitempty"`
	OriginalLanguage       string              `json:"originalLanguage,omitempty"`
	LastVolume             string              `json:"lastVolume,omitempty"`
	LastChapter            string              `json:"lastChapter,omitempty"`
	PublicationDemographic string              `json:"publicationDemographic,omitempty"`
	Status                 string              `json:"status,omitempty"`
	// Year of release
	Year                           int32    `json:"year,omitempty"`
	ContentRating                  string   `json:"contentRating,omitempty"`
	ChapterNumbersResetOnNewVolume bool     `json:"chapterNumbersResetOnNewVolume,omitempty"`
	Tags                           []string `json:"tags,omitempty"`
	PrimaryCover                   string   `json:"primaryCover,omitempty"`
	Version                        int32    `json:"version,omitempty"`
}
