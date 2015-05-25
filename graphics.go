package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

func (page *Page) DrawImage(image *Image, x, y, width, height float32) error {
	C.HPDF_Page_DrawImage(
		page.page, image.image,
		C.HPDF_REAL(x), C.HPDF_REAL(y), C.HPDF_REAL(width), C.HPDF_REAL(height),
	)
	return page.pdf.GetLastError()
}

func (page *Page) SetFontAndSize(font *Font, size float32) error {
	C.HPDF_Page_SetFontAndSize(
		page.page, font.font, C.HPDF_REAL(size),
	)
	return page.pdf.GetLastError()
}

func (page *Page) SetRGBFill(r float32, g float32, b float32) error {
	C.HPDF_Page_SetRGBFill(
		page.page,
		C.HPDF_REAL(r),
		C.HPDF_REAL(g),
		C.HPDF_REAL(b),
	)

	return page.pdf.GetLastError()
}
