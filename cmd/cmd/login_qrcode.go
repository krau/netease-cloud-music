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

package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/chaunsin/netease-cloud-music/api"
	"github.com/chaunsin/netease-cloud-music/api/weapi"
	"github.com/chaunsin/netease-cloud-music/pkg/log"

	"github.com/spf13/cobra"
)

type loginQrcodeCmd struct {
	root *Login
	cmd  *cobra.Command
	l    *log.Logger

	timeout time.Duration // 登录超时时间
	dir     string        // 二维码文件路径
}

func qrcode(root *Login, l *log.Logger) *cobra.Command {
	c := &loginQrcodeCmd{
		root: root,
		l:    l,
	}
	c.cmd = &cobra.Command{
		Use:     "qrcode",
		Short:   "user qrcode login",
		Example: "ncm login qrcode xxx",
		Run: func(cmd *cobra.Command, args []string) {
			if err := c.execute(); err != nil {
				cmd.Println(err)
			}
		},
	}
	c.addFlags()
	return c.cmd
}

func (c *loginQrcodeCmd) addFlags() {
	c.cmd.Flags().DurationVarP(&c.timeout, "timeout", "t", time.Minute*5, "1s 1m")
	c.cmd.Flags().StringVarP(&c.dir, "dir", "d", "./", "./")
}

func (c *loginQrcodeCmd) execute() error {
	cli, err := api.NewWithErr(c.root.root.Cfg.Network, c.l)
	if err != nil {
		return fmt.Errorf("NewWithErr: %w", err)
	}
	request := weapi.New(cli)

	ctx, cancel := context.WithTimeout(c.cmd.Context(), c.timeout)
	defer cancel()

	// 1. 生成key
	key, err := request.QrcodeCreateKey(ctx, &weapi.QrcodeCreateKeyReq{Type: 1})
	if err != nil {
		return fmt.Errorf("QrcodeCreateKey: %w", err)
	}
	if key.UniKey == "" {
		return fmt.Errorf("QrcodeCreateKey resp: %+v\n", key)
	}

	// 2. 生成二维码
	qr, err := request.QrcodeGenerate(ctx, &weapi.QrcodeGenerateReq{CodeKey: key.UniKey})
	if err != nil {
		return fmt.Errorf("QrcodeGenerate: %s", err)
	}

	// 3. 手机扫码
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	p := filepath.Join(dir, c.dir)
	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		return fmt.Errorf("MkdirAll: %w", err)
	}
	p = filepath.Join(p, "qrcode.png")
	if err := os.WriteFile(p, qr.Qrcode, os.ModePerm); err != nil {
		return err
	}
	fmt.Println(">>>>> please scan qrcode in your phone <<<<<")
	fmt.Printf("qrcode file %s\n", p)
	fmt.Printf("qrcode: %s\n", qr.QrcodePrint)

	// 4. 轮训获取扫码状态
	for {
		select {
		case <-ctx.Done():
			log.Warn("timeout retry")
			return ctx.Err()
		default:
		}

		time.Sleep(time.Second * 3)
		resp, err := request.QrcodeCheck(ctx, &weapi.QrcodeCheckReq{Type: 1, Key: key.UniKey})
		if err != nil {
			return fmt.Errorf("QrcodeCheck: %w", err)
		}
		switch resp.Code {
		case 800: // 二维码不存在或已过期
			return fmt.Errorf("current QrcodeCheck resp: %v\n", resp)
		case 801: // 等待扫码
			continue
		case 802: // 正在扫码授权中
			continue
		case 803: // 授权登录成功
			log.Info("current QrcodeCheck resp: %v\n", resp)
			goto ok
		default:
			log.Error("current QrcodeCheck resp: %v\n", resp)
		}
	}
ok:

	// 5. 查询登录信息是否成功
	user, err := request.GetUserInfo(ctx, &weapi.GetUserInfoReq{})
	if err != nil {
		return fmt.Errorf("GetUserInfo: %s", err)
	}
	log.Info("login success: %+v", user)
	return nil
}