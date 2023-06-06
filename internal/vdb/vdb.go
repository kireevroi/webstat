/*
	Простенькая виртуальная база данных на хэшмэпе.
*/

package vdb

import (
	"log"
	"time"

	"github.com/kireevroi/webstat/internal/status"
	"github.com/kireevroi/webstat/internal/urlcache"
)

type DataBase struct {
	m map[string]time.Duration
}

// TODO: Refactor
func (d *DataBase) Init(path string) {
	log.Println("Starting virtual database")

	ul, err := urlcache.ReadFile(path)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	d.m = status.GetTime(ul)

	for k, v := range d.m {
		log.Printf("%v %v", k, v)
	}
	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			tmpul := ul
			for {
				tmpul, err = urlcache.ReadFile("list.txt")
				if err != nil {
					log.Printf("error opening file: %v", err)
					continue
				}
				break
			}
			d.m = status.GetTime(tmpul)
			for k, v := range d.m {
				log.Printf("%v %v", k, v)
			}
			ul = tmpul
		}
	}
}

func (d *DataBase) SearchTime(url string) string {
	return status.SearchTime(d.m, url)
}

func (d *DataBase) MinTime() string {
	return status.MinTime(d.m)
}

func (d *DataBase) MaxTime() string {
	return status.MaxTime(d.m)
}
