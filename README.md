## GL-iNet Spitz AX X3000 Signal Stats server

- This is a simple HTTP server that returns an autorefreshing page with CA stats to help when adjusting and positioning the router and antennas to get the best signal.

- How to setup ?
1. Clone this repo.
2. Upload openwrt/bin/glinet-spitz-ax-signal-stats to /root
3. Upload openwrt/etc/init.d/glinet-spitz-ax-signal-stats to /etc/init.d
4. SSH to the router
5. chmod +x /root/glinet-spitz-ax-signal-stats
6. chmod +x /etc/init.d/glinet-spitz-ax-signal-stats
7. /etc/init.d/glinet-spitz-ax-signal-stats enable
8. /etc/init.d/glinet-spitz-ax-signal-stats start
9. Visit http://router-ip:8080/ default is http://192.168.8.1:8080/
10. Adjust router to get best numbers.
11. Enjoy.