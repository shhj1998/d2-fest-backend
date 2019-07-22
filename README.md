# RSS Search API

[![Build Status](https://travis-ci.org/shhj1998/rss-search-api.svg?branch=master)](https://travis-ci.org/shhj1998/rss-search-api)
[![Coverage Status](https://coveralls.io/repos/github/shhj1998/rss-search-api/badge.svg?branch=master)](https://coveralls.io/github/shhj1998/rss-search-api?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/shhj1998/rss-search-api)](https://goreportcard.com/report/github.com/shhj1998/rss-search-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![](https://godoc.org/github.com/shhj1998/rss-search-api/rsserver?status.svg)](https://godoc.org/github.com/shhj1998/rss-search-api/rsserver)

This repository is a server code for fetching and searching RSS data efficiently. There are also many useful Go packages to help you develop your own RSS project!

## Packages

- [rsserver](https://github.com/shhj1998/rss-search-api/tree/master/rsserver) : Our main package that provides many functionalities to act with our rss database. More information is available on our [documentation](https://godoc.org/github.com/shhj1998/rss-search-api/rsserver).
- [rssapi](https://github.com/shhj1998/rss-search-api/tree/master/rssapi) : A simple REST api server that uses rsserver package. Available apis are described below.

## API
### Channel API

### Item API

## Schema



## Performance

### Test DB Info

We used 8 rss services and 3594 rss items(news) to test our api performance. The details of the rss services are described below.

| Name                           | Link                                                | Rows |
| ------------------------------ | --------------------------------------------------- | ---- |
| 다음뉴스 - 사회Top RSS         | https://media.daum.net/syndication/society.rss      | 387  |
| 다음뉴스 - 시사 주요뉴스 RSS   | https://media.daum.net/syndication/entertain.rss    | 975  |
| 다음뉴스 - 스포츠 주요뉴스 RSS | https://media.daum.net/syndication/today_sports.rss | 810  |
| 다음뉴스 - 정치Top RSS         | https://media.daum.net/syndication/politics.rss     | 312  |
| 다음뉴스 - 경제Top RSS         | https://media.daum.net/syndication/economic.rss     | 431  |
| 다음뉴스 - 국제Top RSS         | https://media.daum.net/syndication/foreign.rss      | 302  |
| 다음뉴스 - 문화/생활Top RSS    | https://media.daum.net/syndication/culture.rss      | 202  |
| 다음뉴스 - Tech Top RSS        | https://media.daum.net/syndication/digital.rss      | 163  |

### Result

Detail of the result is described below. The code used to test performance is [here](https://github.com/shhj1998/rss-search-api/blob/master/rssapi/channel_test.go). It will not work in your local because it must be with a .env file that contains the test db information. We tested three api that were most likely to be used. The **GET /api/v1/channel/items/count/:count** apis show almost same performance although the count value changed.

| API                                | Time(ms)   |
| ---------------------------------- | ---------- |
| GET /api/v1/channel                | 3.9837295  |
| GET /api/v1/channel/items          | 29.8150581 |
| GET /api/v1/channel/items?id=3     | 6.1691366  |
| GET /api/v1/channel/items/count/1  | 27.7327405 |
| GET /api/v1/channel/items/count/2  | 27.3187845 |
| GET /api/v1/channel/items/count/3  | 28.1879731 |
| GET /api/v1/channel/items/count/4  | 34.7423070 |
| GET /api/v1/channel/items/count/5  | 28.3924582 |
| GET /api/v1/channel/items/count/6  | 27.4330836 |
| GET /api/v1/channel/items/count/7  | 27.9774116 |
| GET /api/v1/channel/items/count/8  | 27.0644527 |
| GET /api/v1/channel/items/count/9  | 33.9821699 |
| GET /api/v1/channel/items/count/10 | 34.5249172 |

![Imgur](https://i.imgur.com/ztFsvEJ.png)



## See also
