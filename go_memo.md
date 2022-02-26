# Goの補足事項

## difine
{{difine "layout"}}{{end}}や{{template "content"}}{{end}}などは、Goのテンプレートエンジンの記法
今回は、layoutテンプレートとして定義し、その中で部品となるcontentテンプレートをincludeする。

## Genelate HTMLの中身
template.Must
```
// テンプレートのキャッシュを作成
// template.Must()は独自にエラーチェックを行うため、errorを返り値には持たない
// つまり、ParseFilesがエラーの場合、panicになる
// tempalte.ParseFiles()は可変長引数をとり、その引数として、キャッシュさせたいファイル名を指定
templates := template.Must(template.ParseFiles(files...))
```

template.ExecuteTemplate
```
// difineでテンプレートを定義した場合、ExecuteTemplateでlayoutを明示的に指定する必要がある
tmeplates.ExecuteTemplate(writer, "layout", data)
```