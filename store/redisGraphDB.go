package store

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	rg "github.com/redislabs/redisgraph-go"
	"github.com/subramanian0331/familytree/models"
)

//Graph Database that saves the family tree in graph data structure.
type RedisGraphDB struct {
	host string
}

func (r *RedisGraphDB) AddMember(member *models.Member) error {
	conn, _ := redis.Dial("tcp", r.host)
	defer conn.Close()
	graph := rg.GraphNew("FamilyTree", conn)
	newNode := rg.Node{
		Label: "member",
		Properties: map[string]interface{}{
			"id":         member.Id.String(),
			"firstName":  member.Firstname,
			"lastName":   member.Lastname,
			"dob":        time.Time(member.Dob).Format("03/15/2001"),
			"sex":        member.Sex.String(),
			"occupation": member.Occupation,
			"family":     member.Family,
		},
	}
	graph.AddNode(&newNode)
	graph.Commit()
	return nil
}
func (r *RedisGraphDB) AddRelationship(member1 string, member2 string, Relationship models.Relationship, family string) error {
	return nil
}
func (r *RedisGraphDB) GetMemberDetails(memberName string, family string) (*models.Member, error) {
	conn, _ := redis.Dial("tcp", r.host)
	defer conn.Close()
	graph := rg.GraphNew("FamilyTree", conn)
	member := models.Member{}
	query := fmt.Sprintf(`MATCH (x) WHERE x.name = %s RETURN x`, memberName)
	result, err := graph.Query(query)
	if err != nil {
		return nil, err
	}
	result.PrettyPrint()
	val := result.Record()
	p, ok := val.GetByIndex(0).(rg.Path)
	fmt.Printf("%s %v\n", p, ok)
	return &member, nil
}
func (r *RedisGraphDB) UpdateMemberDetails(member *models.Member, family string) error {
	return nil
}
func (r *RedisGraphDB) GetParents(memberName string, family string) ([]*models.Member, error) {
	return nil, nil
}
func (r *RedisGraphDB) GetChildren(memberName string, family string) ([]*models.Member, error) {
	return nil, nil
}
func (r *RedisGraphDB) GetAllFamilyMembers(family string) ([]*models.Member, error) {
	return nil, nil
}
