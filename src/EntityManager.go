package main

import "github.com/jakecoffman/cp/v2"

type EntityManager struct {
	m_Entities    []*Entity
	m_ToAdd       []*Entity
	m_EntitiesMap map[string][]*Entity // tagged entities
	m_Total       int
	space         *cp.Space
}

func NewEntityManager() *EntityManager {
	em := &EntityManager{
		m_Entities:    []*Entity{},
		m_EntitiesMap: make(map[string][]*Entity),
		space:         cp.NewSpace(),
	}
	return em
}

func (em *EntityManager) CreateEntity(tag string) *Entity {
	// Create a new entity using NewEntity, which returns a pointer.
	en := NewEntity(em.m_Total, tag)
	// Store the pointer in the entities map.

	em.m_ToAdd = append(em.m_ToAdd, en)

	// Return the pointer to the created entity.
	em.m_Total++
	return en
}

func (e *EntityManager) GetEntitiesByTag(tag string) []*Entity {
	return e.m_EntitiesMap[tag]
}

func (e *EntityManager) GetEntities() []*Entity {
	return e.m_Entities
}

func (em *EntityManager) Update() {
	// Add new entities

	for _, e := range em.m_ToAdd {
		em.m_Entities = append(em.m_Entities, e)
		em.m_EntitiesMap[e.tag] = append(em.m_EntitiesMap[e.tag], e)
		if e.Shape != nil {
			em.space.AddShape(e.Shape)
		}
		if e.Body != nil {
			em.space.AddBody(e.Body)
		}
	}

	// Remove dead entities from map
	//entities is a list of pointers to Entity{}
	for tag, entities := range em.m_EntitiesMap {
		updatedEntities := []*Entity{}
		// e is the pointer to the entity
		for _, e := range entities {
			if e.active {
				updatedEntities = append(updatedEntities, e)
			}
		}
		em.m_EntitiesMap[tag] = updatedEntities
	}

	// Remove dead entities from list of entities
	activeEntities := []*Entity{}
	for _, e := range em.m_Entities {
		if e.active {
			activeEntities = append(activeEntities, e)
		} else {
			// remove them from physics space
			if e.Shape != nil {
				em.space.RemoveShape(e.Shape)
			}
			if e.Body != nil {
				em.space.RemoveBody(e.Body)
			}
			// unload the font if its a font entity
			if e.Font != nil {
				e.Font.Close()
			}
		}
	}
	em.m_Entities = activeEntities

	em.m_ToAdd = []*Entity{}
}

func (em *EntityManager) Space() *cp.Space {
	return em.space
}

func (em *EntityManager) GetByID(tag string, id int) *Entity {
	for _, e := range em.m_EntitiesMap[tag] {
		if e.id == id {
			return e
		}
	}
	return nil
}
