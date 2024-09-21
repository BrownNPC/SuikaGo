package main

type EntityManager struct {
	m_Entities    []*Entity
	m_ToAdd       []*Entity
	m_EntitiesMap map[string][]*Entity // tagged entities
	m_Total       int
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		m_Entities:    []*Entity{},
		m_EntitiesMap: make(map[string][]*Entity),
	}
}

func (em *EntityManager) CreateEntity(tag string) *Entity {
	em.m_Total++
	// Create a new entity using NewEntity, which returns a pointer.
	en := NewEntity(em.m_Total, tag)
	// Store the pointer in the entities map.

	em.m_ToAdd = append(em.m_ToAdd, en)

	// Return the pointer to the created entity.
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
	}

	// Remove dead entities

	for _, e := range em.m_Entities {
		if !e.active {
			delete(em.m_EntitiesMap, e.tag)
		}
	}

	// Remove dead entities from list of entities

	activeEntities := []*Entity{}
	for _, e := range em.m_Entities {
		if e.active {
			activeEntities = append(activeEntities, e)
		}
	}
	em.m_Entities = activeEntities

	em.m_ToAdd = []*Entity{}
}
