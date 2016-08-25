Content.

```
drwxr-x--- 3 gonzalo  102 ago 25 12:36 MyGame/
-rw-r--r-- 1 gonzalo 2,2K ago 25 14:52 main.go
-rw-r--r-- 1 gonzalo  260 ago 25 15:20 main_test.go
-rwxr-xr-x 1 gonzalo  503 ago 25 12:27 monster.fbs*
-rwxr-xr-x 1 gonzalo  145 ago 25 15:20 monster_in.json*
-rw-r--r-- 1 gonzalo   84 ago 25 15:16 monster_out.bin
-rw-r--r-- 1 gonzalo  160 ago 25 15:16 monster_out.json
-rw-r--r-- 1 gonzalo 1,8K ago 25 15:23 sample_text.cpp.bak
```

- origin: `monster.fbs`
- compile sample text with make, which is the one that parses a .fb file (monster.bs), reads a json file (monster_in.json) and exports it to a another json file (monster_out.json) and to binary (monster_out.bin)
- execute it `../flatsampletext ./ monster.fbs monster_in.json monster_out`
- monster_in.json and monster_out.json should have the same content
- codegen go `flatc++go --go monster.fbs`
- use main_test.go with the codegen go code to read the bytes from monster_out.bin and compare with go code with that the values match the ones in monster_in.json.
- `go test -v .`
- assert test pass
