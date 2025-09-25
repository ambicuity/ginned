// Copyright 2014 ambicuity Ritesh Rana. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unicode"
)

// BindKey indicates a default bind key.
const BindKey = "_gin-gonic/gin/bindkey"

// Bind is a helper function for given interface object and returns a Gin middleware.
func Bind(val any) HandlerFunc {
	value := reflect.ValueOf(val)
	if value.Kind() == reflect.Ptr {
		panic(`Bind struct can not be a pointer. Example:
	Use: gin.Bind(Struct{}) instead of gin.Bind(&Struct{})
`)
	}
	typ := value.Type()

	return func(c *Context) {
		obj := reflect.New(typ).Interface()
		if c.Bind(obj) == nil {
			c.Set(BindKey, obj)
		}
	}
}

// WrapF is a helper function for wrapping http.HandlerFunc and returns a Gin middleware.
func WrapF(f http.HandlerFunc) HandlerFunc {
	return func(c *Context) {
		f(c.Writer, c.Request)
	}
}

// WrapH is a helper function for wrapping http.Handler and returns a Gin middleware.
func WrapH(h http.Handler) HandlerFunc {
	return func(c *Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// H is a shortcut for map[string]any
type H map[string]any

// MarshalXML allows type H to be used with xml.Marshal.
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// assert1 panics with the given message if the condition is false.
func assert1(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

// filterFlags removes flags from content by finding the first space or semicolon.
func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func chooseData(custom, wildcard any) any {
	if custom != nil {
		return custom
	}
	if wildcard != nil {
		return wildcard
	}
	panic("chooseData: both custom and wildcard data are nil - negotiation config is invalid")
}

func parseAccept(acceptHeader string) []string {
	if acceptHeader == "" {
		return nil
	}
	
	parts := strings.Split(acceptHeader, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		if i := strings.IndexByte(part, ';'); i > 0 {
			part = part[:i]
		}
		if part = strings.TrimSpace(part); part != "" {
			out = append(out, part)
		}
	}
	return out
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("lastChar: string cannot be empty")
	}
	return str[len(str)-1]
}

func nameOfFunction(f any) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// joinPaths combines an absolute path with a relative path.
func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if len(relativePath) > 0 && relativePath[len(relativePath)-1] == '/' && 
	   len(finalPath) > 0 && finalPath[len(finalPath)-1] != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			debugPrint("Environment variable PORT=\"%s\"", port)
			return ":" + port
		}
		debugPrint("Environment variable PORT is undefined. Using port :8080 by default")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

// isASCII checks if a string contains only ASCII characters.
// Reference: https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// generateETag creates a simple ETag from file metadata for caching optimization
func generateETag(modTime time.Time, size int64) string {
	return fmt.Sprintf(`"%x-%x"`, modTime.Unix(), size)
}
