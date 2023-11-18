package main

import (
	// "encoding/json"
	"encoding/base64"
	"html/template"
	"net/http"
	"io"
)

// Data adalah struktur untuk menyimpan data formulir
type Data struct {
	Nik              string		
	NamaLengkap      string		
	JenisKelamin     string		
	GolonganDarah    string		
	Ttl              string		
	Alamat           string		
	RtRw             string		
	KelDesa          string		
	Kecamatan        string		
	Agama            string		
	StatusPerkawinan string		
	Pekerjaan        string		
	Kewarganegaraan  string
	FotoKTP			 string		
}

// JSONData adalah variabel untuk menyimpan data dalam format JSON
// var JSONData []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan untuk tampilan formulir
		tmpl, err := template.ParseFiles("coba.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/cetak", func(w http.ResponseWriter, r *http.Request) {
		// Habis lu ngisi formulirnya, data nya masuk ke sini...
	 if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("userfile")
		if err != nil {
			http.Error(w, "Unable to retrieve image", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Read image data into a byte slice
		imageData, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Unable to read image data", http.StatusInternalServerError)
			return

		}
		imageDataString := base64.StdEncoding.EncodeToString(imageData)
		
		 data := Data{
			 Nik:              r.FormValue("nik"),
			 NamaLengkap:      r.FormValue("nama"),
			 JenisKelamin:     r.FormValue("sex"),
			 GolonganDarah:    r.FormValue("darah"),
			 Ttl:              r.FormValue("tl") + " " + r.FormValue("tgl") + "-" + r.FormValue("bln") + "-" + r.FormValue("thn"),
			 Alamat:           r.FormValue("alamat"),
			 RtRw:             r.FormValue("rt") + "/" + r.FormValue("rw"),
			 KelDesa:          r.FormValue("desa"),
			 Kecamatan:        r.FormValue("kecamatan"),
			 Agama:            r.FormValue("agama"),
			 StatusPerkawinan: r.FormValue("status"),
			 Pekerjaan:        r.FormValue("pekerjaan"),
			 Kewarganegaraan:  r.FormValue("warganegara"),
			 FotoKTP:		   imageDataString,
		 }
		// 	// Mengonversi data ke dalam format JSON
		 // jsonData, err := json.Marshal(data)
		 // if err != nil {
		 // 	http.Error(w, err.Error(), http.StatusInternalServerError)
		 // 	return
		 // }
	
		 // Menyimpan data JSON ke dalam variabel global
		 // JSONData = jsonData
	
		 // Di sini lu nge load si cetak.html nya, sekalian nngasih dia data yang tadi lu simpen
		//  di atas
			tmpl, err := template.ParseFiles("cetak.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, data) // nih yang ada {, data} nya...
			// abis itu lu ke bagian cetak.html buat tau lebih lanjut
		//  w.Header().Set("Content-Type", "application/json")
		//  json.NewEncoder(w).Encode(data)
		//  return
	 }
		// http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Menjalankan server pada port 8080
	http.ListenAndServe(":8080", nil)
}
