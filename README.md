# AutoCTM
Automated Certificate Transparency Monitor written in GO 


## What is Certificate Transparency? 

Certificate Transparency (CT) is a system designed to make TLS/SSL certificate issuance public, auditable, and verifiable. It helps prevent mis-issued or rogue certificates and increases trust in HTTPS.

## The Problem

Certificate Authorities (CAs) issue TLS/SSL certificates for domains. Sometimes, a CA can mistakenly issue a certificate to the wrong party, or a malicious actor could obtain one. Without transparency, domain owners might not know about these certificates until it’s too late.

In short

- CAs can mistakenly or maliciously issue certificates for domains they shouldn't
- Domain owners may have no visibility into certificates issued in their name
- The window of unawareness is the dangerous part. An attacker could use a fraudulent certificate before anyone notices

## How CT works

Certificate Transparency is a system that requires all TLS certificates to be publicly logged, allowing domain owners and browsers to detect and reject unauthorized certificates.