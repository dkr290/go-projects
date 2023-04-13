package instance

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

func GetAndDeleteTopic(ctx context.Context, RG string, NamespaceName string, SubscriptionID string, topics string) {

	fmt.Printf("Please check also that %s namespace is not paired and needs to delete all the topics prior pairing\n", NamespaceName)
	fmt.Printf("Do you want to start deleting of all topics from %s namespace: y\\n :", NamespaceName)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if text == 'y' {
			fmt.Println("starting the topics wipe... from ", NamespaceName)
			topicsClient, err := armservicebus.NewTopicsClient(SubscriptionID, getCreds(), nil)
			if err != nil {
				log.Fatalln(err)
			}
			tps := strings.Split(topics, ",")

			for _, tt := range tps {

				topicsClient.Delete(ctx, RG, NamespaceName, tt, nil)
				fmt.Println("deleting", tt)

			}
			break

		} else if text == 'n' {
			fmt.Println("Abording operation")
			os.Exit(1)
		} else {
			fmt.Println("You have to specify either y or n y\\n: ")

		}
	}

}
