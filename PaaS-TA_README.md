## 01. Check the commit contents
> [enable cce](https://github.com/PaaS-TA/cf-networking-release/commit/7d8c71e60f383a7f48dabf8f08b9578dead3daa7)

## 02. Submodule modify : lib/pq (src/github.com/lib/pq)
> pull v1.10.0 lib/pq
``` 
$ rm -rf src/github.com/lib/pq
$ mkdir -p src/github.com/lib/pq
$ git clone https://github.com/lib/pq.git -b v1.10.0 src/github.com/lib/pq
```

## 03. Submodule modify : go-sql-driver/mysql (src/github.com/go-sql-driver/mysql)
> pull v1.4.1 go-sql-driver/mysql
``` 
$ rm -rf src/github.com/go-sql-driver/mysql
$ mkdir -p src/github.com/go-sql-driver/mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.4.1 src/github.com/go-sql-driver/mysql
```

