# go-file-management-api

## Apis
### Upload
```
curl --location --request PUT 'http://localhost:8080/v1/upload/0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec' \
    --form 'file=@"./integration_test/cat.jpg"'
```

### List
```
curl --location 'http://localhost:8080/v1/list/0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec'
```

### Remove
```
curl --location --request DELETE 'http://localhost:8080/v1/file/0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec/cat.jpg'
```

## Design Pattern
- reference from Arnon Keereena