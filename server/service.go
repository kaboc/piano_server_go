package main

import (
	"fmt"
	"io"
	"log"
	"sync"

	pb "github.com/kaboc/piano_server_go/protos"
)

var streams sync.Map

type pianoService struct {
	pb.UnimplementedPianoServer
}

func (*pianoService) Connect(stream pb.Piano_ConnectServer) error {
	streams.Store(stream, nil)
	defer func() {
		streams.Delete(stream)
		log.Println("Disconnect:", &stream)
	}()

	log.Println("Connected:", &stream)

	for {
		note, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Printf("Request from %v: %d\n", &stream, note.Pitch)

		if note.Pitch > 127 {
			log.Println("pitch out of range")
			continue
		}

		streams.Range(func(key, _ interface{}) bool {
			if key != stream {
				s := key.(pb.Piano_ConnectServer)
				err = s.Send(&pb.Note{Pitch: note.Pitch})
				if err != nil {
					log.Printf("error while broadcasting: %v", err)
				}
			}
			return true
		})
	}
}
