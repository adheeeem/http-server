package banners

import (
	"context"
	"errors"
	"sync"
)

type Service struct {
	mu           sync.RWMutex
	items        []*Banner
	nextBannerID int64
}

func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
}

func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	return s.items, nil
}
func (s *Service) ById(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}
	return nil, errors.New("banner not found")
}
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if item.ID == 0 {
		s.nextBannerID++
		item.ID = s.nextBannerID
		s.items = append(s.items, item)
		return item, nil
	} else {
		for i, banner := range s.items {
			if banner.ID == item.ID {
				item.ID = banner.ID
				s.items[i] = item
				return item, nil
			}
		}
	}
	return nil, errors.New("internal server error")
}
func (s *Service) RemoveById(ctx context.Context, id int64) (*Banner, error) {
	for i, banner := range s.items {
		if banner.ID == id {
			deletedBanner := banner
			s.items = append(s.items[:i], s.items[i+1:]...)
			return deletedBanner, nil
		}
	}
	return nil, errors.New("banner not found")
}
