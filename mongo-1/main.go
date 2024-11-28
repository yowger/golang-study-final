package main

import (
	"context"
	"fmt"
	"log"
	"myapp/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func connectMongoDB(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(opts)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	cfg := config.LoadConfig()

	client, err := connectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Error disconnecting MongoDB:", err)
		}
	}()

	db := client.Database("go-mongo-1")
	collection := db.Collection("users")

	doc := bson.M{"name": "Alice", "age": 30}
	result, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Fatal("Error inserting document:", err)
	}
	fmt.Println("Inserted document ID:", result.InsertedID)

	filter := bson.M{"name": "Alice"}
	var result2 bson.M
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal("Error finding document:", err)
	}
	fmt.Println("Found document:", result2)

	// server := &http.Server{
	// 	Addr:    ":" + cfg.AppPort,
	// 	Handler: nil,
	// }

	// fmt.Println("Starting server on port", cfg.AppPort)
	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatal("failed to start server: ", err)
	// }

}

/*
	why defer
	manual cleanup vs automatic garbage collection in node js

The difference in behavior between Node.js (e.g., in Express) and Go arises from how each language and runtime handles resource management, garbage collection, and concurrency. Here’s a detailed comparison:

---

### **1. Node.js and Express: Automatic Cleanup**
- **Event Loop Architecture**:
  Node.js is built on an event-driven, non-blocking I/O model. Operations like HTTP requests, database queries, or timeouts are managed by the event loop. When an asynchronous operation completes, its callback is executed, and resources are cleaned up automatically by Node.js's garbage collector.

- **No Explicit Cleanup**:
  In most cases, you don’t need to manually clean up things like timeouts or HTTP requests in Express. Node.js automatically:
  - Cancels ongoing timeouts when a request is completed.
  - Frees up memory associated with completed promises.
  - Closes unused network sockets when the server shuts down.

- **Context-Like Behavior in Node.js**:
  In Node.js, you can use constructs like `req` and `res` objects in Express to pass around state or "context." These are scoped to the lifetime of the request. For instance:
  ```javascript
  app.get('/example', async (req, res) => {
      const data = await someAsyncOperation();
      res.send(data); // Automatically ends the request lifecycle.
  });
  ```
  When `res.send()` is called, Node.js automatically handles cleaning up the request-response cycle.

---

### **2. Go: Manual Resource Management**
Go’s design emphasizes explicit control over resources. This is a key reason why manual cleanup is often required:

- **Concurrency with Goroutines**:
  - In Go, goroutines are lightweight threads, and you have fine-grained control over how they start, stop, or share state. However, there’s no automatic way to stop them unless explicitly instructed.
  - `context` provides a mechanism to propagate cancelation signals across goroutines. Without explicitly canceling or timing out a context, goroutines might keep running, consuming unnecessary resources.

- **Manual Cleanup with `defer`**:
  - Go requires you to manually clean up resources such as timers, database connections, or HTTP client requests to avoid resource leaks.
  - Example:
    ```go
    func main() {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel() // Ensures cleanup when the function exits
        result, err := longRunningOperation(ctx)
    }
    ```

- **Why `defer`?**:
  - Go does not provide an implicit mechanism for releasing resources like timers or context-based signals.
  - Without `defer cancel()`, the context timer or derived signals would persist until garbage collection, which could cause unnecessary resource usage.

---

### **3. Key Differences in Resource Management**
| Feature                      | Node.js (Express)                                      | Go                                                |
|------------------------------|-------------------------------------------------------|--------------------------------------------------|
| **Concurrency Model**         | Event Loop with async callbacks/promises              | Goroutines (lightweight threads)                 |
| **Automatic Cleanup**         | Yes, most async operations are cleaned up automatically | No, explicit cleanup (e.g., `defer cancel()`) is required |
| **Garbage Collection**        | Managed by V8, automatically handles memory and objects | Managed by Go runtime, but explicit cleanup is needed for context/goroutines |
| **Context/Request Lifecycle** | Request lifecycle (`req`, `res`) scoped to each request | Contexts (`context.Context`) must be explicitly managed |
| **Timeouts and Cancelation**  | Built into APIs (e.g., `setTimeout`, `AbortController`) | Requires `context.WithTimeout` or `context.WithCancel` |

---

### **4. Why Go Requires Explicit Cleanup?**
1. **Design Philosophy**:
   - Go emphasizes simplicity and predictability. By requiring explicit cleanup, Go developers have fine-grained control over resource usage.
   - This reduces hidden behaviors or surprises that might arise in complex systems.

2. **Efficient Resource Usage**:
   - Without explicit cleanup, timers and goroutines could linger indefinitely, leading to memory leaks or excessive resource consumption.

3. **Concurrency Model Differences**:
   - Node.js relies on its single-threaded event loop to manage the lifecycle of asynchronous operations.
   - Go’s goroutines run concurrently and independently; unless explicitly canceled or stopped, they will continue execution.

---

### **5. Do We Ever Need Manual Cleanup in Node.js?**
Yes, but it's less common than in Go. Examples include:
- **Abort Signals**:
  For cancelable operations like HTTP requests, you might use `AbortController`:
  ```javascript
  const controller = new AbortController();
  const signal = controller.signal;

  setTimeout(() => controller.abort(), 5000); // Cancel after 5 seconds

  fetch('https://example.com', { signal })
      .then(response => console.log(response))
      .catch(err => console.error('Request aborted', err));
  ```

- **Event Listeners**:
  If you add event listeners in Node.js, you should remove them manually to avoid memory leaks:
  ```javascript
  const handler = () => console.log('Event fired');
  emitter.on('event', handler);
  emitter.off('event', handler); // Manual cleanup
  ```

---

### **Conclusion**
- In **Node.js**, resource cleanup is mostly automatic due to the event loop and garbage collection.
- In **Go**, manual cleanup (e.g., with `defer`) is required because Go prioritizes explicit control over resource management, making it more predictable in multi-threaded scenarios.

Understanding this distinction can help you write efficient, resource-conscious applications in both languages!
*/
