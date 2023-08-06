# signage
Sign and Verify binaries in your $PATH for end-point OpSec, written in Go. Uses SHA-256.

Install:
1. install golang on your machine
2. mkdir ~/signage
3. touch ~/signage/signed.xml ~/signage/diff.log
4. git clone https://github.com/sysdizzy/signage to your preffered directory
5. go build src/* OR run from source code (go run src/signage.go) from the cloned directory.

WARNING: Sudo or root access is required since it makes hashes for the sudo command.

Usage: 

signage sign (signs all binaries in $PATH and adds keys to ~/signage/signed.xml)

signage verify (Checks hash of all binaries and compares them to signed.xml, discrepancies added to ~/signage/diff.log)


tested on macOS, FreeBSD, and NetBSD
