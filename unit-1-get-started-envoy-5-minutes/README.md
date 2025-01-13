Article link: https://tetrate.io/blog/get-started-with-envoy-in-5-minutes
Github ref: https://github.com/peterj/color-app

To use envoy for color-app we will follow these step bellow:

Step 1: dockerize color-app
- $ docker build -t color-app:1.0.0 .
- $ docker run -d -p 5177:3000 --env BG_COLOR="#01a3a4" --name color-app-container-01 color-app:1.0.0
- $ docker run -d -p 5188:3000 --env BG_COLOR="#2e86de" --name color-app-container-02 color-app:1.0.0

Step 2: Run envoy config
- $ func-e run -c envoy.yaml


Note:
- To stop envoy process:
  + $ ps aux | grep envoy
  + $ kill <PID>


Explain "ps aux | grep envoy":
- ps stands for “process status”.
- aux are options that define the output format:
  + a: Shows processes from all users.
  + u: Displays detailed information about each process.
  + x: Includes processes not attached to a terminal.
