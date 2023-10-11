package main

import (
	"fmt"
	"net/http"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func (app *application) CreateAndSendInvoice(w http.ResponseWriter, r *http.Request) {

	// receive json post
	var order Order
	err := app.readJSON(w, r, &order)
	if err != nil {
		app.badRequest(w, r, err)
	}

	// generate pdf invoice
	app.CreateInvoicePDF(order)

	// create mail
	//FIXME:

	// send mail with attachment
	//FIXME:

	// send response
	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	resp.Error = false
	resp.Message = fmt.Sprintf("Invoice %d.pdf created and sent to %s .", order.ID, order.Email)
	app.writeJSON(w, http.StatusCreated, resp)
}

func (app *application) CreateInvoicePDF(order Order) error {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(10, 13, 10)
	pdf.SetAutoPageBreak(true, 0)

	importer := gofpdi.NewImporter()
	t := importer.ImportPage(pdf, "./pdf-templates/invoice.pdf", 1, "/MediaBox")
	pdf.AddPage()

	importer.UseImportedTemplate(pdf, t, 0, 0, 215.9, 0)

	// write info
	pdf.SetY(50)
	pdf.SetX(10)
	pdf.SetFont("Times", "", 11)
	pdf.CellFormat(97, 8, fmt.Sprintf("Attention: %s %s", order.FirstName, order.LastName), "", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(97, 8, order.Email, "", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(97, 8, order.CreatedAt.Format("02-01-2006"), "", 0, "L", false, 0, "")

	// write info
	pdf.SetY(93)
	pdf.SetX(58)
	pdf.CellFormat(155, 8, order.Product, "", 0, "L", false, 0, "")

	pdf.SetX(166)
	pdf.CellFormat(20, 8, fmt.Sprintf("%d", order.Quantity), "", 0, "C", false, 0, "")

	pdf.SetX(185)
	pdf.CellFormat(20, 8, fmt.Sprintf("$%.2f", float32(order.Amount/100.0)), "", 0, "R", false, 0, "")

	pdf.SetX(185)
	pdf.CellFormat(20, 8, fmt.Sprintf("$%.2f", float32(order.Amount/100.0)), "", 0, "R", false, 0, "")

	invoicePath := fmt.Sprintf("./invoices/%d.pdf", order.ID)

	err := pdf.OutputFileAndClose(invoicePath)
	if err != nil {
		return err
	}

	return nil

}
