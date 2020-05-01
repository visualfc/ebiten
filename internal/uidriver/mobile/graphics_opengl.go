// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

<<<<<<< HEAD
// +build android ios,386 ios,amd64 ios,ebitengl

package mobile
=======
<<<<<<< HEAD:internal/uidriver/glfw/graphics_opengl.go
// +build freebsd linux windows ebitengl
// +build !js

package glfw
=======
// +build android ios,386 ios,amd64 ios,ebitengl

package mobile
>>>>>>> v1.11.1:internal/uidriver/mobile/graphics_opengl.go
>>>>>>> v1.11.1

import (
	"github.com/hajimehoshi/ebiten/internal/driver"
	"github.com/hajimehoshi/ebiten/internal/graphicsdriver/opengl"
)

func (*UserInterface) Graphics() driver.Graphics {
	return opengl.Get()
}
