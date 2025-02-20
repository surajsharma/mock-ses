# AWS SES Mock Api

## Rules and Limitations
### Email Warming Up Process

- [ ] New accounts start in sandbox mode
- [ ] Daily sending quota: 200 emails/24 hours in sandbox
- [ ] Maximum send rate: 1 email/second in sandbox
- [ ] Production limits:
  - [ ] Week 1: 10,000 emails/24 hours
  - [ ] Week 2: 50,000 emails/24 hours
  - [ ] After review: Up to 100,000+ emails/24 hours

### Verification Requirements
- [ ] Sender email addresses must be verified
- [ ] In sandbox, recipient emails must also be verified
- [ ] Domain verification required for production

## Rate Limiting
- [ ] Soft limit: Warning when approaching quota
- [ ] Hard limit: Rejection when quota exceeded

---

## AWS SES Error Codes

- `MessageRejected`: Generic rejection
- `MailFromDomainNotVerified`: Sender domain not verified
- `MaxSendingRateExceeded`: Rate limit exceeded
- `DailyQuotaExceeded`: Daily quota exceeded
- `AccountSuspended`: Account suspended
- `ValidationError`: Invalid input parameters
- `EmailNotVerified`: Email address not verified
- `ProductionAccessNotGranted`: Still in sandbox mode

---

## Project Structure 

```
mock-ses/
│── main.go
│── handlers/
│   ├── send_email.go
│   ├── stats.go
│── models/
│   ├── email.go
│   ├── stats.go
│── middleware/
│   ├── rate_limit.go
│── config/
│   ├── settings.go
│── storage/
│   ├── memory_store.go
│── go.mod
│── go.sum
```