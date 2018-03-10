# json2msgpack

## Installation

```sh
go get -u github.com/saidie/json2msgpack
go install github.com/saidie/json2msgpack
```

## Usage

```sh
cat <<JSON | json2msgpack > result.msgpack
{
  "hello": "world",
  "array": [
    12345,
    {
      "object": {
        "foo": "bar"
      }
    }
  ]
}
JSON
```
