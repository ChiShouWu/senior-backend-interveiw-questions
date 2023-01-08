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
    <details><summary>Answer</summary>
    TCP will check if the data is received correctly, and UDP will not. UDP is faster than TCP. TCP can used when want data reliable, like email, file transfer,
     UDP can used when want data fast, like streaming, online game

11. How TCP implement flow control?
    <details><summary>Answer</summary>

12. HTTP v.s. HTTPS
    <details><summary>Answer</summary>
    - HTTP: HyperText Transfer Protocol, HTTP is a protocol for transmitting hypermedia documents, such as HTML. It was designed for communication between web browsers and web servers, but it can also be used for other purposes. HTTP follows a classical client-server model, with a client opening a connection to make a request, then waiting until it receives a response. HTTP is a stateless protocol, meaning that the server does not keep any data (state) between two requests. Each request is treated as an independent transaction. HTTP is a text-based protocol in which requests and responses consist of text. The text-based nature of HTTP makes it easy to use in debugging and logging, as well as for testing and automation.

    - HTTPS: HTTP over TLS/SSL, HTTPS is the secure version of HTTP. HTTPS is HTTP over TLS/SSL. HTTPS is often used to protect highly confidential online transactions like online banking and online shopping orders. HTTPS is also used to protect communication between web servers and web browsers, including communication of sensitive data, such as passwords and credit card numbers. HTTPS uses the same TCP port as HTTP, which is port 80. The application protocol used by HTTPS is HTTP. HTTPS is the secure version of HTTP. HTTPS is HTTP over TLS/SSL. HTTPS is often used to protect highly confidential online transactions like online banking and online shopping orders. HTTPS is also used to protect communication between web servers and web browsers, including communication of sensitive data, such as passwords and credit card numbers. HTTPS uses the same TCP port as HTTP, which is port 80. The application protocol used by HTTPS is HTTP.
    <details>

13. Describe the process of https encryption
    <details><summary>Answer</summary>
    - The client sends a ClientHello message to the server, which includes the version of the SSL/TLS protocol that the client supports, a client-generated random value, and the client's list of supported cipher suites.

    - The server responds with a ServerHello message, which includes the version of the SSL/TLS protocol that the server supports, a server-generated random value, and the server's selected cipher suite.

    - The server sends its certificate to the client. The certificate contains the server's public key.

    - The server sends a ServerHelloDone message to the client, which indicates that the server is done with its part of the handshake.

    - The client verifies the server's certificate. If the certificate is not valid, the client will terminate the connection.

    - The client sends a ClientKeyExchange message to the server. The message contains the client's premaster secret, which is encrypted using the server's public key.

    - The client sends a ChangeCipherSpec message to the server. This message notifies the server that subsequent messages will be protected using the session keys.

    - The client sends a Finished message to the server. This message is used to verify that the key exchange and authentication processes were successful.

    - The server sends a ChangeCipherSpec message to the client.

    - The server sends a Finished message to the client.
    <details>


14. HTTP v.s. HTTP2 v.s. HTTP3
    <details><summary>Answer</summary>
    - HTTP: HyperText Transfer Protocol, HTTP is a protocol for transmitting hypermedia documents, such as HTML. It was designed for communication between web browsers and web servers, but it can also be used for other purposes. HTTP follows a classical client-server model, with a client opening a connection to make a request, then waiting until it receives a response. HTTP is a stateless protocol, meaning that the server does not keep any data (state) between two requests. Each request is treated as an independent transaction. HTTP is a text-based protocol in which requests and responses consist of text. The text-based nature of HTTP makes it easy to use in debugging and logging, as well as for testing and automation.

    - HTTP2: HTTP/2 is a major revision of the HTTP network protocol used by the World Wide Web. It was developed by the HTTP Working Group of the Internet Engineering Task Force (IETF) to improve web page performance over HTTP. HTTP/2 allows multiplexing of requests over a single TCP connection, which enables more efficient use of the network. It also allows servers to push content to the client without the client having to request it, which can reduce latency. HTTP/2 is specified in RFC 7540, published in May 2015.

    - HTTP3: HTTP/3 is the third major version of the Hypertext Transfer Protocol (HTTP) used by the World Wide Web. HTTP/3 is a binary protocol that uses the QUIC transport protocol, which is based on the Internet Engineering Task Force (IETF) Transport Layer Security (TLS) protocol. HTTP/3 is the successor to HTTP/2, which was based on the Transport Layer Security (TLS) protocol. HTTP/3 is currently in draft stage, and is expected to be finalized in 2020.
    It use UDP instead of TCP, so it is faster than HTTP2.
    <details>


15. What is socket?
    <details><summary>Answer</summary>
    A socket is an endpoint for communication between two machines. A socket is bound to a port number so that the TCP layer can identify the application that data is destined to be sent to. An endpoint is a combination of an IP address and a port number. A socket is bound to a port number so that the TCP layer can identify the application that data is destined to be sent to. An endpoint is a combination of an IP address and a port number.
    <details>
16. What it 2MSL in TCP?
    <details><summary>Answer</summary>
    2MSL is the maximum segment lifetime. It is the maximum time a segment is allowed to exist in the network. When a endpoint of TCP connection close and send last ACK to the other endpoint, the other endpoint will wait for 2MSL and then close the connection. If the other endpoint does not receive the ACK, it will retransmit the FIN packet. If the other endpoint receives the ACK, it will close the connection.
    <details>

17. What headers usually used in HTTP request
    - 
