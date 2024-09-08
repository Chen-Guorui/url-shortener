package model

import (
	"url-shortener/log"
)

type Url struct {
	Id          string `db:"id" json:"id"`
	ShortUrl    string `db:"short_url" json:"short_url"`
	OriginalUrl string `db:"original_url" json:"original_url"`
}

func (u *Url) IsEmpty() bool {
	return u.Id == "" || u.OriginalUrl == "" || u.ShortUrl == ""
}

func Insert(data *Url) error {
	_, err := db.Exec("INSERT INTO `url` (`id`, `short_url`, `original_url`) VALUES (?, ?, ?)", data.Id, data.ShortUrl, data.OriginalUrl)
	if err != nil {
		log.Logger.Printf("Insert data error: %s, data: %+v", err, data)
		return err
	}

	return nil
}

func FindOneByShortUrl(shortUrl string) (*Url, error) {
	rows, err := db.Query("SELECT * FROM `url` WHERE `short_url` = ? limit 1", shortUrl)
	if err != nil {
		log.Logger.Printf("Find by short url error: %s, url: %+v", err, shortUrl)
		return nil, err
	}
	defer rows.Close()

	url := Url{}
	for rows.Next() {
		var id, shortUrl, originalUrl string

		err = rows.Scan(&id, &shortUrl, &originalUrl)
		if err != nil {
			log.Logger.Println(err)
			return nil, err
		}
		url.Id = id
		url.ShortUrl = shortUrl
		url.OriginalUrl = originalUrl
	}

	return &url, nil
}

func FindOneByOriginalUrl(originalUrl string) (*Url, error) {
	rows, err := db.Query("SELECT * FROM `url` WHERE `original_url` = ? limit 1", originalUrl)
	if err != nil {
		log.Logger.Printf("Find by original url error: %s, url: %+v", err, originalUrl)
		return nil, err
	}
	defer rows.Close()

	url := Url{}
	for rows.Next() {
		var id, shortUrl, originalUrl string

		err = rows.Scan(&id, &shortUrl, &originalUrl)
		if err != nil {
			log.Logger.Println(err)
			return nil, err
		}
		url.Id = id
		url.ShortUrl = shortUrl
		url.OriginalUrl = originalUrl
	}

	return &url, nil
}
