/*
	Обработчики для Gin
*/
package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kireevroi/webstat/internal/statistics"
	"github.com/kireevroi/webstat/internal/urlcache"
	"github.com/kireevroi/webstat/internal/vdb"
)

// ApiMiddleware отвечает за миддлвейр связанный с примитивной авторизацией админа
func ApiMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		api := c.Query("key")
		if api == key {
			c.Set("allowed", true)
		}
		c.Next()
	}
}

// WebsiteTimeHandler возвращает время доступа для сайта в запросе
func WebsiteTimeHandler(d *vdb.DataBase, sm *statistics.StatMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		website := c.Query("website")
		if _, ok := c.Get("allowed"); ok {
			c.String(http.StatusOK, "%v", sm.Get(website))
			return
		}
		sm.Set(website)
		u, err := urlcache.CleanURL(website)
		if err != nil {
			c.String(http.StatusBadRequest, "Bad URL Query")
			return
		}

		st := d.SearchTime(u)
		c.String(http.StatusOK, "%s", st)
	}
}

// MaxHandler возвращает максимальное время доступа
func MaxHandler(d *vdb.DataBase, s *statistics.Stats) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get("allowed"); ok {
			c.String(http.StatusOK, "%v", s.Get())
			return
		}
		s.Set()
		max := d.MaxTime()
		c.String(http.StatusOK, "%s", max)
	}
}

// MinHandler возвращает минимальное время доступа
func MinHandler(d *vdb.DataBase, s *statistics.Stats) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get("allowed"); ok {
			c.String(http.StatusOK, "%v", s.Get())
			return
		}
		s.Set()
		max := d.MinTime()
		c.String(http.StatusOK, "%s", max)
	}
}
