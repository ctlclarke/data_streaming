package main

import (
    "fmt"
    "os"
    "os/signal"

    "github.com/yosssi/gmq/mqtt"
    "github.com/yosssi/gmq/mqtt/client"
)

func main() {
    // Set up channel on which to send signal notifications.
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, os.Interrupt, os.Kill)

    // Create an MQTT Client.
    cli := client.New(&client.Options{
        // Define the processing of the error handler.
        ErrorHandler: func(err error) {
            fmt.Println(err)
        },
    })

    // Terminate the Client.
    defer cli.Terminate()

    // Connect to the MQTT Server.
    err := cli.Connect(&client.ConnectOptions{
        Network:  "tcp",
        Address:  "172.28.42.32:15675",
        ClientID: []byte("dylon"),
    })
    if err != nil {
        panic(err)
    }

  
    // Publish a message.
    err = cli.Publish(&client.PublishOptions{
        QoS:       mqtt.QoS0,
        TopicName: []byte("bar/baz"),
        Message:   []byte("testMessage"),
    })
    if err != nil {
        panic(err)
    }

    // Disconnect the Network Connection.
    if err := cli.Disconnect(); err != nil {
        panic(err)
    }
}