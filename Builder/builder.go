package main

import "fmt"

// The NotificationBuilder has fields exported
type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

// SetTitle modifies the NotificationBuilder.title property
func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

// SetSubTitle modifies the NotificationBuilder.subtitle property
func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.SubTitle = subtitle
}

// SetMessage modifies the NotificationBuilder.message property
func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

// SetImage modifies the NotificationBuilder.image property
func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

// SetIcon modifies the NotificationBuilder.icon property
func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

// SetPriority modifies the NotificationBuilder.priority property
func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

// SetNotType modifies the NotificationBuilder.notType property
func (nb *NotificationBuilder) SetNotType(notType string) {
	nb.NotType = notType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {

	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("you need to specify a subtitle")
	}

	if nb.Priority > 5 {
		return nil, fmt.Errorf("priority must be between 0 and 5")
	}

	return &Notification{
		title:    nb.Title,
		subtitle: nb.SubTitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
