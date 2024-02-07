# ttoy

```text
NAME:
   ttoy - terminal dev toys

USAGE:
   ttoy  [command [command options]] [arguments...]

COMMANDS:
   json       format json or convert json to yaml, toml and so on
   yaml       convert yaml to json, toml and so on
   toml       convert toml to json, yaml and so on
   xml, html  format xml/html or convert to markdown
   encode     encode url, html and so on
   decode     decode url, html and so on
   hash       generate hashes
   uuid       generate uuid
```

## Examples

### Formaters
```shell
echo '{"name":"Tom", "age":27}' | ttoy json
```
or
```shell
cat test.json | ttoy json
```
or

```shell
ttoy json << EOF
{"name":"Tom",
"age":27}
EOF
```
or just
```shell
ttoy json
```
then a tui will appear to take inputs

### Converters

```shell
echo '{"name":"Tom", "age":27}' | ttoy json toml
```

### Encoders/Decoders

```shell
echo 'example.com/pet?name=tom&age=27' | ttoy encode url 
```

### Generaters

```shell
ttoy uuid
```

