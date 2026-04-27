# AutoCTM
Automated Certificate Transparency Monitor written in GO.


## What is Certificate Transparency? 

Certificate Transparency (CT) is a system designed to make TLS/SSL certificate issuance public, auditable, and verifiable. It helps prevent mis-issued or rogue certificates and increases trust in HTTPS.

## The Problem

Certificate Authorities (CAs) issue TLS/SSL certificates for domains. Sometimes, a CA can mistakenly issue a certificate to the wrong party, or a malicious actor could obtain one. Without transparency, domain owners might not know about these certificates until it’s too late.

In short

- CAs can mistakenly or maliciously issue certificates for domains they shouldn't
- Domain owners may have no visibility into certificates issued in their name
- The window of unawareness is the dangerous part. An attacker could use a fraudulent certificate before anyone notices

## How CT works

Certificate Transparency solves the problem of rogue certificate issuance by requiring every CA to submit certificates to public, append-only logs before browsers will trust them. This means anyone can monitor these logs to spot unauthorized certificates for their domain, and browsers will outright reject any certificate that lacks proof of being logged.

# Disclaimer
AutoCTM is provided as-is, without warranty of any kind. It is intended as a monitoring aid only and may not detect all certificates issued for your domains. Do not rely solely on this tool for security decisions. The authors accept no liability for any damages arising from its use or from missed detections.
