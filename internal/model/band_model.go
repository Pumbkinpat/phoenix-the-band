package model

type BandData struct {
	Band              Band              `json:"band"`
	Hero              Hero              `json:"hero"`
	Intro             Intro             `json:"intro"`
	Members           []Member          `json:"members"`
	Partners          Partners          `json:"partners"`
	Events            Events            `json:"events"`
	PerformanceVideos PerformanceVideos `json:"performance_videos"`
	Navigation        []NavItem         `json:"navigation"`
	Footer            Footer            `json:"footer"`
}

type Band struct {
	Name          string       `json:"name"`
	Phone         string       `json:"phone"`
	ContactPerson string       `json:"contact_person"`
	MemberCount   int          `json:"member_count"`
	Genres        []string     `json:"genres"`
	SocialLinks   []SocialLink `json:"social_links"`
}

type SocialLink struct {
	Platform      string `json:"platform"`
	SocialLinkURL string `json:"social_link_url"`
	Logo          string `json:"logo"`
}

type Hero struct {
	Title           string `json:"title"`
	BackgroundImage string `json:"background_image"`
}

type Intro struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs"`
	Closing    string   `json:"closing"`
	Images     []string `json:"images"`
}

type Member struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Role          string         `json:"role"`
	Image         string         `json:"image"`
	Education     Education      `json:"education,omitempty"`
	Description   string         `json:"description"`
	TVShows       []string       `json:"tv_shows,omitempty"`
	YoutubeVideos []YoutubeVideo `json:"youtube_videos,omitempty"`
	SpecialSkill  string         `json:"special_skill,omitempty"`
	TotalSlides   int            `json:"total_slides"`
}

type Education struct {
	School string   `json:"school"`
	Majors []string `json:"majors,omitempty"`
	Genres []string `json:"genres,omitempty"`
}

type YoutubeVideo struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Partners struct {
	Title    string            `json:"title"`
	Logos    []string          `json:"logos"`
	Featured []FeaturedPartner `json:"featured"`
}

type FeaturedPartner struct {
	Name            string `json:"name"`
	BackgroundImage string `json:"background_image"`
}

type Events struct {
	Title             string          `json:"title"`
	FeaturedEvents    []FeaturedEvent `json:"featured_events"`
	OtherEventsTitle  string          `json:"other_events_title"`
	OtherEventsImages []string        `json:"other_events_images"`
}

type FeaturedEvent struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type PerformanceVideos struct {
	Title  string             `json:"title"`
	Videos []PerformanceVideo `json:"youtube_videos"`
}

type PerformanceVideo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	EmbedURL string `json:"embed_url"`
}

type NavItem struct {
	Label  string `json:"label"`
	Anchor string `json:"anchor"`
}

type Footer struct {
	Copyright string `json:"copyright"`
}
