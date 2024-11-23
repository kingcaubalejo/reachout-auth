package repository

import (
	"reachout-auth/model"
)

var users = []*model.User{
	{ ID: "644b8bce-7af6-4a23-b9ce-5582dca26cb6", Username: "gdaniells0", Password: "$2a$04$NKFJGm8Xj37IqQ2b0TnP/uDHfKaM4p5rm9CxahS14BqZwboU9Mxpi", Email: "htargetter0@liveinternet.ru"},
	{ ID: "960ceb6d-be25-4c45-94c8-f85f9fa3670e", Username: "bluxton1", Password: "$2a$04$Ymtm6vN26qpiSHe3R34J3eq3zdiw2cUzccIb8uP1vUXmGd0D1ujEu", Email: "ebengoechea1@photobucket.com"},
	{ ID: "1894c316-0d13-47d1-845c-95654a5fa882", Username: "ctolland2", Password: "$2a$04$YLOJqAgRZUZD5POn6lXi6OFejSA4B6RV6mrsBIy.9mvzIMgfetYQa", Email: "bcrossland2@gmpg.org"},
	{ ID: "0bafafc8-5e72-4319-869b-0b611c8488a7", Username: "jelphinston4", Password: "	$2a$04$tqUMPOGOVi8lLaDuF0RLQuyDYtwCRv7HG1p1TjwmgS4QdrZ35v2l6", Email: "dhollingsby4@etsy.com"},
	{ ID: "90b32715-4cc1-41a2-b8f1-2611748934d4", Username: "cdavids6", Password: "$2a$04$9f1J.5gPrW90MS.6ltV1HuVA44kpxPpt0HuDbtHquMcSXi7BzqHBe", Email: "dlemoir6@unblog.fr"},
}

func GetUsers() []*model.User {
	return users
}

func FindUserById(id string) *model.User {
	for _, a := range users {
		if a.ID == id {
			return a
		}
	}

	return nil
}

func Save(user *model.User) *model.User {
	users = append(users, user)
	return user
}