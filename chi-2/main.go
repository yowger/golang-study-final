package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var corsOptions = cors.Options{
	AllowedOrigins:   []string{"https://www.google.com"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge:           300,
}

func main() {
	r := chi.NewRouter()

	// r.Use(cors.Handler(
	// 	cors.Options{
	// 		AllowedOrigins:   []string{"https://www.google.com"},
	// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 		ExposedHeaders:   []string{"Link"},
	// 		AllowCredentials: false,
	// 		MaxAge:           300,
	// 	},
	// ))
	// r.Use(middleware.CleanPath)
	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.Heartbeat("/health"))
	// r.Use(middleware.AllowContentType("application/json", "text/xml"))
	// r.Use(middleware.Compress(5, "text/html", "text/css"))
	// r.Use(middleware.Timeout(30 * time.Second))
	// r.Use(middleware.GetHead)

	// for dynamic data
	// r.Use(middleware.NoCache)

	r.Use(
		cors.Handler(corsOptions),
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(30*time.Second),
		middleware.Compress(5),
		middleware.Heartbeat("/health"),
	)

	r.Group(func(r chi.Router) {
		r.Get("/", helloWorldHandler)
	})
	r.Group(func(r chi.Router) {
		// r.Use(AuthMiddleware)
		r.Use(MyMiddleware)

		r.Route("/todo", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				user := r.Context().Value("user").(string)

				w.Write([]byte(fmt.Sprintf("Get all todos %s", user)))
			})

			r.Route("/{todoID}", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Get single todo"))
				})
				r.Put("/", func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Update todo"))
				})
				r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Delete dodo"))
				})
			})
		})
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Could not start server:", err)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func getAdminProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello admin!"))
}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

/*
	for middle ware

	Chi provides a variety of middlewares that help streamline common HTTP functionalities in Go applications. Which middlewares are "important" depends on the needs of your application, but here are some **core middlewares commonly used in most projects**:

---

### **1. `middleware.Logger`**
- **Purpose**: Logs HTTP requests, including the method, URL, response status, and execution time.
- **Importance**:
  - Essential for debugging and monitoring.
  - Useful for understanding how your application is being used.
- **Example**:
  ```go
  r.Use(middleware.Logger)
  ```

---

### **2. `middleware.Recoverer`**
- **Purpose**: Recovers from panics and ensures the server doesn't crash.
- **Importance**:
  - Critical for production to handle unexpected errors gracefully.
  - Ensures panics return a 500 Internal Server Error instead of crashing.
- **Example**:
  ```go
  r.Use(middleware.Recoverer)
  ```

---

### **3. `middleware.RequestID`**
- **Purpose**: Adds a unique request ID to each incoming request, which can be used for logging and debugging.
- **Importance**:
  - Helps correlate logs across distributed systems or microservices.
  - Useful for tracing individual requests.
- **Example**:
  ```go
  r.Use(middleware.RequestID)
  ```

---

### **4. `middleware.Timeout`**
- **Purpose**: Sets a maximum duration for handling a request.
- **Importance**:
  - Prevents long-running requests from consuming server resources indefinitely.
  - Critical for APIs that need predictable response times.
- **Example**:
  ```go
  r.Use(middleware.Timeout(30 * time.Second))
  ```

---

### **5. `middleware.Compress`**
- **Purpose**: Compresses responses (e.g., Gzip or Deflate) to reduce payload size.
- **Importance**:
  - Improves performance by reducing bandwidth usage.
  - Useful for APIs serving large responses.
- **Example**:
  ```go
  r.Use(middleware.Compress(5)) // Compression level: 0-9
  ```

---

### **6. `middleware.AllowContentType`**
- **Purpose**: Restricts incoming requests to specific `Content-Type` headers.
- **Importance**:
  - Improves security by rejecting unexpected request formats.
  - Useful for APIs accepting only JSON, XML, etc.
- **Example**:
  ```go
  r.Use(middleware.AllowContentType("application/json"))
  ```

---

### **7. `middleware.Heartbeat`**
- **Purpose**: Adds a health check endpoint (e.g., `/health`).
- **Importance**:
  - Essential for monitoring server uptime.
  - Integrates well with tools like Kubernetes or AWS load balancers.
- **Example**:
  ```go
  r.Use(middleware.Heartbeat("/health"))
  ```

---

### **8. `middleware.GetHead`**
- **Purpose**: Converts `HEAD` requests into `GET` requests automatically.
- **Importance**:
  - Useful for APIs that don't explicitly implement `HEAD` handlers.
  - Ensures compliance with HTTP standards.
- **Example**:
  ```go
  r.Use(middleware.GetHead)
  ```

---

### **9. `middleware.NoCache`**
- **Purpose**: Disables caching for responses.
- **Importance**:
  - Useful for endpoints that should always serve fresh data (e.g., dynamic APIs).
- **Example**:
  ```go
  r.Use(middleware.NoCache)
  ```

---

### **10. `middleware.StripSlashes`**
- **Purpose**: Normalizes URLs by removing trailing slashes.
- **Importance**:
  - Improves consistency in route handling.
  - Prevents duplication issues with URLs like `/users` and `/users/`.
- **Example**:
  ```go
  r.Use(middleware.StripSlashes)
  ```

---

### **Recommended Combination**
For most applications, these middlewares together provide a solid foundation:
```go
r.Use(
  middleware.RequestID,
  middleware.Logger,
  middleware.Recoverer,
  middleware.Timeout(30*time.Second),
  middleware.Compress(5),
  middleware.Heartbeat("/health"),
)
```

---

### **When to Add Additional Middlewares**
- **Security**: Use `AllowContentType` and `NoCache` for stricter API behavior.
- **Performance**: Use `Compress` for bandwidth optimization.
- **Routing Consistency**: Use `StripSlashes` for uniform URL handling.

---

### **Conclusion**
Focus on **`Logger`**, **`Recoverer`**, **`RequestID`**, and **`Timeout`** for critical functionality in most Go projects. Use other middlewares based on your application's specific needs, like content-type validation, compression, or caching. Chi's middleware collection is modular, so you can choose only what you need.
*/

/*
	test in google console

	fetch('https://api.example.com/data', {
	method: 'GET',
	headers: {
		'Content-Type': 'application/json',
	},
	})
	.then(response => response.json())
*/

/*
The content sent to a server in an HTTP request depends on the application's purpose and the format the server expects. Here are the most common types of content clients send to servers:

---

### **1. Form Data**
#### **Type**: `application/x-www-form-urlencoded`
- **Description**: This is the default for HTML forms. Data is sent in key-value pairs, URL-encoded.
- **Example**:
  ```http
  POST /submit-form HTTP/1.1
  Content-Type: application/x-www-form-urlencoded

  name=John+Doe&age=30
  ```
- **Use Case**: Simple form submissions.

---

### **2. JSON (JavaScript Object Notation)**
#### **Type**: `application/json`
- **Description**: A lightweight data-interchange format, widely used in modern APIs.
- **Example**:
  ```http
  POST /api/user HTTP/1.1
  Content-Type: application/json

  {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```
- **Use Case**: REST APIs for structured data.

---

### **3. XML (Extensible Markup Language)**
#### **Type**: `application/xml` or `text/xml`
- **Description**: A markup language used for structured data exchange.
- **Example**:
  ```http
  POST /api/user HTTP/1.1
  Content-Type: application/xml

  <user>
    <name>John Doe</name>
    <email>john.doe@example.com</email>
  </user>
  ```
- **Use Case**: Legacy systems or specific APIs requiring XML.

---

### **4. Multipart Form Data**
#### **Type**: `multipart/form-data`
- **Description**: Used for sending binary data, such as files, along with text fields.
- **Example**:
  ```http
  POST /upload HTTP/1.1
  Content-Type: multipart/form-data; boundary=------WebKitFormBoundary

  ------WebKitFormBoundary
  Content-Disposition: form-data; name="file"; filename="example.txt"
  Content-Type: text/plain

  File content goes here
  ------WebKitFormBoundary
  Content-Disposition: form-data; name="description"

  File description
  ------WebKitFormBoundary--
  ```
- **Use Case**: File uploads.

---

### **5. Raw Text**
#### **Type**: `text/plain`
- **Description**: Plain text data sent directly.
- **Example**:
  ```http
  POST /api/log HTTP/1.1
  Content-Type: text/plain

  This is a log message.
  ```
- **Use Case**: Logging or sending simple, unstructured data.

---

### **6. Binary Data**
#### **Type**: `application/octet-stream`
- **Description**: Raw binary data (e.g., images, videos, or custom file formats).
- **Example**:
  ```http
  POST /upload HTTP/1.1
  Content-Type: application/octet-stream

  <binary data>
  ```
- **Use Case**: Uploading non-text files.

---

### **7. GraphQL Queries**
#### **Type**: `application/graphql` or `application/json`
- **Description**: A query language for APIs that retrieves specific data structures.
- **Example**:
  ```http
  POST /graphql HTTP/1.1
  Content-Type: application/json

  {
    "query": "{ user { name email } }"
  }
  ```
- **Use Case**: APIs using GraphQL.

---

### **8. Protobuf (Protocol Buffers)**
#### **Type**: `application/protobuf`
- **Description**: A compact binary format used for efficient data exchange.
- **Example**:
  - Content is binary and not human-readable.
- **Use Case**: High-performance APIs or gRPC.

---

### **9. Custom Content Types**
#### **Type**: `application/vnd.custom+json` (or similar)
- **Description**: APIs can define custom media types for specific use cases.
- **Example**:
  ```http
  POST /custom-endpoint HTTP/1.1
  Content-Type: application/vnd.example+json

  { "customData": "value" }
  ```
- **Use Case**: APIs with domain-specific data.

---

### **10. URL Parameters (Query Strings)**
#### **Type**: Part of the URL (`GET` requests)
- **Description**: Data sent as part of the URL, typically in key-value pairs.
- **Example**:
  ```http
  GET /search?q=golang&page=2 HTTP/1.1
  ```
- **Use Case**: Non-sensitive data or filtering.

---

### **11. WebSocket Data**
- **Description**: While not a traditional HTTP content type, WebSockets send messages (often JSON or binary) after an initial HTTP handshake.
- **Example**: Sending chat messages in real-time.

---

### **How Servers Handle Content**
1. **Content-Type Header**:
   - The `Content-Type` header tells the server how to parse the incoming data.
   - Example: `application/json` means the server should parse the body as JSON.

2. **Middleware**:
   - Middleware like `middleware.AllowContentType` in `chi` ensures the server only accepts specific content types.

---

### **Key Points to Consider**
- **Security**: Validate and sanitize incoming data to prevent injection attacks.
- **Consistency**: Clearly document what content types your server expects.
- **Efficiency**: Choose formats like JSON or Protobuf for structured data and performance.

By understanding the type of content clients send, you can design robust and secure APIs.
*/

/*
Here’s a detailed explanation of your questions:

---

### **1. Other Examples of Headers (for `AllowedHeaders` and `ExposedHeaders`)**

#### **Common Request Headers (for `AllowedHeaders`)**
- **`X-Requested-With`**: Typically used in AJAX requests to indicate the request originated from JavaScript.
- **`Referer`**: Indicates the URL of the page making the request.
- **`Origin`**: Specifies the origin of the request, such as the domain or scheme (useful for security checks).
- **`Authorization`**: Used for passing credentials like tokens or API keys.
- **`Content-Length`**: Indicates the size of the body of the request in bytes.
- **`Accept-Language`**: Specifies the preferred language of the client.
- **`If-Modified-Since`**: Used for conditional requests, asking the server to return content only if it has been modified since the specified time.

#### **Common Response Headers (for `ExposedHeaders`)**
- **`ETag`**: A unique identifier for a specific version of a resource, useful for caching and conditional requests.
- **`Cache-Control`**: Provides caching instructions to the client.
- **`Retry-After`**: Indicates how long the client should wait before making another request (useful in rate-limiting).
- **`Content-Disposition`**: Provides instructions for how content should be displayed (e.g., as an attachment for file downloads).
- **`X-RateLimit-Limit`**: Indicates the maximum number of requests allowed within a certain period.
- **`X-RateLimit-Remaining`**: Shows how many requests the client can still make in the current period.
- **`Location`**: Used in responses to indicate a URL for redirection or the location of a newly created resource.

---

### **2. Explanation of `MaxAge`**

#### **What It Does**
`MaxAge` sets the time (in seconds) for which the browser will cache the result of a preflight OPTIONS request.

#### **Why Preflight Caching Matters**
When a browser makes a cross-origin request, it often sends a preflight OPTIONS request to check the server's CORS policy (e.g., allowed origins, methods, and headers). If you set `MaxAge`, the browser will cache the server's response for the specified duration. During this time:
- The browser won't re-send the OPTIONS request for the same resource and headers.
- It will reuse the cached CORS policy for subsequent requests.

#### **Practical Example**
1. **First Request**:
   - A `POST` request from `https://example.com` to `https://api.example.com` is sent.
   - The browser sends a preflight OPTIONS request.
   - The server responds with allowed origins, methods, headers, etc.
   - The browser caches this response for the `MaxAge` duration (e.g., 300 seconds).

2. **Second Request** (within 300 seconds):
   - A similar request is made.
   - The browser skips the OPTIONS preflight because it has the cached policy.

#### **Key Benefit**:
This improves performance by reducing unnecessary network traffic.

#### **Important Note**:
If the policy (e.g., allowed methods or headers) changes on the server before the cached duration expires, the client might still use the old policy until the cache is invalidated.

---

### **Summary**
- `AllowedHeaders`: Controls what headers the client can send.
- `ExposedHeaders`: Controls what headers the client can access in the response.
- `MaxAge`: Avoids repetitive preflight requests by caching the server's CORS policy for a set time. Yes, it saves the server's allowed origins, methods, headers, etc., temporarily for better efficiency.

If you have a dynamic or sensitive CORS policy, keep `MaxAge` short (or at default). For stable APIs with consistent policies, a higher `MaxAge` can significantly boost performance.
*/
