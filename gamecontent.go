package gonhl

import "fmt"

const endpointGameContent = "/game/%d/content"

// GetGameContent retrieves all media related content pertaining to a specific NHL game.
// Can contain various articles, videos, and images related to the game.
func GetGameContent(c *Client, id int) GameContent {
	var content GameContent
	status := c.makeRequest(fmt.Sprintf(endpointGameContent, id), nil, &content)
	fmt.Println(status)
	return content
}

type GameContent struct {
	Link      string `json:"link"`
	Editorial struct {
		Preview struct {
			Title     string        `json:"title"`
			TopicList string        `json:"topicList"`
			Items     []PreviewItem `json:"items"`
		} `json:"preview"`
		Articles struct {
			Title     string        `json:"title"`
			TopicList string        `json:"topicList"`
			Items     []interface{} `json:"items"`
		} `json:"articles"`
		Recap struct {
			Title     string      `json:"title"`
			TopicList string      `json:"topicList"`
			Items     []RecapItem `json:"items"`
		} `json:"recap"`
	} `json:"editorial"`
	Media struct {
		Epg        []Epg `json:"epg"`
		Milestones struct {
			Title       string          `json:"title"`
			StreamStart string          `json:"streamStart"`
			Items       []MilestoneItem `json:"items"`
		} `json:"milestones"`
	} `json:"media"`
	Highlights struct {
		Scoreboard Highlight `json:"scoreboard"`
		GameCenter Highlight `json:"gameCenter"`
	} `json:"highlights"`
}

type PreviewItem struct {
	Type            string           `json:"type"`
	State           string           `json:"state"`
	Date            string           `json:"date"`
	ID              string           `json:"id"`
	Headline        string           `json:"headline"`
	Subhead         string           `json:"subhead"`
	SeoTitle        string           `json:"seoTitle"`
	SeoDescription  string           `json:"seoDescription"`
	SeoKeywords     string           `json:"seoKeywords"`
	Slug            string           `json:"slug"`
	Commenting      bool             `json:"commenting"`
	Tagline         string           `json:"tagline"`
	TokenData       map[string]Token `json:"tokenData"`
	Contributor     Contributor      `json:"contributor"`
	KeywordsDisplay []Keyword        `json:"keywordsDisplay"`
	KeywordsAll     []Keyword        `json:"keywordsAll"`
	Approval        string           `json:"approval"`
	URL             string           `json:"url"`
	DataURI         string           `json:"dataURI"`
	PrimaryKeyword  Keyword          `json:"primaryKeyword"`
	Media           struct {
		Type  string `json:"type"`
		Image Image  `json:"image"`
	} `json:"media"`
	Preview string `json:"preview"`
}

type RecapItem struct {
	Type            string           `json:"type"`
	State           string           `json:"state"`
	Date            string           `json:"date"`
	ID              string           `json:"id"`
	Headline        string           `json:"headline"`
	Subhead         string           `json:"subhead"`
	SeoTitle        string           `json:"seoTitle"`
	SeoDescription  string           `json:"seoDescription"`
	SeoKeywords     string           `json:"seoKeywords"`
	Slug            string           `json:"slug"`
	Commenting      bool             `json:"commenting"`
	Tagline         string           `json:"tagline"`
	TokenData       map[string]Token `json:"tokenData"`
	Contributor     Contributor      `json:"contributor"`
	KeywordsDisplay []Keyword        `json:"keywordsDisplay"`
	KeywordsAll     []Keyword        `json:"keywordsAll"`
	Approval        string           `json:"approval"`
	URL             string           `json:"url"`
	DataURI         string           `json:"dataURI"`
	PrimaryKeyword  Keyword          `json:"primaryKeyword"`
	Media           struct {
		Type  string `json:"type"`
		Image Image  `json:"image"`
	} `json:"media"`
	Preview string `json:"preview"`
}

type MediaURLS struct {
	HTTPCLOUDMOBILE   string `json:"HTTP_CLOUD_MOBILE"`
	HTTPCLOUDTABLET   string `json:"HTTP_CLOUD_TABLET"`
	HTTPCLOUDTABLET60 string `json:"HTTP_CLOUD_TABLET_60"`
	HTTPCLOUDWIRED    string `json:"HTTP_CLOUD_WIRED"`
	HTTPCLOUDWIRED60  string `json:"HTTP_CLOUD_WIRED_60"`
	HTTPCLOUDWIREDWEB string `json:"HTTP_CLOUD_WIRED_WEB"`
	FLASH192K320X180  string `json:"FLASH_192K_320X180"`
	FLASH450K400X224  string `json:"FLASH_450K_400X224"`
	FLASH1200K640X360 string `json:"FLASH_1200K_640X360"`
	FLASH1800K960X540 string `json:"FLASH_1800K_960X540"`
}

type Token struct {
	TokenGUID string `json:"tokenGUID"`
	Type      string `json:"type"`
	VideoID   string `json:"videoId"`
	Href      string `json:"href"`
	Tags      []struct {
		Type        string `json:"@type"`
		Value       string `json:"@value"`
		DisplayName string `json:"@displayName"`
	} `json:"tags"`
	Date            string    `json:"date"`
	Headline        string    `json:"headline"`
	Duration        string    `json:"duration"`
	Blurb           string    `json:"blurb"`
	BigBlurb        string    `json:"bigBlurb"`
	MediaPlaybackID string    `json:"mediaPlaybackId"`
	Image           Image     `json:"image"`
	MediaURLS       MediaURLS `json:"mediaURLS"`

	// Token 2
	ID       string `json:"id"`
	TeamID   string `json:"teamId"`
	Position string `json:"position"`
	Name     string `json:"name"`
	SeoName  string `json:"seoName"`

	// Token 3
	HrefMobile string `json:"hrefMobile"`
}

type Contributor struct {
	Contributors []struct {
		Name    string `json:"name"`
		Twitter string `json:"twitter"`
	} `json:"contributors"`
	Source string `json:"source"`
}

type Epg struct {
	Title     string    `json:"title"`
	Platform  string    `json:"platform,omitempty"`
	Items     []EpgItem `json:"items"`
	TopicList string    `json:"topicList,omitempty"`
}

type EpgItem struct {
	GUID            string `json:"guid"`
	MediaState      string `json:"mediaState"`
	MediaPlaybackID string `json:"mediaPlaybackId"`
	MediaFeedType   string `json:"mediaFeedType"`
	CallLetters     string `json:"callLetters"`
	EventID         string `json:"eventId"`
	Language        string `json:"language"`
	FreeGame        bool   `json:"freeGame"`
	FeedName        string `json:"feedName"`
	GamePlus        bool   `json:"gamePlus"`
}

type MilestoneItem struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Type         string    `json:"type"`
	TimeAbsolute string    `json:"timeAbsolute"`
	TimeOffset   string    `json:"timeOffset"`
	Period       string    `json:"period"`
	StatsEventID string    `json:"statsEventId"`
	TeamID       string    `json:"teamId"`
	PlayerID     string    `json:"playerId"`
	PeriodTime   string    `json:"periodTime"`
	OrdinalNum   string    `json:"ordinalNum"`
	Highlight    Highlight `json:"highlight"`
}

type Highlight struct {
	Title     string      `json:"title"`
	TopicList string      `json:"topicList"`
	Items     []MediaItem `json:"items"`
}

type MediaItem struct {
	Type            string     `json:"type"`
	ID              string     `json:"id"`
	Date            string     `json:"date"`
	Title           string     `json:"title"`
	Blurb           string     `json:"blurb"`
	Description     string     `json:"description"`
	Duration        string     `json:"duration"`
	AuthFlow        bool       `json:"authFlow"`
	MediaPlaybackID string     `json:"mediaPlaybackId"`
	MediaState      string     `json:"mediaState"`
	Keywords        []Keyword  `json:"keywords"`
	Image           Image      `json:"image"`
	Playbacks       []PlayBack `json:"playbacks"`
}

type Keyword struct {
	Type        string `json:"type"`
	Value       string `json:"value"`
	DisplayName string `json:"displayName"`
}

type Image struct {
	Title   string `json:"title"`
	AltText string `json:"altText"`
	Cuts    struct {
		One136X640  ImageMetaData `json:"1136x640"`
		One024X576  ImageMetaData `json:"1024x576"`
		Nine60X540  ImageMetaData `json:"960x540"`
		Seven68X432 ImageMetaData `json:"768x432"`
		Six40X360   ImageMetaData `json:"640x360"`
		Five68X320  ImageMetaData `json:"568x320"`
		Three72X210 ImageMetaData `json:"372x210"`
		Three20X180 ImageMetaData `json:"320x180"`
		Two48X140   ImageMetaData `json:"248x140"`
		One24X70    ImageMetaData `json:"124x70"`
		// More Cuts from media
		Two568X1444 ImageMetaData `json:"2568x1444"`
		Two208X1242 ImageMetaData `json:"2208x1242"`
		Two048X1152 ImageMetaData `json:"2048x1152"`
		One704X960  ImageMetaData `json:"1704x960"`
		One536X864  ImageMetaData `json:"1536x864"`
		One284X722  ImageMetaData `json:"1284x722"`
	} `json:"cuts"`
}

type ImageMetaData struct {
	AspectRatio string `json:"aspectRatio"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Src         string `json:"src"`
	At2X        string `json:"at2x"`
	At3X        string `json:"at3x"`
}

type PlayBack struct {
	Name   string `json:"name"`
	Width  string `json:"width"`
	Height string `json:"height"`
	URL    string `json:"url"`
}
