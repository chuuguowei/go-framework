package bootstrap

import (
	"embed"
	"github.com/gin-gonic/gin"
	"log"
	"mime"
	"net/http"

	w "github.com/jhunters/goassist/web"
)

const (
	DefaultPrefixPath = "/___"
)

type WebModule struct {
	prefixPath string
	webDir     string
	htmlFs     embed.FS
	staticFs   embed.FS
}

func NewWebModule(prefixPath string, webroot string, htmlFs, staticFs embed.FS) (*WebModule, error) {
	mime.AddExtensionType(".ttf", "font/ttf")
	mime.AddExtensionType(".woff", "font/woff")

	if len(prefixPath) == 0 {
		prefixPath = DefaultPrefixPath
	}

	return &WebModule{prefixPath, webroot, htmlFs, staticFs}, nil
}

func (wm *WebModule) ServStaticFiles(router *gin.Engine) {
	// static files
	// web could use WebDir to develop by embed mode or direct mode(files modify aware on running)
	webDir := w.WebDir{Prefix: "./web/static", EmbedPrefix: "./web/static", Content: wm.staticFs, Embbed: true} // embed mode files modify not aware on running
	router.StaticFS(wm.getPath("static"), webDir)
	router.HEAD("/images/favicon.ico", func(c *gin.Context) {
		favicon(c, webDir)
	})
	router.GET("/images/favicon.ico", func(c *gin.Context) {
		favicon(c, webDir)
	})
}

func (wm *WebModule) ServHtmlFiles(router *gin.Engine) {

	delimitLeft := "${{"
	delimitRight := "}}"

	// web template files
	webTemplate := w.TemplateFS{Content: wm.htmlFs, Embbed: true, DelimsLeft: delimitLeft, DelimsRigth: delimitRight}
	temp, filenames, err := webTemplate.Parse("./", "*/html/*.html")

	if err != nil {
		log.Fatalf("servHtmlFiles failed: %s\n", err)
		return
	}
	if filenames == nil {
		router.SetHTMLTemplate(temp)
	} else {
		router.Delims(delimitLeft, delimitRight)
		router.LoadHTMLFiles(filenames...)
	}

	// index page
	router.GET(wm.getPath("/"), func(c *gin.Context) {
		// process index page
		c.Header("Content-Type", "text/html")
		c.HTML(
			http.StatusOK, "index.html", map[string]interface{}{
				"Prefix": wm.prefixPath,
			},
		)
	})

	// new index page
	router.GET(wm.getPath("/new"), func(c *gin.Context) {
		// process index page
		c.Header("Content-Type", "text/html")
		c.HTML(
			http.StatusOK, "index_new.html", map[string]interface{}{
				"Prefix": wm.prefixPath,
			},
		)
	})
	// demo page
	router.GET(wm.getPath("/demo"), func(c *gin.Context) {
		// process index page
		c.Header("Content-Type", "text/html")
		c.HTML(
			http.StatusOK, "demo.html", map[string]interface{}{
				"Prefix": wm.prefixPath,
			},
		)

	})
}

func (wm *WebModule) getPath(path string) string {
	return wm.prefixPath + path
}

// favicon
func favicon(c *gin.Context, fs http.FileSystem) {
	c.FileFromFS("favicon.ico", fs)
}

func (wm *WebModule) Close() {
}
