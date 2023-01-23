# HTTP-proxy + vulnerability scanner on top of it

Small project to learn Go language.

See also:

- <https://hub.docker.com/repository/docker/coradead/proxy_server> — Java
- <https://github.com/Natali-Skv/technopark_IS_http_proxy> — Go
- <https://github.com/kirill555101/http-proxy> — Node.js
- <https://github.com/papazloynt/proxy> — Go
- <https://github.com/yutfut/Proxy-Server> — Go

For reference:

- <https://github.com/mitmproxy/mitmproxy> — Python
- <https://github.com/john-pentest/fproxy> — Node.js
- <https://github.com/kr/mitm> — Go
- <https://github.com/zaproxy/zaproxy> — Java

Resources

- [RFC 9112 HTTP/1.1](https://datatracker.ietf.org/doc/html/rfc9112)
- [RFC 9293 Transmission Control Protocol (TCP)](https://datatracker.ietf.org/doc/html/rfc9293)

## Usage

1. Start server. `go run .`
2. Make a request to proxy.
    - `curl --include --proxy http://127.0.0.1:8080 http://mail.ru`
    - `curl --include --proxy http://127.0.0.1:8080 http://www.google.com`
