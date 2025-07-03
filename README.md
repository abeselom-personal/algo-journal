# Algo Journal Template

Track your daily algorithm practice by pushing solutions to GitHub with automated reminders via Telegram and Email.

---

## Repo Structure

```

/leetcode
/easy
/medium
/hard
/codeforces
/div2
/div3
/notes

````

---

## How to Use

1. Solve problems and add code files in appropriate folders.
2. Commit and push daily with meaningful commit messages.
3. Use the provided `commit.sh` script for quick commits.
4. Automated reminders will notify you via Telegram and Email every day.

---

## Automation

- GitHub Action runs daily at 8 AM UTC (~11 AM Ethiopia time).
- Sends a Telegram message to remind you.
- Sends an email reminder through Gmail SMTP.

---

## Environment Variables

Create a `.env` file (or set GitHub Secrets) with:

```env
# Telegram
TELEGRAM_BOT_TOKEN=your_telegram_bot_token_here
TELEGRAM_CHAT_ID=your_telegram_chat_id_here

# Gmail (App Password recommended)
GMAIL_PASSKEY=your_gmail_app_password
EMAIL_ADDRESS=your_email@example.com

# Optional (for future use)
NEXT_PUBLIC_GTM=your_google_tag_manager_id
NEXT_PUBLIC_APP_URL=your_app_url
````

---

## commit.sh

```bash
#!/bin/bash
msg="Update: $(date '+%Y-%m-%d')"
git add .
git commit -m "$msg"
git push
```

Make executable with `chmod +x commit.sh`.

Run daily to commit and push your changes quickly.

---

## Notes

* Use Gmail App Password for `GMAIL_PASSKEY` (2FA enabled accounts).
* Telegram Chat ID can be found by sending a message to your bot and checking `https://api.telegram.org/bot<YourBotToken>/getUpdates`.
* Customize folders and scripts as needed.

````

---

```env
# env.example

# Telegram bot token from BotFather
TELEGRAM_BOT_TOKEN=

# Your Telegram chat ID (where bot sends messages)
TELEGRAM_CHAT_ID=

# Gmail SMTP app password (not your Gmail password)
GMAIL_PASSKEY=

# Gmail email address used to send reminders
EMAIL_ADDRESS=

````
