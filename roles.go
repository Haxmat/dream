package dream

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

//Error variables
var (
	ErrRoleNotFound = errors.New("Role not found")
	ErrInvalidIndex = errors.New("invalid index") // Returned if the specified index is less than or greater than the length of the array
)

//Roles is a sortable list of roles
type Roles []*discordgo.Role

// UpdatePositions Updates the positions of the roles according to their position in the slice
func (r Roles) UpdatePositions() {
	for i, v := range r {
		v.Position = i
	}
}

//Move moves a role from one position to another in the slice
func (r Roles) Move(from, to int) error {
	if from < 0 || to < 0 || from >= len(r) || to >= len(r) {
		return ErrInvalidIndex
	}

	// Save the value of r[from] and remove it from the slice
	value := r[from]
	r = append(r[:from], r[from+1:]...)

	// Cut the slice in half to make room for inserting 'value'
	start := r[:to]

	// Make a copy of end so it no longer refers to the underlying array and cannot
	// Be modified by appending 'value' to 'start'. Otherwise it would overwrite the
	// First index in 'end'
	end := make([]*discordgo.Role, len(r[to:]))
	copy(end, r[to:])

	// Insert 'value' into the slice at index 'to'
	r = append(start, value)
	r = append(r, end...)

	return nil
}

// MoveByID finds a role by id and moves its position in the slice to 'to'
func (r Roles) MoveByID(ID string, to int) error {
	_, index, err := r.GetByID(ID)
	if err != nil {
		return err
	}
	return r.Move(index, to)
}

// MoveByName moves a role by name
func (r Roles) MoveByName(name string, to int) error {
	_, index, err := r.GetByName(name)
	if err != nil {
		return err
	}
	return r.Move(index, to)
}

//GetByID Returns a role by ID
func (r Roles) GetByID(ID string) (*discordgo.Role, int, error) {
	for i, v := range r {
		if v.ID == ID {
			return v, i, nil
		}
	}
	return nil, -1, ErrRoleNotFound
}

//GetByName Returns a role by Name
func (r Roles) GetByName(name string) (*discordgo.Role, int, error) {
	for i, v := range r {
		if v.Name == name {
			return v, i, nil
		}
	}
	return nil, -1, ErrRoleNotFound
}

// Less used to satisfy sort.Sort
func (r Roles) Less(i, j int) bool {
	return r[i].Position < r[j].Position
}

// Swap used to satisfy sort.Sort Swap
func (r Roles) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Len used to satisfy sort.Sort Len
func (r Roles) Len() int {
	return len(r)
}
