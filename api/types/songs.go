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

package types

type Artist struct {
	Id    int64         `json:"id"`
	Name  string        `json:"name"`
	Tns   []interface{} `json:"tns"`
	Alias []interface{} `json:"alias"`
}

type Album struct {
	Id     int64         `json:"id"`
	Name   string        `json:"name"`
	PicUrl string        `json:"picUrl"`
	Tns    []interface{} `json:"tns"`
	PicStr string        `json:"pic_str"`
	Pic    int64         `json:"pic"`
}

type ChargeInfo struct {
	Rate          int64       `json:"rate"`
	ChargeUrl     interface{} `json:"chargeUrl"`
	ChargeMessage interface{} `json:"chargeMessage"`
	ChargeType    int64       `json:"chargeType"`
}

type FreeTrialPrivilege struct {
	CannotListenReason interface{} `json:"cannotListenReason"`
	ListenType         interface{} `json:"listenType"`
	PlayReason         interface{} `json:"playReason"`
	ResConsumable      bool        `json:"resConsumable"`
	UserConsumable     bool        `json:"userConsumable"`
}

type FreeTimeTrialPrivilege struct {
	RemainTime     int64 `json:"remainTime"`
	ResConsumable  bool  `json:"resConsumable"`
	Type           int64 `json:"type"`
	UserConsumable bool  `json:"userConsumable"`
}

type Privileges struct {
	Id    int64 `json:"id"`
	Fee   int64 `json:"fee"`
	Payed int64 `json:"payed"`
	// St 小于0时为灰色歌曲, 使用上传云盘的方法解灰后 st == 0
	St   int64 `json:"st"`
	Pl   int64 `json:"pl"`
	Dl   int64 `json:"dl"`
	Sp   int64 `json:"sp"`
	Cp   int64 `json:"cp"`
	Subp int64 `json:"subp"`
	// Cs 是否为云盘歌曲
	Cs    bool  `json:"cs"`
	Maxbr int64 `json:"maxbr"`
	Fl    int64 `json:"fl"`
	// Toast 是否「由于版权保护，您所在的地区暂时无法使用。」
	Toast         bool  `json:"toast"`
	Flag          int64 `json:"flag"`
	PreSell       bool  `json:"preSell"`
	PlayMaxbr     int64 `json:"playMaxbr"`
	DownloadMaxbr int   `json:"downloadMaxbr"`
	// MaxBrLevel 歌曲最高音质
	MaxBrLevel         string `json:"maxBrLevel"`
	PlayMaxBrLevel     string `json:"playMaxBrLevel"`
	DownloadMaxBrLevel string `json:"downloadMaxBrLevel"`
	// PlLevel 当前用户的该歌曲最高试听音质
	PlLevel string `json:"plLevel"`
	// DlLevel 当前用户的该歌曲最高下载音质
	DlLevel string `json:"dlLevel"`
	// FlLevel 免费用户的该歌曲播放音质
	FlLevel            string             `json:"flLevel"`
	Rscl               interface{}        `json:"rscl"`
	FreeTrialPrivilege FreeTrialPrivilege `json:"freeTrialPrivilege"`
	RightSource        int64              `json:"rightSource"`
	ChargeInfoList     []ChargeInfo       `json:"chargeInfoList"`
}
