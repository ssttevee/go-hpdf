package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

func (pdf *PDF) LoadPngImageFromFile(filename string) (*Image, error) {
	cfilename := C.CString(filename)
	cimage := C.HPDF_LoadPngImageFromFile(pdf.doc, cfilename)
	C.free(unsafe.Pointer(cfilename))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadPngImageFromFile2(filename string) (*Image, error) {
	cfilename := C.CString(filename)
	cimage := C.HPDF_LoadPngImageFromFile2(pdf.doc, cfilename)
	C.free(unsafe.Pointer(cfilename))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadRaw1BitImageFromMem(mem []byte, width, height, stride uint32, blackIs1, topIsFirst bool) (*Image, error) {
	ptr := (*C.HPDF_BYTE)((unsafe.Pointer(&mem[0])))
	var blackIs1Int C.int = 0
	if blackIs1 {
		blackIs1Int = 1
	}
	var topIsFirstInt C.int = 0
	if topIsFirst {
		topIsFirstInt = 1
	}
	cimage := C.HPDF_Image_LoadRaw1BitImageFromMem(pdf.doc, ptr, C.HPDF_UINT(width), C.HPDF_UINT(height), C.HPDF_UINT(stride), blackIs1Int, topIsFirstInt)

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadPngImageFromMem(mem []byte) (*Image, error) {
	ptr := (*C.HPDF_BYTE)((unsafe.Pointer(&mem[0])))
	cimage := C.HPDF_LoadPngImageFromMem(pdf.doc, ptr, C.HPDF_UINT(uint32(len(mem))))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadRawImageFromFile(
	filename string, width, height uint32, colorSpace ColorSpace,
) (*Image, error) {
	cfilename := C.CString(filename)
	cimage := C.HPDF_LoadRawImageFromFile(
		pdf.doc, cfilename,
		C.HPDF_UINT(width), C.HPDF_UINT(height),
		C.HPDF_ColorSpace(colorSpace),
	)
	C.free(unsafe.Pointer(cfilename))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadRawImageFromMem(
	mem []byte,
	width uint32,
	height uint32,
	colorSpace ColorSpace,
	bitsPerComponent uint32,
) (*Image, error) {
	ptr := (*C.HPDF_BYTE)((unsafe.Pointer(&mem[0])))
	cimage := C.HPDF_LoadRawImageFromMem(
		pdf.doc,
		ptr,
		C.HPDF_UINT(width),
		C.HPDF_UINT(height),
		C.HPDF_ColorSpace(colorSpace),
		C.HPDF_UINT(bitsPerComponent),
	)

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadJpegImageFromFile(filename string) (*Image, error) {
	cfilename := C.CString(filename)
	cimage := C.HPDF_LoadJpegImageFromFile(pdf.doc, cfilename)
	C.free(unsafe.Pointer(cfilename))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadJpegImageFromMem(mem []byte) (*Image, error) {
	ptr := (*C.HPDF_BYTE)((unsafe.Pointer(&mem[0])))
	cimage := C.HPDF_LoadJpegImageFromMem(pdf.doc, ptr, C.HPDF_UINT(uint32(len(mem))))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

type Image struct {
	image C.HPDF_Image
	pdf   *PDF
}

func newImage(image C.HPDF_Image, pdf *PDF) *Image {
	return &Image{image, pdf}
}

func (image *Image) GetSize() *Point {
	return pointFromHPDFPoint(C.HPDF_Image_GetSize(image.image))
}

func (image *Image) GetWidth() uint32 {
	return uint32(C.HPDF_Image_GetWidth(image.image))
}

func (image *Image) GetHeight() uint32 {
	return uint32(C.HPDF_Image_GetHeight(image.image))
}

func (image *Image) GetBitsPerComponent() uint32 {
	return uint32(C.HPDF_Image_GetBitsPerComponent(image.image))
}

func (image *Image) GetColorSpace() ColorSpace {
	ccolorSpace := C.HPDF_Image_GetColorSpace(image.image)

	if ccolorSpace == nil {
		return CS_UNKNOWN
	} else {
		colorSpace := C.GoString(ccolorSpace)

		switch colorSpace {
		case "DeviceGray":
			return CS_DEVICE_GRAY
		case "DeviceRGB":
			return CS_DEVICE_RGB
		case "DeviceCMYK":
			return CS_DEVICE_CMYK
		case "Indexed":
			return CS_INDEXED
		default:
			return CS_UNKNOWN
		}
	}
}

func (image *Image) SetColorMask(
	rmin, rmax, gmin, gmax, bmin, bmax uint32,
) error {
	C.HPDF_Image_SetColorMask(
		image.image,
		C.HPDF_UINT(rmin), C.HPDF_UINT(rmin),
		C.HPDF_UINT(gmin), C.HPDF_UINT(gmin),
		C.HPDF_UINT(bmin), C.HPDF_UINT(bmin),
	)
	return image.pdf.GetLastError()
}

func (image *Image) SetMaskImage(mask *Image) error {
	C.HPDF_Image_SetMaskImage(image.image, mask.image)
	return image.pdf.GetLastError()
}
