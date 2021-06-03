package data

import "encoding/xml"

type Action struct {
	Type   ActionType `xml:"type,attr,omitempty"`
	Ttmurl string     `xml:"ttm-url,attr,omitempty"`
}

// May be one of link, none
type ActionType string

// May be one of start, center, end
type AnchorPosition string

type Anon1 struct {
	Area   []string `xml:"area,omitempty"`
	Width  float64  `xml:"width,attr,omitempty"`
	Height float64  `xml:"height,attr,omitempty"`
	Url    string   `xml:"url,attr,omitempty"`
}

type Anon2 struct {
	Action         string     `xml:"action,omitempty"`
	Beaconmeta     string     `xml:"beacon-meta,omitempty"`
	Sourceid       string     `xml:"source-id,attr,omitempty"`
	Source         SourceType `xml:"source,attr,omitempty"`
	Targetanchorid int        `xml:"target-anchor-id,attr,omitempty"`
	Label          string     `xml:"label,attr,omitempty"`
	Rect           RectType   `xml:"rect,attr,omitempty"`
}

type Area struct {
	Action         string     `xml:"action,omitempty"`
	Beaconmeta     string     `xml:"beacon-meta,omitempty"`
	Sourceid       string     `xml:"source-id,attr,omitempty"`
	Source         SourceType `xml:"source,attr,omitempty"`
	Targetanchorid int        `xml:"target-anchor-id,attr,omitempty"`
	Label          string     `xml:"label,attr,omitempty"`
	Rect           RectType   `xml:"rect,attr,omitempty"`
}

// Must match the pattern [a-z_]+
type AssetName string

type Badge struct {
	Url              string           `xml:"url,attr,omitempty"`
	Width            float64          `xml:"width,attr,omitempty"`
	Height           float64          `xml:"height,attr,omitempty"`
	Hpos             float64          `xml:"h-pos,attr,omitempty"`
	Vpos             float64          `xml:"v-pos,attr,omitempty"`
	X                AnchorPosition   `xml:"x,attr,omitempty"`
	Y                AnchorPosition   `xml:"y,attr,omitempty"`
	Badgeapplication BadgeApplication `xml:"badge-application,attr,omitempty"`
}

func (t *Badge) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Badge
	var overlay struct {
		*T
		X *AnchorPosition `xml:"x,attr,omitempty"`
		Y *AnchorPosition `xml:"y,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.X = (*AnchorPosition)(&overlay.T.X)
	overlay.Y = (*AnchorPosition)(&overlay.T.Y)
	return d.DecodeElement(&overlay, &start)
}

// May be one of image, text
type BadgeApplication string

type Beaconmeta struct {
	Sku    string `xml:"sku,attr,omitempty"`
	Itemid string `xml:"item-id,attr,omitempty"`
}

type Body struct {
	Linearlayout Linearlayout `xml:"linear-layout,omitempty"`
	// Image             Image             `xml:"image,omitempty"`
	// Flexbox           Flexbox           `xml:"flexbox,omitempty"`
	// Itematom          Itematom          `xml:"item-atom,omitempty"`
	// Collapsiblelayout Collapsiblelayout `xml:"collapsible-layout,omitempty"`
}

// May be one of solid, dotted, dashed
type BorderStyleType string

type Category struct {
	Name     NonEmptyString `xml:"name,attr,omitempty"`
	Anchorid int            `xml:"anchor-id,attr,omitempty"`
}

// May be one of replace, expose
type CollapseType string

type Collapsiblelayout struct {
	Image             string         `xml:"image,omitempty"`
	Linearlayout      string         `xml:"linear-layout,omitempty"`
	Flexbox           string         `xml:"flexbox,omitempty"`
	Itematom          string         `xml:"item-atom,omitempty"`
	Collapsiblelayout string         `xml:"collapsible-layout,omitempty"`
	Width             DimensionType  `xml:"width,attr,omitempty"`
	Height            DimensionType  `xml:"height,attr,omitempty"`
	Aspectratio       float64        `xml:"aspect-ratio,attr,omitempty"`
	Flexgrow          int            `xml:"flex-grow,attr,omitempty"`
	Anchorid          int            `xml:"anchor-id,attr,omitempty"`
	Title             NonEmptyString `xml:"title,attr,omitempty"`
	Exposedtitle      string         `xml:"exposed-title,attr,omitempty"`
	Icon              AssetName      `xml:"icon,attr,omitempty"`
	Iconurl           string         `xml:"icon-url,attr,omitempty"`
	Collapsetype      CollapseType   `xml:"collapse-type,attr,omitempty"`
}

func (t *Collapsiblelayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Collapsiblelayout
	var overlay struct {
		*T
		Aspectratio  *float64      `xml:"aspect-ratio,attr,omitempty"`
		Collapsetype *CollapseType `xml:"collapse-type,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Aspectratio = (*float64)(&overlay.T.Aspectratio)
	overlay.Collapsetype = (*CollapseType)(&overlay.T.Collapsetype)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern #[\dA-Fa-f]{6}
type Color string

// May be one of match-parent, wrap-content, preserve-aspect
type DimensionByString string

type DimensionType string

type Flexbox struct {
	Image             string          `xml:"image,omitempty"`
	Linearlayout      string          `xml:"linear-layout,omitempty"`
	Flexbox           string          `xml:"flexbox,omitempty"`
	Itematom          string          `xml:"item-atom,omitempty"`
	Collapsiblelayout string          `xml:"collapsible-layout,omitempty"`
	Badge             string          `xml:"badge,omitempty"`
	Width             DimensionType   `xml:"width,attr,omitempty"`
	Height            DimensionType   `xml:"height,attr,omitempty"`
	Aspectratio       float64         `xml:"aspect-ratio,attr,omitempty"`
	Flexgrow          int             `xml:"flex-grow,attr,omitempty"`
	Anchorid          int             `xml:"anchor-id,attr,omitempty"`
	Orientation       OrientationType `xml:"orientation,attr,omitempty"`
	Styleid           string          `xml:"style-id,attr,omitempty"`
}

func (t *Flexbox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Flexbox
	var overlay struct {
		*T
		Aspectratio *float64 `xml:"aspect-ratio,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Aspectratio = (*float64)(&overlay.T.Aspectratio)
	return d.DecodeElement(&overlay, &start)
}

type Flyersource struct {
	Area        []Area         `xml:"area,omitempty"`
	Width       float64        `xml:"width,attr,omitempty"`
	Height      float64        `xml:"height,attr,omitempty"`
	Rect        RectType       `xml:"rect,attr,omitempty"`
	Path        string         `xml:"path,attr,omitempty"`
	Resolutions ResolutionType `xml:"resolutions,attr,omitempty"`
}

// May be one of bold, normal
type Fontweight string

type Head struct {
	Wayfinder Wayfinder `xml:"wayfinder,omitempty"`
	Styles    Styles    `xml:"styles,omitempty"`
}

type Image struct {
	Urlsource   Urlsource     `xml:"url-source,omitempty"`
	Flyersource Flyersource   `xml:"flyer-source,omitempty"`
	Badge       string        `xml:"badge,omitempty"`
	Width       DimensionType `xml:"width,attr,omitempty"`
	Height      DimensionType `xml:"height,attr,omitempty"`
	Aspectratio float64       `xml:"aspect-ratio,attr,omitempty"`
	Flexgrow    int           `xml:"flex-grow,attr,omitempty"`
	Anchorid    int           `xml:"anchor-id,attr,omitempty"`
	Styleid     string        `xml:"style-id,attr,omitempty"`
}

func (t *Image) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Image
	var overlay struct {
		*T
		Aspectratio *float64 `xml:"aspect-ratio,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Aspectratio = (*float64)(&overlay.T.Aspectratio)
	return d.DecodeElement(&overlay, &start)
}

// May be one of image-top, image-left, image-right
type ImagePositionType string

type Itematom struct {
	Action         string            `xml:"action,omitempty"`
	Beaconmeta     string            `xml:"beacon-meta,omitempty"`
	Badge          string            `xml:"badge,omitempty"`
	Text           string            `xml:"text,omitempty"`
	Width          DimensionType     `xml:"width,attr,omitempty"`
	Height         DimensionType     `xml:"height,attr,omitempty"`
	Aspectratio    float64           `xml:"aspect-ratio,attr,omitempty"`
	Flexgrow       int               `xml:"flex-grow,attr,omitempty"`
	Anchorid       int               `xml:"anchor-id,attr,omitempty"`
	Sourceid       string            `xml:"source-id,attr,omitempty"`
	Source         SourceType        `xml:"source,attr,omitempty"`
	Targetanchorid int               `xml:"target-anchor-id,attr,omitempty"`
	Styleid        string            `xml:"style-id,attr,omitempty"`
	Imageurl       string            `xml:"image-url,attr,omitempty"`
	Imagewidth     float64           `xml:"image-width,attr,omitempty"`
	Imageheight    float64           `xml:"image-height,attr,omitempty"`
	Imageposition  ImagePositionType `xml:"image-position,attr,omitempty"`
}

func (t *Itematom) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Itematom
	var overlay struct {
		*T
		Aspectratio   *float64           `xml:"aspect-ratio,attr,omitempty"`
		Imageposition *ImagePositionType `xml:"image-position,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Aspectratio = (*float64)(&overlay.T.Aspectratio)
	overlay.Imageposition = (*ImagePositionType)(&overlay.T.Imageposition)
	return d.DecodeElement(&overlay, &start)
}

// May be one of en, fr, es
type Language string

type Linearlayout struct {
	Image             []Image           `xml:"image,omitempty"`
	Linearlayout      string            `xml:"linear-layout,omitempty"`
	Flexbox           Flexbox           `xml:"flexbox,omitempty"`
	Itematom          Itematom          `xml:"item-atom,omitempty"`
	Collapsiblelayout Collapsiblelayout `xml:"collapsible-layout,omitempty"`
	Badge             Badge             `xml:"badge,omitempty"`
	Width             DimensionType     `xml:"width,attr,omitempty"`
	Height            DimensionType     `xml:"height,attr,omitempty"`
	Aspectratio       float64           `xml:"aspect-ratio,attr,omitempty"`
	Flexgrow          int               `xml:"flex-grow,attr,omitempty"`
	Anchorid          int               `xml:"anchor-id,attr,omitempty"`
	Orientation       OrientationType   `xml:"orientation,attr,omitempty"`
	Styleid           string            `xml:"style-id,attr,omitempty"`
}

func (t *Linearlayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Linearlayout
	var overlay struct {
		*T
		Aspectratio *float64 `xml:"aspect-ratio,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Aspectratio = (*float64)(&overlay.T.Aspectratio)
	return d.DecodeElement(&overlay, &start)
}

// Must be at least 1 items long
type NonEmptyString string

// May be one of horizontal, vertical
type OrientationType string

// Must match the pattern ([-+]?\d+(\.\d+){0,1} ){3}[-+]?\d+(\.\d+){0,1}
type RectType string

// Must match the pattern \d+(\.\d+){0,1}( \d+(\.\d+){0,1})*
type ResolutionType string

// May be one of tiled
type SourceFormatType string

// May be one of action, coupon, ecom, ecom-search, flyer, offer
type SourceType string

type Storefront struct {
	XMLName     xml.Name       `xml:"storefront"`
	XMLNS       string         `xml:"xmlns,attr"`
	Head        Head           `xml:"head"`
	Body        Body           `xml:"body"`
	Title       NonEmptyString `xml:"title,attr,omitempty"`
	Subtitle    string         `xml:"subtitle,attr,omitempty"`
	Uuid        NonEmptyString `xml:"uuid,attr,omitempty"`
	Revision    int            `xml:"revision,attr,omitempty"`
	Sfmlversion string         `xml:"sfml-version,attr,omitempty"`
	Runid       string         `xml:"run-id,attr,omitempty"`
	Segmentid   string         `xml:"segment-id,attr,omitempty"`
	Storeid     int            `xml:"store-id,attr,omitempty"`
	Languageid  Language       `xml:"language-id,attr,omitempty"`
}

type Style struct {
	Bordertopshow    bool            `xml:"border-top-show,attr,omitempty"`
	Borderbottomshow bool            `xml:"border-bottom-show,attr,omitempty"`
	Borderleftshow   bool            `xml:"border-left-show,attr,omitempty"`
	Borderrightshow  bool            `xml:"border-right-show,attr,omitempty"`
	Bordercolor      Color           `xml:"border-color,attr,omitempty"`
	Borderstyle      BorderStyleType `xml:"border-style,attr,omitempty"`
	Borderwidth      int             `xml:"border-width,attr,omitempty"`
	Padding          RectType        `xml:"padding,attr,omitempty"`
	Margin           RectType        `xml:"margin,attr,omitempty"`
	Backgroundurl    string          `xml:"background-url,attr,omitempty"`
	Backgroundcolor  Color           `xml:"background-color,attr,omitempty"`
	Fontfacesrc      string          `xml:"font-face-src,attr,omitempty"`
	Fontfamily       string          `xml:"font-family,attr,omitempty"`
	Fontweight       Fontweight      `xml:"font-weight,attr,omitempty"`
	Fontcolor        Color           `xml:"font-color,attr,omitempty"`
	Fontsize         int             `xml:"font-size,attr,omitempty"`
}

func (t *Style) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Style
	var overlay struct {
		*T
		Bordertopshow    *bool       `xml:"border-top-show,attr,omitempty"`
		Borderbottomshow *bool       `xml:"border-bottom-show,attr,omitempty"`
		Borderleftshow   *bool       `xml:"border-left-show,attr,omitempty"`
		Borderrightshow  *bool       `xml:"border-right-show,attr,omitempty"`
		Fontweight       *Fontweight `xml:"font-weight,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Bordertopshow = (*bool)(&overlay.T.Bordertopshow)
	overlay.Borderbottomshow = (*bool)(&overlay.T.Borderbottomshow)
	overlay.Borderleftshow = (*bool)(&overlay.T.Borderleftshow)
	overlay.Borderrightshow = (*bool)(&overlay.T.Borderrightshow)
	overlay.Fontweight = (*Fontweight)(&overlay.T.Fontweight)
	return d.DecodeElement(&overlay, &start)
}

type Styles struct {
	Style string `xml:"style,omitempty"`
}

type Text struct {
	Styleid  string `xml:"style-id,attr,omitempty"`
	Value    string `xml:"value,attr,omitempty"`
	Maxlines int    `xml:"max-lines,attr,omitempty"`
}

func (t *Text) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Text
	var overlay struct {
		*T
		Maxlines *int `xml:"max-lines,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Maxlines = (*int)(&overlay.T.Maxlines)
	return d.DecodeElement(&overlay, &start)
}

type Urlsource struct {
	Area   []Area  `xml:"area,omitempty"`
	Width  float64 `xml:"width,attr,omitempty"`
	Height float64 `xml:"height,attr,omitempty"`
	Url    string  `xml:"url,attr,omitempty"`
}

type Wayfinder struct {
	Category []Category `xml:"category,omitempty"`
}
