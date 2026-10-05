# PITCHING_SCRIPT.md

# AI Issue Publisher

AI와 대화하다 보면 좋은 아이디어가 정말 많이 나온다.

문제는 대부분 GitHub에 등록되지 않는다는 것이다.

하지만 내가 만들고 싶은 프로젝트는 단순히

```text
ChatGPT → GitHub Issue
```

도구가 아니다.

---

## 내가 발견한 문제

AI와 대화하다 보면 이런 일이 자주 발생한다.

```text
"좋은 아이디어네."
"나중에 작업하자."
"이건 이슈로 남겨야겠다."
```

그리고 결국 잊어버린다.

그래서 많은 사람들이 말한다.

```text
AI 결과를 GitHub Issue로 쉽게 등록하자.
```

하지만 나는 조금 다른 문제를 발견했다.

---

## GitHub의 Author는 단순한 작성자가 아니다

GitHub에서

```text
opened by alice
```

라고 적혀 있으면

우리는 자연스럽게 생각한다.

```text
alice가 작성했다.
alice가 이해하고 있다.
alice가 설명할 수 있다.
alice가 수정할 수 있다.
```

즉 Author는 책임 주체다.

---

## 그런데 AI가 생성한 이슈는 다르다

예를 들어

```text
나:
현재 아키텍처를 리뷰해줘.

AI:
CQRS로 분리하는 것이 좋겠습니다.
...
```

라는 대화를 했다고 하자.

나는 단지

```text
이건 나중에 검토할 가치가 있다.
```

고 판단했을 뿐이다.

반드시

```text
나는 이 내용을 완전히 이해한다.
```

또는

```text
나는 이 내용에 책임진다.
```

는 아니다.

---

## 그런데 왜 Author는 항상 사람일까?

현재 대부분의 도구는

```text
opened by alice
```

를 만든다.

하지만 실제 생성자는 AI다.

그 결과

```text
생성 주체
=
책임 주체
```

가 되어 버린다.

나는 이것이 AI 시대의 협업 모델과 맞지 않는다고 생각한다.

---

## AI Issue Publisher

AI Issue Publisher는

AI와의 대화 결과를

사람이 아닌

별도의 AI Author로 GitHub에 등록한다.

예:

```text
opened by ai-backlog-bot
```

---

사람이 한 일은

```text
이 이슈를 저장할 가치가 있다고 판단했다.
```

이다.

AI가 한 일은

```text
문제를 발견했다.
아이디어를 제안했다.
Issue 초안을 작성했다.
```

이다.

이 둘은 다른 역할이다.

---

## 왜 중요한가?

AI Author는 심리적 소유권을 제거한다.

예:

```text
opened by alice
```

↓

수정하기 어렵다.

삭제하기 어렵다.

반박하기 어렵다.

---

반면

```text
opened by ai-backlog-bot
```

↓

수정 가능

삭제 가능

폐기 가능

재작성 가능

---

AI Issue는 초안이다.

사람의 주장이나 결정이 아니다.

---

## 워크플로우

AI에게 요청한다.

```text
현재 로깅 시스템 문제점을 분석해줘.
```

AI가 답한다.

```markdown
# Add timestamps to logging system

Current logs do not contain timestamps...
```

복사한다.

그리고 실행한다.

```bash
ai-issue
```

끝.

---

현재 Git Repository를 자동 인식한다.

클립보드를 자동 읽는다.

Issue 제목을 자동 추출한다.

그리고

```text
opened by ai-backlog-bot
```

으로 등록한다.

---

## 이 프로젝트는 생산성 도구가 아니다

사실

```bash
gh issue create
```

로도 이슈는 만들 수 있다.

문제는 속도가 아니다.

---

내가 해결하려는 문제는

```text
AI가 생성한 작업을
사람의 책임 체계와
어떻게 분리할 것인가?
```

이다.

---

## GitHub는 사람만을 위한 공간이 아니게 되었다

이제 우리는

* ChatGPT
* Claude
* Gemini
* Copilot
* Coding Agents

와 함께 일한다.

그런데 GitHub에는 여전히 사람만 Author가 될 수 있다고 가정한다.

AI Issue Publisher는

AI를 GitHub의 1급 참여자로 다룬다.

---

## 철학

AI는 제안한다.

사람은 등록 여부를 결정한다.

하지만 그 둘은 같은 역할이 아니다.

AI Issue Publisher는

AI가 생성한 작업을

사람의 책임 체계와 분리된 상태로

GitHub 백로그에 기록한다.
