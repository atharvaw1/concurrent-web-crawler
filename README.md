# Concurrent Web Crawler

A simple concurrent web crawler implemented in Go. The crawler is designed to perform concurrent crawling of web pages up to a specified depth.

## Project Structure

The project is organized into the following directories:

- **`crawler`**: Contains the main logic for crawling web pages concurrently.
- **`urlset`**: Defines a data structure for managing a set of visited URLs.
- **`worker`**: Implements the worker Goroutines that process URLs concurrently.
- **`main.go`**: The entry point of the program that orchestrates the crawling process.

## Usage

To run the concurrent web crawler, use the following command:

```bash
go run main.go 
```

## Dependencies

The project uses the following dependencies:

- `golang.org/x/net/html`: Go package for HTML parsing.

## How It Works

1. The main program starts a crawling process with the specified starting URL and depth.
2. The crawler fetches the HTML content of the page and extracts links.
3. Links are sent to worker Goroutines for further crawling.
4. Workers process links concurrently until the specified depth is reached.
5. The program waits for all Goroutines to finish before completion.


## Contributing

Contributions are welcome! If you have ideas for improvements or new features, please submit a pull request.


Feel free to customize the content further based on your specific project details and preferences.
