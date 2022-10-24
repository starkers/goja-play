# can we use javascript to create json for golang consumption?
It'd be nice to create the json in regular `commonJS` style..

EG:
```
const foo = {
    waz: "baz"
  }
```

The golang in here should be able to run each `example-*/runme.js` file..

EG:

```bash
for a in example*/runme.js ; do
  go run . --file "${a}"
  echo "============================================"
done
```
