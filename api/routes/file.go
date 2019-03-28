package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

// CreateFile ...
func CreateFile(w http.ResponseWriter, r *http.Request) {
	var body models.File // to save File JSON body
	var err error

	// Decoding request body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		WriteError(w, 422, err, "Unable to read request body")
		return
	}

	// Inserting Document
	nFile, err := body.Create()
	if err != nil {
		WriteError(w, 422, err, "Failed to insert new document")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(nFile)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// ReadOneFile - ...
func ReadOneFile(w http.ResponseWriter, r *http.Request) {
	var err error
	var mFile models.File

	fIDStr := mux.Vars(r)["id"] // File ObjectId String

	// Read File
	mFile, err = mFile.ReadSingle(bson.M{"_id": bson.ObjectIdHex(fIDStr)})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(mFile)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// ReadFiles - read all files, both Public and Private
func ReadFiles(w http.ResponseWriter, r *http.Request) {
	var err error
	var mFiles models.Files

	// Read File
	mFiles, err = mFiles.Read(bson.M{})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(mFiles)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// EditFile - ...
func EditFile(w http.ResponseWriter, r *http.Request) {
	var err error
	var nFile models.File

	fIDStr := mux.Vars(r)["id"] // File ObjectId String

	// Decoding request body
	var body map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		WriteError(w, 422, err, "Unable to read request body")
		return
	}

	fmt.Printf("body %+v\n", body)

	// Update File Document
	nFile, err = nFile.Update(
		bson.M{"_id": bson.ObjectIdHex(fIDStr)},
		body,
	)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(nFile)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// DeleteFile - ...
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	var err error
	var dFile models.File

	fIDStr := mux.Vars(r)["id"] // File ObjectId String

	// Read File
	_, err = dFile.Delete(bson.ObjectIdHex(fIDStr))
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
