package services

import (
	"Gononymous/internal/core/domains/dao"
	drivenports "Gononymous/internal/core/ports/driven_ports"
	"Gononymous/utils"
	"context"
	rnd "math/rand"
	"sync"
)

type SessionService struct {
	SessionRepo drivenports.SessionRepoInterface
	Character   drivenports.CharacterRepoInterface
	Picker      Picker
}

type Picker struct {
	mu    sync.Mutex
	arr   []int
	right int
}

func NewPicker() *Picker {
	arr := make([]int, 0, 826)
	for i := 1; i <= 826; i++ {
		arr = append(arr, i)
	}
	return &Picker{mu: sync.Mutex{}, arr: arr, right: len(arr) - 1}
}

func NewSessionService(SessionRepo drivenports.SessionRepoInterface, Character drivenports.CharacterRepoInterface) *SessionService {
	return &SessionService{SessionRepo: SessionRepo, Character: Character, Picker: *NewPicker()}
}

func (p *Picker) Pick() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.right < 0 {
		p.right = len(p.arr) - 1
	}
	idx := rnd.Intn(p.right + 1)
	id := p.arr[idx]
	p.arr[idx] = p.arr[p.right]
	p.arr[p.right] = id
	p.right--
	return id
}

func (s *SessionService) CreateSession(ctx context.Context) error {
	ch, err := s.Character.GetCharacter(ctx, s.Picker.Pick())
	if err != nil {
		return err
	}
	id := utils.UUID()
	session := dao.Session{UsersId: id, Name: ch.Name, AvatarURL: ch.AvatarURL}
	err = s.SessionRepo.AddSession(ctx, session)
	return err
}
