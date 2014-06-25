pingmon
=======

Attempts to ping an ip at an interval, and logs it's success

run it, and it prints it's things to stderr(?). Ping 8.8.8.8 and 10.10.0.1 every
second. You're welcome to hard code different ips.

Needs root. Too lazy to figure out why.

sudo go run ./pingmon.go &> pinglog
