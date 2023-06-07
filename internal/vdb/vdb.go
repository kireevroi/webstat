/*
	Простенькая виртуальная база данных на хэшмэпе.
*/

package vdb

import (
	"fmt"
	"log"
	"time"

	"github.com/kireevroi/webstat/internal/status"
	"github.com/kireevroi/webstat/internal/urlcache"
)

type DataBase struct {
	m   map[string]time.Duration
	max string
	min string
}

// RunVDB запускает тикер на заданное время t и парсит заданный файл в path
func (d *DataBase) RunVDB(path string, t time.Duration) {
	log.Println("Starting virtual database")

	ticker := time.NewTicker(t)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		var ul []string
		var err error
		for {
			ul, err = urlcache.ReadFile(path)
			if err != nil {
				log.Printf("error opening file: %v", err)
				continue
			}
			break
		}
		d.m, d.max, d.min = status.GetTime(ul)
		// for k, v := range d.m { // опциональная печать полученных данных
		// 	log.Printf("%v %v", k, v)
		// }
	}
}

func (d *DataBase) SearchTime(url string) string {
	return status.SearchTime(d.m, url)
}

func (d *DataBase) MinTime() string {
	return fmt.Sprint(d.min)
}

func (d *DataBase) MaxTime() string {
	return fmt.Sprint(d.max)
}
