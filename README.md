
# Parsep

Docker-Parsep is a web scraping and rendering service that extracts precious metal prices from various sources and displays them in an HTML table. The project is written in Go and uses Colly for web scraping, Docker for containerization, and a JavaScript module for HTML to PNG conversion.

## Functionality

1. Scrape precious metal prices from various sources using Colly.
2. Process and format the scraped data.
3. Generate an HTML table displaying the prices.
4. Convert the HTML table to a PNG image using a JavaScript module.

## Main Components

- `model`: Contains data structures and constants.
- `repository`: Handles web scraping using Colly.
- `service`: Implements the core logic, including rendering HTML and managing data.
- `render_client`: Provides an HTTP client for interacting with an external rendering service.
- `html-to-png`: JavaScript module for converting HTML to PNG.

## Technologies


-   **Go**: The main language used for the backend logic of the application. Go is used to define the data structures, handle web scraping using Colly, and implement the core logic of the service.
-   **Colly**: A web scraping framework for Go used to extract data from websites. Colly is used to scrape precious metal prices from various sources and process the scraped data.
-   **Docker**: A containerization technology used to package the application and its dependencies for easy deployment. Docker is used to create a container image of the application, which can be easily deployed on any platform.
-   **JavaScript**: A programming language used for the HTML to PNG conversion module. The JavaScript module is used to convert the HTML table generated by the service into a PNG image, which can be easily shared or embedded in other applications.
