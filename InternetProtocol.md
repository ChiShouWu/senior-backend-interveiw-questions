# Internet Protocol

1. What's the different between Unicode and UTF-8？
    <details><summary>Answer</summary>
    - Unicode is a character set, UTF-8 is a kind of it.
2. What's the different between `session` and `cookie`?
    <details><summary>Answer</summary>
    Session is a server-side storage, cookie is a client-side storage.
    <details>
3. Please briefly explain how the CORS mechanism works and its purpose.
    <details><summary>Answer</summary>
    When a request is send from users browser to a server, the browser will add a `Origin` header to the request. Which contains `protocol`, `host` and `port` of the request means which web this request from. The server will check if the `Origin` is in the `Access-Control-Allow-Origin` list. If not, the server will return a `403` error.
    <details>

4. Please briefly explain how the HTTP Cache mechanism(Expires, Cache-Control,
   <details><summary>Answer</summary>
    - Expires: The time when the cache will be expired. But if user computer's time is not correct, the cache will be expired at wrong time.
    - Cache-control: The cache control header is used to specify the cacheability of a response. It can be used to specify the maximum age of the response, whether it can be cached by a public or private cache, and whether it can be cached at all.
    - Last-Modified: The time when the resource was last modified.
    - If-Modified-Since: The time when the resource was last modified. If the resource is expired, browser will ask the server check if the resource is modified. If not, the server will return a `304` status code.
    - Etag: The unique identifier of the resource. Like hash value of the resource.
    - If-None-Match: Browser send to server the Etag as `If-None-Match` header. If the resource was changed, the server will return a new resource with a new Etag. If not, the server will return a `304` status code.
5. Explain in detail what happens if you visit a website
   <details><summary>Answer</summary>
    - DNS lookup
    - TCP connection
    - HTTP request
    - HTTP response
    - TCP connection closed
    - DNS cache
    <details>
6. What are the seven layers in the OSI system model?
    <details><summary>Answer</summary>
    - Application: The application layer is the top layer of the OSI model. It is responsible for the presentation of data to the user and the application programs.
    - Presentation: The presentation layer is responsible for the translation of data between the application layer and the session layer.
    - Session： The session layer is responsible for the establishment and termination of sessions between the communicating devices.
    - Transport: The transport layer is responsible for the reliable transmission of data between the communicating devices.
    - Network: The network layer is responsible for the routing of data between different networks.
    - Data Link: The data link layer is responsible for the reliable transmission of data between two directly connected devices.
    - Physical: The physical layer is responsible for the transmission of raw data between two directly connected devices.
    <details>
7. What’s a reverse proxy?
    <details><summary>Answer</summary>
    A reverse proxy is a type of proxy server that retrieves resources on behalf of a client from one or more servers. These resources are then returned to the client, appearing as if they originated from the proxy server itself.
    <details>
8. Which method would you use to handle the API versioning of web services?
    <details><summary>Answer</summary>
    - URL versioning: `https://api.example.com/v1/`
    - Header versioning: `Accept: application/vnd.example.v1+json`
    - Media type versioning: `Accept: application/json; version=1`
    <details>
9.  Why TCP hand shake need 3 steps? And what is 4 way Wavehand?
    <details><summary>Answer</summary>

    - SYN: The client sends a SYN packet to the server, and the client enters the SYN_SENT state.

    - SYN/ACK: The server receives the SYN packet and sends a SYN/ACK packet to the client, and the server enters the SYN_RCVD state.

    - ACK: The client receives the SYN/ACK packet and sends an ACK packet to the server, and the client enters the ESTABLISHED state. The server receives the ACK packet and enters the ESTABLISHED state.

    - 4 way wavehand: The client sends a FIN packet to the server, and the client enters the FIN_WAIT_1 state. The server receives the FIN packet and sends an ACK packet to the client, and the server enters the CLOSE_WAIT state. The client receives the ACK packet and enters the FIN_WAIT_2 state. The server sends a FIN packet to the client, and the server enters the LAST_ACK state. The client receives the FIN packet and sends an ACK packet to the server, and the client enters the TIME_WAIT state. The server receives the ACK packet and enters the CLOSED state. The client waits for 2MSL and enters the CLOSED state.
      <details>
10. TCP v.s. UDP
11. How TCP implement flow control?
    <details><summary>Answer</summary>
    
12. HTTP v.s. HTTPS
13. HTTPS 加密過程
14. HTTP v.s. HTTP2 v.s. HTTP3
15. What is socket?
16. What it 2MSL in TCP?
17. Header usually used in HTTP request