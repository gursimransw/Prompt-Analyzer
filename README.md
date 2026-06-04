# Prompt Analyzer API

A lightweight AI Security API written in Go for detecting risky or malicious prompts using rule-based detection, weighted risk scoring, severity mapping, and configurable verdict/action policies.
This project is the foundation of a larger open-source AI Security Platform aimed at helping security teams defend enterprise AI applications against prompt injection, jailbreak attempts, data exfiltration attempts, and other LLM-related abuse patterns.
---

## Current Status

This is an early-stage prototype.
Current version includes:

- Regex-based prompt detection
- Rule-based category matching
- Detection reasons
- Matched rule IDs
- Rule weights
- Risk score calculation
- Severity mapping through policy config
- Verdict/action mapping through policy config
- Duplicate category removal
- Config-driven rules and policies
- HTTP API endpoint for prompt analysis

This is not yet using machine learning, embeddings, or semantic similarity. The current risk score is deterministic and based on matched rule weights.

---
## Why This Project Exists

Most AI applications now rely on LLMs, RAG systems, agents, and tool-calling workflows. These systems introduce new security risks such as:

- Prompt injection
- Jailbreak attempts
- System prompt extraction
- Data exfiltration
- Sensitive information leakage
- Role manipulation
- Tool misuse
- Unsafe model behavior

The goal of this project is to build an open-source AI security engine that can eventually act as a security gateway and detection layer for enterprise AI systems.

---

## Current Architecture

```text
User Prompt
   ↓
HTTP API Endpoint
   ↓
Regex Rule Matching Engine
   ↓
Matched Rules + Categories + Reasons
   ↓
Weight-Based Risk Scoring
   ↓
Policy Config
   ↓
Severity + Verdict
   ↓
JSON Response
