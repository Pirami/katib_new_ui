# AutoDL UI 가이드

위 AutoDL UI는 기존 katib UI의 new-ui를 기반으로 Accuinsight+ 내제화시킨 것이다.

AutoDL UI 의 frontend and backend를 각각 실행하는 방법을 아래와 같이 참고하기.

frontend는 Angular프레임워크와 typescript 언어로 구성되어있다.
backend는 go 언어로 구성되어있다.

참고로 AutoDL은 go 언어 기반으로 구축되어있다.

## AutoDL frontend
먼저, frontend의 위치는 pkg/new-ui/v1beta1/frontend에 있다.

frontend를 실행시키기 위한 환경셋팅
node -v: v12.18.1
npm -v: 6.14.5
nvm -v: 0.39.1

frontend 개발 후에 npm run build:prod 를 실행시켜서 frontend/dist/static 구성이 만들어지게 한다.

## AutoDL backend
먼저, backend의 위치는 cmd/new-ui/v1beta1에 있다.

backend를 실행시키기 위한 환경셋팅
(mac 기준) go version: go1.18.3 darwin/amd64

실행 방법
go run main.go --build-dir=/frontend/dist --port=포트번호 --db-manager-address=주소

예시>
go run main.go --build-dir=/frontend/dist --port=8081 --db-manager-address=20.196.232.2:6789

*실행시키기 위한 옵션 설명
Run `main.go` file with appropriate flags, where:

   - `--build-dir` - directory with the frontend artifacts.
   - `--port` - port to access Katib UI.
   - `--db-manager-address` - Katib DB manager external IP and port address.

   For example, if you use port-forwarding to expose `katib-db-manager`, run this command:

   ```
   go run main.go --build-dir=../../../pkg/new-ui/v1beta1/frontend/dist --port=8080 --db-manager-address=localhost:6789
   ```
   After that, you can access the UI using this URL: `http://localhost:8080/katib/`.*

## AutoDL UI image 빌드하는 방법

docker build . -f cmd/new-ui/v1beta1/Dockerfile -t <name of your image>
