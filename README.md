garble
======

garble is a daemon that does a separate evil thing over each port it
listens on. It can be used to test your network clients for resiliency
when faced with failure and evil.

garble was inspired by the Killer Test Harness from Michael Nygard's
talk at Velocity 2012, ["Stablity
Patterns"](http://velocityconf.com/velocity2012/public/schedule/detail/24841)
(slides 44 and 45).

Presently, six not-so-evil ideas are implemented:
- Fast Close: Accept a connection then close it immediately
- Timeout: Accept a connection, do nothing, then close after 30 seconds
- Never Close: Accept a connection and never close it
- Slow but Steady: Accept a connection and return one byte every second
- Random: Accept a connection and return random binary data
- Slow but Steady Proxy: Proxy to a backend server and return the response one byte per second

Ideas to be implemented:
- Read the request at 1 byte/sec
