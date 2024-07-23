# URL_Shortener

step 1.............................

You can download and install Go from the official Go website. https://go.dev/

step 2.............................

git clone https://github.com/bhavesh848785/URL_Shortener.git

step 3.............................
Running the Server
Navigate to the project directory and run the following command to start the server:

> > go run main.go

step 4.............................

Shortening a URL
To shorten a URL, send a POST request to the /shorten endpoint with a JSON payload containing the original URL:

curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'

The response will contain a JSON object with the shortened URL:

{
"short_url": "https://github.com/bhavesh848785"
}

step 5.............................

Redirecting to the Original URL
To redirect to the original URL, visit the shortened URL in your browser or send a GET request to the /redirect/{id} endpoint, where {id} is the shortened URL ID:

curl http://localhost:8080/redirect/4007c7c3
This will redirect you to the original URL associated with the shortened URL.
