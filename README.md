你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Ebiten

[![Build Status](https://travis-ci.org/hajimehoshi/ebiten.svg?branch=master)](https://travis-ci.org/hajimehoshi/ebiten)
[![Go Report Card](https://goreportcard.com/badge/github.com/hajimehoshi/ebiten)](https://goreportcard.com/report/github.com/hajimehoshi/ebiten)

**A dead simple 2D game library in Go**

Ebiten is an open-source game library, with which you can develop 2D games with simple API for multi platforms in the Go programming language.

* [Website (ebiten.org)](https://ebiten.org)
* [API Reference](https://pkg.go.dev/github.com/hajimehoshi/ebiten)
* [Cheat Sheet](https://ebiten.org/documents/cheatsheet.html)

![Overview](https://ebiten.org/images/overview.png)

## Platforms

* Windows (No Cgo!)
* macOS
* Linux
* FreeBSD
* Android
* iOS
* Web browsers (Chrome, Firefox, Safari and Edge)
  * [GopherJS](https://github.com/hajimehoshi/ebiten/wiki/GopherJS)
  * [WebAssembly](https://github.com/hajimehoshi/ebiten/wiki/WebAssembly) (Experimental)

Note: Gamepad and keyboard are not available on Android/iOS.

For installation on desktops, see [the installation instruction](https://ebiten.org/documents/install.html).

## Features

* 2D Graphics (Geometry/Color matrix transformation, Various composition modes, Offscreen rendering, Fullscreen, Text rendering, Automatic batches, Automatic texture atlas)
* Input (Mouse, Keyboard, Gamepads, Touches)
* Audio (Ogg/Vorbis, MP3, WAV, PCM)

## Packages

* [ebiten](https://pkg.go.dev/github.com/hajimehoshi/ebiten)
  * [audio](https://pkg.go.dev/github.com/hajimehoshi/ebiten/audio)
    * [mp3](https://pkg.go.dev/github.com/hajimehoshi/ebiten/audio/mp3)
    * [vorbis](https://pkg.go.dev/github.com/hajimehoshi/ebiten/audio/vorbis)
    * [wav](https://pkg.go.dev/github.com/hajimehoshi/ebiten/audio/wav)
  * [ebitenutil](https://pkg.go.dev/github.com/hajimehoshi/ebiten/ebitenutil)
  * [inpututil](https://pkg.go.dev/github.com/hajimehoshi/ebiten/inpututil)
  * [mobile](https://pkg.go.dev/github.com/hajimehoshi/ebiten/mobile)
  * [text](https://pkg.go.dev/github.com/hajimehoshi/ebiten/text)

## Community

### Slack

`#ebiten` channel in [Gophers Slack](https://blog.gopheracademy.com/gophers-slack-community/)

## License

Ebiten is licensed under Apache license version 2.0. See LICENSE file.
