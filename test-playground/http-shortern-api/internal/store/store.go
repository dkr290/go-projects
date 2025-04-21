package store

import (
	"context"
	"fmt"
	"http-shortern-api/internal/utils"
	"sync"
)

type Store struct {
	mu    sync.RWMutex
	links map[string]Link
}

func NewStore() *Store {
	return &Store{
		links: make(map[string]Link),
	}
}

func (s *Store) Create(ctx context.Context, link Link) error {
	if err := validateNewLink(link); err != nil {
		return fmt.Errorf("%w: %w", utils.ErrInvalidRequest, err)
	}

	if link.Key == "fortesting" {
		return fmt.Errorf("%w: db at IP ... failed", utils.ErrInternal)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.links[link.Key]; ok {
		return utils.ErrExists
	}
	s.links[link.Key] = link
	return nil
}

func (s *Store) Retrieve(ctx context.Context, key string) (Link, error) {
	if err := validateLinkKey(key); err != nil {
		return Link{}, fmt.Errorf("%w: %w", utils.ErrInvalidRequest, err)
	}
	if key == "fortesting" {
		return Link{}, fmt.Errorf("%w: db at IP ... failed", utils.ErrInternal)
	}
	// holds the read-lock until the function returns
	s.mu.RLock()

	defer s.mu.RUnlock()

	link, ok := s.links[key]
	if !ok {
		return Link{}, utils.ErrNotExist
	}
	return link, nil
}
