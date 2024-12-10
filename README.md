# RSS Aggregator

This project is a simple RSS aggregator built with Go. It fetches RSS feeds, parses them, and stores the posts in a PostgreSQL database.  Users can follow feeds and view the aggregated posts.

## Features

* **RSS Feed Scraping:**  Fetches and parses RSS feeds concurrently using goroutines.  Handles duplicate entries gracefully. The scraper intelligently fetches only new feeds based on a timed schedule.
* **Post Storage:** Stores fetched posts in a PostgreSQL database, including title, description, publication date, and URL.
* **User Authentication:**  Uses API keys for user authentication and authorization.
* **Feed Following:** Allows users to follow specific feeds.
* **Post Retrieval:** Enables users to retrieve posts from the feeds they follow.
* **REST API:** Exposes a RESTful API for interacting with the application.

## Key Components and Files

* **`main.go`:** The main application entry point. Sets up the database connection, starts the scraping process, and configures the HTTP router.
* **`scraper.go`:** Contains the logic for scraping RSS feeds.  The `startScraping` function manages the scraping process, utilizing goroutines for concurrency and a ticker for scheduled fetching. The `scrapeFeed` function handles the fetching and parsing of individual feeds and inserts posts into the database.  It also handles potential errors and duplicate entries.
* **`rss.go`:**  Provides functions for fetching and parsing RSS feeds using the `encoding/xml` package.
* **`models.go`:** Defines the data structures (User, Feed, FeedFollow, Post) used throughout the application. It also includes helper functions to convert between database representations and API representations of these structures.
* **`database` directory:** Contains the SQL queries and database interaction logic.
* **`handler_*.go`:** Files containing HTTP handlers for various API endpoints (user creation, feed creation, feed following, post retrieval, etc.).
* **`middleware_auth.go`:** Implements API key authentication middleware.
* **`json.go`:** Provides helper functions for responding with JSON.

## Scraper Details

The scraper is a core component of this application.  Here's a breakdown of how it works:

1. **Scheduled Scraping:** The `startScraping` function uses a `time.Ticker` to fetch feeds at regular intervals.
2. **Concurrent Fetching:**  It fetches multiple feeds concurrently using goroutines, improving efficiency. The number of concurrent goroutines is configurable.
3. **Database Interaction:** Uses the `database` package to interact with the PostgreSQL database.  The `GetNextFeesdToFetch` function retrieves feeds that haven't been fetched recently.  The `CreatePost` function inserts new posts into the database. `MarkFeedAsFetched` function updates the database to mark the feed as fetched, and prevents it from being selected next time by `GetNextFeesdToFetch` until next time interval.
4. **Duplicate Handling:** The scraper checks for duplicate posts before inserting them into the database, preventing redundant entries.
5. **Error Handling:**  Implements error handling to gracefully handle issues such as network errors or invalid RSS feeds.


## Getting Started

1. Clone the repository.
2. Set up a PostgreSQL database.
3. Create a `.env` file with the following environment variables:
    * `PORT`: The port to run the server on.
    * `DB_URL`: The connection string for your PostgreSQL database.
