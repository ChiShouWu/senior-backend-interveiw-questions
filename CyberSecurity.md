# Cyber Security

1. Explain SQL injection
    <details>
    <summary>Answer</summary>
    When our service is using SQL database, attacker may insert some SQL command to the input field, and the SQL command will be executed by the database. For example, attacker may insert `select * from users where id = 1; drop table users;` to the input field, and the database will return the user info, and drop the users table.
    </details>
2. Which method would you use to prevent SQL injection risks?
    <details>
    <summary>Answer</summary>

    - Use prepared statements and parameterized queries and make sure the interpreter can tell the difference between data and code

    - Use stored procedures and calling them needed to avoid dynamic generation in SQL
    - Use white list input validation and avoid blacklist methods since these are not as secure
    </details>

3. International Organization for Standardization (ISO) 27034
4. [Center for Internet Security (CIS) Control 16: Application Software Security](https://blog.netwrix.com/2022/04/05/cis-control-application-software-security/)
5. [Payment Card Industry (PCI) Payment Application Data Security Standard (PA-DSS)](https://listings.pcisecuritystandards.org/minisite/en/docs/PA-DSS_v3.pdf)
6. [OWASP Application Security Verification Standard (ASVS)](https://owasp.org/www-project-application-security-verification-standard/)

7. [What are disadvantages of using session based authentication?](https://github.com/Gauthamjm007/Backend-NodeJS-Golang-Interview_QA#what-are-disadvantages-of-using-session-based-authentication)

8. [What are different authentication methods?](https://github.com/Gauthamjm007/Backend-NodeJS-Golang-Interview_QA#what-are-disadvantages-of-using-session-based-authentication

9.  Man-in-the-middle attack
    <details>
    <summary>Answer</summary>
    The attacker spy on the transport of communication. Insert a layer into the line and exchange the data of 2 sides. Thus, the attacker can fabricate the data and send to the other side.
    So attacker can steal your data.
    </details>

10. What is XSS? How to prevent it?
    <details>
    <summary>Answer</summary>
    XSS is a type of injection, in which malicious scripts are injected into otherwise benign and trusted websites. XSS attacks occur when an attacker uses a web application to send malicious code, generally in the form of a browser side script, to a different end user. Flaws that allow these attacks to succeed are quite widespread and occur anywhere a web application uses input from a user within the output it generates without validating or encoding it.
    </details>

12. Define Traceroute.
    <details>
    <summary>Answer</summary>
    Tracerroute is a network tool depends on ICMP. It can show the route like router and ip address of the packet from the source to the destination.
    It use TTL to implement. When packet reach the router, the router will decrease the TTL by 1. If the TTL is 0, the packet will be dropped. So the router can know the route of the packet.
    Most of the firewall will block the ICMP packet, so traceroute will not work. Or the router will not decrease the TTL, so the packet will not be dropped.
    <details>

13. 防火牆那邊可能會期望知道 WAF 跟iptable 之類的ㄅ，然後 NAT 運作方式