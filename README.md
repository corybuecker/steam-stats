# steam-stats-fetcher

A Go-based client to fetch a user's games from the Steam API and store the results in MongoDB.

Setting up the configuration
----------------------------

    use steam_stats_fetcher

    db.configuration.update({"id": "steam"}, {"steam_api_key": "<your steam api key>", id: "steam", steam_id: "<your steam id>"}, {"upsert": true})

Usage
-----

    steam-stats-fetcher --host <url of Mongo DB> steam

[![Build Status](https://travis-ci.org/corybuecker/steam-stats-fetcher.svg?branch=master)](https://travis-ci.org/corybuecker/steam-stats-fetcher)
