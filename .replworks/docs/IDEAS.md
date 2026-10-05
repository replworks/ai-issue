# AI Issue Publisher

AI와의 대화에서 나온 아이디어를 사람의 작업과 분리된 GitHub Issue로 등록하는 CLI 도구.

---

# 문제

AI(ChatGPT, Claude, Gemini 등)와 협업하다 보면 다음과 같은 결과물이 자주 나온다.

* 버그 발견
* 개선 아이디어
* 리팩토링 제안
* 신규 기능 제안
* 아키텍처 제안
* 기술 부채 정리
* 작업 계획

좋은 내용이 많지만 대부분 GitHub에 등록되지 않는다.

등록 과정이 번거롭기 때문이다.

---

하지만 더 중요한 문제가 있다.

AI가 제안한 내용을 GitHub Issue로 등록하면 일반적으로 Author는 사람이다.

예:

```text
opened by alice
```

그러면 GitHub 사용자들은 자연스럽게 다음을 기대한다.

```text
alice가 내용을 이해한다.
alice가 설명할 수 있다.
alice가 수정할 수 있다.
alice가 책임진다.
```

하지만 AI와의 대화에서 생성된 이슈는 다르다.

사람은 단지

```text
이 이슈를 백로그에 남길 가치가 있다.
```

고 판단했을 뿐일 수 있다.

반드시 내용을 완전히 이해하거나 책임질 수 있는 것은 아니다.

---

# 핵심 문제

GitHub에서 Author는 단순한 작성자가 아니다.

Author는 암묵적으로 책임 주체로 인식된다.

AI가 생성한 이슈를 사람 이름으로 등록하면

```text
생성 주체
=
책임 주체
```

로 오해된다.

---

# 목표

AI가 생성한 작업과 사람이 생성한 작업을 GitHub 상에서 명확히 분리한다.

예:

```text
Human Issue

opened by alice
```

---

```text
AI Issue

opened by ai-backlog-bot
```

---

사람은

```text
Issue를 생성할지 결정한다.
```

AI는

```text
Issue 내용을 생성한다.
```

이 둘을 구분한다.

---

# 핵심 철학

이 프로젝트는

```text
누가 등록했는가
```

보다

```text
누가 생성했는가
```

를 중요하게 생각한다.

---

사람이 한 일

```text
AI 결과를 검토했다.
Issue로 남길 가치가 있다고 판단했다.
```

---

AI가 한 일

```text
문제를 발견했다.
아이디어를 제안했다.
Issue 초안을 작성했다.
```

---

따라서 AI가 생성한 Issue는 AI Author로 등록한다.

---

# 사용 예시

AI와 대화

```text
현재 로깅 시스템의 문제점을 분석해줘
```

↓

AI 응답

```markdown
# Add timestamps to logging system

Current logs do not contain timestamps...
```

↓

복사

↓

CLI 실행

```bash
ai-issue
```

↓

Issue 생성

```text
opened by ai-backlog-bot
```

---

# MVP

## CLI

```bash
ai-issue
```

---

## Clipboard 자동 읽기

클립보드 내용을 자동으로 가져온다.

```bash
Cmd+C
```

↓

```bash
ai-issue
```

---

## Repository 자동 감지

현재 Git Repository를 자동 인식한다.

```bash
git remote get-url origin
```

↓

```text
company/backend
```

---

## 제목 자동 추출

첫 번째 Markdown Header를 제목으로 사용한다.

예:

```markdown
# Add timestamps to logging system
```

↓

```text
Add timestamps to logging system
```

---

## Preview

```text
Repository: company/backend

Author: ai-backlog-bot

Title:
Add timestamps to logging system

Create Issue? (Y/n)
```

---

## AI Author 생성

Issue는 별도 GitHub 계정으로 생성한다.

예:

```text
ai-backlog-bot
```

---

GitHub 상에서는

```text
opened by ai-backlog-bot
```

으로 표시된다.

---

# 왜 별도 Author를 사용하는가

라벨은 충분하지 않다.

예:

```text
opened by alice
labels: ai
```

여전히 사람의 작업처럼 보인다.

---

반면

```text
opened by ai-backlog-bot
```

은 GitHub 전체 UI에서 의미가 명확하다.

* Issue List
* Activity
* Notifications
* Search

모든 곳에서

```text
AI가 생성한 작업
```

임을 알 수 있다.

---

# 협업 관점의 장점

AI Author는 심리적 소유권을 최소화한다.

예:

```text
opened by alice
```

↓

수정하기 부담스러움

닫기 부담스러움

---

```text
opened by ai-backlog-bot
```

↓

수정 가능

닫기 가능

삭제 가능

---

AI 이슈는 초안으로 취급된다.

---

# 아키텍처

```text
Clipboard
↓
AI Issue Publisher
↓
GitHub API
↓
ai-backlog-bot
```

---

## 인증

MVP에서는 단일 AI 계정을 사용한다.

예:

```text
ai-backlog-bot
```

---

향후 확장 가능

```text
chatgpt-bot
claude-bot
gemini-bot
```

하지만 기본 철학은

```text
Human
vs
AI
```

분리이다.

---

# 비목표

현재 버전에서는 하지 않는다.

* AI 자동 요약
* PR 생성
* Commit 생성
* Slack 연동
* Jira 연동
* AI 성과 분석
* 프로젝트 관리 기능

---

# 성공 조건

* AI 대화 결과가 GitHub에 누락되지 않는다.
* 사람이 만든 이슈와 AI가 만든 이슈가 명확히 구분된다.
* Issue 생성 시간이 1분 이내다.
* AI 이슈를 부담 없이 수정하거나 폐기할 수 있다.

---

# 프로젝트 철학

AI는 제안한다.

사람은 등록 여부를 결정한다.

하지만 그 둘은 동일한 역할이 아니다.

GitHub는 오랫동안 사람 중심의 협업 도구였다.

AI Issue Publisher는

AI가 생성한 작업을

사람의 책임 체계와 분리된 상태로

GitHub 백로그에 기록한다.
