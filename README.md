# ttoy

terminal dev toys

- Formaters
- Converters
- Generaters
- Encoders/Decoders

## Examples

### convert json to yaml

```sh
ttoy -i tt.json
```

a tui appers, and select "converter" -> "yaml"

> ttoy will print the yaml result on stdin, if you pass -o flag, the result will be written into the output file.

```sh
cat tt.json | ttoy
```



### generate svg graph from json

```sh
ttoy -i tt.json -o tt.svg
```

select "generators" -> "json --> svg".