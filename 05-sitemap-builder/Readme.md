# Sitemap Builder

Sitemap Builder is a Go-based tool that generates a sitemap for a specified website. It starts from the root URL and explores all reachable internal pages, creating an XML sitemap.

## Features

- Crawls and maps internal links starting from a specified root URL.
- Generates a standard XML sitemap suitable for search engines.
- Prevents infinite loops by tracking visited pages.
- Supports setting a maximum depth for link exploration.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.13 or later)
- Basic command-line knowledge

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/sitemap-builder.git
cd sitemap-builder
make run
```

## Usage