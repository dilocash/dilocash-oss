#!/bin/bash
# ./scripts/ai-context.sh | pbcopy  # Copies everything to your clipboard
# Project: dilocash-oss Context Exporter
# Purpose: Bundles project "truth" for AI coding assistants.

echo "=== SYSTEM DIRECTIVE ==="
echo "You are an expert Go and TypeScript engineer working on 'dilocash-oss'."
echo "You must strictly follow the Clean Architecture and ADRs provided below."
echo "Always prioritize 'shopspring/decimal' for money and Protobuf for API contracts."
echo ""

echo "=== ARCHITECTURE DECISION RECORDS ==="
cat docs/adr/*.md | grep -v "TEMPLATE.md"
echo ""

echo "=== API CONTRACT (Source of Truth) ==="
cat proto/v1/dilocash.proto
echo ""

echo "=== DATABASE SCHEMA ==="
cat apps/api/db/schema.sql
echo ""

echo "=== PROJECT RULES ==="
cat .cursorrules 2>/dev/null || echo "No .cursorrules found."