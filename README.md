# signage
Sign and Verify binaries in your $PATH for end-point OpSec.

Usage: 
signage sign (signs all binaries in $PATH and adds keys to ~/signage/signed.xml)

signage verify (Checks hash of all binaries and compares them to signed.xml, discrepancies added to ~/signage/diff.log)

WARNING: Sudo or root access is required since it makes hashes for the sudo command.

tested on macOS, FreeBSD, and NetBSD
