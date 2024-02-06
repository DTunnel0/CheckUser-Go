package http

const CERT_CONTENT = `
-----BEGIN CERTIFICATE-----
MIIEpjCCA46gAwIBAgIUHjmghw6aVaAYFolDPUmXDWjTP2gwDQYJKoZIhvcNAQEL
BQAwgYsxCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBDbG91ZEZsYXJlLCBJbmMuMTQw
MgYDVQQLEytDbG91ZEZsYXJlIE9yaWdpbiBTU0wgQ2VydGlmaWNhdGUgQXV0aG9y
aXR5MRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRMwEQYDVQQIEwpDYWxpZm9ybmlh
MB4XDTI0MDIwNjA0NDIwMFoXDTM5MDIwMjA0NDIwMFowYjEZMBcGA1UEChMQQ2xv
dWRGbGFyZSwgSW5jLjEdMBsGA1UECxMUQ2xvdWRGbGFyZSBPcmlnaW4gQ0ExJjAk
BgNVBAMTHUNsb3VkRmxhcmUgT3JpZ2luIENlcnRpZmljYXRlMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzYpGVpz8P2EOSDUozIY6tpqpUc8cQGgLieJf
+pTrjkTnc3Ug7j/aBT89rCw1DvfZky1E+74L+t5wp9OCegDyGwSXJvAK7IXFUEoP
n1VwzmGKDMZaL97f+nHhSAz+TCWAdS+3Y/vqOnOUwO0ZZabbihQzg/MxROB2BMT+
pEEeemsuCAA5FWEbq5wbOiytg81Mgl9mY117Vk0rgapHrH9YtToAaLbc4LWnbqNK
fgN/TqS+sfMUssQ88ds6E2YXJBZ5MbzEGrkbBqXk+6pxn7VelxbqLK+vpgdVN123
uwbG36nV5ssJSxxFtlgCQBrHvgO4Np8wodicpedDHoEGW4zAaQIDAQABo4IBKDCC
ASQwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcD
ATAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQBBN8Q9nWySQtaYzxfLliMUV3igDAf
BgNVHSMEGDAWgBQk6FNXXXw0QIep65TbuuEWePwppDBABggrBgEFBQcBAQQ0MDIw
MAYIKwYBBQUHMAGGJGh0dHA6Ly9vY3NwLmNsb3VkZmxhcmUuY29tL29yaWdpbl9j
YTApBgNVHREEIjAggg8qLmR0dW5uZWwuY2xvdWSCDWR0dW5uZWwuY2xvdWQwOAYD
VR0fBDEwLzAtoCugKYYnaHR0cDovL2NybC5jbG91ZGZsYXJlLmNvbS9vcmlnaW5f
Y2EuY3JsMA0GCSqGSIb3DQEBCwUAA4IBAQCIujyUFna6n4YfGrMzRqdvhocWidgd
PBD/hwSHYh+ZaJJzfA7zeqGC7jxOrJrlqhqv2rdu55t05yuHBG7G6Dcny+cZ2r8E
NBUM+o3DWRa/eG9IZzz9r37lofx/3gn+5x8Z+JGjyUUNrnd6FaUtwv7MSGilbd5I
XEMxs3838Z+BDDeIq3bG7T2tWhX3tbsb3BYtoC7e/R/zjxPvHelgBhASnsy1c7+I
Jn8t92QeSW1dXyM0rPElUELoBU09Gm7F/eJh3+l2m5mszdGHKnO+TxGKCgI4XSPV
31Axx5Y5PHnG21aBJ3o6TmGL6IP7CRTgsrkOiPIox391w/x4WcAy6M7h
-----END CERTIFICATE-----`

const KEY_CONTENT = `
-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDNikZWnPw/YQ5I
NSjMhjq2mqlRzxxAaAuJ4l/6lOuOROdzdSDuP9oFPz2sLDUO99mTLUT7vgv63nCn
04J6APIbBJcm8ArshcVQSg+fVXDOYYoMxlov3t/6ceFIDP5MJYB1L7dj++o6c5TA
7RllptuKFDOD8zFE4HYExP6kQR56ay4IADkVYRurnBs6LK2DzUyCX2ZjXXtWTSuB
qkesf1i1OgBottzgtaduo0p+A39OpL6x8xSyxDzx2zoTZhckFnkxvMQauRsGpeT7
qnGftV6XFuosr6+mB1U3Xbe7BsbfqdXmywlLHEW2WAJAGse+A7g2nzCh2Jyl50Me
gQZbjMBpAgMBAAECggEACo2wAOB8nzB3kEOSfbXiq9+TFA1DURdGiGTkMrSUx7BI
BgickT8cDarqmr2GV+dn94OaiCaA7Peg5y0YBPWpeLOqoyguF2ji8bVrye2UJjSh
5vgER3L1IyHXxGBOalB+oQW7L9oOc8Pdfm5uIGcJ3LQe1QaWoAe6Z5JJ1Ns3Gmfw
vjU/bV95JmnIVXMR74Ow219Jo5WvnaiwZbziijcqCSmuUqYDuFQp3AaNfA1zNXtT
o74jaQ2b5QE9Y0kk1wVU4DutDnVUwFrsR4glldVQV9HYYokEvuDI9tOZqLSU+Iwk
/EWY+dgaS2x8YMDacgZNxSjNbhDhiFa9cgHxSadPAQKBgQD+FNnetocMmZ7RweFb
YSQPTqk4r9RYNw5Z9pCEN1cnwufDKOFjyoI9NPZeTSFhYXCjzExj4vZsXwNEivkz
Db5I6xyLJ92TjkoEd6y9PU2B5CKPpI6T+z7TY8RBeMUET4seQc22F9XGDSC6/dyS
D56+qms/NWkWbltwtyzO71xQYQKBgQDPF5due9hE8Q2WNtZI4rViGDybL6tH11RE
naRa63rKv/rKHo/8hamzZbFfU21k/Oko49eYwbprRKA9nTrmDYGollAXT/JWHrBG
Rbo3Fi8YVV16Eyr6DICMz600gTsDheq3TvpW0YKW0xM7mcpUMo9KFAiIwFNBC8fw
eS14f9kNCQKBgBiOb1glD/xZxI3FTUCFrPSFx7kg1UcJWyu6ttDwgE3penjUNKRu
aBP+UGlgzv8YaciK8D8fKm3i6O/w9pDGnUNy5blVSwb9042G+3z2tcz9/ZEgeF85
AyNvGwKw52m5PlrYRUd6GkEf96/a5TyAofkPg3oCcXungtLsATqmy6dBAoGAc7UB
tghaIMLyTXCcL6MDpyhVjHuI3p3wBlpx/x68v9WeARosZvIjjAmQnetWHuu0NlV/
G2l6h/6S7XoQ84KuZAx/+VaA1x9UbB7/WVH6xETF8rQM+iLMHGDYrJJb767+Iqds
9d8fcLfEcjOOOZb6OnCRCB81JQ25C6IZBs+f7UkCgYA73GAHtHCNJfzS4szyJL8+
Ohj4Pj8JqkUrEYGS+u0wkXjIG2xwMpaydrEaiO2vG8vYmZ1N1TG7iOTi9Vt9XhY3
nfvrzl3uxz2rELFknvMwnEylNVdfIWh1BE8mQFVA676+I0nGTs4Z2dhGaQtIKx7G
15BqeXcwx/7xcM4BW96Xdg==
-----END PRIVATE KEY-----
`