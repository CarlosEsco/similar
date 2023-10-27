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

type AuthorCreate struct {
	Name      string             `json:"name"`
	Biography *map[string]string `json:"biography,omitempty"`
	Twitter   string             `json:"twitter,omitempty"`
	Pixiv     string             `json:"pixiv,omitempty"`
	MelonBook string             `json:"melonBook,omitempty"`
	FanBox    string             `json:"fanBox,omitempty"`
	Booth     string             `json:"booth,omitempty"`
	NicoVideo string             `json:"nicoVideo,omitempty"`
	Skeb      string             `json:"skeb,omitempty"`
	Fantia    string             `json:"fantia,omitempty"`
	Tumblr    string             `json:"tumblr,omitempty"`
	Youtube   string             `json:"youtube,omitempty"`
	Weibo     string             `json:"weibo,omitempty"`
	Naver     string             `json:"naver,omitempty"`
	Website   string             `json:"website,omitempty"`
	Version   int32              `json:"version,omitempty"`
}
