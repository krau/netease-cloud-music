// MIT License
//
// Copyright (c) 2024 chaunsin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package weapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chaunsin/netease-cloud-music/api/types"
)

type ArtistSongsReq struct {
	Id           int64  `json:"id"`            // 歌手id
	PrivateCloud string `json:"private_cloud"` // boolean
	WorkType     int64  `json:"work_type"`     // 通常为1
	Order        string `json:"order"`         // hot,time
	Offset       int64  `json:"offset"`        // 第几页
	Limit        int64  `json:"limit"`         // 每页条数
}

type ArtistSongsResp struct {
	types.RespCommon[any]
	More  bool                   `json:"more"`
	Total int64                  `json:"total"`
	Songs []ArtistSongsRespSongs `json:"songs"`
}

type ArtistSongsRespSongs struct {
	Id              int64          `json:"id"`
	A               interface{}    `json:"a"`
	Al              types.Album    `json:"al"`
	Alia            []string       `json:"alia"`
	Ar              []types.Artist `json:"ar"`
	Cd              string         `json:"cd"`
	Cf              string         `json:"cf"`
	Cp              int64          `json:"cp"`
	Crbt            interface{}    `json:"crbt"`
	DjId            int64          `json:"djId"`
	Dt              int64          `json:"dt"`
	Fee             int64          `json:"fee"`
	Ftype           int64          `json:"ftype"`
	H               *types.Quality `json:"h"`
	Hr              *types.Quality `json:"hr"`
	L               *types.Quality `json:"l"`
	M               *types.Quality `json:"m"`
	Sq              *types.Quality `json:"sq"`
	Mst             int64          `json:"mst"`
	Mv              int64          `json:"mv"`
	Name            string         `json:"name"`
	No              int64          `json:"no"`
	NoCopyrightRcmd interface{}    `json:"noCopyrightRcmd"`
	Pop             float64        `json:"pop"`
	Pst             int64          `json:"pst"`
	Rt              string         `json:"rt"`
	RtUrl           interface{}    `json:"rtUrl"`
	RtUrls          []interface{}  `json:"rtUrls"`
	Rtype           int64          `json:"rtype"`
	Rurl            interface{}    `json:"rurl"`
	SongJumpInfo    interface{}    `json:"songJumpInfo"`
	St              int64          `json:"st"`
	T               int64          `json:"t"`
	V               int64          `json:"v"`
	Tns             []string       `json:"tns,omitempty"`
	Privilege       struct {
		types.Privileges
		Code    int         `json:"code"`
		Message interface{} `json:"message"`
	} `json:"privilege"`
}

// ArtistSongs 歌手所有歌曲
// url:
// needLogin:
func (a *Api) ArtistSongs(ctx context.Context, req *ArtistSongsReq) (*ArtistSongsResp, error) {
	var (
		url   = "https://music.163.com/weapi/v1/artist/songs"
		reply ArtistSongsResp
	)
	if req.Order == "" {
		req.Order = "hot"
	}
	if req.Limit == 0 {
		req.Limit = 100
	}

	resp, err := a.client.Request(ctx, http.MethodPost, url, "weapi", req, &reply)
	if err != nil {
		return nil, fmt.Errorf("Request: %w", err)
	}
	_ = resp
	return &reply, nil
}
