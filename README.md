# signage
Sign and Verify binaries in your $PATH for end-point OpSec, written in Go. Uses SHA-256.

Install:
1. install golang on your machine
2. git clone https://github.com/sysdizzy/signage.git to your preffered directory
3. cd signage && mkdir signatures logs
4. touch signatures/signed.xml logs/diff.log
5. go build signage.go (OR run from source code) go run signage.go

WARNING: Sudo or root access is required since it makes hashes for the sudo command.

Usage: 

signage sign (signs all binaries in $PATH and adds keys to ~/signage/signed.xml)

signage verify (Checks hash of all binaries and compares them to signed.xml, discrepancies added to ~/signage/diff.log)


tested on macOS, FreeBSD, and NetBSD
