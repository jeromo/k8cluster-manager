package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	/*	employees := []model.Employee{}
		db.Find(&employees)
		respondJSON(w, http.StatusOK, employees)
	*/
}

func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	/*	employee := model.Employee{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&employee); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		if err := db.Save(&employee).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusCreated, employee)
	*/
}

func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	/*	vars := mux.Vars(r)

		name := vars["name"]
		employee := getEmployeeOr404(db, name, w, r)
		if employee == nil {
			return
		}
		respondJSON(w, http.StatusOK, employee)
	*/
}

func UpdateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	/*	vars := mux.Vars(r)

		name := vars["name"]
		employee := getEmployeeOr404(db, name, w, r)
		if employee == nil {
			return
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&employee); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		if err := db.Save(&employee).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, employee)
	*/
}

func DeleteEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	/*	vars := mux.Vars(r)

		name := vars["name"]
		employee := getEmployeeOr404(db, name, w, r)
		if employee == nil {
			return
		}
		if err := db.Delete(&employee).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusNoContent, nil)
	*/
}

func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	/*	vars := mux.Vars(r)

		name := vars["name"]
		employee := getEmployeeOr404(db, name, w, r)
		if employee == nil {
			return
		}
		employee.Disable()
		if err := db.Save(&employee).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, employee)
	*/
}

func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	/*	vars := mux.Vars(r)

		name := vars["name"]
		employee := getEmployeeOr404(db, name, w, r)
		if employee == nil {
			return
		}
		employee.Enable()
		if err := db.Save(&employee).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, employee)
	*/
}
