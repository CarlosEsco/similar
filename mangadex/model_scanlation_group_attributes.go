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

type ScanlationGroupAttributes struct {
	Name            string              `json:"name,omitempty"`
	AltNames        []map[string]string `json:"altNames,omitempty"`
	Website         string              `json:"website,omitempty"`
	IrcServer       string              `json:"ircServer,omitempty"`
	IrcChannel      string              `json:"ircChannel,omitempty"`
	Discord         string              `json:"discord,omitempty"`
	ContactEmail    string              `json:"contactEmail,omitempty"`
	Description     string              `json:"description,omitempty"`
	Twitter         string              `json:"twitter,omitempty"`
	MangaUpdates    string              `json:"mangaUpdates,omitempty"`
	FocusedLanguage []string            `json:"focusedLanguage,omitempty"`
	Locked          bool                `json:"locked,omitempty"`
	Official        bool                `json:"official,omitempty"`
	Inactive        bool                `json:"inactive,omitempty"`
	// Should respected ISO 8601 duration specification: https://en.wikipedia.org/wiki/ISO_8601#Durations
	PublishDelay string `json:"publishDelay,omitempty"`
	Version      int32  `json:"version,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
}
