# Hello API

This is an improved version of the current hello-api we use in production. It will use less resources and scale better.

## Dependencies

- Go version 1.21

## Setup

Development expects to run in a unix-like environment. If you are using Windows, please consider following these directions [here](https://docs.microsoft.com/en-us/windows/wsl/install-win10) to install the Windows Subsystem for Linux (WSL).

### Install Go (1.21)

```bash
sudo make setup
```

### Upgrade Go (1.21)

```bash
sudo make install-go
```

## Release Milestones

### V0 (1 day)

- [ ] Onboarding documentation
- [ ] Simple API response (Hello World)
- [ ] unit tests
- [ ] running somewhere other than the developer's machine

### V1 (7 days)

- [ ] Create translation endpoint
- [ ] Store translations in short-tern storage
- [ ] Call existing service to translate
- [ ] Move towards long-term storage
