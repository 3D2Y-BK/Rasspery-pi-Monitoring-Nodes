package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ANSI terminal colors
const (
	Reset      = "\033[0m"
	Bold       = "\033[1m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Magenta    = "\033[35m"
	Cyan       = "\033[36m"
	White      = "\033[37m"
	BgRed      = "\033[41m"
	BgGreen    = "\033[42m"
	BgBlack    = "\033[40m"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	printBanner()
	time.Sleep(1 * time.Second)

	// Phase 1: Reconnaissance (Nmap scan)
	runRecon()

	// Phase 2: LLMNR Poisoning & Relay
	runResponder()

	// Phase 3: Zerologon Attack on Domain Controller
	runZerologon()

	// Phase 4: Golden Ticket & LSASS Extraction
	runDomainTakeover()
}

func printBanner() {
	fmt.Printf("%s%s", Bold, Cyan)
	fmt.Println("======================================================================")
	fmt.Println("              SENTINEL-X SURVEILLANCE TAKEOVER EMULATOR v4.2")
	fmt.Println("======================================================================")
	fmt.Printf("%s%s[i] Target Domain:   %ssentinel.net (AD_LAB_DOMAIN_v4)%s\n", Reset, Bold, White, Reset)
	fmt.Printf("%s[i] Audit Interface: %seth0 (192.168.1.100)%s\n", Bold, White, Reset)
	fmt.Printf("%s[i] Current Mode:    %sACTIVE PENETRATION SURVEILLANCE (SAFE)%s\n\n", Bold, Yellow, Reset)
}

func showSpinner(prefix string, duration time.Duration) {
	chars := []string{"|", "/", "-", "\\"}
	start := time.Now()
	i := 0
	for time.Since(start) < duration {
		fmt.Printf("\r%s %s%s%s ", prefix, Cyan, chars[i%len(chars)], Reset)
		time.Sleep(100 * time.Millisecond)
		i++
	}
	fmt.Printf("\r%s %s[DONE]%s\n", prefix, Green, Reset)
}

func showProgressBar(prefix string, duration time.Duration) {
	totalSteps := 40
	sleepInterval := duration / time.Duration(totalSteps)
	for i := 0; i <= totalSteps; i++ {
		pct := (i * 100) / totalSteps
		bar := ""
		for j := 0; j < totalSteps; j++ {
			if j < i {
				bar += "█"
			} else {
				bar += "░"
			}
		}
		fmt.Printf("\r%s [%s%s%s%s] %d%%", prefix, Cyan, bar, Reset, Reset, pct)
		time.Sleep(sleepInterval)
	}
	fmt.Println()
}

func runRecon() {
	fmt.Printf("%s[*] PHASE 1: RECONNAISSANCE & PORT MAPPING%s\n", Bold+Blue, Reset)
	time.Sleep(500 * time.Millisecond)

	showProgressBar("Scanning Subnet [192.168.1.0/24]", 2*time.Second)
	fmt.Printf("%s[+] Active Host Discovered: 192.168.1.10   (DC-AD-01.sentinel.net)%s\n", Green, Reset)
	fmt.Printf("%s[+] Active Host Discovered: 192.168.1.50   (CL-WIN-10.sentinel.net)%s\n\n", Green, Reset)

	time.Sleep(800 * time.Millisecond)
	fmt.Printf("%s[i] Running detailed port probe on hosts...%s\n", Bold+White, Reset)
	time.Sleep(500 * time.Millisecond)

	// DC-AD-01 ports
	fmt.Printf("\n%sPORT SCAN REPORT FOR DC-AD-01 (192.168.1.10):%s\n", Bold, Reset)
	portsDC := []string{
		"53/tcp    open  domain      (Microsoft DNS)",
		"88/tcp    open  kerberos-sec (Microsoft Windows Kerberos)",
		"135/tcp   open  msrpc       (Microsoft Windows RPC)",
		"389/tcp   open  ldap        (Microsoft Windows Active Directory)",
		"445/tcp   open  microsoft-ds (SMBv2/v3 Signing Enabled)",
		"3268/tcp  open  ldap        (Global Catalog)",
	}
	for _, p := range portsDC {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("  " + p)
	}

	// CL-WIN-10 ports
	fmt.Printf("\n%sPORT SCAN REPORT FOR CL-WIN-10 (192.168.1.50):%s\n", Bold, Reset)
	portsCL := []string{
		"135/tcp   open  msrpc       (Microsoft Windows RPC)",
		"445/tcp   open  microsoft-ds (SMBv2/v3 Signing NOT Enforced)",
		"3389/tcp  open  ms-wbt-server (RDP Enabled)",
	}
	for _, p := range portsCL {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("  " + p)
	}
	fmt.Println()
	time.Sleep(1 * time.Second)
}

func runResponder() {
	fmt.Printf("%s[*] PHASE 2: POISONING BROADCASTS & AUTHENTICATION RELAY%s\n", Bold+Blue, Reset)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("%s[+] Launching LLMNR/NBT-NS Spoofer (Responder mode)...%s\n", Green, Reset)
	time.Sleep(800 * time.Millisecond)

	fmt.Printf("%s[i] Listening for LLMNR name resolution queries...%s\n", Cyan, Reset)
	time.Sleep(1500 * time.Millisecond)

	// Poisoning trigger
	fmt.Printf("%s[!] Poisoned query: CL-WIN-10 requested 'shares.local' (LLMNR captured)%s\n", Yellow, Reset)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%s[+] Poisoned answer sent. Redirecting CL-WIN-10 authentication attempt to our relay.%s\n", Green, Reset)
	time.Sleep(1 * time.Second)

	fmt.Printf("%s[+] Intercepting incoming NetNTLMv2 Challenge authentication from 192.168.1.50:%s\n", Green, Reset)
	fmt.Println("    User:      sentinel\\mahdi")
	fmt.Println("    Challenge: 1122334455667788")
	fmt.Println("    Response:  a1b2c3d4e5f6g7h8a1b2c3d4e5f6g7h8...")
	time.Sleep(1 * time.Second)

	fmt.Printf("\n%s[i] Relaying credentials to DC-AD-01 (192.168.1.10)...%s\n", Cyan, Reset)
	time.Sleep(800 * time.Millisecond)
	fmt.Printf("%s[-] Relaying failed on DC-AD-01: SMB Signing Enforced (Status Access Denied)%s\n", Red, Reset)

	time.Sleep(800 * time.Millisecond)
	fmt.Printf("%s[i] Relaying credentials to CL-WIN-10 (192.168.1.50)...%s\n", Cyan, Reset)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s[+] Relaying SUCCESSFUL! SMB Signing Not Enforced on CL-WIN-10.%s\n", Green, Reset)
	fmt.Printf("%s[+] Logged in: Administrative Session spawned on CL-WIN-10 (192.168.1.50).%s\n", Green, Reset)
	time.Sleep(500 * time.Millisecond)

	showSpinner("Dumping local SAM database on CL-WIN-10", 2*time.Second)
	fmt.Printf("%s[+] Dumped SAM hashes: %s\n", Green, Reset)
	fmt.Println("    Administrator:500:aad3b435b51404eeaad3b435b51404ee:31d6cfe0d16ae931b73c59d7e0c089c0 (NTLM)")
	fmt.Println("    mahdi:1001:aad3b435b51404eeaad3b435b51404ee:8846f7eaee2d9333a3c59d7e0c089c03 (NTLM)")
	fmt.Println()
	time.Sleep(1200 * time.Millisecond)
}

func runZerologon() {
	fmt.Printf("%s[*] PHASE 3: ZEROLOGON NETLOGON BYPASS (CVE-2020-1472)%s\n", Bold+Blue, Reset)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("%s[i] Launching Zerologon handshake exploit against DC-AD-01 (192.168.1.10)...%s\n", Cyan, Reset)
	time.Sleep(1 * time.Second)

	// Simulation loop showing brute force attempts
	totalAttempts := 256
	successfulAttempt := 84
	fmt.Printf("Brute forcing Netlogon authentication bypass:\n")
	for k := 0; k <= successfulAttempt; k += 4 {
		fmt.Printf("\r  -> Testing spoofed Netlogon tokens: %d/%d...", k, totalAttempts)
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Printf("\r  -> Testing spoofed Netlogon tokens: %d/%d... %s[SUCCESS]%s\n", successfulAttempt, totalAttempts, Green, Reset)
	time.Sleep(600 * time.Millisecond)

	fmt.Printf("%s[+] Cryptographic bypass matched! Null challenge accepted by Domain Controller.%s\n", Green, Reset)
	time.Sleep(800 * time.Millisecond)

	fmt.Printf("%s[i] Resetting computer account password of DC-AD-01$ to blank...%s\n", Cyan, Reset)
	time.Sleep(1200 * time.Millisecond)
	fmt.Printf("%s[+] Account 'DC-AD-01$' password changed successfully to clean NT hash (31d6cfe0d16a...)%s\n", Green, Reset)
	time.Sleep(800 * time.Millisecond)

	// secretsdump
	fmt.Printf("\n%s[i] Launching Impacket secretsdump to extract Active Directory database...%s\n", Cyan, Reset)
	showProgressBar("Dumping NTDS.dit via DRSUAPI", 2500*time.Millisecond)

	fmt.Printf("\n%sEXTRACTED DOMAIN HASHES (sentinel.net):%s\n", Bold, Reset)
	hashes := []struct {
		User string
		Rid  int
		Hash string
	}{
		{"Administrator", 500, "2b576ac214f923b73c59d7e0c089c03d"},
		{"krbtgt", 502, "a9e52579dfd93b33a3c59d7e0c089c03f"},
		{"mahdi", 1104, "8846f7eaee2d9333a3c59d7e0c089c033"},
		{"DC-AD-01$", 1000, "31d6cfe0d16ae931b73c59d7e0c089c0"}, // Empty password hash
	}
	for _, h := range hashes {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("  %s%s:%d:aad3b435b51404eeaad3b435b51404ee:%s%s\n", White, h.User, h.Rid, h.Hash, Reset)
	}
	fmt.Println()
	time.Sleep(1 * time.Second)
}

func runDomainTakeover() {
	fmt.Printf("%s[*] PHASE 4: GOLDEN TICKET FORGERY & FULL DOMAIN TAKEOVER%s\n", Bold+Blue, Reset)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("%s[i] Forging Kerberos Golden Ticket for domain administrator...%s\n", Cyan, Reset)
	time.Sleep(800 * time.Millisecond)
	fmt.Println("    Domain:     sentinel.net")
	fmt.Println("    Domain SID: S-1-5-21-3623819074-94610234-29713007")
	fmt.Println("    User:       Administrator")
	fmt.Println("    Group RIDs: 513, 512, 520, 518, 519 (Domain Admins)")
	fmt.Printf("    krbtgt Key: a9e52579dfd93b33a3c59d7e0c089c03f\n")
	time.Sleep(1200 * time.Millisecond)

	fmt.Printf("%s[+] Golden Ticket generated successfully! Saving ticket to cache memory.%s\n", Green, Reset)
	time.Sleep(600 * time.Millisecond)

	fmt.Printf("%s[i] Injecting Golden Ticket in memory (Pass-the-Ticket)...%s\n", Cyan, Reset)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s[+] Ticket injected successfully! Local session privilege upgraded to DOMAIN ADMIN clearance.%s\n", Green, Reset)
	time.Sleep(1 * time.Second)

	fmt.Printf("%s[i] Connecting to DC-AD-01 (192.168.1.10) using PsExec via Ticket...%s\n", Cyan, Reset)
	showSpinner("Establishing admin shell connection", 2*time.Second)

	// Simulated Interactive shell takeover success message
	fmt.Println()
	fmt.Printf("%s%s", BgGreen, BgBlack)
	fmt.Println("======================================================================")
	fmt.Println("     DOMAIN TAKEOVER COMPLETE: COMPROMISED DC-AD-01 (DOMAIN ADMIN)   ")
	fmt.Println("======================================================================")
	fmt.Print(Reset)
	fmt.Printf("\n%sC:\\Windows\\system32> %swhoami /groups%s\n", Bold+White, Reset, Reset)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  Group Name                                 Type             SID")
	fmt.Println("  ========================================== ================ =========================================")
	fmt.Printf("  %ssentinel\\Domain Admins                   Group            S-1-5-21-3623819074-94610234-29713007-512%s\n", Green, Reset)
	fmt.Println("  sentinel\\Schema Admins                   Group            S-1-5-21-3623819074-94610234-29713007-518")
	fmt.Println("  sentinel\\Enterprise Admins               Group            S-1-5-21-3623819074-94610234-29713007-519")
	fmt.Println("  BUILTIN\\Administrators                    Alias            S-1-5-32-544")
	fmt.Println()
	time.Sleep(1 * time.Second)
	fmt.Printf("%sC:\\Windows\\system32> %s_ %s(Active AD Domain Takeover established)%s\n\n", Bold+White, Reset, Yellow, Reset)
}
