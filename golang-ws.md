# Golang Web Servers

In this section we will be looking at how to create a web server in Golang. We will be using the `net/http` package to create a web server. In addition to tat, we explore the fundamentals of building web servers.

## What is a web server?

A Web Server is a software application that handles HTTP requests sent by HTTP clients like web browsers, and returns web pages in response to the clients. Web servers usually deliver html documents along with images, style sheets, and scripts.

![Image](Image)

## What is a Handler Function?

A handler function is a function that is called when a request is received by the server. The handler function is responsible for writing the response headers and body. The handler function is also responsible for closing the connection after the response is sent.

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

Essentially, whenever we hit the endpoint '/', the handler is called. Thus logic behind the enpoint is placed. 

## What does the CORS mdiddleware do?

- CORS stands for Cross-Origin Resource Sharing. It is a mechanism that allows restricted resources on a web page to be requested from another domain outside the domain from which the first resource was served. A web page may freely embed cross-origin images, stylesheets, scripts, iframes, and videos. Certain "cross-domain" requests, notably Ajax requests, are forbidden by default by the same-origin security policy. CORS defines a way in which the browser and the server can interact to determine whether or not to allow the cross-origin request. It is more of a security feature than a feature of the language itself.

Example
    
```go

```