# can we use javascript to create json for golang consumption?
It'd be nice to create the json in regular `commonJS` style..

This is just experimenting with https://github.com/dop251/goja (native golang javascript engine) to see how far we can go

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
![image](https://user-images.githubusercontent.com/462087/198570020-d9aed9e9-1c9f-4d6c-925b-1ed102452845.png)
