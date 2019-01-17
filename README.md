# ssh-mysql-sample
踏み台サーバにSSHしてDBにリクエストを投げるサンプル。
前提として、アクセス先のパラメータをtomlファイルに指定してあるものとします。

## USAGE
- default  
`go run ./cmd/app/main.go`

- パラメータ指定  
`go run cmd/app/main.go -s {id} -e {env}`

| parameter | description|
|:----|:---- |
|id| store_id|
|env| environment<br>- dev<br>- stg<br>- prd|
||
