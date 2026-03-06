# Security Policy

## Supported Versions

We release patches for security vulnerabilities for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 1.3.x   | :white_check_mark: |
| 1.2.x   | :x:                |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of CloudWeGo Template seriously. If you have discovered a security vulnerability, we appreciate your help in disclosing it to us in a responsible manner.

### Please do NOT:

- Open a public GitHub issue
- Disclose the vulnerability publicly
- Exploit the vulnerability

### Please DO:

**Report security vulnerabilities to:**

📧 **Email**: security@example.com

Or create a private security advisory on GitHub:
https://github.com/yi-nology/cloudwego-template/security/advisories/new

### What to include:

Please include the following information in your report:

1. **Description** of the vulnerability
2. **Steps to reproduce** the issue
3. **Possible impact** of the vulnerability
4. **Suggested fix** (if any)
5. **Your name/handle** (optional, for credit)

### Our commitment:

- ⏱️ We will acknowledge receipt of your report within **48 hours**
- 🔍 We will investigate and provide an initial assessment within **7 days**
- 🛠️ We will release a fix as soon as possible, depending on complexity
- 📢 We will publicly disclose the issue after the fix is released
- 🙏 We will credit you in our security advisories (unless you prefer to remain anonymous)

## Security Best Practices

When using CloudWeGo Template:

1. **Keep dependencies up to date**
   ```bash
   go mod tidy
   go get -u
   ```

2. **Use environment variables for sensitive data**
   ```yaml
   database:
     password: "${DB_PASSWORD}"
   redis:
     password: "${REDIS_PASSWORD}"
   ```

3. **Enable security scanning**
   ```bash
   go install github.com/securego/gosec/v2/cmd/gosec@latest
   gosec ./...
   ```

4. **Review configuration files**
   - Never commit secrets to version control
   - Use `.gitignore` for sensitive files

5. **Keep Go version updated**
   ```bash
   # Check for updates
   go version
   ```

## Known Security Issues

No known security vulnerabilities in the current version.

## Security Updates

Security updates will be announced via:

- [GitHub Security Advisories](https://github.com/yi-nology/cloudwego-template/security/advisories)
- [GitHub Releases](https://github.com/yi-nology/cloudwego-template/releases)
- [CHANGELOG.md](CHANGELOG.md)

## Contact

For general security questions:

- 💬 [GitHub Discussions](https://github.com/yi-nology/cloudwego-template/discussions)
- 📧 Email: security@example.com

---

**Thank you for helping keep CloudWeGo Template secure!** 🔒
