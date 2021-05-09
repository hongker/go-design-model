package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudent_Clone(t *testing.T) {
	student := &Student{
		Name:  "user1",
		Books: []string{"三国演义","红楼梦","算法导论"},
	}

	studentClone := student.Clone()
	assert.Equal(t, student.Name, studentClone.Name)
	assert.Equal(t, student.Books, studentClone.Books)

	student.Books = append(student.Books, "计算机网络")
	assert.Equal(t, student.Books, studentClone.Books)
}



func TestStudent_DeepClone(t *testing.T) {
	student := &Student{
		Name:  "user1",
		Books: []string{"三国演义","红楼梦","算法导论"},
	}

	studentClone := student.DeepClone()
	assert.Equal(t, student.Name, studentClone.Name)
	assert.Equal(t, student.Books, studentClone.Books)

	student.Books = append(student.Books, "计算机网络")
	assert.NotEqual(t, student.Books, studentClone.Books)
}

