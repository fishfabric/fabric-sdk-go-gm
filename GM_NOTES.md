## 01 replace x509 & tls import

```
-       "crypto/x509"
+       "github.com/Hyperledger-TWGC/tjfoc-gm/x509"
```

```
-       tls "crypto/tls"
+       tls "github.com/Hyperledger-TWGC/tjfoc-gm/gmtls"
```

```
-       "crypto/tls"
+       tls "github.com/Hyperledger-TWGC/tjfoc-gm/gmtls"
```

```
-       "google.golang.org/grpc/credentials"
+       credentials "github.com/Hyperledger-TWGC/tjfoc-gm/gmtls/gmcredentials"
```