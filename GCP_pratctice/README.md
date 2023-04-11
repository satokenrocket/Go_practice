## Dockerの立ち上げとローカル環境での疎通確認
# dockerの立ち上げと立ち上げたMysqlの接続
1. docker-compose up -d
2. mysql -utest_user -h 127.0.0.1 --port 13306 -p

# mysqlの構築
1. source ファイル名(ファイルパス);

# Dockerの立ち上げとローカル環境での疎通確認
1. go build
1. go run

# Docker build でタグをつけてイメージ作成
docker build . -t asia-northeast1-docker.pkg.dev/sage-shard-380406/test/go-test:latest
docker push asia-northeast1-docker.pkg.dev/sage-shard-380406/test/go-test:latest
呪文内容
https://cloud.google.com/artifact-registry/docs/docker/pushing-and-pulling?hl=ja

# タグ付きイメージをContainer Registry に push 


GCP  URL
docker run asia-northeast1-docker.pkg.dev/sage-shard-380406/test/go-test:latest


参考
https://my0shym.hatenablog.com/entry/2022/10/16/172507

