package notifications

import (
	"fmt"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

func CreateContactList(name string) error {
	ctx := context.Background()
	client := GetSesClient()

	input := &sesv2.CreateContactListInput {
		ContactListName: &name,
	}
	_, err := client.CreateContactList(ctx, input)
	if err != nil {
		fmt.Printf("Error creating contact list: %v\n", err)
		return err
	}

	return nil
}

func CreateContact(name string, email string) error {
	ctx := context.Background()
	client := GetSesClient()

	contact := &sesv2.CreateContactInput {
		ContactListName: &name,
		EmailAddress: &email,
	}
	_, err := client.CreateContact(ctx, contact)
	if err != nil {
		fmt.Printf("Error creating contact: %v\n", err)
		return err
	}
	return nil
}

func ListContact(name string) ([]string, error) {
	ctx := context.Background()
	client := GetSesClient()

	emails := []string{}
	list := &sesv2.ListContactsInput{
		ContactListName: &name,
	}
	output, err := client.ListContacts(ctx, list)
	if err != nil {
		fmt.Printf("Error fetching list: %v\n", err)
		return emails, err
	}
	for _, member := range output.Contacts {
		emails = append(emails, *member.EmailAddress)
	}
	return emails, nil
}

func DeleteContact(name string, email string) error {
	ctx := context.Background()
	client := GetSesClient()

	contact := &sesv2.DeleteContactInput {
		ContactListName: &name,
		EmailAddress: &email,
	}
	_, err := client.DeleteContact(ctx, contact)
	if err != nil {
		fmt.Printf("Error deleting contact: %v\n", err)
		return nil
	}
	return nil
}

func DeleteList(name string) error {
	ctx := context.Background()
	client := GetSesClient()

	input := &sesv2.DeleteContactListInput {
		ContactListName: &name,
	}
	_, err := client.DeleteContactList(ctx, input)
	if err != nil {
		fmt.Printf("Error deleting contact list: %v\n", err)
		return err
	}

	return nil
}
