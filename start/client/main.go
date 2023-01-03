package main

import (
    "context"
    "fmt"
    "log"

    petv1 "buf.build/gen/go/johnvajda/petapis/protocolbuffers/go/pet/v1"
    "buf.build/gen/go/johnvajda/petapis/grpc/go/pet/v1/petv1grpc"

    "google.golang.org/grpc"
)

func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}
func run() error {
    connectTo := "127.0.0.1:8080"
    conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
    if err != nil {
        return fmt.Errorf("failed to connect to PetStoreService on %s: %w", connectTo, err)
    }
    log.Println("Connected to", connectTo)

    petStore := petv1.NewPetStoreServiceClient(conn)
    if _, err := petStore.PutPet(context.Background(), &petv1.PutPetRequest{
        PetType: petv1.PetType_PET_TYPE_SNAKE,
        Name:    "Ekans",
    }); err != nil {
        return fmt.Errorf("failed to PutPet: %w", err)
    }

    log.Println("Successfully PutPet")
    return nil
}