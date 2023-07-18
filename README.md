## Introduction

This is a minimal reproduction of the issue described [here][gh-discussion].

I was trying to perform a nested, idempotent upsert (as shown in the GraphQL
mutation below), and it was failing.

It turns out I was missing update permissions! Once added, things started 
working.

## Getting started

### Start Remote Webhook

```shell
go run main.go
```

### Start Hasura GraphQL Engine

```shell
(cd hasura && make start)
```

### Start Hasura Console

```shell
(cd hasura && make console)
```

### Execute bulk insertion

#### HTTP Headers

It's fine to use the x-hasura-admin-secret header, but to reproduce the bug, 
make sure you're also using the `x-hasura-role` header:

```
x-hasura-role: writer
```

#### Mutation

```graphql
mutation CreatePetAndOwner(
  $owner: OwnerInsertInput!
  $pet: PetInsertInput!
) {
  owner: insertOwnerOne(
    object: $owner
    onConflict: {
      constraint: owner_pkey
      updateColumns: [id]
    }
  ) {
    id
  }
  
  pet: insertPetOne(
    object: $pet
    onConflict: {
      constraint: pet_pkey
      updateColumns: [id]
    }
  ) {
    id
  }
}
```

#### Variables

```json
{
  "owner": {
    "id": "6b683c4d-0fb0-4f47-b3f4-822177753b1d",
    "name": "Kevin"
  },
  "pet": {
    "id": "5cd3dc17-1b05-433c-b8a4-9c2f4a09c603",
    "name": "Sparky",
    "ownerId": "6b683c4d-0fb0-4f47-b3f4-822177753b1d"
  }
}
```

[gh-discussion]: https://github.com/hasura/graphql-engine/discussions/9767