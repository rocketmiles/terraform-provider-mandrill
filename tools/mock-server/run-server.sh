#!/bin/bash
cd "$(dirname "$0")"

MANDRILL_TEST_PORT="${MANDRILL_TEST_PORT:-8080}"
PROCESS_NAME="terraform-mandrill-provider-wiremock"
MOCK_SERVER="wiremock-jre8-standalone.jar"
JVM_ARGS="--verbose --disable-banner --global-response-templating --port=$MANDRILL_TEST_PORT"

echo "Downloading Wiremock Server... $MOCK_SERVER"

[ -f "$MOCK_SERVER" ] && echo "$MOCK_SERVER exists. Skipping download." || \
     wget -O "$MOCK_SERVER" "https://repo1.maven.org/maven2/com/github/tomakehurst/wiremock-jre8-standalone/2.32.0/wiremock-jre8-standalone-2.32.0.jar"

echo "Running Wiremock Server..."
echo "..."
echo "..."
java -jar -Dname=$PROCESS_NAME $MOCK_SERVER $JVM_ARGS
