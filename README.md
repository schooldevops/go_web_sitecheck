# 시스템 점검중 

- port: 8888
- 환경변수 파일: /bin/envfile
- 메인페이지 파일: /bin/index.html
- 실행파일: /bin/demo

## docker build

```
docker build -t scale_check:v1.0 .
```

## docker run 

```
docker run --name scale_check -p 8888:8888 -v /var/envfile:/bin/envfile scale_check
```
