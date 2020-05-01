// Copyright 2013 The Ebiten Authors
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

// Code generated by genkeys.go using 'go generate'. DO NOT EDIT.

// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package glfw

import (
	"github.com/hajimehoshi/ebiten/internal/driver"
	"github.com/hajimehoshi/ebiten/internal/glfw"
)

var glfwKeyToDriverKey = map[glfw.Key]driver.Key{
	glfw.Key0:            driver.Key0,
	glfw.Key1:            driver.Key1,
	glfw.Key2:            driver.Key2,
	glfw.Key3:            driver.Key3,
	glfw.Key4:            driver.Key4,
	glfw.Key5:            driver.Key5,
	glfw.Key6:            driver.Key6,
	glfw.Key7:            driver.Key7,
	glfw.Key8:            driver.Key8,
	glfw.Key9:            driver.Key9,
	glfw.KeyA:            driver.KeyA,
	glfw.KeyB:            driver.KeyB,
	glfw.KeyC:            driver.KeyC,
	glfw.KeyD:            driver.KeyD,
	glfw.KeyE:            driver.KeyE,
	glfw.KeyF:            driver.KeyF,
	glfw.KeyG:            driver.KeyG,
	glfw.KeyH:            driver.KeyH,
	glfw.KeyI:            driver.KeyI,
	glfw.KeyJ:            driver.KeyJ,
	glfw.KeyK:            driver.KeyK,
	glfw.KeyL:            driver.KeyL,
	glfw.KeyM:            driver.KeyM,
	glfw.KeyN:            driver.KeyN,
	glfw.KeyO:            driver.KeyO,
	glfw.KeyP:            driver.KeyP,
	glfw.KeyQ:            driver.KeyQ,
	glfw.KeyR:            driver.KeyR,
	glfw.KeyS:            driver.KeyS,
	glfw.KeyT:            driver.KeyT,
	glfw.KeyU:            driver.KeyU,
	glfw.KeyV:            driver.KeyV,
	glfw.KeyW:            driver.KeyW,
	glfw.KeyX:            driver.KeyX,
	glfw.KeyY:            driver.KeyY,
	glfw.KeyZ:            driver.KeyZ,
	glfw.KeyApostrophe:   driver.KeyApostrophe,
	glfw.KeyBackslash:    driver.KeyBackslash,
	glfw.KeyBackspace:    driver.KeyBackspace,
	glfw.KeyCapsLock:     driver.KeyCapsLock,
	glfw.KeyComma:        driver.KeyComma,
	glfw.KeyDelete:       driver.KeyDelete,
	glfw.KeyDown:         driver.KeyDown,
	glfw.KeyEnd:          driver.KeyEnd,
	glfw.KeyEnter:        driver.KeyEnter,
	glfw.KeyEqual:        driver.KeyEqual,
	glfw.KeyEscape:       driver.KeyEscape,
	glfw.KeyF1:           driver.KeyF1,
	glfw.KeyF2:           driver.KeyF2,
	glfw.KeyF3:           driver.KeyF3,
	glfw.KeyF4:           driver.KeyF4,
	glfw.KeyF5:           driver.KeyF5,
	glfw.KeyF6:           driver.KeyF6,
	glfw.KeyF7:           driver.KeyF7,
	glfw.KeyF8:           driver.KeyF8,
	glfw.KeyF9:           driver.KeyF9,
	glfw.KeyF10:          driver.KeyF10,
	glfw.KeyF11:          driver.KeyF11,
	glfw.KeyF12:          driver.KeyF12,
	glfw.KeyGraveAccent:  driver.KeyGraveAccent,
	glfw.KeyHome:         driver.KeyHome,
	glfw.KeyInsert:       driver.KeyInsert,
	glfw.KeyKP0:          driver.KeyKP0,
	glfw.KeyKP1:          driver.KeyKP1,
	glfw.KeyKP2:          driver.KeyKP2,
	glfw.KeyKP3:          driver.KeyKP3,
	glfw.KeyKP4:          driver.KeyKP4,
	glfw.KeyKP5:          driver.KeyKP5,
	glfw.KeyKP6:          driver.KeyKP6,
	glfw.KeyKP7:          driver.KeyKP7,
	glfw.KeyKP8:          driver.KeyKP8,
	glfw.KeyKP9:          driver.KeyKP9,
	glfw.KeyKPAdd:        driver.KeyKPAdd,
	glfw.KeyKPDecimal:    driver.KeyKPDecimal,
	glfw.KeyKPDivide:     driver.KeyKPDivide,
	glfw.KeyKPEnter:      driver.KeyKPEnter,
	glfw.KeyKPEqual:      driver.KeyKPEqual,
	glfw.KeyKPMultiply:   driver.KeyKPMultiply,
	glfw.KeyKPSubtract:   driver.KeyKPSubtract,
	glfw.KeyLeft:         driver.KeyLeft,
	glfw.KeyLeftAlt:      driver.KeyLeftAlt,
	glfw.KeyLeftBracket:  driver.KeyLeftBracket,
	glfw.KeyLeftControl:  driver.KeyLeftControl,
	glfw.KeyLeftShift:    driver.KeyLeftShift,
	glfw.KeyMenu:         driver.KeyMenu,
	glfw.KeyMinus:        driver.KeyMinus,
	glfw.KeyNumLock:      driver.KeyNumLock,
	glfw.KeyPageDown:     driver.KeyPageDown,
	glfw.KeyPageUp:       driver.KeyPageUp,
	glfw.KeyPause:        driver.KeyPause,
	glfw.KeyPeriod:       driver.KeyPeriod,
	glfw.KeyPrintScreen:  driver.KeyPrintScreen,
	glfw.KeyRight:        driver.KeyRight,
	glfw.KeyRightAlt:     driver.KeyRightAlt,
	glfw.KeyRightBracket: driver.KeyRightBracket,
	glfw.KeyRightControl: driver.KeyRightControl,
	glfw.KeyRightShift:   driver.KeyRightShift,
	glfw.KeyScrollLock:   driver.KeyScrollLock,
	glfw.KeySemicolon:    driver.KeySemicolon,
	glfw.KeySlash:        driver.KeySlash,
	glfw.KeySpace:        driver.KeySpace,
	glfw.KeyTab:          driver.KeyTab,
	glfw.KeyUp:           driver.KeyUp,
}

var driverKeyToGLFWKey = map[driver.Key]glfw.Key{
	driver.Key0:            glfw.Key0,
	driver.Key1:            glfw.Key1,
	driver.Key2:            glfw.Key2,
	driver.Key3:            glfw.Key3,
	driver.Key4:            glfw.Key4,
	driver.Key5:            glfw.Key5,
	driver.Key6:            glfw.Key6,
	driver.Key7:            glfw.Key7,
	driver.Key8:            glfw.Key8,
	driver.Key9:            glfw.Key9,
	driver.KeyA:            glfw.KeyA,
	driver.KeyB:            glfw.KeyB,
	driver.KeyC:            glfw.KeyC,
	driver.KeyD:            glfw.KeyD,
	driver.KeyE:            glfw.KeyE,
	driver.KeyF:            glfw.KeyF,
	driver.KeyG:            glfw.KeyG,
	driver.KeyH:            glfw.KeyH,
	driver.KeyI:            glfw.KeyI,
	driver.KeyJ:            glfw.KeyJ,
	driver.KeyK:            glfw.KeyK,
	driver.KeyL:            glfw.KeyL,
	driver.KeyM:            glfw.KeyM,
	driver.KeyN:            glfw.KeyN,
	driver.KeyO:            glfw.KeyO,
	driver.KeyP:            glfw.KeyP,
	driver.KeyQ:            glfw.KeyQ,
	driver.KeyR:            glfw.KeyR,
	driver.KeyS:            glfw.KeyS,
	driver.KeyT:            glfw.KeyT,
	driver.KeyU:            glfw.KeyU,
	driver.KeyV:            glfw.KeyV,
	driver.KeyW:            glfw.KeyW,
	driver.KeyX:            glfw.KeyX,
	driver.KeyY:            glfw.KeyY,
	driver.KeyZ:            glfw.KeyZ,
	driver.KeyApostrophe:   glfw.KeyApostrophe,
	driver.KeyBackslash:    glfw.KeyBackslash,
	driver.KeyBackspace:    glfw.KeyBackspace,
	driver.KeyCapsLock:     glfw.KeyCapsLock,
	driver.KeyComma:        glfw.KeyComma,
	driver.KeyDelete:       glfw.KeyDelete,
	driver.KeyDown:         glfw.KeyDown,
	driver.KeyEnd:          glfw.KeyEnd,
	driver.KeyEnter:        glfw.KeyEnter,
	driver.KeyEqual:        glfw.KeyEqual,
	driver.KeyEscape:       glfw.KeyEscape,
	driver.KeyF1:           glfw.KeyF1,
	driver.KeyF2:           glfw.KeyF2,
	driver.KeyF3:           glfw.KeyF3,
	driver.KeyF4:           glfw.KeyF4,
	driver.KeyF5:           glfw.KeyF5,
	driver.KeyF6:           glfw.KeyF6,
	driver.KeyF7:           glfw.KeyF7,
	driver.KeyF8:           glfw.KeyF8,
	driver.KeyF9:           glfw.KeyF9,
	driver.KeyF10:          glfw.KeyF10,
	driver.KeyF11:          glfw.KeyF11,
	driver.KeyF12:          glfw.KeyF12,
	driver.KeyGraveAccent:  glfw.KeyGraveAccent,
	driver.KeyHome:         glfw.KeyHome,
	driver.KeyInsert:       glfw.KeyInsert,
	driver.KeyKP0:          glfw.KeyKP0,
	driver.KeyKP1:          glfw.KeyKP1,
	driver.KeyKP2:          glfw.KeyKP2,
	driver.KeyKP3:          glfw.KeyKP3,
	driver.KeyKP4:          glfw.KeyKP4,
	driver.KeyKP5:          glfw.KeyKP5,
	driver.KeyKP6:          glfw.KeyKP6,
	driver.KeyKP7:          glfw.KeyKP7,
	driver.KeyKP8:          glfw.KeyKP8,
	driver.KeyKP9:          glfw.KeyKP9,
	driver.KeyKPAdd:        glfw.KeyKPAdd,
	driver.KeyKPDecimal:    glfw.KeyKPDecimal,
	driver.KeyKPDivide:     glfw.KeyKPDivide,
	driver.KeyKPEnter:      glfw.KeyKPEnter,
	driver.KeyKPEqual:      glfw.KeyKPEqual,
	driver.KeyKPMultiply:   glfw.KeyKPMultiply,
	driver.KeyKPSubtract:   glfw.KeyKPSubtract,
	driver.KeyLeft:         glfw.KeyLeft,
	driver.KeyLeftAlt:      glfw.KeyLeftAlt,
	driver.KeyLeftBracket:  glfw.KeyLeftBracket,
	driver.KeyLeftControl:  glfw.KeyLeftControl,
	driver.KeyLeftShift:    glfw.KeyLeftShift,
	driver.KeyMenu:         glfw.KeyMenu,
	driver.KeyMinus:        glfw.KeyMinus,
	driver.KeyNumLock:      glfw.KeyNumLock,
	driver.KeyPageDown:     glfw.KeyPageDown,
	driver.KeyPageUp:       glfw.KeyPageUp,
	driver.KeyPause:        glfw.KeyPause,
	driver.KeyPeriod:       glfw.KeyPeriod,
	driver.KeyPrintScreen:  glfw.KeyPrintScreen,
	driver.KeyRight:        glfw.KeyRight,
	driver.KeyRightAlt:     glfw.KeyRightAlt,
	driver.KeyRightBracket: glfw.KeyRightBracket,
	driver.KeyRightControl: glfw.KeyRightControl,
	driver.KeyRightShift:   glfw.KeyRightShift,
	driver.KeyScrollLock:   glfw.KeyScrollLock,
	driver.KeySemicolon:    glfw.KeySemicolon,
	driver.KeySlash:        glfw.KeySlash,
	driver.KeySpace:        glfw.KeySpace,
	driver.KeyTab:          glfw.KeyTab,
	driver.KeyUp:           glfw.KeyUp,
}
