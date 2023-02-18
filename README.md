## DESCRIPTION

Ascii-art-web-stylize consists in creating and running a server, in which it will be possible to use a<br> user-friendly, responsive web GUI (graphical user interface) to print graphical representation of given text.

## AUTHORS

Creata21

## USAGE

Open command line and go to project's folder path something like "student/ascii-art-web-stylize", <br>then run command `go run cmd/main.go`. Next, open your browser and go to "localhost:8080/". 

## iMPLEMENTATION

**Webpage can allow to use different style banners:**

- *shadow*<br>
- *standard*<br>
- *thinkertoy*<br>

**Web-Page implements the following HTTP endpoints:**

GET /: Sends HTML response, the main page.
1.1. GET Tip: go templates to receive and display data from the server.

POST /ascii-art: that sends data to Go server (text and a banner)
2.1. POST Tip: use form and other types of tags to make the post request.


**The main page has:**

text input
radio buttons, select object or anything else to switch between banners
button, which sends a POST request to '/ascii-art' and outputs the result on the page.


**Endpoints return appropriate HTTP status codes.**

OK (200), if everything went without errors.
Not Found, if nothing is found, for example templates or banners.
Bad Request, for incorrect requests.
Internal Server Error, for unhandled errors.
