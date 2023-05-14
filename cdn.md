# CDN (Content delivery network)

## What is a CDN?
Is a group of servers distributed in different geographical locations, store static data. When a user requests a resource, the CDN will find the nearest server to the user and return the resource to the user. This can reduce the time of the user to get the resource.


## How CDN work?

1. The user makes a request for a piece of content, such as a webpage, image, or video, by entering a URL into their web browser.

2. The request is sent to a domain name server (DNS), which looks up the IP address of the website or CDN that the content is hosted on.

3. The request is then sent to the CDN's edge servers, which are distributed around the world and are designed to serve content as quickly and efficiently as possible.

4. The edge server closest to the user's location receives the request and checks to see if it has a cached copy of the requested content.

5. If the content is cached on the edge server, it is immediately returned to the user's browser. This is known as a cache hit.

6. If the content is not cached on the edge server, the edge server retrieves it from the origin server, which is the server that hosts the original copy of the content.

7. Once the content is retrieved from the origin server, it is cached on the edge server for a specified period of time, so that it can be served more quickly in the future.

8. The content is then returned to the user's browser, completing the request.
 

## Benefits of CDN
- Reduce the time of the user to get the resource. Better user experience. CDN distributed globally and the user can get the resource from the nearest server. Also have good load balance ability and optimal hardware or software configuration.
- Reduce the bandwidth cost of the server.
  It can reduce the bandwidth cost of the server. Because the user can get the resource from the nearest server. The server does not need to send the resource to the user directly.
- Enhance the server content availability and backup ability.
  CDN can cache the resource and store it in different locations. If the server is down, the user can still get the resource from the other CDN.

- Enhance the server security. CDN can prevent DDos attack since we distribute the server to different locations. And other security features.

# DFS (Distributed file system)

## What is a DFS?
A file system that stores data on multiple servers. It can be used to store large amounts of data and provide high availability and high performance.

## How DFS work?
When user upload a file to the DFS, the DFS will split the file into multiple chunks and store them on different servers. When the user wants to download the file, the DFS will find the chunks of the file and merge them together.

Also DFS can replicate the file to different servers. If one server is down, the user can still get the file chunk from the other server.

## Benefits of DFS
- Improved the capacity to change the size of the data and also improves the ability to exchange the data.
- It improved the availability of file, access time, and network efficiency.

# CDN vs DFS
DFS is for numerous users in a team or applications working huge file at the same time. 

CDN provide data to end users, it pay to work with ISP provider at specific location.


