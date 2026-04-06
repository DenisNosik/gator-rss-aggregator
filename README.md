### Gator
The rss blog aggregator

Prerequisites

- Go (version 1.25 or higher)
- PostgreSQL

## Installation

### 1. Install the project:

```bash
> go install github.com/DenisNosik/gator-rss-aggregator@latest
```
### 2. Create a configuration file:

Create file *~/.gatorconfig.json* in your home directory with the following content:
```bash
{
  "db_url": "postgres://username:password@host:5432/database_name",
  "current_user_name": "your_username"
}
```
*Replace **username**, **password**, **host**, **database_name** and **your_username** with your actual data.*

## Database Setup
### 1. Create the database
### 2. Run migrations manually (e.g., goose postgres "connection-string" up)

## Usage
### 1. Register a new user
```bash
> gator-rss-aggregator register your_username
```
*Now you can login with*
```bash
> gator-rss-aggregator login your_username
```
### 2. To add new feed use
```bash
> gator-rss-aggregator addfeed "feed_name" "feed_url"
```

### 3. Aggregate (fetch) new posts
The agg command starts an infinite aggregation loop — it will continuously scan all added RSS feeds and add new posts to the database.
```bash
# start fetching posts every 3 min.
> gator-rss-aggregator agg 3m

# start fetching posts every 30 sec.
> gator-rss-aggregator agg 30s
```
Supported time formats: s (seconds), m (minutes), h (hours).

**Press Ctrl + C in the terminal to stop the loop.**

### 4. Browse posts
After fetching posts, you can view them with the browse command:
```bash
# show default number of posts (defaul 2)
> gator-rss-aggregator browse

# Show latest 10 posts
> gator-rss-aggregator browse 10
```
