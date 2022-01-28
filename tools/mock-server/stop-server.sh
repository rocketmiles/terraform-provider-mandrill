MOCK_SERVER="wiremock-jre8-standalone.jar"

WM_PID=$(pgrep -f "$MOCK_SERVER")

[[ $WM_PID != "" ]] && kill $WM_PID && echo "Wiremock running with process id: $WM_PID killed successfully."
[[ -f "wiremock.out" ]] && rm wiremock.out

[[ $WM_PID == "" ]] && echo "Wiremock Process not running. Could not stop" && exit 0

exit 0