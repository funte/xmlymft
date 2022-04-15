package common

type AlbumInfo struct {
	PlayCount   int    `json:"playCount"`
	Cover       string `json:"coverPath"`
	Title       string `json:"title"`
	Intro       string `json:"intro"`
	Id          int    `json:"albumId"`
	Category    string `json:"categoryTitle"`
	CreatedTime int    `json:"careatedAt"`
	UpdatedTime int    `json:"updatedAt"`
	Author      string `json:"nickname"`
	TracksCount int    `json:"tracksCount"`
}

type SearchAlbumResult struct {
	PageNum   int         `json:"currentPage"`
	TotalPage int         `json:"totalPage"`
	Albums    []AlbumInfo `json:"docs"`
}

// Search album parameters, bind for query string and post data.
type SearchAlbumParams struct {
	Keyword  string `form:"kw"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}

type TrackInfo struct {
	Index    int    `json:"index"`
	Id       int    `json:"trackId"`
	Name     string `json:"trackName"`
	Duration int    `json:"duration"`
}

type QueryPlayListResult struct {
	Tracks  []TrackInfo `json:"tracksAudioPlay"`
	PageNum int         `json:"pageNum"`
}

// Query album play list paramters, bind for query string and post data.
type QueryPlayListParams struct {
	AlbumId  int `form:"id"`
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}

type QueryTrackAddressResult struct {
	Type     string
	ByteSize int
	Address  string
}
