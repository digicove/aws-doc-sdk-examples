// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
// snippet-start:[iam.go.create_account_alias]
package main

// snippet-start:[iam.go.create_account_alias.imports]
import (
    "flag"
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/iam"
    "github.com/aws/aws-sdk-go/service/iam/iamiface"
)
// snippet-end:[iam.go.create_account_alias.imports]

// MakeAccountAlias creates an alias for your IAM account
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
//     alias is the alias for the account
// Output:
//     If success, nil
//     Otherwise, an error from the call to CreateAccountAlias
func MakeAccountAlias(svc iamiface.IAMAPI, alias *string) error {
    // snippet-start:[iam.go.create_account_alias.call]
    _, err := svc.CreateAccountAlias(&iam.CreateAccountAliasInput{
        AccountAlias: alias,
    })
    // snippet-end:[iam.go.create_account_alias.call]

    return err
}

func main() {
    // snippet-start:[iam.go.create_account_alias.args]
    alias := flag.String("a", "", "The account alias")
    flag.Parse()

    if *alias == "" {
        fmt.Println("You must supply an account alias (-a ALIAS)")
    }
    // snippet-end:[iam.go.create_account_alias.args]

    // snippet-start:[iam.go.create_account_alias.session]
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := iam.New(sess)
    // snippet-end:[iam.go.create_account_alias.session]

    err := MakeAccountAlias(svc, alias)
    if err != nil {
        fmt.Println("Got an error creating an account alias")
        fmt.Println(err)
        return
    }

    fmt.Printf("Created account alias " + *alias)
}
// snippet-end:[iam.go.create_account_alias]
