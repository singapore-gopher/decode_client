# decode_json
JSON Decode challenge

# how-to

## Challenge3

### Decoding TCP stream

You are connecting to a TCP server which sends you a stream of JSON data. The format is the following:
```json
{"timestamp": 1429004470, "value":"SOME ENCRYPTED VALUE"}
```
The task is to send back 3 "magic" values to the challenge's web server at the url: /stage3/submit.json.
To make it easier for you, there is a method already in the client to do so.
You can use stream.IsMagicValue(string)bool to decide whether a value is of interest, and you'll find some
hints in the code itself.

