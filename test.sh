#!/bin/bash

# Unified test script for BP repository (Linux/macOS)

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo "Starting BP Test Suite..."
FAILED=0

# 1. Go Tests
echo -e "\n[1/3] Running Go Tests..."
if (cd go && go test ./...); then
    echo -e "${GREEN}Go Tests Passed${NC}"
else
    echo -e "${RED}Go Tests Failed${NC}"
    FAILED=1
fi

# 2. JavaScript Tests
echo -e "\n[2/3] Running JavaScript Tests..."
if (cd js && npm test); then
    echo -e "${GREEN}JS Tests Passed${NC}"
else
    echo -e "${RED}JS Tests Failed${NC}"
    FAILED=1
fi

# 3. PowerShell Tests
echo -e "\n[3/3] Running PowerShell Tests..."
if pwsh -File pwsh.test.ps1; then
    echo -e "${GREEN}PowerShell Tests Passed${NC}"
else
    echo -e "${RED}PowerShell Tests Failed${NC}"
    FAILED=1
fi

echo -e "\n----------------------------"
if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}ALL TESTS PASSED SUCCESSFULLY${NC}"
    exit 0
else
    echo -e "${RED}SOME TESTS FAILED${NC}"
    exit 1
fi
