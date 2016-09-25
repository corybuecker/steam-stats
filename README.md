# steamfetcher

A Go-based client to fetch a user's games from the Steam API and store the results in MongoDB.

[![Build Status](https://travis-ci.org/corybuecker/steamfetcher.svg?branch=master)](https://travis-ci.org/corybuecker/steamfetcher)

Setting up the configuration
----------------------------

Get a Steam API key from http://steamcommunity.com/dev/registerkey. You can find your Steam ID at http://steamrep.com.

    use steamfetcher

    db.configuration.update({"id": "steam"}, {"steam_api_key": "<your steam api key>", id: "steam", steam_id: "<your steam id>"}, {"upsert": true})

Usage
-----

    steamfetcher --host <url of Mongo DB> steam
