# decode_json
JSON Decode challenge

## Challenge 1

### The Warmup

Simply add the first and the second values in the data you get from the server. Return it in the specified format, along with your team's name.

### GET http://forgot_to_update_hostname/stage1/data.json
```bash
> GET /stage1/data.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 10:25:30 GMT
< Content-Length: 109
< Content-Type: text/plain; charset=utf-8
<

{"inputs":[{"first":5,"second":10},{"first":7,"second":234},{"first":9,"second":8},{"first":14,"second":84}]}
```

### POST http://forgot_to_update_hostname/stage1/submit.json
```bash
> POST /stage1/submit.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
> Content-type: application/json
> Content-Length: 78
>

{"team":"YOURTEAM","solutions":[{"sum":15},{"sum":241},{"sum":17},{"sum":9811}]}

< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 08:14:12 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<

{"passed":3}
```