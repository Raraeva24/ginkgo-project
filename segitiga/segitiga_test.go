package segitiga_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "segitiga"
)

var _ = Describe("KalkulatorSegitiga", func() {
    Describe("Luas", func() {
        It("Harus menghitung luas segitiga dengan betul", func() {
            segitiga := segitiga.InitSegitiga(4, 3, 4, 4, 5) // alas = 4, Tinggi = 3, Sisi1 = 4, Sisi2 = 4, Sisi3 = 5
            Expect(segitiga.Area()).To(Equal(6.0))            // Seharusnya 6
        })
    })

    Describe("Keliling", func() {
        It("Harus menghitung keliling segitiga dengan benar", func() {
            segitiga := segitiga.InitSegitiga(4, 3, 4, 4, 5) // alas = 4, Tinggi = 3, Sisi1 = 4, Sisi2 = 4, Sisi3 = 5
            Expect(segitiga.Perimeter()).To(Equal(13.0))      // Seharusnya 13
        })
    })

    Describe("Validasi", func() {
        Context("Ketika sisi-sisi dan tinggi membentuk segitiga yang valid", func() {
            It("Tidak seharusnya mengembalikan error", func() {
                segitiga := segitiga.InitSegitiga(4, 3, 4, 4, 5) // alas = 4, Tinggi = 3, Sisi1 = 4, Sisi2 = 4, Sisi3 = 5
                err := segitiga.ValidateTriangle()
                Expect(err).To(BeNil())
            })
        })

        Context("Ketika sisi-sisi dan tinggi tidak membentuk segitiga yang valid", func() {
            It("Harus mengembalikan error ketika alas bernilai negatif", func() {
                segitiga := segitiga.InitSegitiga(-1, 3, 4, 4, 5) // alas = -1 (tidak valid)
                err := segitiga.ValidateTriangle()
                Expect(err).ToNot(BeNil())
            })
        })
    })
})
