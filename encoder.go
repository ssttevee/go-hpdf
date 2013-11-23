package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"

#if HPDF_MAJOR_VERSION >= 2 && HPDF_MINOR_VERSION < 3
#define HPDF_UTF_SUPPORT 0

HPDF_STATUS HPDF_UseUTFEncodings(HPDF_Doc pdf) {
	return HPDF_OK;
}
#else
#define HPDF_UTF_SUPPORT 1
#endif
*/
import "C"
import (
	"errors"
	"unsafe"
)

func (pdf *PDF) GetEncoder(encodingName string) (*Encoder, error) {
	cencodingName := C.CString(encodingName)
	encoder := C.HPDF_GetEncoder(pdf.doc, cencodingName)
	C.free(unsafe.Pointer(cencodingName))

	if encoder != nil {
		return &Encoder{encoder}, nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) GetCurrentEncoder() *Encoder {
	encoder := C.HPDF_GetCurrentEncoder(pdf.doc)

	if encoder != nil {
		return &Encoder{encoder}
	} else {
		return nil
	}
}

func (pdf *PDF) SetCurrentEncoder(encodingName string) error {
	cencodingName := C.CString(encodingName)
	C.HPDF_SetCurrentEncoder(pdf.doc, cencodingName)
	C.free(unsafe.Pointer(cencodingName))

	return pdf.GetLastError()
}

func (pdf *PDF) UseJPEncodings() error {
	C.HPDF_UseJPEncodings(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) UseKREncodings() error {
	C.HPDF_UseKREncodings(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) UseCNTEncodings() error {
	C.HPDF_UseCNTEncodings(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) UseCNSEncodings() error {
	C.HPDF_UseCNSEncodings(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) UseUTFEncodings() error {
	if C.HPDF_UTF_SUPPORT == 0 {
		return errors.New("UTF Encoding is not supported")
	}
	C.HPDF_UseUTFEncodings(pdf.doc)
	return pdf.GetLastError()
}

type Encoder struct {
	encoding C.HPDF_Encoder
}
