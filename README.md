# Flight scraper

![Build](https://github.com/antodippo/flight-scraper/workflows/Build/badge.svg?branch=master)
![Release](https://github.com/antodippo/flight-scraper/workflows/Release/badge.svg)
[![Go Report](https://goreportcard.com/badge/github.com/antodippo/flight-scraper)](https://goreportcard.com/report/github.com/antodippo/flight-scraper)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

It's a simple flight prices scraper written in Go. Given a route and a date, it's gonna query [Kayak.com](https://www.kayak.com) website, parse the page, and send the information to the provided email recipient.

The simplest way to use it is to run:

```bash
$ ./flight-scraper -departure=FCO -arrival=AMS -date=2020-06-15 -recipient=test@test.com
```

where: 

- `departure` and `arrival` are [IATA airport codes](https://en.wikipedia.org/wiki/IATA_airport_code)
- `date` is in YYYY-MM-DD format
- `recipient` is the email recipient

At the moment you also need to set up an email SMTP server, adding the environment variables in the `.env` file. 

Use the `.env.dist` file as a reference:

```bash
SMTP_FROM="Flights scraper <noreply@antodippo.com>"
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=587
SMTP_USER=myusername
SMTP_PWD=mypassword
```

This is similar to what you get as a result in your email:

| Route     | Time          | Airline  | Price |
| --------- | ------------- | -------- | ----- |
| FCO - AMS | 15:30 - 18:00 | easyJet  | 76€   |
| FCO - AMS | 10:30 - 13:05 | LEVEL    | 81€   |
| FCO - AMS | 19:35 - 22:10 | LEVEL    | 97€   |
| FCO - AMS | 8:35 - 11:05  | Alitalia | 109€  |
| FCO - AMS | 14:05 - 16:35 | Alitalia | 109€  |
| FCO - AMS | 6:30 - 8:50   | KLM      | 143€  |
| FCO - AMS | 19:40 - 22:05 | KLM      | 143€  |
| FCO - AMS | 12:40 - 15:10 | KLM      | 143€  |
| FCO - AMS | 10:20 - 12:55 | KLM      | 143€  |
| FCO - AMS | 17:25 - 20:00 | KLM      | 143€  |

A possible use is to run it as a cron job, so you get flight prices in your inbox every chosen amount of time.

This was build just for learning and experimentation purposes, before using it please read Kayak.com [terms and conditions](https://www.kayak.com/terms-of-use).