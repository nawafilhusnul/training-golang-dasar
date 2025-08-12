package main

import (
	"errors"
	"fmt"
)

/*
===============================================================================
TASK: STUDENT ENHANCEMENT
===============================================================================
Instruksi:
1. Tambahkan method AddGrade(grade int) untuk menambah nilai ke slice Grades
2. Tambahkan method AddCourse(courseName string, grade float64) untuk menambah course
3. Tambahkan method GetCourseGrade(courseName string) float64 untuk mengambil nilai course
4. Validasi input: grade harus 0-100, course name tidak boleh kosong

STARTER CODE:
*/

type Student struct {
	Name    string
	Grades  []int
	Courses map[string]float64
}

// TODO: Implementasikan method-method berikut:

// AddGrade menambahkan nilai baru ke slice Grades
// Parameter: grade (int) - nilai antara 0-100
// Return: error jika grade tidak valid
func (s *Student) AddGrade(grade int) error {
	// IMPLEMENTASI DI SINI
	if grade < 0 || grade > 100 {
		return errors.New("grade harus antara 0-100")
	}

	return nil
}

// AddCourse menambahkan course baru dengan nilainya
// Parameter: courseName (string), grade (float64)
// Return: error jika input tidak valid
func (s *Student) AddCourse(courseName string, grade float64) error {
	// IMPLEMENTASI DI SINI
	if courseName == "" {
		return errors.New("nama course tidak boleh kosong")
	}
	if grade < 0 || grade > 100 {
		return errors.New("grade harus antara 0-100")
	}

	return nil
}

// GetCourseGrade mengambil nilai untuk course tertentu
// Parameter: courseName (string)
// Return: grade (float64), exists (bool)
func (s Student) GetCourseGrade(courseName string) (float64, bool) {
	// IMPLEMENTASI DI SINI
	return 0, false
}

func testStudentEnhancement() {
	// Test Task 1: Student Enhancement
	fmt.Println("TASK 1: Student Enhancement")
	student := Student{Name: "Budi"}

	student.AddGrade(85)
	student.AddGrade(90)
	student.AddCourse("Matematika", 87.5)
	student.AddCourse("Fisika", 92.0)

	if grade, exists := student.GetCourseGrade("Matematika"); exists {
		fmt.Printf("Nilai Matematika: %.1f\n", grade)
	}
	fmt.Printf("Grades: %v\n\n", student.Grades)
}
