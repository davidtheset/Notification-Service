# Notification-Service
## Exercise Description

We have a Notification system that sends out email notifications of various types (supdatesupdate, daily news, project invitations, etc). We need to protect recipients from getting too many emails, either due to system errors or due to abuse, so let's limit the number of emails sent to them by implementing a rate-limited version of NotificationService. The system must reject requests that are over the limit.

Some sample notification types and rate limit rules, e.g.:
- Status: not more than 2 per minute for each recipient
- News: not more than 1 per day for each recipient
- Marketing: not more than 3 per hour for each recipient
- Etc. these are just samples, the system might have several rate limit rules!

NOTES:

- Your solution will be evaluated on code quality, clarity and development best practices.
- Feel free to use the programming language, frameworks, technologies, etc that you feel more comfortable with.
- Below you'll find a code snippet that can serve as a guidance of one of the implementation alternatives in Java. Feel free to use it if you find it useful or ignore it otherwise; it is not required to use it at all nor translate this code to your programming language.


The system should be able to:

1. Receive requests to send notifications.
2. Check if a recipient has reached the rate limit for the requested notification type.
3. Send the notification if the rate limit is not exceeded, or reject it if the rate limit is exceeded.

Languaje and framework:
The exercise includes a base implementation in Go using the Gin framework and unit tests.

## Running Tests

To run the unit tests, use the following command:

```bash
go test ./test
