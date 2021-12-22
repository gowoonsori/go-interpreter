# Go를 이용하여 Interpreter 만들기

## 파일 구조
- token : 어휘 분석을 위한 최소 단위의 토큰들
- lexer : 어휘 분석 (lexical analysis)
- repl : 사용자의 입력물을 토큰으로 분석한 결과물 출력
- ast : 추상 구문 트리
  - ast.go : AST를 위한 interface와 부모노드인 Program 정의
  - statement.go : 구문들 정의한 파일로 (let 구문, return 구문, 표현식 등등)
  - expression.go : 표현식을 정의한 파일로 변수에 대한 정보들
- parse : AST를 만들기 위한 파서