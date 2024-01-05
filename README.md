# Prisma Client with Go

Special thanks to [Luca Steeb](https://github.com/steebchen), for developing an elegant package.

This package is dedicated to [Goravel](https://github.com/goravel/goravel) as an extend package.

## Version

| goravel/prisma | goravel/framework |
|-------------|-------------------|
| v1.1.x      | v1.13.x           |

## Install

1. Add package:

```bash
go get github.com/goravel/prisma
```

2. Register Service Provider:

```
// config/app.go
import "github.com/goravel/prisma"

"providers": []foundation.ServiceProvider{
    ...
    &prisma.ServiceProvider{},
}
```

3. See artisan list

```bash
go run . artisan list
go run . artisan prisma:version
```

## Info on Commands

Be sure to check [Prisma Client with Go](https://github.com/steebchen/prisma-client-go), as there are more information there as well.
