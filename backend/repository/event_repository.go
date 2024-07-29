package repository

import (
	config "com.electricity.online/db"
	pb "com.electricity.online/pb"
)

type EventRepository struct{}

func (er *EventRepository) GetEvents() ([]*pb.Event, error) {
	var events []*pb.Event
	err := config.DatabaseClient.Table("event").Find(&events).Error
	return events, err
}

func (er *EventRepository) CreateEvent(event *pb.Event) error {
	return config.DatabaseClient.Table("event").Create(event).Error
}

func (er *EventRepository) GetEventById(id string) (*pb.Event, error) {
	var event *pb.Event
	err := config.DatabaseClient.Table("event").First(&event, id).Error
	return event, err
}

func (er *EventRepository) UpdateEvent(event *pb.Event) error {
	var ev *pb.Event
	err := config.DatabaseClient.Table("event").First(&ev, event.Id).Error
	if err != nil {
		return err
	}
	return config.DatabaseClient.Table("event").Model(&ev).Updates(event).Error
}

func (er *EventRepository) DeleteEvent(id string) error {
	var event *pb.Event
	result := config.DatabaseClient.Table("event").Where("id = ?", id).Delete(&event)
	return result.Error
}
