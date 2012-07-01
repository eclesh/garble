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

License
=======
Copyright (C) 2012 Eric Lesh

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
