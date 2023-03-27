# go-transpose-csv

既存のCSVから特定の列を抽出して、転置したcsvを生成するスクリプト

例:
```csv:input.csv
id,content
a,apple
b,baby
```

```csv:output.csv
id,a
content,apple
id,b
content,b
```

